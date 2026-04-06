package users

import (
	"time"
)

type User struct {
	ID        int64
	Name      string
	Email     string
	Password  string
	CreatedAt time.Time
}

type UserRepository interface {
	Create(name, email, password string) (*User, error)
	GetUserByEmail(email string) (*User, error)
	GetUserByID(id int64) (*User, error)
}
