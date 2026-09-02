// API 响应通用类型
export interface ApiResponse<T> {
  code: number
  data?: T
  msg?: string
}

// 用户相关类型
export interface User {
  id: string
  display_name: string
  uid: string
  title?: string
  gender?: string
  department?: string
  department_name?: string
  status: number
  is_deleted: number
  disabled_at: number
}

export interface UserSearchParams {
  name: string
  status?: number
  page_number?: number
  page_size?: number
}

// 项目相关类型
export interface Project {
  id: string
  name: string
}

// 任务相关类型
export interface Task {
  id: string
  title: string
}

// 工时相关类型
export interface WorkloadEntry {
  id: string
  description: string
  duration: number
  created_at: number
  created_by: string
  updated_at: number
  reported_at: number
  project_id: string
  task_id: string
  project_info?: Project
  task_info?: Task
  user_info?: User
  project_name?: string
  task_title?: string
  user_name?: string
}

export interface WorkloadParams {
  uid: string
  pageSize?: number
  pageNumber?: number
}

export interface PaginatedWorkload {
  data: WorkloadEntry[]
  total: number
  page_size: number
  page_number: number
}

export interface PaginatedUsers {
  data: User[]
  total: number
  page_size: number
  page_number: number
}
