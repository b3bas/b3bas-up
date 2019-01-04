/*
	https://www.youtube.com/watch?v=iOGIKG3EptI
	https://github.com/awslabs/aws-go-wordfreq-sample/blob/master/cmd/uploads3/main.go

	https://docs.aws.amazon.com/sdk-for-go/api/aws/
	- first configure your aws credentials run: aws configure
	- go get -u github.com/aws/aws-sdk-go/aws
	- login to UI web aws s3 interface
	- go to S3 service
	- create a Bucket called `b3bas-up` in the desired region (default region: `ap-southeast-1`)
	- run:  go run main.go b3bas-up [file_name]
*/

package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/dimiro1/banner"
	_ "github.com/dimiro1/banner/autoload"
	colorable "github.com/mattn/go-colorable"
	"github.com/sirupsen/logrus"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

var (
	TEST_AWS_ACCESS_KEY string = os.Getenv("AWS_ACCESS_KEY")
	TEST_AWS_SECRET_KEY string = os.Getenv("AWS_SECRET_KEY")
	TEST_AWS_TOKEN      string = os.Getenv("AWS_TOKEN")
	TEST_S3_REGION      string = "ap-southeast-1"
	TEST_S3_BUCKET      string = "b3bas-up"
)

const B3BAS_LOG_TEST = "/var/log/b3bas-up/b3bas-up.log"

func InitLogoTest() {
	isEnabled := true
	isColorEnabled := true
	banner.Init(colorable.NewColorableStdout(), isEnabled, isColorEnabled, bytes.NewBufferString(" B3BAS Golang Framework {{ .AnsiColor.Green }}(Running){{ .AnsiColor.Default }}\n"))
}

func SaveLogTest() {
	// create the logger
	logger := logrus.New()
	logger.Formatter = &logrus.JSONFormatter{}

	fName, err := os.OpenFile(B3BAS_LOG_TEST, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0755)
	if err != nil {
		logger.Fatal(err)
	}
	defer fName.Close()
	// multiwriter simultaneously
	logger.SetOutput(io.MultiWriter(os.Stdout, fName))
	println(" Log file saved to:", B3BAS_LOG_TEST)
}

func UploadTest() {
	if len(os.Args) != 3 {
		fmt.Printf("usage: %s <bucket> <filename>\n", filepath.Base(os.Args[0]))
		os.Exit(1)
	}

	bucket := os.Args[1]
	if bucket == "" {
		bucket = TEST_S3_BUCKET
	}
	filename := os.Args[2]

	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Failed to open file", filename, err)
		os.Exit(1)
	}
	defer file.Close()

	//select Region to use.
	conf := aws.Config{Region: aws.String(TEST_S3_REGION)}
	sess := session.New(&conf)
	svc := s3manager.NewUploader(sess)

	fmt.Println("Uploading file to S3...")
	result, err := svc.Upload(&s3manager.UploadInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(filepath.Base(filename)),
		Body:   file,
	})
	if err != nil {
		fmt.Println("error", err)
		os.Exit(1)
	}

	fmt.Printf("Successfully uploaded %s to %s\n", filename, result.Location)
}
