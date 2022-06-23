package entity

import (
	"context"
	"time"
)

type User struct {
	ID        string    `json:"_id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Nickname  string    `json:"nickname"`
	Password  string    `json:"password"`
	Email     string    `json:"email"`
	Country   string    `json:"country"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserService interface {
	CreateUser(ctx context.Context, user User) (err error)
	UpdateUser(ctx context.Context, userID string, user User) (err error)
}

type UserRepository interface {
	Insert(ctx context.Context, user User) (err error)
	Update(ctx context.Context, userID string, user User) (err error)
}
