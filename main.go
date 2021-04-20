package main

import (
	"context"
	cors "github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/streadway/amqp"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
	"log"
	"pixstall-user/app/file/repo"
	"pixstall-user/app/middleware"
	"strings"
	"time"
)

func main() {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	//MongoDB
	dbClient, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		panic(err)
	}
	defer cancel()
	defer func() {
		if err = dbClient.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()
	db := dbClient.Database("pixstall-user")

	//RabbitMQ
	rabbitmqConn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ %v", err)
	}
	defer rabbitmqConn.Close()
	ch, err := rabbitmqConn.Channel()
	if err != nil {
		log.Fatalf("Failed to create channel %v", err)
	}
	err = ch.ExchangeDeclare(
		"user",
		"topic",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("Failed to create exchange %v", err)
	}

	commMsgBroker := InitCommissionMessageBroker(db, rabbitmqConn)
	go commMsgBroker.StartCommUsersValidateQueue()
	defer commMsgBroker.StopAllQueue()

	// gRPC - Auth
	authGRPCConn, err := grpc.Dial("localhost:50051", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatal(err)
	}
	defer authGRPCConn.Close()

	//gRPC - File
	fileGRPCConn, err := grpc.Dial("localhost:50052", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatal(err)
	}
	defer fileGRPCConn.Close()

	r := gin.Default()
	//r.Use(cors.Default())
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Access-Control-Allow-Origin", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowWildcard: true,
		AllowFiles: true,
		MaxAge: 12 * time.Hour,
	}))

	apiGroup := r.Group("/api")

	authGroup := apiGroup.Group("/auth")
	{
		ctr := InitAuthController(authGRPCConn, db)
		authGroup.GET("/url", ctr.GetAuthURL)
		authGroup.GET("/callback", ctr.AuthCallback)
	}

	regGroup := apiGroup.Group("/reg")
	{
		authIDExtractor := middleware.NewJWTPayloadsExtractor([]string{"authId"})
		ctr := InitRegController(authGRPCConn, (*repo.FileGRPCClientConn)(fileGRPCConn), db, ch, rabbitmqConn)
		regGroup.POST("", authIDExtractor.ExtractPayloadsFromJWT, ctr.Registration)
	}

	userGroup := apiGroup.Group("/users")
	{
		userIDExtractor := middleware.NewJWTPayloadsExtractor([]string{"userId"})
		ctr := InitUserController(authGRPCConn, (*repo.FileGRPCClientConn)(fileGRPCConn), db)
		userGroup.GET("/:id", func(c *gin.Context) {
			if strings.HasSuffix(c.Request.RequestURI, "/me") {
				userIDExtractor.ExtractPayloadsFromJWT(c)
				if c.IsAborted() {
					return
				}
				ctr.GetUserDetails(c)
				return
			}
			ctr.GetUser(c)
		})
		userGroup.PATCH("/:id", userIDExtractor.ExtractPayloadsFromJWT, ctr.UpdateUser)
	}

	err = r.Run(":9001")
	print(err)
}
