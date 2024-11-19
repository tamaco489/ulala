package s3

import (
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

func GenerateSession() *session.Session {
	s := session.Must(session.NewSessionWithOptions(session.Options{
		Profile:           "default",
		SharedConfigState: session.SharedConfigEnable,
	}))

	return s
}

func S3DownLoader(s *session.Session, f *os.File, bucketName, objPath string) (int64, error) {
	d := s3manager.NewDownloader(s)
	m, err := d.Download(f, &s3.GetObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(objPath),
	})

	if err != nil {
		return 0, err
	}

	return m, nil
}

func S3UpLoader(s *session.Session, bucketName, objPath string, f *os.File) (*s3manager.UploadOutput, error) {
	u := s3manager.NewUploader(s)
	m, err := u.Upload(&s3manager.UploadInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(objPath),
		Body:   f,
	})

	if err != nil {
		return nil, err
	}

	return m, nil
}
