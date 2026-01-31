import { Routes, Route } from 'react-router-dom'
import { MainLayout } from '@/layouts'
import { HomePage, PlaceholderPage } from '@/pages'

export function AppRoutes() {
  return (
    <Routes>
      <Route path="/" element={<MainLayout />}>
        <Route index element={<HomePage />} />
        <Route path="workload" element={<HomePage />} />
        <Route path="users" element={<PlaceholderPage title="人员管理" />} />
        <Route path="settings" element={<PlaceholderPage title="系统设置" />} />
      </Route>
    </Routes>
  )
}
