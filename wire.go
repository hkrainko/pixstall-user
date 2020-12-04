//+build wireinject

package main

import (
	"github.com/google/wire"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
	token_repo "pixstall-user/app/token/repo/jwt"
	auth_deliv "pixstall-user/app/auth/delivery/http"
	auth_repo "pixstall-user/app/auth/repository/grpc"
	auth_ucase "pixstall-user/app/auth/usecase"
	reg_deliv "pixstall-user/app/reg/delivery/http"
	reg_ucase "pixstall-user/app/reg/usecase"
	user_deliv "pixstall-user/app/user/delivery/http"
	user_ucase "pixstall-user/app/user/usecase"
	user_repo "pixstall-user/app/user/repository/mongo"
)

func InitAuthController(grpcConn *grpc.ClientConn, db *mongo.Database) auth_deliv.AuthController {
	wire.Build(
		auth_deliv.NewAuthController,
		auth_ucase.NewAuthUseCase,
		auth_repo.NewGRPCAuthRepository,
		user_repo.NewMongoUserRepo,
		token_repo.NewJWTTokenRepo,
		)
	return auth_deliv.AuthController{}
}

func InitRegController(grpcConn *grpc.ClientConn, db *mongo.Database) reg_deliv.RegController {
	wire.Build(
		reg_deliv.NewRegController,
		user_repo.NewMongoUserRepo,
		reg_ucase.NewRegUseCase,
	)
	return reg_deliv.RegController{}
}

func InitUserController(grpcConn *grpc.ClientConn, db *mongo.Database) user_deliv.UserController {
	wire.Build(
		user_deliv.NewUserController,
		user_ucase.NewUserUseCase,
		user_repo.NewMongoUserRepo,
	)
	return user_deliv.UserController{}
}