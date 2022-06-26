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

//go:generate mockgen -destination=../mocks/mock_user.go -package=mocks -source=user.go UserService
type UserService interface {
	FindUsers(ctx context.Context, filters map[string]string, page int) (users []User, err error)
	CreateUser(ctx context.Context, user User) (err error)
	UpdateUser(ctx context.Context, userID string, user User) (err error)
	DeleteUser(ctx context.Context, userID string) (err error)
}

//go:generate mockgen -destination=../mocks/mock_user.go -package=mocks -source=user.go UserRepository
type UserRepository interface {
	FindAll(ctx context.Context, filterParams map[string]string, page int) (users []User, err error)
	Insert(ctx context.Context, user User) (err error)
	Update(ctx context.Context, userID string, user User) (err error)
	Delete(ctx context.Context, userID string) (err error)
}
