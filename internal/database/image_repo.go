package database

import (
	"context"
	"fmt"

	// Internal
	"github.com/BogdanStaziyev/jungle-test/internal/entity"

	// External
	"github.com/BogdanStaziyev/jungle-test/pkg/postgres"
)

type imageRepo struct {
	*postgres.Postgres
}

func NewImageRepo(ir *postgres.Postgres) *imageRepo {
	return &imageRepo{
		ir,
	}
}

// The SaveImage method of imageRepo takes an entity.Image parameter.
// Attempts to insert its properties into the images table in the database.
func (i *imageRepo) SaveImage(image entity.Image) error {
	sql := `INSERT INTO images (user_id, image_path, image_url, created_date) VALUES ($1, $2, $3, now())`
	result, err := i.Pool.Exec(context.Background(), sql, image.UserID, image.Path, image.URL)
	if err != nil {
		return fmt.Errorf("image repository SaveImage error: %w", err)
	}
	rows := result.RowsAffected()
	if rows == 0 {
		return fmt.Errorf("image repository SaveImage error")
	}
	return nil
}

// GetImages is a method of the imageRepo struct that retrieves all images associated with a given user ID.
func (i *imageRepo) GetImages(id int64) (images []entity.Image, err error) {
	sql := `SELECT id, user_id, image_path, image_url, created_date FROM images WHERE user_id=$1`
	rows, err := i.Pool.Query(context.Background(), sql, id)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var image entity.Image
		err = rows.Scan(&image.ID, &image.UserID, &image.Path, &image.URL, &image.CreatedDate)
		if err != nil {
			return nil, err
		}
		images = append(images, image)
	}
	if rows.Err() != nil {
		return nil, err
	}
	return images, nil
}
