package repo

import (
	"context"
	"errors"
	"google.golang.org/grpc"
	"io"
	"pixstall-user/app/file/repo/grpc/proto"
	"pixstall-user/domain/file"
	"pixstall-user/domain/file/model"
)

type FileGRPCClientConn grpc.ClientConn

type grpcFileRepository struct {
	grpcConn *FileGRPCClientConn
}

func NewGRPCFileRepository(grpcConn *FileGRPCClientConn) file.Repo {
	return &grpcFileRepository{
		grpcConn: grpcConn,
	}
}

func (g grpcFileRepository) SaveFile(ctx context.Context, file model.File, fileType model.FileType, ownerID string, acl []string) (*string, error) {
	conn := (*grpc.ClientConn)(g.grpcConn)
	client := proto.NewFileServiceClient(conn)

	stream, err := client.SaveFile(ctx)
	if err != nil {
		return nil, err
	}
	gFileType, err := g.gRPCFileTypeFormDomain(fileType)
	if err != nil {
		return nil, err
	}
	req := &proto.SaveFileRequest{
		Data: &proto.SaveFileRequest_MetaData{
			MetaData: &proto.MetaData{
				FileType: gFileType,
				Name:     file.Name,
				Owner: ownerID,
				Acl: acl,
			},
		},
	}
	err = stream.SendMsg(req)
	if err != nil {
		return nil, err
	}
	buffer := make([]byte, 1024)

	for {
		n, err := file.File.Read(buffer)
		if err == io.EOF {
			break
		}
		req := &proto.SaveFileRequest{
			Data: &proto.SaveFileRequest_File{
				File: buffer[:n],
			},
		}
		err = stream.SendMsg(req)
	}
	resp, err := stream.CloseAndRecv()
	if err != nil {
		return nil, err
	}
	return &resp.Path, nil
}

func (g grpcFileRepository) SaveFiles(ctx context.Context, files []model.File, fileType model.FileType, ownerID string, acl []string) ([]string, error) {
	panic("implement me")
}

func (g grpcFileRepository) gRPCFileTypeFormDomain(dFileType model.FileType) (proto.MetaData_FileType, error) {
	switch dFileType {
	case model.FileTypeMessage:
		return proto.MetaData_Message, nil
	case model.FileTypeCompletion:
		return proto.MetaData_Completion, nil
	case model.FileTypeCommissionRef:
		return proto.MetaData_CommissionRef, nil
	case model.FileTypeCommissionProofCopy:
		return proto.MetaData_CommissionProofCopy, nil
	case model.FileTypeArtwork:
		return proto.MetaData_Artwork, nil
	case model.FileTypeRoof:
		return proto.MetaData_Roof, nil
	case model.FileTypeOpenCommission:
		return proto.MetaData_OpenCommission, nil
	case model.FileTypeProfile:
		return proto.MetaData_Profile, nil
	default:
		return -1, errors.New("not found")
	}
}