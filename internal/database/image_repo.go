package database

import (
	"context"
	"fmt"
	"github.com/BogdanStaziyev/jungle-test/internal/entity"
	"github.com/BogdanStaziyev/jungle-test/pkg/postgres"
	"time"
)

type imageRepo struct {
	*postgres.Postgres
}

func NewImageRepo(ir *postgres.Postgres) *imageRepo {
	return &imageRepo{
		ir,
	}
}

func (i *imageRepo) SaveImage(image entity.Image) (id int64, err error) {
	sql := `INSERT INTO images (user_id, path, url, created_date, updated_date) VALUES ($1, $2, $3, $4, $5) RETURNING id`
	err = i.Pool.QueryRow(context.Background(), sql, image.UserID, image.Path, image.URL, time.Now(), time.Now()).Scan(&id)
	if err != nil {
		return id, fmt.Errorf("image repo Save error: %w", err)
	}
	return id, nil
}

func (i *imageRepo) GetImages(id int64) ([]entity.Image, error) {
	return nil, nil
}
