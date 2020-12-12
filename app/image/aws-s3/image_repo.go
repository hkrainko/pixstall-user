package aws_s3

import (
	"bytes"
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"image"
	"image/png"
	domainImage "pixstall-user/domain/image"
)
import "github.com/aws/aws-sdk-go/service/s3"

type awsS3ImageRepository struct {
	s3 *s3.S3
}

const (
	BucketName = "pixstall-store-dev"
)

func NewAWSS3ImageRepository(s3 *s3.S3) domainImage.Repo {
	return &awsS3ImageRepository{
		s3: s3,
	}
}

func (a awsS3ImageRepository) SaveImage(ctx context.Context, path string, imageName string, image image.Image) error {
	// create buffer
	buff := new(bytes.Buffer)

	// encode image to buffer
	err := png.Encode(buff, image)
	if err != nil {
		fmt.Println("failed to create buffer", err)
	}

	// convert buffer to reader
	reader := bytes.NewReader(buff.Bytes())

	// use it in `PutObjectInput`
	_, err = a.s3.PutObjectWithContext(ctx, &s3.PutObjectInput{
		Bucket: aws.String(BucketName),
		Key:    aws.String(path + imageName),
		Body:   reader,
		ContentType: aws.String("image"),
		ACL: aws.String("public-read"),  //profile should be public accessible
	})

	if err != nil {
		return err
	}
	return nil
}
