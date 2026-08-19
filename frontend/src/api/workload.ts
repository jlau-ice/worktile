import request from './request'
import type { ApiResponse, WorkloadParams, PaginatedWorkload, WorkloadEntry } from '@/types'

/**
 * 根据用户 UID 获取工时记录
 */
export function getWorkload(params: WorkloadParams): Promise<ApiResponse<PaginatedWorkload>> {
  return request.get('/workload', { params })
}

/**
 * 根据工时内容模糊查询
 */
export function searchWorkloadByContent(params: { description: string }): Promise<ApiResponse<WorkloadEntry[]>> {
  return request.get('/workload/search', { params })
}
