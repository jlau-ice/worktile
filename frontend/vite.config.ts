import { defineConfig } from 'vite'
import react from '@vitejs/plugin-react'
import path from 'path'

export default defineConfig({
  resolve: {
    alias: {
      '@': path.resolve(__dirname, './src'),
    },
  },
  server: {
    port: 28881,
    strictPort: true, // 端口被占用时报错，而不是自动尝试下一个端口
    host: true,
    allowedHosts: [
      'web.upiei.cn',
      'localhost',
      '192.168.172.241',
    ],
    proxy: {
      '/api': {
        target: 'http://127.0.0.1:28882',
        changeOrigin: true,
        xfwd: true,
      },
    },
  },
  plugins: [react()],
})
