package service

import (
	"context"
	"mime/multipart"
	"net/textproto"
	"testing"

	"github.com/iqunlim/easyblog/repository"
	"github.com/stretchr/testify/assert"
)


func CreateMockFileHeader(t *testing.T, filename string, content []byte) *multipart.FileHeader {
	v := make(textproto.MIMEHeader)
	v.Set("Content-Type", "image/jpeg")
	fileHeader := &multipart.FileHeader{
			Filename:       filename,
			Header:        v,
			Size:           int64(len(content)),
	}
	return fileHeader
}


func TestImageHandlerServiceImpl_Upload_EXPECTED(t *testing.T) {
	// This test is bad, rewrite it
	/*
	repo := repository.NewMockImageRepository(t)
	sv := NewImageService(repo)
	v := make(textproto.MIMEHeader)
	v.Set("Content-Type", "image/jpeg")
	fileHeader := &multipart.FileHeader{
			Filename:       "test.jpg",
			Header:        v,
			Size:           int64(len([]byte("Hello World"))),
	}
	repo.EXPECT().Upload(mock.Anything, mock.Anything, mock.AnythingOfType("string")).Return("/static/upload.txt", nil)
	_, err := sv.Upload(context.Background(), fileHeader)
	fmt.Println(err.Error())
	assert.NoError(t, err)
	*/
}
func TestImageHandlerServiceImpl_Upload_BADFILETYPE(t *testing.T) {
	repo := repository.NewMockImageRepository(t)
	sv := NewImageService(repo)
	v := make(textproto.MIMEHeader)
	v.Set("Content-Type", "text/plain")
	fileHeader := &multipart.FileHeader{
			Filename:       "test.jpg",
			Header:        	v,
			Size:           int64(len([]byte("Hello World"))),
	}
	_, err := sv.Upload(context.Background(), fileHeader)
	assert.Error(t, err)
}