import { defineConfig } from 'vite';
import react from '@vitejs/plugin-react';

// https://vite.dev/config/
export default defineConfig({
  plugins: [react()],
  build: {
    outDir: './build',
  },
  server: {
    port: 3002,
    proxy: {
      '/api': {
        target: 'http://localhost:3001/',
        changeOrigin: true,
        followRedirects: true,
        secure: false,
      },
    },
  },
});
