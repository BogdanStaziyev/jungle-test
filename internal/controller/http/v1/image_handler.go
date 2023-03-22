package v1

import (
	"fmt"
	"net/http"
	"strings"

	// Echo
	"github.com/labstack/echo/v4"

	// Internal
	"github.com/BogdanStaziyev/jungle-test/internal/controller/http/responses"
	"github.com/BogdanStaziyev/jungle-test/internal/entity"

	// External
	"github.com/BogdanStaziyev/jungle-test/pkg/jwt"
	"github.com/BogdanStaziyev/jungle-test/pkg/logger"
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

	imageRouter := router.Group("/images")
	{
		imageRouter.POST("/upload-picture", r.Upload)
		imageRouter.GET("", r.DownloadAll)
	}
}

// The Upload function is an HTTP handler that takes an image file submitted as a form data and checks format.
// Returning an error message if any issues occur or a success message with the image URL if the upload is successful.
func (i *imageHandler) Upload(ctx echo.Context) error {
	var entityImage entity.Image
	img, err := ctx.FormFile("image")
	if err != nil {
		i.l.Error(err, "image handler upload form file", "err:", err)
		return response.ErrorResponse(ctx, http.StatusBadRequest, "The image was not uploaded. Please add an image to the field and try again.")
	}
	//Check file format
	contentType := img.Header.Get("Content-Type")
	if contentType != "image/png" && contentType != "image/jpeg" {
		i.l.Error(err, "image handler upload get check content", "err: ", err)
		return response.ErrorResponse(ctx, http.StatusBadRequest, "The format of the submitted file is not supported. The file should be in the format of: .png or .jpeg")
	}
	// Get user id and name from JWT
	user := i.t.GetUserFromContext(ctx)
	entityImage.UserID = user.ID
	entityImage.URL = ctx.Scheme() + "://" + ctx.Request().Host
	//Upload image to storage and write to DB
	url, err := i.is.UploadImage(img, entityImage)
	if err != nil {
		i.l.Error(err, "image handler upload image id", "err: ", err)
		return response.ErrorResponse(ctx, http.StatusInternalServerError, "Could not save image, try again later")
	}
	return response.MessageResponse(ctx, http.StatusOK, fmt.Sprintf("Image successful upload, url: %s", url))
}

// The DownloadAll function is an HTTP handler that retrieves all images uploaded by the current user from the database.
func (i *imageHandler) DownloadAll(ctx echo.Context) error {
	// Get user ID and Name from token
	user := i.t.GetUserFromContext(ctx)

	// Getting all images from db
	images, err := i.is.DownloadImages(user.ID)
	if err != nil {
		if strings.HasSuffix(err.Error(), "no rows in result set") {
			i.l.Error("image controller service response", "err: ", err)
			return response.ErrorResponse(ctx, http.StatusNotFound, "Could not find, images not exists")
		} else {
			i.l.Error("auth controller login find user service response", "err: ", err)
			return response.ErrorResponse(ctx, http.StatusInternalServerError, "Could not find images, server error")
		}
	}
	return response.Response(ctx, http.StatusOK, images)
}
