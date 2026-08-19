package interfaces

import (
	"context"
	"worktile/worktile-query-server/internal/types"
)

type WorkloadService interface {
	SearchWorkload(ctx context.Context, dto types.WorkloadDTO) (types.PaginatedWorkload, error)
	SearchWorkloadByContent(ctx context.Context, dto types.WorkloadSearchDTO) ([]types.MissionAddonWorkloadEntries, error)
}

type WorkloadRepository interface {
	WorkloadByUid(ctx context.Context, dto types.WorkloadDTO) ([]types.MissionAddonWorkloadEntries, int64, error)
	WorkloadByContent(ctx context.Context, dto types.WorkloadSearchDTO) ([]types.MissionAddonWorkloadEntries, error)
}
