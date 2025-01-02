package service

import (
	"context"
	"fmt"
	"io"
	"mime/multipart"

	"github.com/google/uuid"
	"github.com/iqunlim/easyblog/repository"
)



type ImageHandlerService interface {
	Upload(ctx context.Context, fileReader *multipart.FileHeader) (string, error)
	Delete(ctx context.Context, fileName string) error
	Download(ctx context.Context, fileName string) (io.ReadCloser, error)
}


func NewImageService(imagerepository repository.ImageRepository) ImageHandlerService {

	return &ImageHandlerServiceImpl {
		imagerepository: imagerepository,
	}
}

type ImageHandlerServiceImpl struct {
	imagerepository repository.ImageRepository
}
func (s *ImageHandlerServiceImpl) Upload(ctx context.Context, fileHeader *multipart.FileHeader) (string, error) {

	u1, err := uuid.NewUUID()
	if err != nil {
		return "Failure", err
	}

	// Process file content type here
	fileType := fileHeader.Header.Get("Content-Type")

	extension, err := processFileTypeToExtension(fileType)
	if err != nil {
		return "", fmt.Errorf("Invalid Content-Type")
	}

	fileBodyReader, err := fileHeader.Open()
	if err != nil {
		return "", err
	}

	return s.imagerepository.Upload(ctx, fileBodyReader, u1.String() + extension)
}

func (s *ImageHandlerServiceImpl) Delete(ctx context.Context, fileName string) error {return nil}

func (s *ImageHandlerServiceImpl) Download(ctx context.Context, fileName string) (io.ReadCloser, error) {
	return nil, nil
}


func processFileTypeToExtension(mimetype string) (string, error) {
	switch mimetype {
	case "image/jpeg":
		return ".jpg", nil
	case "image/png":
		return ".png", nil
	default:
		return "", fmt.Errorf("Invalid Type. This can not be uploaded to the repository.")
	}
}