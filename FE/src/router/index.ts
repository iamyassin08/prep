import { createRouter, createWebHistory } from "vue-router";
import HomeView from "@/views/HomeView.vue";

import "preline/preline";
import { type IStaticMethods } from "preline/preline";
import AppLayout from "@/layouts/AppLayout.vue";
import DashboardLayout from "@/layouts/DashboardLayout.vue";

import Prep from "@/components/Prep.vue";
import User from "@/components/User.vue";


declare global {
  interface Window {
    HSStaticMethods: IStaticMethods;
  }
}
const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: "/",
      name: "home",
      meta: {layout: AppLayout},
      component: HomeView,
    },


    {
      path: '/prep/',
      name: 'prep',
      meta: { layout: DashboardLayout },
      component: Prep, 
    },
    
    {
      path: '/user/',
      name: 'user',
      meta: { layout: DashboardLayout },
      component: User, 
    },

  ],
});

router.afterEach((to, from, failure) => {
  if (!failure) {
    setTimeout(() => {
      window.HSStaticMethods.autoInit();
    }, 100);
  }
});
export default router;
