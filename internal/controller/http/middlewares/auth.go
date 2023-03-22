package middlewares

import (
	"github.com/BogdanStaziyev/jungle-test/pkg/jwt"
	"github.com/labstack/echo/v4"
	MW "github.com/labstack/echo/v4/middleware"
	"net/http"
)

type authMiddleware struct {
	secret string
}

func NewMiddleware(secret string) *authMiddleware {
	return &authMiddleware{
		secret: secret,
	}
}

func (a *authMiddleware) ValidateJWT() echo.MiddlewareFunc {
	config := MW.JWTConfig{
		ErrorHandler: func(err error) error {
			return &echo.HTTPError{
				Code:    http.StatusUnauthorized,
				Message: "Not authorized",
			}
		},
		SigningKey: []byte(a.secret),
		Claims:     &jwt.Claim{},
	}
	return MW.JWTWithConfig(config)
}
