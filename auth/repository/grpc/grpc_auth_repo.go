package grpc

import (
	"context"
	"google.golang.org/grpc"
	"log"
	"pixstall_server/app/domain"
	pb "github.com/red-auth/proto"
)

type grpcAuthRepository struct {
	grpcConn *grpc.ClientConn
}

func NewGRPCAuthRepository(grpcConn *grpc.ClientConn) domain.AuthRepository {
	return &grpcAuthRepository{
		grpcConn: grpcConn,
	}
}

func (g grpcAuthRepository) GetAuthURL(ctx context.Context, authType string) (string, error) {
	client := pb.NewAuthServiceClient(g.grpcConn)

	result, err := client.GetAuthUrl(ctx, &pb.GetAuthUrlRequest{
		AuthType: authType,
	})
	if err != nil {
		log.Println(err)
		return "", err
	}
	return result.AuthUrl, nil
}