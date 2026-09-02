import request from './request'
import type { ApiResponse, User, UserSearchParams, PaginatedUsers } from '@/types'

/**
 * 根据姓名搜索用户
 */
export function searchUsers(params: UserSearchParams): Promise<ApiResponse<User[]>> {
  return request.get('/users', { params })
}

export interface UserPageParams {
  display_name?: string
  status?: number
  page_number: number
  page_size: number
}

export function getUsersByPage(
  params: UserPageParams
): Promise<ApiResponse<PaginatedUsers>> {
  return request.get('/users/page', { params })
}
