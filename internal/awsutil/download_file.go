package awsutil

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

func DownloadFile(awsSession *session.Session, bucket, key string) ([]byte, error) {
	// create a new AWS S3 downloader
	downloader := s3manager.NewDownloader(awsSession)

	// init file buffer
	buf := aws.NewWriteAtBuffer([]byte{})

	// download the item from the bucket
	_, err := downloader.Download(
		buf,
		&s3.GetObjectInput{
			Bucket: aws.String(bucket),
			Key:    aws.String(key),
		},
	)

	// return file bytes and possible error
	return buf.Bytes(), err
}
