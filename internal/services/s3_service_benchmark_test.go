package services_test

import (
	"bytes"
	"testing"
	"todo-service/internal/mocks"
	"todo-service/internal/services"
)

func BenchmarkUploadFileToS3(b *testing.B) {
	mockS3 := mocks.NewMockS3Client()
	s3Service := services.NewS3Service(mockS3)

	file := bytes.NewReader([]byte("benchmark file content"))

	for i := 0; i < b.N; i++ {
		s3Service.UploadFile("benchmark.txt", file)
	}
}

