import { useState } from 'react'
import { Box, Typography, TextField, Button, Alert, CircularProgress } from '@mui/material'
import { Search as SearchIcon } from '@mui/icons-material'
import { Table, Tooltip } from 'antd'
import type { ColumnsType } from 'antd/es/table'
import { searchWorkloadByContent } from '@/api'
import { formatTimestamp } from '@/utils'
import type { WorkloadEntry } from '@/types'

export function ContentSearchPage() {
  const [searchText, setSearchText] = useState('')
  const [workload, setWorkload] = useState<WorkloadEntry[]>([])
  const [loading, setLoading] = useState(false)
  const [error, setError] = useState<string | null>(null)

  const handleSearch = async () => {
    if (!searchText.trim()) {
      setError('请输入搜索内容')
      return
    }

    setLoading(true)
    setError(null)
    setWorkload([])

    try {
      const response = await searchWorkloadByContent({ description: searchText })
      setWorkload(response.data || [])
    } catch (err) {
      const errorMessage = err instanceof Error ? err.message : '查询失败'
      setError(errorMessage)
    } finally {
      setLoading(false)
    }
  }

  const handleKeyDown = (e: React.KeyboardEvent) => {
    if (e.key === 'Enter') {
      handleSearch()
    }
  }

  const columns: ColumnsType<WorkloadEntry> = [
    {
      title: '序号',
      key: 'index',
      align: 'center',
      width: 70,
      fixed: 'left',
      render: (_, __, index) => index + 1,
    },
    {
      title: '人员名称',
      dataIndex: 'user_name',
      key: 'user_name',
      width: 120,
      align: 'center',
      ellipsis: true,
      render: (text) => text || '-',
    },
    {
      title: '工时内容',
      dataIndex: 'description',
      key: 'description',
      ellipsis: true,
      render: (text) => (
        <Tooltip title={text} placement="topLeft">
          <span>{text || '无描述'}</span>
        </Tooltip>
      ),
    },
    {
      title: '时长(小时)',
      dataIndex: 'duration',
      key: 'duration',
      width: 100,
      align: 'center',
    },
    {
      title: '工作日期',
      dataIndex: 'reported_at',
      key: 'reported_at',
      width: 160,
      align: 'center',
      render: formatTimestamp,
    },
    {
      title: '创建时间',
      dataIndex: 'created_at',
      key: 'created_at',
      width: 160,
      align: 'center',
      render: formatTimestamp,
    },
    {
      title: '项目名称',
      dataIndex: 'project_name',
      key: 'project_name',
      width: 200,
      align: 'center',
      ellipsis: true,
      render: (text) => {
        const name = text || '-'
        return (
          <Tooltip title={name} placement="topLeft">
            <span>{name}</span>
          </Tooltip>
        )
      },
    },
    {
      title: '任务名称',
      dataIndex: 'task_title',
      key: 'task_title',
      width: 150,
      align: 'center',
      ellipsis: true,
      render: (text) => {
        const title = text || '-'
        return (
          <Tooltip title={title} placement="topLeft">
            <span>{title}</span>
          </Tooltip>
        )
      },
    },
  ]

  return (
    <Box
      sx={{
        display: 'flex',
        flexDirection: 'column',
        height: 'calc(100vh - 64px - 48px)',
        overflow: 'hidden',
      }}
    >
      {/* 固定顶部区域 */}
      <Box sx={{ flexShrink: 0 }}>
        <Typography variant="h5" component="h1" gutterBottom>
          内容查询
        </Typography>

        <Box sx={{ display: 'flex', gap: 2, mb: 2 }}>
          <TextField
            label="输入工时内容进行搜索"
            variant="outlined"
            fullWidth
            value={searchText}
            onChange={(e) => setSearchText(e.target.value)}
            onKeyDown={handleKeyDown}
            size="small"
          />
          <Button
            variant="contained"
            onClick={handleSearch}
            disabled={loading}
            startIcon={<SearchIcon />}
            sx={{ minWidth: 100 }}
          >
            查询
          </Button>
        </Box>

        {loading && (
          <Box sx={{ display: 'flex', justifyContent: 'center', my: 2 }}>
            <CircularProgress size={24} />
          </Box>
        )}

        {error && (
          <Alert severity="error" sx={{ mb: 2 }}>
            {error}
          </Alert>
        )}

        {!loading && workload.length === 0 && searchText && !error && searchText.length > 0 && (
          <Alert severity="info" sx={{ mb: 2 }}>
            未找到匹配的工时记录
          </Alert>
        )}
      </Box>

      {/* 可滚动的表格区域 */}
      <Box sx={{ flexGrow: 1, overflow: 'hidden' }}>
        <Table
          columns={columns}
          dataSource={workload}
          rowKey="id"
          loading={loading}
          bordered
          size="middle"
          scroll={{ x: 1000, y: 'calc(100vh - 350px)' }}
          pagination={false}
        />
      </Box>
    </Box>
  )
}
