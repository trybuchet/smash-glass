import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'
import path from "path";

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [vue()],
  resolve: {
    alias: {
      "@": path.resolve(__dirname, "src"),
      "@bindings": path.resolve(__dirname, "bindings/SmashGlass/services/"),
    },
  },
  css: {
    preprocessorOptions: {
        scss: {
            api: 'modern',
            additionalData: `
                @use "sass:color";
                @use "@/styles/scss/flare" as *;
            `,
        },
    },
  },
  esbuild: {
    tsconfigRaw: `{
      "compilerOptions": {
        "noImplicitThis": false
      }
    }`,
  },
})