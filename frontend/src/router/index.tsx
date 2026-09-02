import { Routes, Route } from 'react-router-dom'
import { MainLayout } from '@/layouts'
import { HomePage, PlaceholderPage, ContentSearchPage, UserManagementPage } from '@/pages'

export function AppRoutes() {
  return (
    <Routes>
      <Route path="/" element={<MainLayout />}>
        <Route index element={<HomePage />} />
        <Route path="workload" element={<HomePage />} />
        <Route path="content-search" element={<ContentSearchPage />} />
        <Route path="users" element={<UserManagementPage />} />
        <Route path="settings" element={<PlaceholderPage title="系统设置" />} />
      </Route>
    </Routes>
  )
}
