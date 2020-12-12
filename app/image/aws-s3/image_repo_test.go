package aws_s3

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/endpoints"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/stretchr/testify/assert"
	"image"
	"image/png"
	"log"
	"os"
	domainImage "pixstall-user/domain/image"
	"testing"
)

var repo domainImage.Repo
var sess *session.Session
var awsS3 *s3.S3
var ctx context.Context

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	teardown()
	os.Exit(code)
}

func setup() {
	fmt.Println("Before all tests")
	ctx = context.Background()

	awsAccessKey := "AKIA5BWICLKRWX6ARSEF"
	awsSecret := "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"
	token := ""
	creds := credentials.NewStaticCredentials(awsAccessKey, awsSecret, token)
	_, err := creds.Get()
	if err != nil {
		log.Fatal(err)
	}

	sess = session.Must(session.NewSessionWithOptions(session.Options{
		Config: aws.Config{
			Region:                        aws.String(endpoints.ApEast1RegionID),
			CredentialsChainVerboseErrors: aws.Bool(true),
			Credentials:                   creds,
		},
		//Profile:                 "default", //[default], use [prod], [uat]
		//SharedConfigState:       session.SharedConfigEnable,
	}))
	awsS3 = s3.New(sess)
	repo = NewAWSS3ImageRepository(awsS3)
}

func teardown() {
	dropAll()
	fmt.Println("After all tests")
}

func TestAwsS3ImageRepository_SaveImage(t *testing.T) {
	dropAll()
	testImage, err := loadImageFromPath("./test_image.png")
	assert.NoError(t, err)
	assert.NotNil(t, testImage)

	err = repo.SaveImage(ctx, "test/02/", "test.png", *testImage)
	assert.NoError(t, err)

	resultImage, err := downloadImage("test/02/", "test.png")
	assert.NoError(t, err)
	assert.NotNil(t, resultImage)
}

//Private methods
func dropAll() {

}

func loadImageFromPath(path string) (*image.Image, error) {
	imageFile, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer imageFile.Close()

	fileInfo, err := imageFile.Stat()
	log.Printf("fileInfo name:%v size:%v", fileInfo.Name(), fileInfo.Size())

	//imagedData, imageType, err := image.Decode(imageFile)
	//if err != nil {
	//	return nil, err
	//}
	pngImage, err := png.Decode(imageFile)
	if err != nil {
		return nil, err
	}
	return &pngImage, nil
}

func downloadImage(path string, fileName string) (*image.Image, error) {

	file, err := os.Create(fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	downloader := s3manager.NewDownloader(sess)
	numBytes, err := downloader.Download(file, &s3.GetObjectInput{
		Bucket: aws.String(BucketName),
		Key:    aws.String(path + fileName),
	})
	if err != nil {
		return nil, err
	}
	log.Printf("downloadImage bytes:%v", numBytes)
	pngImage, err := png.Decode(file)
	if err != nil {
		return nil, err
	}
	return &pngImage, nil
}
