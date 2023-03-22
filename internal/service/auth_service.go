package service

import (
	"fmt"
	"github.com/BogdanStaziyev/jungle-test/internal/entity"
	"github.com/BogdanStaziyev/jungle-test/pkg/jwt"
	"github.com/BogdanStaziyev/jungle-test/pkg/passwords"
)

type authService struct {
	repo     AuthRepo
	token    jwt.Token
	password passwords.Generator
}

func NewAuthService(t jwt.Token, p passwords.Generator, r AuthRepo) *authService {
	return &authService{
		token:    t,
		password: p,
		repo:     r,
	}
}

func (a *authService) Register(user entity.User) (id int64, err error) {
	if _, err = a.repo.FindByName(user.Name); err == nil {
		return id, fmt.Errorf("auth service register, error user already exists: %w", err)
	}
	user.Password, err = a.password.GeneratePasswordHash(user.Password)
	if err != nil {
		return id, fmt.Errorf("auth service register, could not generate password hash error: %w", err)
	}
	id, err = a.repo.Save(user)
	if err != nil {
		return id, fmt.Errorf("auth service register error: %w", err)
	}
	return id, nil
}

func (a *authService) Login(user entity.User) (accessToken string, err error) {
	userFromDB, err := a.repo.FindByName(user.Name)
	if err != nil {
		return "", fmt.Errorf("auth service login, user not exists error: %w", err)
	}
	valid := a.password.CheckPasswordHash(user.Password, userFromDB.Password)
	if !valid {
		return "", fmt.Errorf("auth service login, invalid password: %w", err)
	}
	accessToken, err = a.token.CreateToken(userFromDB.Name, userFromDB.ID)
	if err != nil {
		return "", fmt.Errorf("auth service login, cteate token error: %w", err)
	}
	return accessToken, nil
}
