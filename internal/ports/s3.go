package ports

import "io"

type S3Client interface {
	UploadFile(bucketName, key string, file io.Reader) (string, error)
}

