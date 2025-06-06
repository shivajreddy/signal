import { defineConfig } from 'vite'
import react from '@vitejs/plugin-react'

export default defineConfig({
    plugins: [react()],
    server: {
        port: 80,
        host: true,
        watch: {
            usePolling: true,
            interval: 1000
        },
        proxy: {
            '/api': {
                target: 'http://server:8080',
                changeOrigin: true,
                rewrite: (path) => path.replace(/^\/api/, '')
            },
            '/ws': {
                target: 'ws://server:8000',
                ws: true,
                changeOrigin: true
            }
        }
    },
    css: {
        postcss: './postcss.config.js'
    }
}) 
