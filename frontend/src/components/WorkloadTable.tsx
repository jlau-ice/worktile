import { useCallback, useRef, useState, useEffect } from 'react'
import { Table } from 'antd'
import type { ColumnsType, TablePaginationConfig } from 'antd/es/table'
import { formatTimestamp } from '@/utils'
import { useWorkload } from '@/hooks'
import type { WorkloadEntry, User } from '@/types'

interface WorkloadTableProps {
  selectedUser: User | null
}

export function WorkloadTable({ selectedUser }: WorkloadTableProps) {
  const containerRef = useRef<HTMLDivElement>(null)
  const [tableHeight, setTableHeight] = useState(400)

  const { workload, total, loading, page, pageSize, fetchWorkload, setPage, setPageSize } =
    useWorkload(selectedUser)

  // 计算表格高度
  useEffect(() => {
    const updateHeight = () => {
      if (containerRef.current) {
        // 容器高度 - 分页高度(56px) - 表头高度(55px) - padding
        const height = containerRef.current.clientHeight - 120
        setTableHeight(Math.max(height, 200))
      }
    }

    updateHeight()
    window.addEventListener('resize', updateHeight)
    return () => window.removeEventListener('resize', updateHeight)
  }, [])

  const handleTableChange = useCallback(
    (pagination: TablePaginationConfig) => {
      const newPage = pagination.current || 1
      const newPageSize = pagination.pageSize || 10

      if (newPage !== page || newPageSize !== pageSize) {
        setPage(newPage)
        setPageSize(newPageSize)
        fetchWorkload(newPage, newPageSize)
      }
    },
    [page, pageSize, fetchWorkload, setPage, setPageSize]
  )

  const columns: ColumnsType<WorkloadEntry> = [
    {
      title: '序号',
      key: 'index',
      align: 'center',
      width: 70,
      fixed: 'left',
      render: (_, __, index) => (page - 1) * pageSize + index + 1,
    },
    {
      title: '工时内容',
      dataIndex: 'description',
      key: 'description',
      ellipsis: true,
      render: (text) => text || '无描述',
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
      dataIndex: 'project_info',
      key: 'project_name',
      width: 150,
      align: 'center',
      render: (projectInfo) => projectInfo?.name || '-',
    },
    {
      title: '任务名称',
      dataIndex: 'task_info',
      key: 'task_title',
      width: 150,
      align: 'center',
      render: (taskInfo) => taskInfo?.title || '-',
    },
  ]

  if (!selectedUser) {
    return null
  }

  return (
    <div ref={containerRef} style={{ height: '100%' }}>
      <Table
        columns={columns}
        dataSource={workload}
        rowKey="id"
        loading={loading}
        bordered
        size="middle"
        scroll={{ x: 900, y: tableHeight }}
        pagination={{
          current: page,
          pageSize: pageSize,
          total: total,
          showSizeChanger: true,
          showQuickJumper: true,
          showTotal: (total, range) => `第 ${range[0]}-${range[1]} 条，共 ${total} 条`,
          pageSizeOptions: [10, 20, 50, 100],
        }}
        onChange={handleTableChange}
      />
    </div>
  )
}
