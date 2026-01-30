package interfaces

import (
	"context"
	"worktile/worktile-query-server/internal/types"
)

type UserService interface {
	SearchUsers(ctx context.Context, name string) ([]types.User, error)
}

type UserRepository interface {
	FetchByName(ctx context.Context, name string) ([]types.User, error)
}
