package main

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/endpoints"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	cors "github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/streadway/amqp"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
	"log"
	"pixstall-user/app/middleware"
	"time"
)

func main() {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	//AWS s3
	awsAccessKey := "AKIA5BWICLKRWX6ARSEF"
	awsSecret := "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"
	token := ""
	creds := credentials.NewStaticCredentials(awsAccessKey, awsSecret, token)
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		Config: aws.Config{
			Region:                        aws.String(endpoints.ApEast1RegionID),
			CredentialsChainVerboseErrors: aws.Bool(true),
			Credentials:                   creds,
		},
		//Profile:                 "default", //[default], use [prod], [uat]
		//SharedConfigState:       session.SharedConfigEnable,
	}))
	awsS3 := s3.New(sess)

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

	//gRPC
	grpcConn, err := grpc.Dial("localhost:50051", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatal(err)
	}
	defer grpcConn.Close()

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
		ctr := InitAuthController(grpcConn, dbClient.Database("pixstall-user"))
		authGroup.GET("/url", ctr.GetAuthURL)
		authGroup.GET("/callback", ctr.AuthCallback)
	}

	regGroup := apiGroup.Group("/reg")
	{
		authIDExtractor := middleware.NewJWTPayloadsExtractor([]string{"authId"})
		ctr := InitRegController(grpcConn, dbClient.Database("pixstall-user"), ch, awsS3)
		regGroup.POST("/registration", authIDExtractor.ExtractPayloadsFromJWT, ctr.Registration)
	}

	userGroup := apiGroup.Group("/users")
	{
		userIDExtractor := middleware.NewJWTPayloadsExtractor([]string{"userId"})
		ctr := InitUserController(grpcConn, dbClient.Database("pixstall-user"), awsS3)
		userGroup.GET("/:id", ctr.GetUser)
		userGroup.GET("/:id/details", userIDExtractor.ExtractPayloadsFromJWT, ctr.GetUserDetails)
		userGroup.PATCH("/:id", userIDExtractor.ExtractPayloadsFromJWT, ctr.UpdateUser)
	}

	err = r.Run(":9001")
	print(err)
}
