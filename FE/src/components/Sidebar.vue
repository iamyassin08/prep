<script setup lang="ts">
import { useUserStore } from '@/stores/userStore'; 
import NoAccess from '../NoAccess.vue';

const userStore = useUserStore();
</script>

<template>
   <!-- Navigation Toggle -->
   <RouterLink
   to="/"
   type="button"
   class="text-gray-500 hover:text-gray-500 mt-6 transition duration-500 ease-in-out transform hover:scale-110 sm:hidden"
   data-hs-overlay="#docs-sidebar"
   aria-controls="docs-sidebar"
   aria-label="Toggle navigation"
 >
   <span class="sr-only">Toggle Navigation</span>
   <svg
     xmlns="http://www.w3.org/2000/svg"
     width="24"
     height="24"
     viewBox="0 0 24 24"
     fill="none"
     stroke="currentColor"
     stroke-width="2"
     stroke-linecap="round"
     stroke-linejoin="round"
     class="lucide lucide-ellipsis-vertical"
   >
     <circle cx="12" cy="12" r="1" />
     <circle cx="12" cy="5" r="1" />
     <circle cx="12" cy="19" r="1" />
   </svg>
 </RouterLink>
 <!-- End Navigation Toggle -->

  <!-- <div
    id="docs-sidebar"
    class="hs-overlay [--auto-close:lg]  hs-overlay-open:translate-x-0 -translate-x-full transition-transform duration-300 ease-in-out transform fixed top-0 start-0 bottom-0 z-[60] w-64 bg-white border-e border-gray-200 pt-7 pb-10 overflow-y-auto lg:block lg:translate-x-0 lg:end-auto lg:bottom-0 [&::-webkit-scrollbar]:w-2 [&::-webkit-scrollbar-thumb]:rounded-full [&::-webkit-scrollbar-track]:bg-gray-100 [&::-webkit-scrollbar-thumb]:bg-gray-300 dark:[&::-webkit-scrollbar-track]:bg-neutral-700 dark:[&::-webkit-scrollbar-thumb]:bg-neutral-500 dark:bg-neutral-800 dark:border-neutral-700"
  > -->
    <div class="px-4 mt-12"></div>

    <nav
      class="hs-accordion-group p-6 w-full flex flex-col flex-wrap"
      data-hs-accordion-always-open
    >
      <ul class="space-y-1.5" role="tablist">
        <!-- Admin Dashboard -->
        <li v-if="userStore.isAdmin" class="hs-accordion lg:block hidden" id="dashboard-accordion">
          <RouterLink to="/admindashboard">
            <button
              type="button"
              id="dashboard"
              role="tab"
              aria-controls="basic-tabs-1"
              aria-selected="false" data-hs-tab="#basic-tabs-1"
              class="flex hs-tab-active:bg-red-500 items-center gap-x-3.5 py-2 px-2.5 dark:text-neutral-300 bg-gray-100  text-sm text-gray-700 rounded-lg hover:bg-gray-200 dark:bg-neutral-700 dark:text-neutral-300 dark:hover:bg-neutral-600 transition duration-500 ease-in-out transform hover:scale-105"
            >
              <svg
                xmlns="http://www.w3.org/2000/svg"
                width="24"
                height="24"
                viewBox="0 0 24 24"
                fill="none"
                stroke="currentColor"
                stroke-width="2"
                stroke-linecap="round"
                stroke-linejoin="round"
                class="lucide lucide-layout-dashboard"
              >
                <rect width="7" height="9" x="3" y="3" rx="1" />
                <rect width="7" height="5" x="14" y="3" rx="1" />
                <rect width="7" height="9" x="14" y="12" rx="1" />
                <rect width="7" height="5" x="3" y="16" rx="1" />
              </svg>
              Admin Dashboard
            </button>
          </RouterLink>
        </li>





        <!-- Profile -->
        <li class="hs-accordion lg:block hidden" id="profile-accordion">
          <RouterLink to="/user/dashboard">
            <button
              id="profile"
              type="button"
              role="tab"
              aria-controls="basic-tabs-2"
              aria-selected="false" data-hs-tab="#basic-tabs-2"
              class="hs-accordion-toggle hs-tab-active:bg-red-500 hs-accordion-active:text-blue-600 dark:text-neutral-300 hs-accordion-active:hover:bg-transparent w-full text-start flex items-center gap-x-3.5 py-2 px-2.5 text-sm text-gray-700 rounded-lg hover:bg-gray-200 dark:bg-neutral-800 dark:hover:bg-neutral-700  dark:hover:text-neutral-300 dark:hs-accordion-active:text-white transition duration-500 ease-in-out transform hover:scale-105"
            >
              <svg
                xmlns="http://www.w3.org/2000/svg"
                width="24"
                height="24"
                viewBox="0 0 24 24"
                fill="none"
                stroke="currentColor"
                stroke-width="2"
                stroke-linecap="round"
                stroke-linejoin="round"
                class=" size-4 lucide lucide-user"
              >
                <path d="M20 21v-2a4 4 0 0 0-4-4H8a4 4 0 0 0-4 4v2" />
                <circle cx="12" cy="7" r="4" />
              </svg>
              Dashboard
            </button>
          </RouterLink>
        </li>
        <!-- Admin Category -->
        
       

      </ul>
    </nav>
  <!-- </div> -->

  <!-- NoAccess Component -->
  <div v-if="!userStore.isAdmin && !userStore.isBuyer && !userStore.isSeller">
    <NoAccess />
  </div>
</template>
