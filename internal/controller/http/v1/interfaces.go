package v1

import (
	"mime/multipart"

	// Echo
	"github.com/labstack/echo/v4"

	// Internal
	"github.com/BogdanStaziyev/jungle-test/internal/entity"
)

// Services structure that includes all services
type Services struct {
	AuthService
	ImageService
}

// Middleware structure that includes all middlewares
type Middleware struct {
	AuthMiddleware
}

type AuthService interface {
	Login(user entity.User) (string, error)
	Register(user entity.User) (int64, error)
}

type ImageService interface {
	UploadImage(image *multipart.FileHeader, entityImage entity.Image) (string, error)
	DownloadImages(id int64) ([]entity.Image, error)
}

type AuthMiddleware interface {
	ValidateJWT() echo.MiddlewareFunc
}
