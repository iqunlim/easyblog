package repository

import (
	"context"
	"io"
	"os"
)


type ImageRepository interface {
	Upload(ctx context.Context, fileReader io.Reader, fileName string) (string, error)
	Delete(ctx context.Context, fileName string) error
	Download(ctx context.Context, fileName string) (io.ReadCloser, error)
}

type ImageRepositoryLocalhost struct {
	basePath string
}

func NewImageRepositoryLocalhost(basePath string) ImageRepository {
	if basePath[len(basePath)-1] != '/' {
		basePath = basePath + "/"
	}
	return &ImageRepositoryLocalhost{
		basePath: basePath,
	}
}

func (r *ImageRepositoryLocalhost) Upload(ctx context.Context, fileReader io.Reader, fileName string) (string, error) {

	var filePath = r.basePath + fileName

	file, err := os.Create(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	_, err = io.Copy(file, fileReader)
	if err != nil {
		return "", err
	}
	
	return filePath, nil

}

func (r *ImageRepositoryLocalhost) Delete(ctx context.Context, fileName string) error {
	
	var filePath = r.basePath + fileName
	return os.Remove(filePath)
}

func (r *ImageRepositoryLocalhost) Download(ctx context.Context, fileName string) (io.ReadCloser, error) {

	var filePath = r.basePath + fileName
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	return file, nil
}