package service

import (
	"context"

	"github.com/Slamadalius/faceit/internal/entity"
)

const (
	CountryFilter   = "country"
	FirstNameFilter = "first_name"
)

type service struct {
	repository entity.UserRepository
}

func NewUserService(userRepository entity.UserRepository) entity.UserService {
	return &service{
		repository: userRepository,
	}
}

func (s *service) FindUsers(ctx context.Context, filterParams map[string]string, page int) (users []entity.User, err error) {
	return s.repository.FindAll(ctx, filterParams, page)
}

func (s *service) CreateUser(ctx context.Context, user entity.User) (err error) {
	return s.repository.Insert(ctx, user)
}

func (s *service) UpdateUser(ctx context.Context, userID string, user entity.User) (err error) {
	return s.repository.Update(ctx, userID, user)
}

func (s *service) DeleteUser(ctx context.Context, userID string) (err error) {
	return s.repository.Delete(ctx, userID)
}
