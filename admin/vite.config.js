import {defineConfig} from 'vite'
import vue from '@vitejs/plugin-vue'
import path from 'path'

const baseUrl = {
    development: './',
    beta: './',
    release: './'
}

// https://vitejs.dev/config/
export default defineConfig({
    plugins: [
        vue(),
    ],
    resolve: {
        alias: {
            '~': path.resolve(__dirname, './'),
            '@': path.resolve(__dirname, 'src'),
        }
    },
    css: {
        sourceMap: true,
        preprocessorOptions: {
            less: {
                javascriptEnabled: true,
            }
        },
    },
    server: {
        host: 'localhost',
        port: 8080,
        open: true,
        hmr: {overlay: false}
    },
})