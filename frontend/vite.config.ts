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
    port: 1710,
    host: true,
    proxy: {
      '/api': {
        target: 'http://127.0.0.1:1323',
        changeOrigin: true,
        xfwd: true,
      },
    },
  },
  plugins: [react()],
})
