package grpc

import (
	"context"
	pb "github.com/red-auth/proto"
	"google.golang.org/grpc"
	"log"
	"pixstall-user/domain/auth"
)

type grpcAuthRepository struct {
	grpcConn *grpc.ClientConn
}

func NewGRPCAuthRepository(grpcConn *grpc.ClientConn) auth.Repo {
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

func (g grpcAuthRepository) GetAuthorizedUserInfo(ctx context.Context, authCallBack string) (*string, error) {
	panic("implement me")
}