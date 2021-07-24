package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func main() {
	sess := s3.New(session.New(), &aws.Config{Region: aws.String("eu-central-1")})

	params := &s3.ListObjectsInput{
		Bucket: aws.String("s17-playground"),
	}

	resp, err := sess.ListObjects(params)
	if err != nil {
		fmt.Println(err)
	}

	for _, key := range resp.Contents {
		fmt.Println(*key.Key)
	}
}
