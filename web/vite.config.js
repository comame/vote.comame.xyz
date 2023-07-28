import { defineConfig } from 'vite'
import * as path from 'node:path'

export default defineConfig({
    build: {
        outDir: './static',
        rollupOptions: {
            input: {
                topic: path.resolve(__dirname, 'front/topic/index.html'),
                vote: path.resolve(__dirname, 'front/vote/index.html'),
            },
        },
    },
})
