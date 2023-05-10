package domain

import (
	"time"
)

type User struct {
	UserID    any       `json:"user_id" db:"user_id" form:"user_id"`
	Name      string    `json:"name" db:"name" form:"name"`
	CreatedAt time.Time `json:"created_at" db:"created_at" form:"created_at"`
}

type Users []User
