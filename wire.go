//+build wireinject

package main

import (
	"github.com/google/wire"
	"google.golang.org/grpc"
	auth_deliv "pixstall-user/app/auth/delivery/http"
	auth_repo "pixstall-user/app/auth/repository/grpc"
	auth_ucase "pixstall-user/app/auth/usecase"
)

func InitAuthController(grpcConn *grpc.ClientConn) auth_deliv.AuthController {
	wire.Build(auth_deliv.NewAuthController, auth_ucase.NewAuthUseCase, auth_repo.NewGRPCAuthRepository)
	return auth_deliv.AuthController{}
}