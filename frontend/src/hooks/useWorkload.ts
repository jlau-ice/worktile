import { useState, useCallback, useEffect } from 'react'
import { message } from 'antd'
import { getWorkload } from '@/api'
import type { WorkloadEntry, User } from '@/types'

interface UseWorkloadReturn {
  workload: WorkloadEntry[]
  total: number
  loading: boolean
  page: number
  pageSize: number
  fetchWorkload: (page: number, pageSize: number) => Promise<void>
  setPage: (page: number) => void
  setPageSize: (size: number) => void
}

export function useWorkload(selectedUser: User | null): UseWorkloadReturn {
  const [workload, setWorkload] = useState<WorkloadEntry[]>([])
  const [total, setTotal] = useState(0)
  const [loading, setLoading] = useState(false)
  const [page, setPage] = useState(1)
  const [pageSize, setPageSize] = useState(10)

  const fetchWorkload = useCallback(
    async (currentPage: number, currentPageSize: number) => {
      if (!selectedUser) return

      setLoading(true)
      try {
        const response = await getWorkload({
          uid: selectedUser.uid,
          pageNumber: currentPage,
          pageSize: currentPageSize,
        })
        setWorkload(response.data?.data || [])
        setTotal(response.data?.total || 0)
      } catch (err) {
        const errorMessage = err instanceof Error ? err.message : '获取工时失败'
        message.error(errorMessage)
      } finally {
        setLoading(false)
      }
    },
    [selectedUser]
  )

  // 当用户变化时重置并获取数据
  useEffect(() => {
    if (selectedUser) {
      setPage(1)
      setPageSize(10)
      fetchWorkload(1, 10)
    } else {
      setWorkload([])
      setTotal(0)
    }
  }, [selectedUser, fetchWorkload])

  return {
    workload,
    total,
    loading,
    page,
    pageSize,
    fetchWorkload,
    setPage,
    setPageSize,
  }
}
