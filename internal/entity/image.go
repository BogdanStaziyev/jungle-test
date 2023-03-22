package entity

import (
	"fmt"
	response "github.com/BogdanStaziyev/jungle-test/internal/controller/http/responses"
	"github.com/google/uuid"
	"path/filepath"
	"time"
)

type Image struct {
	ID          int64     `json:"id"`
	UserID      int64     `json:"user_id"`
	Path        string    `json:"image_path"`
	URL         string    `json:"image_url"`
	CreatedDate time.Time `json:"created_date"`
}

func (i *Image) CreatePath(fileName, storage string) {
	// Create a new file name by combining the uuid and the default name. And use "name=" as a delimiter.
	newFileName := fmt.Sprintf("%s%s", uuid.New().String(), fileName)

	// Create file path
	path := filepath.Join(storage, newFileName)
	i.Path = filepath.FromSlash(path)
}

func (i *Image) ImageToResponse() response.Image {
	return response.Image{
		ID:          i.ID,
		UserID:      i.UserID,
		Path:        i.Path,
		URL:         i.URL,
		CreatedDate: i.CreatedDate,
	}
}
