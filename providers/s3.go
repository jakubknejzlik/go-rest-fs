package providers

import (
	"context"
	"io"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
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
	ctx := context.Background()
	uploader := s3manager.NewUploaderWithClient(p.S3)

	_, err := uploader.UploadWithContext(ctx, &s3manager.UploadInput{
		Bucket: aws.String(p.Bucket),
		Key:    aws.String(name),
		Body:   r,
	})

	return err
}

// DeleteFile ...
func (p *S3StorageProvider) DeleteFile(name string) error {
	_, err := p.S3.DeleteObject(&s3.DeleteObjectInput{
		Bucket: aws.String(p.Bucket),
		Key:    aws.String(name),
	})
	return err
}
