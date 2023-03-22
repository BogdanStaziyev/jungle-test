package service

import (
	"github.com/BogdanStaziyev/jungle-test/internal/entity"
	"mime/multipart"
)

// Databases structure that includes all databases
type Databases struct {
	AuthRepo
	ImageRepo
	FileStorage
}

type AuthRepo interface {
	FindByName(name string) (entity.User, error)
	Save(user entity.User) (int64, error)
}

type ImageRepo interface {
	SaveImage(image entity.Image) (int64, error)
	GetImages(id int64) ([]entity.Image, error)
}

type FileStorage interface {
	Save(image *multipart.FileHeader, domainImage *entity.Image) error
}
