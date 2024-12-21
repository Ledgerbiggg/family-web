import {defineConfig} from 'vite'
import vue from '@vitejs/plugin-vue'
import {resolve} from 'path'
import vueDevTools from 'vite-plugin-vue-devtools'

// https://vite.dev/config/
export default defineConfig({
    server: {
        proxy: {
            '/api': {
                target: 'http://localhost:8001/v1',
                changeOrigin: true,
                rewrite: (path: string) => {
                    console.log('Rewriting path:', path); // 打印路径
                    return path.replace(/^\/api/, '');
                },
                secure: false,
            },
        }

    },
    resolve: {
        alias: [
            {
                find: '@',
                replacement: resolve(__dirname, 'src')
            }
        ]
    },
    plugins: [vue(), vueDevTools(),],
})