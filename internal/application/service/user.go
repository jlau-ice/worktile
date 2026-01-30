package service

import (
	"context"
	"errors"
	"worktile/worktile-query-server/internal/types"
	"worktile/worktile-query-server/internal/types/interfaces"
)

type userService struct {
	repo interfaces.UserRepository
}

func NewUserService(
	repo interfaces.UserRepository) interfaces.UserService {
	return &userService{
		repo: repo,
	}
}

func (s *userService) SearchUsers(ctx context.Context, name string) ([]types.User, error) {
	if name == "" {
		return nil, errors.New("name参数不能为空")
	}
	return s.repo.FetchByName(ctx, name)
}
