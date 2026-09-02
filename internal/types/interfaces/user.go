package interfaces

import (
	"context"

	"worktile/worktile-query-server/internal/types"
)

type UserService interface {
	SearchUsers(ctx context.Context, name string) ([]types.User, error)
	GetUsersByPage(ctx context.Context, dto types.UserPageDTO) (*types.PaginatedUsers, error)
}

type UserRepository interface {
	FetchByName(ctx context.Context, name string) ([]types.User, error)
	FetchByPage(ctx context.Context, dto types.UserPageDTO) (*types.PaginatedUsers, error)
}
