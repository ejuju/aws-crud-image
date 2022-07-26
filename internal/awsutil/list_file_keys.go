package awsutil

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func ListFileKeys(awsSession *session.Session, region, bucket string) ([]string, error) {
	keys := []string{}

	params := &s3.ListObjectsInput{Bucket: aws.String(bucket)}
	resp, err := s3.New(awsSession, &aws.Config{Region: aws.String(region)}).ListObjects(params)
	if err != nil {
		return nil, err
	}

	for _, key := range resp.Contents {
		keys = append(keys, *key.Key)
	}
	return keys, nil
}
