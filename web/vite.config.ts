/** @type {import('vite').UserConfig} */
import { defineConfig } from 'vite'
import { resolve } from 'path'
import vue from '@vitejs/plugin-vue'
import UnoCSS from 'unocss/vite'
import VueDevTools from 'vite-plugin-vue-devtools'
import config from './global.config'
import AutoImport from 'unplugin-auto-import/vite'

// https://vitejs.dev/config/
export default defineConfig({
  server: {
    open: true,
    port: config.PORT,
    proxy: {
      '/api': {
        target: config.BASE_URL,
        changeOrigin: true,
        rewrite: (path) => path.replace(/^\/api/, ''),
      },
    },
  },
  plugins: [
    vue(),
    VueDevTools(),
    UnoCSS(),
    AutoImport({
      imports: [
        {
          vue: [
            'ref', // import { ref } from 'vue',
            'reactive', // import { ref } from 'vue',
            'onMounted',
            'Ref'
          ],
        },
      ],
    }),
  ],
  resolve: {
    alias: [{ find: '@', replacement: resolve(__dirname, './src') }],
  },
})
