import { useState, useEffect } from 'react'
import { Box, Typography, CircularProgress, Alert } from '@mui/material'
import { UserSearchInput, UserTabs, WorkloadTable } from '@/components'
import { useUserSearch } from '@/hooks'
import type { User } from '@/types'

export function HomePage() {
  const [searchName, setSearchName] = useState('')
  const [selectedUser, setSelectedUser] = useState<User | null>(null)
  const [tabIndex, setTabIndex] = useState(0)

  const { users, loading, error, search } = useUserSearch()

  const handleSearch = () => {
    setSelectedUser(null)
    search(searchName)
  }

  const handleUserSelect = (user: User, index: number) => {
    setSelectedUser(user)
    setTabIndex(index)
  }

  // 当用户列表更新时，自动选择第一个用户
  useEffect(() => {
    if (users.length > 0 && !selectedUser) {
      setSelectedUser(users[0])
      setTabIndex(0)
    }
  }, [users, selectedUser])

  return (
    <Box
      sx={{
        display: 'flex',
        flexDirection: 'column',
        height: 'calc(100vh - 64px - 48px)', // 减去 header 高度和 padding
        overflow: 'hidden',
      }}
    >
      {/* 固定顶部区域 */}
      <Box sx={{ flexShrink: 0 }}>
        <Typography variant="h5" component="h1" gutterBottom>
          工时查询
        </Typography>

        <UserSearchInput
          value={searchName}
          onChange={setSearchName}
          onSearch={handleSearch}
          loading={loading}
        />

        {loading && (
          <Box sx={{ display: 'flex', justifyContent: 'center', my: 2 }}>
            <CircularProgress size={24} />
          </Box>
        )}

        {error && (
          <Alert severity="error" sx={{ mt: 2 }}>
            {error}
          </Alert>
        )}

        <UserTabs users={users} selectedIndex={tabIndex} onSelect={handleUserSelect} />
      </Box>

      {/* 可滚动的表格区域 */}
      <Box sx={{ flexGrow: 1, overflow: 'hidden', mt: 2 }}>
        <WorkloadTable selectedUser={selectedUser} />
      </Box>
    </Box>
  )
}
