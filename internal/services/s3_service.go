package services

import (
	"bytes"
	"fmt"
	"io"
	"todo-service/internal/ports"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

type S3Service struct {
	client *s3.S3
}

func NewS3Service(cfg *Config) ports.S3Client {
	sess := session.Must(session.NewSession())
	return &S3Service{
		client: s3.New(sess),
	}
}

func (s *S3Service) UploadFile(bucketName, key string, file io.Reader) (string, error) {
	buf := new(bytes.Buffer)
	_, err := io.Copy(buf, file)
	if err != nil {
		return "", err
	}
	_, err = s.client.PutObject(&s3.PutObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(key),
		Body:   bytes.NewReader(buf.Bytes()),
	})
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s/%s", bucketName, key), nil
}

