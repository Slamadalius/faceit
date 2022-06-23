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

	user.CreatedAt = time.Now()

	return s.repository.Insert(ctx, user)
}
