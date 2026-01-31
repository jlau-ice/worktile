import { Box, Tabs, Tab } from '@mui/material'
import type { User } from '@/types'

interface UserTabsProps {
  users: User[]
  selectedIndex: number
  onSelect: (user: User, index: number) => void
}

export function UserTabs({ users, selectedIndex, onSelect }: UserTabsProps) {
  if (users.length === 0) return null

  return (
    <Box sx={{ borderBottom: 1, borderColor: 'divider', mt: 2 }}>
      <Tabs
        value={selectedIndex}
        onChange={(_, newValue) => onSelect(users[newValue], newValue)}
        aria-label="用户列表"
        variant="scrollable"
        scrollButtons="auto"
      >
        {users.map((user, index) => (
          <Tab key={user.id} label={user.display_name} id={`user-tab-${index}`} />
        ))}
      </Tabs>
    </Box>
  )
}
