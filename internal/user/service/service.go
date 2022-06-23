package service

import (
	"context"
	"time"

	"github.com/Slamadalius/faceit/internal/entity"
)

type service struct {
	repository     entity.UserRepository
	contextTimeout time.Duration
}

func NewUserService(userRepository entity.UserRepository, timeout time.Duration) entity.UserService {
	return &service{
		repository:     userRepository,
		contextTimeout: timeout,
	}
}

func (s *service) CreateUser(ctx context.Context, user entity.User) (err error) {
	ctx, cancel := context.WithTimeout(ctx, s.contextTimeout)
	defer cancel()

	return s.repository.Insert(ctx, user)
}

func (s *service) UpdateUser(ctx context.Context, userID string, user entity.User) (err error) {
	ctx, cancel := context.WithTimeout(ctx, s.contextTimeout)
	defer cancel()

	return s.repository.Update(ctx, userID, user)
}

func (s *service) DeleteUser(ctx context.Context, userID string) (err error) {
	ctx, cancel := context.WithTimeout(ctx, s.contextTimeout)
	defer cancel()

	return s.repository.Delete(ctx, userID)
}
