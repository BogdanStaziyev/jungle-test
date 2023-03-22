package entity

import (
	"github.com/BogdanStaziyev/jungle-test/internal/controller/http/responses"
	"time"
)

type User struct {
	ID          int64
	Name        string
	Password    string
	CreatedDate time.Time
	UpdatedDate time.Time
	DeletedDate *time.Time
}

func (u User) DomainToResponse() response.UserResponse {
	return response.UserResponse{
		ID:   u.ID,
		Name: u.Name,
	}
}
