package service

import (
	"context"
	"worktile/worktile-query-server/internal/types"
	"worktile/worktile-query-server/internal/types/interfaces"
)

type workloadService struct {
	repo interfaces.WorkloadRepository
}

func NewWorkloadService(
	repo interfaces.WorkloadRepository) interfaces.WorkloadService {
	return &workloadService{
		repo: repo,
	}
}

func (s *workloadService) SearchWorkload(ctx context.Context, dto types.WorkloadDTO) (types.PaginatedWorkload, error) {
	// 可以在这里做一些业务检查，比如最大页码限制
	if dto.PageSize > 100 {
		dto.PageSize = 100
	}
	// 调用 Repo 获取原始数据和总数
	entries, total, err := s.repo.WorkloadByUid(ctx, dto)
	if err != nil {
		return types.PaginatedWorkload{}, err
	}
	// 组装返回结构
	return types.PaginatedWorkload{
		Data:       entries,
		Total:      total,
		PageSize:   dto.PageSize,
		PageNumber: dto.PageNumber,
	}, nil
}
