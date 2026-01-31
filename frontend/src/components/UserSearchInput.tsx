import { Box, TextField, Button } from '@mui/material'
import { Search as SearchIcon } from '@mui/icons-material'

interface UserSearchInputProps {
  value: string
  onChange: (value: string) => void
  onSearch: () => void
  loading?: boolean
}

export function UserSearchInput({ value, onChange, onSearch, loading }: UserSearchInputProps) {
  const handleKeyDown = (e: React.KeyboardEvent) => {
    if (e.key === 'Enter') {
      onSearch()
    }
  }

  return (
    <Box sx={{ display: 'flex', gap: 2 }}>
      <TextField
        label="输入姓名进行搜索"
        variant="outlined"
        fullWidth
        value={value}
        onChange={(e) => onChange(e.target.value)}
        onKeyDown={handleKeyDown}
        size="small"
      />
      <Button
        variant="contained"
        onClick={onSearch}
        disabled={loading}
        startIcon={<SearchIcon />}
        sx={{ minWidth: 100 }}
      >
        搜索
      </Button>
    </Box>
  )
}
