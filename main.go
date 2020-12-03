package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
	"log"
	"time"
)

func main() {
	//lis, err := net.Listen("tcp", ":50051")
	//r := gin.Default()
	//
	//s := grpc.NewServer()
	//
	//pb.RegisterAuthServiceServer(s, InitAuthController())
	//
	//if err := s.Serve(lis); err != nil {
	//	log.Fatalf("failed to listen: %v", err)
	//}
	//
	//err = r.Run(":9002")
	//fmt.Println(err)

	ctx, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
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

	//gRPC
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	r := gin.Default()

	authGroup := r.Group("/auth")
	{
		ctr := InitAuthController(conn, dbClient.Database("pixstall-user"))
		authGroup.POST("/getAuthUrl", ctr.GetAuthURL)
		authGroup.GET("/authCallback", ctr.AuthCallback)
	}

	err = r.Run(":9001")
	print(err)
}