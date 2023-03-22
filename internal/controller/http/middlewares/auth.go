package middlewares

import (
	"net/http"

	// Echo
	"github.com/labstack/echo/v4"
	MW "github.com/labstack/echo/v4/middleware"

	// External
	"github.com/BogdanStaziyev/jungle-test/pkg/jwt"
)

type authMiddleware struct {
	secret string
}

func NewMiddleware(secret string) *authMiddleware {
	return &authMiddleware{
		secret: secret,
	}
}

// The ValidateJWT function creates an Echo middleware that uses JWT authentication
// Returns an error message if the token is not valid.
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
