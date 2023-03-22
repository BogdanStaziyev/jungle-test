package database

import (
	"context"
	"fmt"
	"github.com/BogdanStaziyev/jungle-test/internal/entity"
	"github.com/BogdanStaziyev/jungle-test/pkg/postgres"
	"time"
)

type authRepo struct {
	*postgres.Postgres
}

func NewAuthRepo(ar *postgres.Postgres) *authRepo {
	return &authRepo{ar}
}

func (a *authRepo) FindByName(name string) (user entity.User, err error) {
	sql := `SELECT id, password_hash, username FROM users WHERE name=$1 AND deleted_date IS NULL`
	err = a.Pool.QueryRow(context.Background(), sql, name).Scan(&user.ID, &user.Password, &user.Name)
	if err != nil {
		return entity.User{}, fmt.Errorf("user repository FindByName error: %w", err)
	}
	return user, nil
}
func (a *authRepo) Save(user entity.User) (id int64, err error) {
	sql := `INSERT INTO users (password, name, created_date, updated_date) VALUES ($1, $2, $3, $4) RETURNING id`
	err = a.Pool.QueryRow(context.Background(), sql, user.Password, user.Name, time.Now(), time.Now()).Scan(&id)
	if err != nil {
		return id, fmt.Errorf("user repo Save error: %w", err)
	}
	return id, nil
}
