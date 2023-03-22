package v1

import (
	"github.com/BogdanStaziyev/jungle-test/internal/entity"
	"github.com/labstack/echo/v4"
	"mime/multipart"
)

// Services structure that includes all services
type Services struct {
	AuthService
	ImageService
}

type Middleware struct {
	AuthMiddleware
}

type AuthService interface {
	Login(user entity.User) (string, error)
	Register(user entity.User) (int64, error)
}

type ImageService interface {
	UploadImage(image *multipart.FileHeader, domainImage *entity.Image) (int64, error)
	DownloadImages(id int64) ([]entity.Image, error)
}

type AuthMiddleware interface {
	ValidateJWT() echo.MiddlewareFunc
}
