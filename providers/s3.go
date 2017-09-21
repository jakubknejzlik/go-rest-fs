package providers

import (
	"io"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

// S3StorageProvider ...
type S3StorageProvider struct {
	S3     *s3.S3
	Bucket string
}

// NewS3StorageProvider ...
func NewS3StorageProvider(bucket string) *S3StorageProvider {
	sess := session.Must(session.NewSession())
	svc := s3.New(sess)
	return &S3StorageProvider{S3: svc, Bucket: bucket}
}

// UploadFile ...
func (p *S3StorageProvider) UploadFile(r io.ReadCloser, name string) error {
	// implement this using s3manager.Uploader()

	// ctx := context.Background()
	// _, err := p.S3.PutObjectWithContext(ctx, &s3.PutObjectInput{
	// 	Bucket: aws.String(p.Bucket),
	// 	Key:    aws.String(name),
	// 	Body:   r,
	// })
	// return err
	return nil
}
