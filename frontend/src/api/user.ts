import request from './request'
import type { ApiResponse, User, UserSearchParams } from '@/types'

/**
 * 根据姓名搜索用户
 */
export function searchUsers(params: UserSearchParams): Promise<ApiResponse<User[]>> {
  return request.get('/users', { params })
}
