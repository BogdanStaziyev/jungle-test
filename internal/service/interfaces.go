package service

import (
	"mime/multipart"

	// Internal
	"github.com/BogdanStaziyev/jungle-test/internal/entity"
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
	SaveImage(image entity.Image) error
	GetImages(id int64) ([]entity.Image, error)
}

type FileStorage interface {
	Save(image *multipart.FileHeader, domainImage *entity.Image) error
}
