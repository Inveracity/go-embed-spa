import { defineConfig } from 'vite'
// import { svelte } from '@sveltejs/vite-plugin-svelte'
import { sveltekit } from '@sveltejs/kit/vite';
import path from "path";

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [sveltekit()],
  resolve: {
    alias: {
      $lib: path.resolve("./src/lib"),
    },
  },
})
