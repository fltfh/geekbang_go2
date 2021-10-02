package biz

import (
	"context"
	"template/internal/domain"
)

type UserRepo interface {
	GetById(ctx context.Context, id int64) (*domain.User, error)
}

type UserService struct {
	user	UserRepo
}

func NewUserService(repo UserRepo) * UserService {
	return &UserService{
		user: repo,
	}
}

func (s *UserService) GetById(ctx context.Context, id int64) (*domain.User, error){
	return s.user.GetById(ctx, id)
}
