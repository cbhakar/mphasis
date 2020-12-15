package s3

import (
	"errors"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/cbhakar/mphasis/store"
	"log"
	"mime/multipart"
	"time"
)

const (
	bucketName      = ""
	accessKeyId     = ""
	accessSecretKey = ""
	region          = ""
)

var (
	s *session.Session
)

func init() {
	var err error
	// Create AWS Session
	s, err = session.NewSession(&aws.Config{
		Region:      aws.String(region),
		Credentials: credentials.NewStaticCredentials(accessKeyId, accessSecretKey, ""),
	})
	if err != nil {
		panic(err)
	}
}

func UploadFileToS3(fileName string, fileContent multipart.File) error {
	var err error
	// Create AWS Session
	s3Svc := s3.New(s)
	_, err = s3Svc.PutObject(&s3.PutObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(fileName),
		Body:   fileContent,
	})
	if err != nil {
		log.Printf("error updating image to s3, err :- %s", err.Error())
		return errors.New("error updating image to s3")
	}
	err = store.AddImage(store.Image{
		ImageID:   0,
		ImageName: fileName,
		CreatedAt: time.Now().Format("2006-01-02 15:04:05"),
	})
	if err != nil {
		log.Printf("error updating record to db :- %s", err.Error())
		return errors.New("error updating record to data base")
	}
	return err
}
