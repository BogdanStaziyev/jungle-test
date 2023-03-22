package requests

import (
	"github.com/BogdanStaziyev/jungle-test/internal/entity"
)

type RequestUser struct {
	Password string `json:"password" validate:"required,gte=8" example:"01234567890"`
	Name     string `json:"name" validate:"required,gte=3"`
}

func (r RequestUser) RegisterToUser() entity.User {
	return entity.User{
		Name:     r.Name,
		Password: r.Password,
	}
}
