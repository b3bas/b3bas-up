package config

import (
	"time"
)

var (
	S3_REGION string = "ap-southeast-1"
	S3_BUCKET string = "b3bas-up"
)

const (
	// Environment
	ENV_DEVELOPMENT = "development"
)

var CF *Config

type Config struct {
	AwsConfig AWSConfig
	Grace     GraceConfig
	Timeout   time.Duration
}

type AWSConfig struct {
	AwsAccessKey string
	AwsSecretKey string
	AwsToken     string
	S3Region     string
	S3Bucket     string
}

type GraceConfig struct {
	Timeout          string
	HTTPReadTimeout  string
	HTTPWriteTimeout string
}
