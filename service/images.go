package service

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"mime/multipart"
	"path"
	"slices"

	"github.com/google/uuid"
	"github.com/iqunlim/easyblog/repository"
)


var (
	WhitelistedFileExtensions = []string{".jpg", ".png"}
	WhitelistedFileMIMETypes = []string{"image/jpeg", "image/png"}
	MaxFileSize int64 = 8 << 20
)


type ImageHandlerService interface {
	Upload(ctx context.Context, fileBody io.Reader, fileReader *multipart.FileHeader) (string, error)
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
func (s *ImageHandlerServiceImpl) Upload(ctx context.Context, fileBody io.Reader, fileHeader *multipart.FileHeader) (string, error) {

	if fileBody == nil {
		return "", fmt.Errorf("Error: No file content")
	}
	// TODO: Image validation?
	// TODO: Image Transformation?
	// Check valid Content-Type header
	if !slices.Contains(WhitelistedFileMIMETypes, fileHeader.Header.Get("Content-Type")) {
		return "", fmt.Errorf("Error: Not whitelisted content type")
	}

	// Check valid file extension
 	fileExtension := path.Ext(fileHeader.Filename)
	if !slices.Contains(WhitelistedFileExtensions, fileExtension) {
		return "", fmt.Errorf("Error: Not whitelisted file extension")
	}

	// Relying solely on fileHeader.size is a poor way to do validation.
	// Here we do it a little bit better, reading the whole file and getting its proper size
	fileBytes, err := io.ReadAll(fileBody)
	if err != nil {
		return "", fmt.Errorf("Error validating image file")
	}

	// Hard-coded 20MB for now
	//config.MaxImageSize
	if int64(len(fileBytes)) > MaxFileSize {
		return "", fmt.Errorf("File image is too large. Greater than %v Bytes", MaxFileSize)
	}

	//reset reader with this handy trick so we can send it down to be written to a file
	fileBody = io.NopCloser(bytes.NewBuffer(fileBytes))

	u1, err := uuid.NewUUID()
	if err != nil {
		return "", err
	}

	return s.imagerepository.Upload(ctx, fileBody, u1.String() + fileExtension)
}

func (s *ImageHandlerServiceImpl) Delete(ctx context.Context, fileName string) error {return nil}

func (s *ImageHandlerServiceImpl) Download(ctx context.Context, fileName string) (io.ReadCloser, error) {
	return nil, nil
}