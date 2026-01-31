import { useState, useCallback } from 'react'
import { message } from 'antd'
import { searchUsers } from '@/api'
import type { User } from '@/types'

interface UseUserSearchReturn {
  users: User[]
  loading: boolean
  error: string | null
  search: (name: string) => Promise<void>
  reset: () => void
}

export function useUserSearch(): UseUserSearchReturn {
  const [users, setUsers] = useState<User[]>([])
  const [loading, setLoading] = useState(false)
  const [error, setError] = useState<string | null>(null)

  const search = useCallback(async (name: string) => {
    if (!name.trim()) {
      message.warning('请输入搜索关键词')
      return
    }

    setLoading(true)
    setError(null)

    try {
      const response = await searchUsers({ name })
      setUsers(response.data || [])
    } catch (err) {
      const errorMessage = err instanceof Error ? err.message : '搜索失败'
      setError(errorMessage)
      message.error(errorMessage)
    } finally {
      setLoading(false)
    }
  }, [])

  const reset = useCallback(() => {
    setUsers([])
    setError(null)
  }, [])

  return { users, loading, error, search, reset }
}
