package database

import (
	"context"
	"fmt"

	// Internal
	"github.com/BogdanStaziyev/jungle-test/internal/entity"

	// External
	"github.com/BogdanStaziyev/jungle-test/pkg/postgres"
)

type authRepo struct {
	*postgres.Postgres
}

func NewAuthRepo(ar *postgres.Postgres) *authRepo {
	return &authRepo{ar}
}

// FindByName is a repository method for finding a user by their name in the database.
func (a *authRepo) FindByName(name string) (user entity.User, err error) {
	sql := `SELECT id, password_hash, username FROM users WHERE username=$1 AND deleted_date IS NULL`
	err = a.Pool.QueryRow(context.Background(), sql, name).Scan(&user.ID, &user.Password, &user.Name)
	if err != nil {
		return entity.User{}, fmt.Errorf("user repository FindByName error: %w", err)
	}
	return user, nil
}

// The Save method of authRepo is responsible for saving a new user to the database.
func (a *authRepo) Save(user entity.User) (id int64, err error) {
	sql := `INSERT INTO users (password_hash, username, created_date, updated_date) VALUES ($1, $2, now(), now()) RETURNING id`
	err = a.Pool.QueryRow(context.Background(), sql, user.Password, user.Name).Scan(&id)
	if err != nil {
		return id, fmt.Errorf("user repo Save error: %w", err)
	}
	return id, nil
}
