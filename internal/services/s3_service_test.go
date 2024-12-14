package services_test

import (
	"bytes"
	"testing"
	"todo-service/internal/mocks"
	"todo-service/internal/services"

	"github.com/stretchr/testify/assert"
)

func TestUploadFileToS3(t *testing.T) {
	mockS3 := mocks.NewMockS3Client()
	s3Service := services.NewS3Service(mockS3)

	file := bytes.NewReader([]byte("test file content"))
	fileID, err := s3Service.UploadFile("test.txt", file)

	assert.NoError(t, err)
	assert.NotEmpty(t, fileID)
	assert.Equal(t, 1, len(mockS3.Uploads))
	assert.Equal(t, "test.txt", mockS3.Uploads[0].Key)
}

