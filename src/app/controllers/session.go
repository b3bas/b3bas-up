package controllers

import (
	"os"

	"github.com/aws/aws-sdk-go/aws/credentials"
)

var (
	AWS_ACCESS_KEY string = os.Getenv("AWS_ACCESS_KEY")
	AWS_SECRET_KEY string = os.Getenv("AWS_SECRET_KEY")
	AWS_TOKEN      string = os.Getenv("AWS_TOKEN")
)

func GetSession() {
	creds := credentials.NewStaticCredentials(AWS_ACCESS_KEY, AWS_SECRET_KEY, AWS_TOKEN)
	_, err := creds.Get()

	if err != nil {
		// handle error
		println("Error Credential... Application Terminated !")
	}
}
