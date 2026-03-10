import { defineConfig } from 'vite';
import react from '@vitejs/plugin-react';
import tailwindcss from '@tailwindcss/vite';
import svgr from 'vite-plugin-svgr';

// https://vite.dev/config/
export default defineConfig({
  plugins: [
    react(),
    tailwindcss(),
    svgr({
      // ✅ ensures all .svg imports can be used as React components
      include: '**/*.svg',
      svgrOptions: {
        exportType: 'named',
        ref: true,
        titleProp: true,
      },
    }),
  ],
  resolve: {
    alias: {
      '@mui/styled-engine': '@mui/styled-engine-sc',
    },
  },
  server: {
    host: '127.0.0.1', // force IPv4 binding
    port: 5173,
    strictPort: true,  // error if port is taken
    open: true,        // auto-open browser

    // --- ADD THIS SECTION ---
    proxy: {
      // Proxy API requests to your local Nginx Docker container
      '/api': {
        target: 'http://localhost:80',
        changeOrigin: true,
      },
      // Proxy Video/MinIO requests to Nginx
      '/videos': {
        target: 'http://localhost:80',
        changeOrigin: true,
      }
    }
    // ------------------------
  },
});
