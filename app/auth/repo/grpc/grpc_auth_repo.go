package grpc

import (
	"context"
	"google.golang.org/grpc"
	"log"
	"pixstall-user/domain/auth"
	authModel "pixstall-user/domain/auth/model"
	pb "pixstall-user/proto"
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

func (g grpcAuthRepository) GetAuthorizedUserInfo(ctx context.Context, authCallBack authModel.AuthCallback) (*authModel.AuthUserInfo, error) {
	client := pb.NewAuthServiceClient(g.grpcConn)

	result, err := client.CallBack(ctx, &pb.CallBackRequest{
		State:    authCallBack.State,
		Code:     authCallBack.Code,
		AuthType: authCallBack.AuthType,
	})

	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &authModel.AuthUserInfo{
		ID:       result.AuthType + "@" + result.Id,
		AuthType: result.AuthType,
		UserName: result.Name,
		Email:    result.Email,
		Birthday: result.Birthday,
		Gender:   result.Gender,
		PhotoURL: result.PhotoUrl,
	}, nil
}
