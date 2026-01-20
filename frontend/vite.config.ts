import tailwindcss from "@tailwindcss/vite";
import { sveltekit } from "@sveltejs/kit/vite";
import { defineConfig } from "vite";
import wails from "@wailsio/runtime/plugins/vite";

export default defineConfig({
  plugins: [tailwindcss(), sveltekit(), wails("./bindings")],
  server: {
    fs: {
      allow: ["./bindings"],
    },
  },
});
