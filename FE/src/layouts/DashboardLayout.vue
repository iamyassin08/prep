<script setup lang="ts">
import AppFooter from "@/components/layout/Footer.vue";
import AppHeader from "@/components/layout/TheHeader.vue";
import Sidebar from "@/components/user/Sidebar.vue";
import { useUserStore } from '@/stores/userStore';
const userStore = useUserStore();


</script>
<template>
    <AppHeader />
    <!-- ========== MAIN CONTENT ========== -->
    <main id="content"
        class="dark:bg-neutral-900 bg-slate-500 pt-[59px] relative before:absolute before:top-0 before:start-0 before:-z-[1] before:w-full before:h-96 before:bg-neutral-800 dark:before:bg-neutral-950">
        <!-- Container -->
        <div class="max-w-[90rem] mx-auto pt-8">
            <!-- Header -->
            <div class="px-2 sm:px-5">
                <div class="pb-2 sm:pb-0 flex justify-between items-center">
                    <h1 class="text-xl font-semibold text-white">
                        Dashboard
                    </h1>

                    <!-- Calendar Button -->
          <div class="flex space-x-2">
            <RouterLink
              v-if="userStore.isSeller "
              :to="'/user/categoryselection'"
              type="button"
              class="py-2 px-3 inline-flex justify-center items-center text-start bg-white/10 border border-transparent text-white text-sm font-medium rounded-lg shadow-sm align-middle hover:bg-white/20 focus:outline-none focus:ring-1 focus:ring-neutral-600"
            >
              Sell Something New
            </RouterLink>

            <RouterLink
              v-if="userStore.isAdmin"
              :to="'/admin/createcategory'"
              type="button"
              class="py-2 px-3 inline-flex justify-center items-center text-start bg-white/10 border border-transparent text-white text-sm font-medium rounded-lg shadow-sm align-middle hover:bg-white/20 focus:outline-none focus:ring-1 focus:ring-neutral-600"
            >
              Create a Category
            </RouterLink>
               <!-- End Calendar Button -->
              </div>
            </div>
            <!-- End Header -->

            <!-- Content -->
            <div class="p-2 sm:p-5 md:pt-5">
                <RouterLinkside id="hs-pro-sidebar"
                    class="w-[260px] hs-overlay [--auto-close:lg] hs-overlay-open:translate-x-0 -translate-x-full transition-all duration-300 transform hidden sticky inset-y-0 start-0 z-[60] h-full top-0 bg-white border border-neutral-200 lg:block lg:absolute lg:start-auto lg:z-40 lg:translate-x-0 lg:inset-y-auto lg:h-[700px] lg:rounded-xl dark:bg-neutral-800 dark:border-neutral-700"
                    tabindex="-1" aria-label="Mini Sidebar">
                    <Sidebar />
                    <div class="lg:hidden absolute top-3 -end-3 z-10">
                        <!-- Sidebar Close -->
                        <button type="button"
                            class="w-6 h-7 inline-flex justify-center items-center gap-x-2 text-sm font-medium rounded-md border border-neutral-200 bg-white text-neutral-500 hover:bg-neutral-50 focus:outline-none focus:bg-neutral-100 disabled:opacity-50 disabled:pointer-events-none dark:bg-neutral-800 dark:border-neutral-700 dark:text-neutral-300 dark:hover:bg-neutral-700 dark:focus:bg-neutral-700"
                            data-hs-overlay="#hs-pro-sidebar">
                            <svg class="shrink-0 size-4" xmlns="http://www.w3.org/2000/svg" width="24" height="24"
                                viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"
                                stroke-linecap="round" stroke-linejoin="round">
                                <polyline points="7 8 3 12 7 16" />
                                <line x1="21" x2="11" y1="12" y2="12" />
                                <line x1="21" x2="11" y1="6" y2="6" />
                                <line x1="21" x2="11" y1="18" y2="18" />
                            </svg>
                        </button>
                        <!-- End Sidebar Close -->
                    </div>
                </RouterLinkside>
                <div class="lg:ps-[280px] space-y-5">
                    <div
                        class="p-4 flex flex-col min-h-screen  bg-white border border-neutral-200 shadow-sm rounded-xl dark:bg-neutral-800 dark:border-neutral-700">
                        <slot></slot>
                    </div>
                </div>
            </div>
        </div>
        </div>
    </main>
    <AppFooter />
    <AppScrollUp />
</template>