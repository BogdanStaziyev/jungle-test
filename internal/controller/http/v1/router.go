package v1

import (
	"github.com/BogdanStaziyev/jungle-test/pkg/validators"
	// echo
	"github.com/labstack/echo/v4"
	MW "github.com/labstack/echo/v4/middleware"

	// external
	"github.com/BogdanStaziyev/jungle-test/pkg/logger"
)

// Router create routes using the Echo router.
func Router(e *echo.Echo, services Services, l logger.Interface) {
	//Options
	e.Use(MW.Logger())
	e.Use(MW.Recover())
	e.Validator = validators.NewValidator()

	//Routes
	//v1 := e.Group("api/v1")
	//{
	//}
}
