import request from './request'
import type { ApiResponse, WorkloadParams, PaginatedWorkload } from '@/types'

/**
 * 根据用户 UID 获取工时记录
 */
export function getWorkload(params: WorkloadParams): Promise<ApiResponse<PaginatedWorkload>> {
  return request.get('/workload', { params })
}
