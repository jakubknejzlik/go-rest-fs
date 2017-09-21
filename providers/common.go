package providers

import (
	"io"
	"net/url"
)

// StorageProvider ...
type StorageProvider interface {
	UploadFile(r io.ReadCloser, name string) error
}

// GetProvider ...
func GetProvider(storageURL string) (*StorageProvider, error) {
	var sp StorageProvider
	s, err := url.Parse(storageURL)

	if err != nil {
		return &sp, err
	}

	switch s.Scheme {
	case "file":
		sp = NewFileStorageProvider(s.Path)
	case "s3":
		sp = NewS3StorageProvider(s.Path)
	}
	return &sp, nil
}
