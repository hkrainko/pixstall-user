//+build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/streadway/amqp"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
	auth_deliv "pixstall-user/app/auth/delivery/http"
	auth_repo "pixstall-user/app/auth/repo/grpc"
	auth_ucase "pixstall-user/app/auth/usecase"
	comm_ucase "pixstall-user/app/commission/usercase"
	comm_deliv_rabbitmq "pixstall-user/app/commission/delivery/rabbitmq"
	file_repo "pixstall-user/app/file/repo"
	reg_deliv "pixstall-user/app/reg/delivery/http"
	reg_ucase "pixstall-user/app/reg/usecase"
	token_repo "pixstall-user/app/token/repo/kong-jwt"
	user_deliv "pixstall-user/app/user/delivery/http"
	user_repo "pixstall-user/app/user/repo/mongo"
	user_ucase "pixstall-user/app/user/usecase"
	msg_broker_repo "pixstall-user/app/msg-broker/repo/rabbitmq"
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

func InitRegController(grpcConn *grpc.ClientConn, fileGRPCConn *file_repo.FileGRPCClientConn,  db *mongo.Database, ch *amqp.Channel, conn *amqp.Connection) reg_deliv.RegController {
	wire.Build(
		reg_deliv.NewRegController,
		user_repo.NewMongoUserRepo,
		reg_ucase.NewRegUseCase,
		msg_broker_repo.NewRabbitMQMsgBrokerRepo,
		file_repo.NewGRPCFileRepository,
		token_repo.NewKongJWTTokenRepo,
	)
	return reg_deliv.RegController{}
}

func InitUserController(grpcConn *grpc.ClientConn, fileGRPCConn *file_repo.FileGRPCClientConn, db *mongo.Database) user_deliv.UserController {
	wire.Build(
		user_deliv.NewUserController,
		user_ucase.NewUserUseCase,
		user_repo.NewMongoUserRepo,
		file_repo.NewGRPCFileRepository,
		token_repo.NewKongJWTTokenRepo,
	)
	return user_deliv.UserController{}
}

func InitCommissionMessageBroker(db *mongo.Database, conn *amqp.Connection) comm_deliv_rabbitmq.CommissionMessageBroker {
	wire.Build(
		comm_deliv_rabbitmq.NewRabbitMQCommissionMessageBroker,
		comm_ucase.NewCommissionUseCase,
		msg_broker_repo.NewRabbitMQMsgBrokerRepo,
		user_repo.NewMongoUserRepo,
		)
	return comm_deliv_rabbitmq.CommissionMessageBroker{}
}
