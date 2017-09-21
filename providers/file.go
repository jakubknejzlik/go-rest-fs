package providers

import (
	"io"
	"os"
	"path"
)

// FileStorageProvider ...
type FileStorageProvider struct {
	Dir string
}

// NewFileStorageProvider ...
func NewFileStorageProvider(dir string) *FileStorageProvider {
	return &FileStorageProvider{Dir: dir}
}

// UploadFile ...
func (p *FileStorageProvider) UploadFile(r io.ReadCloser, name string) error {
	outFile, err := os.Create(path.Join(p.Dir, name))
	if err != nil {
		return err
	}
	defer outFile.Close()
	_, err = io.Copy(outFile, r)
	return err
}
