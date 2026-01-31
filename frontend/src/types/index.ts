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
}

export interface UserSearchParams {
  name: string
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
  updated_at: number
  reported_at: number
  project_id: string
  task_id: string
  project_info?: Project
  task_info?: Task
  project_name?: string
  task_title?: string
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
