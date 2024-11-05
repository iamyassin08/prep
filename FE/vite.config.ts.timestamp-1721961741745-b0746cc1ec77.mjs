// vite.config.ts
import { sentryVitePlugin } from "file:///C:/Users/abdin/Documents/work/nuri/prep-app/node_modules/.pnpm/@sentry+vite-plugin@2.20.1/node_modules/@sentry/vite-plugin/dist/esm/index.mjs";
import { defineConfig } from "file:///C:/Users/abdin/Documents/work/nuri/prep-app/node_modules/.pnpm/vite@5.3.1_@types+node@20.14.6/node_modules/vite/dist/node/index.js";
import { fileURLToPath, URL } from "node:url";
import vue from "file:///C:/Users/abdin/Documents/work/nuri/prep-app/node_modules/.pnpm/@vitejs+plugin-vue@5.0.5_vite@5.3.1_@types+node@20.14.6__vue@3.4.29_typescript@5.4.5_/node_modules/@vitejs/plugin-vue/dist/index.mjs";
var __vite_injected_original_import_meta_url = "file:///C:/Users/abdin/Documents/work/nuri/prep-app/vite.config.ts";
var vite_config_default = defineConfig({
  plugins: [
    vue(),
    sentryVitePlugin({
      org: "nuri-softworks",
      project: "harmony-haul"
    })
  ],
  server: {
    cors: false,
    host: true,
    port: 5173,
    proxy: {
      "*": {
        target: "keycloak.yawmozer.io",
        ws: true,
        changeOrigin: true
      }
    }
  },
  resolve: {
    alias: [
      {
        find: /@\/components\/((?!.*[.](ts|js|tsx|jsx|vue)$).*$)/,
        replacement: fileURLToPath(
          new URL("./src/components/$1/index.vue", __vite_injected_original_import_meta_url)
        )
      },
      {
        find: "@",
        replacement: fileURLToPath(new URL("./src", __vite_injected_original_import_meta_url))
      }
    ]
  },
  build: {
    sourcemap: true
  }
});
export {
  vite_config_default as default
};
//# sourceMappingURL=data:application/json;base64,ewogICJ2ZXJzaW9uIjogMywKICAic291cmNlcyI6IFsidml0ZS5jb25maWcudHMiXSwKICAic291cmNlc0NvbnRlbnQiOiBbImNvbnN0IF9fdml0ZV9pbmplY3RlZF9vcmlnaW5hbF9kaXJuYW1lID0gXCJDOlxcXFxVc2Vyc1xcXFxhYmRpblxcXFxEb2N1bWVudHNcXFxcd29ya1xcXFxudXJpXFxcXGdhcmFnZS1zYWxlLWFwcFwiO2NvbnN0IF9fdml0ZV9pbmplY3RlZF9vcmlnaW5hbF9maWxlbmFtZSA9IFwiQzpcXFxcVXNlcnNcXFxcYWJkaW5cXFxcRG9jdW1lbnRzXFxcXHdvcmtcXFxcbnVyaVxcXFxnYXJhZ2Utc2FsZS1hcHBcXFxcdml0ZS5jb25maWcudHNcIjtjb25zdCBfX3ZpdGVfaW5qZWN0ZWRfb3JpZ2luYWxfaW1wb3J0X21ldGFfdXJsID0gXCJmaWxlOi8vL0M6L1VzZXJzL2FiZGluL0RvY3VtZW50cy93b3JrL251cmkvZ2FyYWdlLXNhbGUtYXBwL3ZpdGUuY29uZmlnLnRzXCI7aW1wb3J0IHsgc2VudHJ5Vml0ZVBsdWdpbiB9IGZyb20gXCJAc2VudHJ5L3ZpdGUtcGx1Z2luXCI7XHJcbmltcG9ydCB7IGRlZmluZUNvbmZpZyB9IGZyb20gJ3ZpdGUnXHJcbmltcG9ydCB7IGZpbGVVUkxUb1BhdGgsIFVSTCB9IGZyb20gJ25vZGU6dXJsJ1xyXG5pbXBvcnQgdnVlIGZyb20gJ0B2aXRlanMvcGx1Z2luLXZ1ZSdcclxuXHJcbi8vIGh0dHBzOi8vdml0ZWpzLmRldi9jb25maWcvXHJcbmV4cG9ydCBkZWZhdWx0IGRlZmluZUNvbmZpZyh7XHJcbiAgcGx1Z2luczogW1xyXG4gICAgdnVlKCksIFxyXG4gICAgc2VudHJ5Vml0ZVBsdWdpbih7XHJcbiAgICAgIG9yZzogXCJudXJpLXNvZnR3b3Jrc1wiLFxyXG4gICAgICBwcm9qZWN0OiBcImhhcm1vbnktaGF1bFwiXHJcbiAgfSldLFxyXG5cclxuICBzZXJ2ZXI6IHtcclxuICAgIGNvcnM6IGZhbHNlLFxyXG4gICAgaG9zdDogdHJ1ZSxcclxuICAgIHBvcnQ6IDUxNzMsXHJcbiAgICBwcm94eToge1xyXG4gICAgICAnKic6IHtcclxuICAgICAgICB0YXJnZXQ6ICdrZXljbG9hay55YXdtb3plci5pbycsXHJcbiAgICAgICAgd3M6IHRydWUsXHJcbiAgICAgICAgY2hhbmdlT3JpZ2luOiB0cnVlXHJcbiAgICAgIH1cclxuICAgIH1cclxuICB9LFxyXG5cclxuICByZXNvbHZlOiB7XHJcblxyXG4gICAgYWxpYXM6IFtcclxuICAgICAge1xyXG4gICAgICAgIGZpbmQ6IC9AXFwvY29tcG9uZW50c1xcLygoPyEuKlsuXSh0c3xqc3x0c3h8anN4fHZ1ZSkkKS4qJCkvLFxyXG4gICAgICAgIHJlcGxhY2VtZW50OiBmaWxlVVJMVG9QYXRoKFxyXG4gICAgICAgICAgICBuZXcgVVJMKFwiLi9zcmMvY29tcG9uZW50cy8kMS9pbmRleC52dWVcIiwgaW1wb3J0Lm1ldGEudXJsKVxyXG4gICAgICAgICksXHJcbiAgICB9LFxyXG4gICAgICB7XHJcbiAgICAgICAgZmluZDogJ0AnLFxyXG4gICAgICAgIHJlcGxhY2VtZW50OiBmaWxlVVJMVG9QYXRoKG5ldyBVUkwoJy4vc3JjJywgaW1wb3J0Lm1ldGEudXJsKSlcclxuICAgICAgfVxyXG4gIF1cclxuICB9LFxyXG5cclxuICBidWlsZDoge1xyXG4gICAgc291cmNlbWFwOiB0cnVlXHJcbiAgfVxyXG59KSJdLAogICJtYXBwaW5ncyI6ICI7QUFBc1YsU0FBUyx3QkFBd0I7QUFDdlgsU0FBUyxvQkFBb0I7QUFDN0IsU0FBUyxlQUFlLFdBQVc7QUFDbkMsT0FBTyxTQUFTO0FBSHlNLElBQU0sMkNBQTJDO0FBTTFRLElBQU8sc0JBQVEsYUFBYTtBQUFBLEVBQzFCLFNBQVM7QUFBQSxJQUNQLElBQUk7QUFBQSxJQUNKLGlCQUFpQjtBQUFBLE1BQ2YsS0FBSztBQUFBLE1BQ0wsU0FBUztBQUFBLElBQ2IsQ0FBQztBQUFBLEVBQUM7QUFBQSxFQUVGLFFBQVE7QUFBQSxJQUNOLE1BQU07QUFBQSxJQUNOLE1BQU07QUFBQSxJQUNOLE1BQU07QUFBQSxJQUNOLE9BQU87QUFBQSxNQUNMLEtBQUs7QUFBQSxRQUNILFFBQVE7QUFBQSxRQUNSLElBQUk7QUFBQSxRQUNKLGNBQWM7QUFBQSxNQUNoQjtBQUFBLElBQ0Y7QUFBQSxFQUNGO0FBQUEsRUFFQSxTQUFTO0FBQUEsSUFFUCxPQUFPO0FBQUEsTUFDTDtBQUFBLFFBQ0UsTUFBTTtBQUFBLFFBQ04sYUFBYTtBQUFBLFVBQ1QsSUFBSSxJQUFJLGlDQUFpQyx3Q0FBZTtBQUFBLFFBQzVEO0FBQUEsTUFDSjtBQUFBLE1BQ0U7QUFBQSxRQUNFLE1BQU07QUFBQSxRQUNOLGFBQWEsY0FBYyxJQUFJLElBQUksU0FBUyx3Q0FBZSxDQUFDO0FBQUEsTUFDOUQ7QUFBQSxJQUNKO0FBQUEsRUFDQTtBQUFBLEVBRUEsT0FBTztBQUFBLElBQ0wsV0FBVztBQUFBLEVBQ2I7QUFDRixDQUFDOyIsCiAgIm5hbWVzIjogW10KfQo=
