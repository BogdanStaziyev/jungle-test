package v1

import (
	"fmt"
	"github.com/BogdanStaziyev/jungle-test/internal/controller/http/requests"
	"github.com/BogdanStaziyev/jungle-test/internal/controller/http/responses"
	"github.com/BogdanStaziyev/jungle-test/pkg/logger"
	"github.com/labstack/echo/v4"
	"net/http"
	"strings"
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

func (r registerHandler) Register(ctx echo.Context) error {
	var registerUser requests.RequestUser
	if err := ctx.Bind(&registerUser); err != nil {
		return response.ErrorResponse(ctx, http.StatusBadRequest, "Could not decode user data")
	}
	if err := ctx.Validate(&registerUser); err != nil {
		return response.ErrorResponse(ctx, http.StatusUnprocessableEntity, "Could not validate user data")
	}
	id, err := r.as.Register(registerUser.RegisterToUser())
	if err != nil {
		return response.ErrorResponse(ctx, http.StatusInternalServerError, fmt.Sprintf("Could not save new user: %s", err))
	}
	return response.MessageResponse(ctx, http.StatusCreated, fmt.Sprintf("User successfully created: %d", id))
}

func (r registerHandler) Login(ctx echo.Context) error {
	var authUser requests.RequestUser
	if err := ctx.Bind(&authUser); err != nil {
		return response.ErrorResponse(ctx, http.StatusBadRequest, "Could not decode user data")
	}
	if err := ctx.Validate(&authUser); err != nil {
		return response.ErrorResponse(ctx, http.StatusUnprocessableEntity, "Could not validate user data")
	}
	accessToken, err := r.as.Login(authUser.RegisterToUser())
	if err != nil {
		if strings.HasSuffix(err.Error(), "upper: no more rows in this result set") {
			return response.ErrorResponse(ctx, http.StatusNotFound, fmt.Sprintf("Could not login, user not exists: %s", err))
		} else {
			return response.ErrorResponse(ctx, http.StatusInternalServerError, fmt.Sprintf("Could not login user: %s", err))
		}
	}
	return response.Response(ctx, http.StatusOK, accessToken)
}
