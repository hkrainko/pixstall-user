package main

import (
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"log"
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

	//gRPC
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	r := gin.Default()

	authGroup := r.Group("/auth")
	{
		ctr := InitAuthController(conn)
		authGroup.POST("/getAuthUrl", ctr.GetAuthURL)
		authGroup.GET("/authCallback", ctr.AuthCallback)
	}

	err = r.Run(":9001")
	print(err)
}