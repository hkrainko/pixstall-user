package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/streadway/amqp"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
	"log"
	"time"
)

func main() {
	//MongoDB
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
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

	authGroup := r.Group("/auth")
	{
		ctr := InitAuthController(grpcConn, dbClient.Database("pixstall-user"))
		authGroup.POST("/getAuthUrl", ctr.GetAuthURL)
		authGroup.GET("/authCallback", ctr.AuthCallback)
	}

	regGroup := r.Group("/reg")
	{
		ctr := InitRegController(grpcConn, dbClient.Database("pixstall-user"), ch)
		regGroup.POST("/register", ctr.Registration)
	}

	//userGroup := r.Group("/user")
	//{
	//	ctr := InitUserController(conn, dbClient.Database("pixstall-user"))
	//	regGroup.POST("/")
	//}

	err = r.Run(":9001")
	print(err)
}
