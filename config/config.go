package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Config struct {
	DBURL       string
	S3Bucket    string
	SQSQueueURL string
}

func LoadConfig() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, using default values.")
	}
	return &Config{
		DBURL:       os.Getenv("DB_URL"),
		S3Bucket:    os.Getenv("S3_BUCKET"),
		SQSQueueURL: os.Getenv("SQS_QUEUE_URL"),
	}, nil
}

func InitDB(cfg *Config) *gorm.DB {
	db, err := gorm.Open(postgres.Open(cfg.DBURL), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to the database: %v", err)
	}
	return db
}

