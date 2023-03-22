package database

import (
	"github.com/BogdanStaziyev/jungle-test/internal/entity"
	"io"
	"mime/multipart"
	"os"
)

type storage struct {
	path string
}

func NewStorage(path string) *storage {
	return &storage{
		path: path,
	}
}

func (s *storage) Save(image *multipart.FileHeader, entityImage *entity.Image) error {
	//Create current path to image
	entityImage.CreatePath(image.Filename, s.path)

	src, err := image.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	//Destination
	dst, err := os.Create(entityImage.Path)
	if err != nil {
		return err
	}
	if _, err = io.Copy(dst, src); err != nil {
		return err
	}
	return nil
}
