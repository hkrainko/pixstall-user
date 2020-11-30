//+build wireinject

package main

import (
	"github.com/google/wire"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
	auth_deliv "pixstall-user/app/auth/delivery/http"
	auth_repo "pixstall-user/app/auth/repository/grpc"
	auth_ucase "pixstall-user/app/auth/usecase"
	user_repo "pixstall-user/app/user/repository/mongo"
)

func InitAuthController(grpcConn *grpc.ClientConn, dbClient *mongo.Client) auth_deliv.AuthController {
	wire.Build(auth_deliv.NewAuthController, auth_ucase.NewAuthUseCase, auth_repo.NewGRPCAuthRepository, user_repo.NewMongoUserRepo)
	return auth_deliv.AuthController{}
}