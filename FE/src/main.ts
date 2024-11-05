import '@/style.css'
import "preline/preline";

import { createApp } from "vue";
import { createPinia } from "pinia";
import * as Sentry from "@sentry/vue";
import App from "./App.vue";
import router from "./router";
import piniaPluginPersistedstate from 'pinia-plugin-persistedstate';
import VueApexCharts from "vue3-apexcharts";
import UserAuthStorePlugin from './plugins/authStore';
import {MotionPlugin} from '@vueuse/motion'
const pinia = createPinia();
const useSentry = import.meta.env.VITE_SENTRY_URL ?? ""
pinia.use(piniaPluginPersistedstate);

const options = {
  url: import.meta.env.VITE_KEYCLOAK_URL,
  clientId: import.meta.env.VITE_KEYCLOAK_CLIENT_ID,
  realm: import.meta.env.VITE_KEYCLOAK_REALM,
};

const renderApp = () => {
  const app = createApp(App);
  if (useSentry != ""){
    Sentry.init({
      app,
      dsn: import.meta.env.VITE_SENTRY_URL,
      integrations: [
        Sentry.browserTracingIntegration(),
        Sentry.replayIntegration(),
      ],
      // Performance Monitoring
      tracesSampleRate: 1.0, //  Capture 100% of the transactions
      // Set 'tracePropagationTargets' to control for which URLs distributed tracing should be enabled
      tracePropagationTargets: ["localhost", /^https:\/\/yourserver\.io\/api/],
      // Session Replay
      replaysSessionSampleRate: 0.1, // This sets the sample rate at 10%. You may want to change it to 100% while in development and then sample at a lower rate in production.
      replaysOnErrorSampleRate: 1.0, // If you're not already sampling the entire session, change the sample rate to 100% when sampling sessions where errors occur.
    });
  }
  app.use(UserAuthStorePlugin, {pinia});
  app.use(pinia);
  app.use(MotionPlugin)
  app.use(VueApexCharts);
  app.use(router);
  app.mount("#app");
};
renderApp();
