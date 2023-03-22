package v1

import (
	"fmt"
	"github.com/BogdanStaziyev/jungle-test/internal/controller/http/responses"
	"github.com/BogdanStaziyev/jungle-test/pkg/jwt"
	"github.com/BogdanStaziyev/jungle-test/pkg/logger"
	"net/http"

	// echo
	"github.com/labstack/echo/v4"
)

type imageHandler struct {
	is ImageService
	l  logger.Interface
	t  jwt.Token
}

func newImageHandler(router *echo.Group, imageService ImageService, l logger.Interface, token jwt.Token) {
	r := &imageHandler{
		is: imageService,
		l:  l,
		t:  token,
	}

	imageRouter := router.Group("/image")
	{
		imageRouter.POST("/upload", r.Upload)
		imageRouter.POST("/download", r.Download)
	}
}

func (i *imageHandler) Upload(ctx echo.Context) error {
	return response.MessageResponse(ctx, http.StatusOK, fmt.Sprintf("Image successful upload, id"))
}

func (i *imageHandler) Download(ctx echo.Context) error {
	user := i.t.GetUserFromContext(ctx)

	images, err := i.is.DownloadImages(user.ID)
	if err != nil {
		return response.ErrorResponse(ctx, http.StatusNotFound, "")
	}
	return response.Response(ctx, http.StatusOK, images)
}
