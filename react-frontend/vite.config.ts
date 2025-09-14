import { defineConfig } from 'vite'
import react from '@vitejs/plugin-react'
import tailwindcss from '@tailwindcss/vite'

// https://vite.dev/config/
export default defineConfig({
  plugins: [react(),
    tailwindcss(),
  ],
  resolve: {
    alias: {
      '@mui/styled-engine': '@mui/styled-engine-sc',
    },
  },
  server: {
    host: '127.0.0.1',   // force IPv4 binding
    port: 5173,
    strictPort: true,    // error if port is taken
    open: true           // auto-open browser
  }
})
