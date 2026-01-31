import axios, { AxiosError, AxiosResponse } from 'axios'
import type { ApiResponse } from '@/types'

const request = axios.create({
  baseURL: '/api',
  timeout: 10000,
})

// 请求拦截器
request.interceptors.request.use(
  (config) => {
    // 可以在这里添加 token 等
    return config
  },
  (error) => {
    return Promise.reject(error)
  }
)

// 响应拦截器
request.interceptors.response.use(
  (response: AxiosResponse<ApiResponse<unknown>>) => {
    const res = response.data

    // 业务错误处理
    if (res.code !== 200) {
      return Promise.reject(new Error(res.msg || '请求失败'))
    }

    // 返回整个响应数据（包含 code, data, msg）
    return response.data as unknown as AxiosResponse
  },
  (error: AxiosError) => {
    // HTTP 错误处理
    const message = error.response?.status
      ? `请求失败: ${error.response.status}`
      : '网络错误，请检查网络连接'
    return Promise.reject(new Error(message))
  }
)

export default request
