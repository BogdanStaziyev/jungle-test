package v1

import (
	"fmt"
	"net/http"
	"strings"

	// Echo
	"github.com/labstack/echo/v4"

	// Internal
	"github.com/BogdanStaziyev/jungle-test/internal/controller/http/requests"
	"github.com/BogdanStaziyev/jungle-test/internal/controller/http/responses"

	// External
	"github.com/BogdanStaziyev/jungle-test/pkg/logger"
)

type registerHandler struct {
	as AuthService
	l  logger.Interface
}

func newRegisterHandler(router *echo.Group, a AuthService, l logger.Interface) {
	r := &registerHandler{
		as: a,
		l:  l,
	}
	usersRouter := router.Group("/users")
	usersRouter.POST("/register", r.Register)
	usersRouter.POST("/login", r.Login)
}

// The Register function is an HTTP handler that takes a request with user registration data, binds and validates it.
// returning an error message if any issues occur, or a success message if the user was registered successfully.
func (r registerHandler) Register(ctx echo.Context) error {
	var registerUser requests.RequestUser
	if err := ctx.Bind(&registerUser); err != nil {
		r.l.Error("auth controller register bind", "err: ", err)
		return response.ErrorResponse(ctx, http.StatusBadRequest, "Could not decode user data")
	}
	if err := ctx.Validate(&registerUser); err != nil {
		r.l.Error("auth controller register validate", "err: ", err)
		return response.ErrorResponse(ctx, http.StatusUnprocessableEntity, "Could not validate user data")
	}
	id, err := r.as.Register(registerUser.RegisterToUser())
	if err != nil {
		r.l.Error("auth controller register register response", "err: ", err)
		return response.ErrorResponse(ctx, http.StatusConflict, "Could not save new user, user already exists")
	}
	return response.MessageResponse(ctx, http.StatusCreated, fmt.Sprintf("User successfully created: %d", id))
}

// The Login function is an HTTP handler that takes a request with user authentication data, binds and validates it.
// Returning an error message if any issues occur or the user does not exist, or a JSON response with the access token.
func (r registerHandler) Login(ctx echo.Context) error {
	var authUser requests.RequestUser
	if err := ctx.Bind(&authUser); err != nil {
		r.l.Error("auth controller login bind", "err: ", err)
		return response.ErrorResponse(ctx, http.StatusBadRequest, "Could not decode user data")
	}
	if err := ctx.Validate(&authUser); err != nil {
		r.l.Error("auth controller login validate", "err: ", err)
		return response.ErrorResponse(ctx, http.StatusUnprocessableEntity, "Could not validate user data")
	}
	accessToken, err := r.as.Login(authUser.RegisterToUser())
	if err != nil {
		if strings.HasSuffix(err.Error(), "no rows in result set") {
			r.l.Error("auth controller login find user service response", "err: ", err)
			return response.ErrorResponse(ctx, http.StatusNotFound, "Could not login, user not exists")
		} else {
			r.l.Error("auth controller login find user service response", "err: ", err)
			return response.ErrorResponse(ctx, http.StatusInternalServerError, "Could not login user, server error")
		}
	}
	return response.MessageResponse(ctx, http.StatusOK, fmt.Sprintf("token: %s", accessToken))
}
