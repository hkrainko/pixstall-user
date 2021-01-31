//+build wireinject

package main

import (
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/google/wire"
	"github.com/streadway/amqp"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
	auth_deliv "pixstall-user/app/auth/delivery/http"
	auth_repo "pixstall-user/app/auth/repo/grpc"
	auth_ucase "pixstall-user/app/auth/usecase"
	image_repo "pixstall-user/app/image/aws-s3"
	reg_deliv "pixstall-user/app/reg/delivery/http"
	reg_ucase "pixstall-user/app/reg/usecase"
	token_repo "pixstall-user/app/token/repo/kong-jwt"
	user_deliv "pixstall-user/app/user/delivery/http"
	user_msg_broker "pixstall-user/app/user/msg-broker/rabbitmq"
	user_repo "pixstall-user/app/user/repo/mongo"
	user_ucase "pixstall-user/app/user/usecase"
)

func InitAuthController(grpcConn *grpc.ClientConn, db *mongo.Database) auth_deliv.AuthController {
	wire.Build(
		auth_deliv.NewAuthController,
		auth_ucase.NewAuthUseCase,
		auth_repo.NewGRPCAuthRepository,
		user_repo.NewMongoUserRepo,
		token_repo.NewKongJWTTokenRepo,
	)
	return auth_deliv.AuthController{}
}

func InitRegController(grpcConn *grpc.ClientConn, db *mongo.Database, ch *amqp.Channel, awsS3 *s3.S3) reg_deliv.RegController {
	wire.Build(
		reg_deliv.NewRegController,
		user_repo.NewMongoUserRepo,
		reg_ucase.NewRegUseCase,
		user_msg_broker.NewRabbitMQUserMsgBroker,
		image_repo.NewAWSS3ImageRepository,
		token_repo.NewKongJWTTokenRepo,
	)
	return reg_deliv.RegController{}
}

func InitUserController(grpcConn *grpc.ClientConn, db *mongo.Database, awsS3 *s3.S3) user_deliv.UserController {
	wire.Build(
		user_deliv.NewUserController,
		user_ucase.NewUserUseCase,
		user_repo.NewMongoUserRepo,
		image_repo.NewAWSS3ImageRepository,
		token_repo.NewKongJWTTokenRepo,
	)
	return user_deliv.UserController{}
}
