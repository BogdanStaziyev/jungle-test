package entity

import (
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
