package jwt

import (
	"github.com/BogdanStaziyev/jungle-test/internal/entity"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"time"
)

const expireTime = 12

type Token interface {
	CreateToken(name string, id int64) (string, error)
	GetUserFromContext(ctx echo.Context) entity.User
}

type tokens struct {
	expireTime int
	secret     string
}

func NewTokenConstructor(secret string) *tokens {
	return &tokens{
		secret: secret,
	}
}

func (t *tokens) CreateToken(name string, id int64) (string, error) {
	exp := time.Now().Add(time.Hour * time.Duration(expireTime)).Unix()
	uid := uuid.New().String()
	claimsAccess := Claim{
		Name: name,
		ID:   id,
		UID:  uid,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: exp,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claimsAccess)
	stringToken, err := token.SignedString([]byte(t.secret))

	return stringToken, err
}

func (t *tokens) GetUserFromContext(ctx echo.Context) (user entity.User) {
	token := ctx.Get("user").(*jwt.Token)
	claims := token.Claims.(*Claim)
	user.ID = claims.ID
	user.Name = claims.Name
	return user
}

type Claim struct {
	Name string `json:"name"`
	ID   int64  `json:"id"`
	UID  string `json:"uid"`
	jwt.StandardClaims
}
