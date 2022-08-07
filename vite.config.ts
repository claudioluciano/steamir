// eslint-disable-next-line
// @ts-nocheck
import { defineConfig } from 'vite'
import { resolve as resolvePath } from 'path'
import vue from '@vitejs/plugin-vue'
import WindiCSS from 'vite-plugin-windicss'

// https://vitejs.dev/config/
export default defineConfig({
  server: {
    watch: {
      usePolling: true
    }
  },
  resolve: {
    alias: {
      '@': resolvePath(__dirname, './src')
    }
  },
  plugins: [vue(), WindiCSS()]
})
