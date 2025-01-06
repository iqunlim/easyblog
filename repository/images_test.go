package repository

import (
	"bytes"
	"context"
	"io"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestImageRepositoryLocalhostUpload_EXPECTED(t *testing.T) {
	//setup
	repo := NewImageRepositoryLocalhost(".")
	datastream := io.NopCloser(bytes.NewReader([]byte("test_img_data")))

	ret, err := repo.Upload(context.Background(), datastream, "file.txt")
	
	defer os.Remove(ret)
	assert.NoError(t, err)
	assert.Equal(t, ret, "/static/files/file.txt")

}

func TestImageRepositoryLocalhostUpload_ERROR(t *testing.T) {

	repo := NewImageRepositoryLocalhost("badpath")
	datastream := io.NopCloser(bytes.NewReader([]byte("test_img_data")))
	_, err := repo.Upload(context.Background(), datastream, "file.txt")
	defer os.Remove("file.txt")
	assert.Error(t, err)
}

func TestImageRepositoryLocalhostDelete_EXPECTED(t *testing.T) {

	repo := NewImageRepositoryLocalhost(".")
	file, err := os.Create("./file")
	file.Close();
	if err != nil {
		t.Fatal(err)
	}

	err = repo.Delete(context.Background(), "file")
	assert.NoError(t, err)
}
func TestImageRepositoryLocalhostDelete_ERROR(t *testing.T) {

	/* File not found error */
	repo := NewImageRepositoryLocalhost(".")

	err := repo.Delete(context.Background(), "file")
	assert.Error(t, err)
}

func TestImageRepostiroyLocalHostDownload_EXPECTED(t *testing.T) {

	expectedString := "test_img_data"
	repo := NewImageRepositoryLocalhost(".")
	file, err := os.Create("./file")
	file.Write([]byte(expectedString))
	file.Close();
	if err != nil {
		t.Fatal(err)
	}

	reader, err := repo.Download(context.Background(), "./file")

	actualBytes, err := io.ReadAll(reader)
	if err != nil {
		t.Fatal(err)
	}
	actualString := string(actualBytes)
	

	assert.Equal(t, expectedString, actualString)
	if err := os.Remove("./file"); err != nil {
		t.Error(err)
	}
}

func TestImageRepostiroyLocalHostDownload_ERROR(t *testing.T) {

	repo := NewImageRepositoryLocalhost(".")
	reader, err := repo.Download(context.Background(), "./file")
	assert.Equal(t,reader, nil)
	assert.Error(t, err)
}