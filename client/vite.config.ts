import { defineConfig } from 'vite';
import react from '@vitejs/plugin-react';

export default defineConfig({
  plugins: [react()],
  build: {
    outDir: './build',
  },
  server: {
    port: 3002,
    proxy: {
      '/api/orders': {
        target: 'http://localhost:8083/',
        changeOrigin: true,
        followRedirects: true,
        secure: false,
        rewrite: (path) => {
          return path.replace(/^\/api/, '');
        },
      },
    },
  },
});
