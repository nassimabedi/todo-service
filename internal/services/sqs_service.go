package services

import (
	"encoding/json"
	"fmt"
	"log"
	"todo-service/internal/domain"
	"todo-service/internal/ports"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

type SQSService struct {
	client   *sqs.SQS
	queueURL string
}

func NewSQSService(cfg *Config) ports.SQSClient {
	sess := session.Must(session.NewSession())
	client := sqs.New(sess)

	return &SQSService{
		client:   client,
		queueURL: cfg.SQSQueueURL,
	}
}

func (s *SQSService) SendMessage(queueURL string, todo *domain.Todo) error {
	message, err := json.Marshal(todo)
	if err != nil {
		return fmt.Errorf("failed to marshal todo: %v", err)
	}

	_, err = s.client.SendMessage(&sqs.SendMessageInput{
		QueueUrl:    aws.String(queueURL),
		MessageBody: aws.String(string(message)),
	})
	if err != nil {
		return fmt.Errorf("failed to send message to SQS: %v", err)
	}

	log.Printf("Message sent to SQS: %s", string(message))
	return nil
}

