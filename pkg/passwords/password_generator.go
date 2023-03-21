package passwords

import (
	"golang.org/x/crypto/bcrypt"
	"strconv"
)

type Generator interface {
	// GeneratePasswordHash generates a password hash for storage in the database
	GeneratePasswordHash(password string) (string, error)
	// CheckPasswordHash checks if the received password matches the one stored in the database
	CheckPasswordHash(password, hash string) bool
}

type generatePasswordHash struct {
	cost string
}

func NewGeneratePasswordHash(cost string) *generatePasswordHash {
	return &generatePasswordHash{
		cost: cost,
	}
}

func (g generatePasswordHash) GeneratePasswordHash(password string) (string, error) {
	cost, err := strconv.Atoi(g.cost)
	if err != nil {
		cost = bcrypt.DefaultCost
	}
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), cost)
	return string(bytes), err
}

func (g generatePasswordHash) CheckPasswordHash(password, hash string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)) == nil
}
