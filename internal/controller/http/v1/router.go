package v1

import (
	// echo
	"github.com/labstack/echo/v4"
	MW "github.com/labstack/echo/v4/middleware"

	// external
	"github.com/BogdanStaziyev/jungle-test/pkg/jwt"
	"github.com/BogdanStaziyev/jungle-test/pkg/logger"
	"github.com/BogdanStaziyev/jungle-test/pkg/validators"
)

// Router create routes using the Echo router.
func Router(e *echo.Echo, mid Middleware, services Services, token jwt.Token, l logger.Interface) {
	//Options
	e.Use(MW.Logger())
	e.Use(MW.Recover())
	e.Validator = validators.NewValidator()

	//Routes
	e.Static("/file_storage", "file_storage")
	v1 := e.Group("api/v1")
	{
		newRegisterHandler(v1, services.AuthService, l)
		v1.Use(mid.ValidateJWT())
		newImageHandler(v1, services.ImageService, l, token)
	}
}
