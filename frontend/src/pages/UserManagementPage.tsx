import { useState, useEffect } from 'react'
import {
  Box,
  Typography,
  TextField,
  Button,
  Table,
  TableBody,
  TableCell,
  TableContainer,
  TableHead,
  TableRow,
  TablePagination,
  Paper,
  CircularProgress,
  Select,
  MenuItem,
  FormControl,
  InputLabel,
  Chip,
} from '@mui/material'
import { getUsersByPage } from '@/api/user'
import type { User } from '@/types'

export function UserManagementPage() {
  const [displayName, setDisplayName] = useState('')
  const [status, setStatus] = useState<number | ''>('')
  const [users, setUsers] = useState<User[]>([])
  const [total, setTotal] = useState(0)
  const [page, setPage] = useState(0)
  const [rowsPerPage, setRowsPerPage] = useState(10)
  const [loading, setLoading] = useState(false)

  const fetchUsers = async (pageNumber: number, pageSize: number) => {
    setLoading(true)
    try {
      const response = await getUsersByPage({
        display_name: displayName,
        status: status === '' ? undefined : status,
        page_number: pageNumber + 1,
        page_size: pageSize,
      })
      if (response.data) {
        setUsers(response.data.data)
        setTotal(response.data.total)
      }
    } catch (error) {
      console.error('Failed to fetch users:', error)
    } finally {
      setLoading(false)
    }
  }

  useEffect(() => {
    fetchUsers(page, rowsPerPage)
  }, [])

  const handleSearch = () => {
    setPage(0)
    fetchUsers(0, rowsPerPage)
  }

  const handleChangePage = (_: unknown, newPage: number) => {
    setPage(newPage)
    fetchUsers(newPage, rowsPerPage)
  }

  const handleChangeRowsPerPage = (event: React.ChangeEvent<HTMLInputElement>) => {
    const newRowsPerPage = parseInt(event.target.value, 10)
    setRowsPerPage(newRowsPerPage)
    setPage(0)
    fetchUsers(0, newRowsPerPage)
  }

  const getStatusDisplay = (statusValue: number) => {
    if (statusValue === 1) {
      return <Chip label="在职" color="success" size="small" />
    }
    if (statusValue === 2) {
      return <Chip label="离职" color="error" size="small" />
    }
    return <Chip label="未知" size="small" />
  }

  return (
    <Box>
      <Typography variant="h5" component="h1" gutterBottom>
        人员管理
      </Typography>

      <Box sx={{ mb: 3, display: 'flex', gap: 2, alignItems: 'center' }}>
        <TextField
          label="用户名"
          value={displayName}
          onChange={(e) => setDisplayName(e.target.value)}
          size="small"
          onKeyPress={(e) => {
            if (e.key === 'Enter') {
              handleSearch()
            }
          }}
        />
        <FormControl size="small" sx={{ minWidth: 120 }}>
          <InputLabel>状态</InputLabel>
          <Select
            value={status}
            label="状态"
            onChange={(e) => setStatus(e.target.value as number | '')}
          >
            <MenuItem value="">全部</MenuItem>
            <MenuItem value={1}>在职</MenuItem>
            <MenuItem value={2}>离职</MenuItem>
          </Select>
        </FormControl>
        <Button variant="contained" onClick={handleSearch} disabled={loading}>
          查询
        </Button>
      </Box>

      {loading ? (
        <Box sx={{ display: 'flex', justifyContent: 'center', my: 4 }}>
          <CircularProgress />
        </Box>
      ) : (
        <Paper>
          <TableContainer>
            <Table>
              <TableHead>
                <TableRow>
                  <TableCell>序号</TableCell>
                  <TableCell>用户名</TableCell>
                  <TableCell>职位</TableCell>
                  <TableCell>性别</TableCell>
                  <TableCell>部门</TableCell>
                  <TableCell>状态</TableCell>
                  <TableCell>是否删除</TableCell>
                  <TableCell>禁用时间</TableCell>
                </TableRow>
              </TableHead>
              <TableBody>
                {users.map((user, index) => (
                  <TableRow key={user.id}>
                    <TableCell>{page * rowsPerPage + index + 1}</TableCell>
                    <TableCell>{user.display_name}</TableCell>
                    <TableCell>{user.title || '-'}</TableCell>
                    <TableCell>{user.gender || '-'}</TableCell>
                    <TableCell>{user.department_name || user.department || '-'}</TableCell>
                    <TableCell>{getStatusDisplay(user.status)}</TableCell>
                    <TableCell>{user.is_deleted}</TableCell>
                    <TableCell>
                      {user.disabled_at
                        ? new Date(user.disabled_at * 1000).toLocaleString()
                        : '-'}
                    </TableCell>
                  </TableRow>
                ))}
              </TableBody>
            </Table>
          </TableContainer>
          <TablePagination
            component="div"
            count={total}
            page={page}
            onPageChange={handleChangePage}
            rowsPerPage={rowsPerPage}
            onRowsPerPageChange={handleChangeRowsPerPage}
            rowsPerPageOptions={[10, 20, 50]}
            labelRowsPerPage="每页显示"
            labelDisplayedRows={({ from, to, count }) => `${from}-${to} / 共 ${count}`}
          />
        </Paper>
      )}
    </Box>
  )
}
