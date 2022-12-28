package auth

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
)

const (
	DefaultRegion = "us-west-2"
)

func Session() (*session.Session, error) {
	return session.NewSession(&aws.Config{
		//TODO: make this real
		Region: aws.String(DefaultRegion),
	})

}
