package users

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID        primitive.ObjectID `bson:"_id"`
	Name      string             `bson:"name"`
	Email     string             `bson:"email"`
	Password  string             `bson:"password"`
	CreatedAt time.Time          `bson:"created_at"`
}

type UserRepository interface {
	Create(name, email, password string) (*User, error)
	GetUserByEmail(email string) (*User, error)
	GetUserByID(id primitive.ObjectID) (*User, error)
}
