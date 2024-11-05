<script setup lang="ts">
import { ref } from 'vue';
import { useSettingStore } from "@/stores/settings";
import AccountDropDown from "@/components/AccountDropDown.vue";


const showMenu = ref(false);

const toggleMenu = () => {
  showMenu.value = !showMenu.value;
};
</script>

<template>
  <header class="flex flex-wrap md:justify-between items-center z-40 w-full py-4 px-6 dark:bg-neutral-800">
    <nav class="relative max-w-7xl w-full flex flex-wrap items-center mx-auto">

      <div class="flex items-center justify-between sm:hidden">
        <div @click="toggleMenu" class="cursor-pointer text-white">
          <svg v-if="!showMenu" xmlns="http://www.w3.org/2000/svg" class="h-6 w-6" fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 6h16M4 12h16m-7 6h7" />
          </svg>
          <svg v-else xmlns="http://www.w3.org/2000/svg" class="h-6 w-6" fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12" />
          </svg>
        </div>
      </div>

      <div class="flex items-center flex-shrink-0 w-full md:w-auto mb-4 md:mb-0 mr-4">
        <RouterLink
          to="/"
          class="flex items-center gap-x-1 text-xl font-bold text-red-500 transition-transform duration-300 transform-gpu hover:scale-105"
          aria-label="Brand"
        >
          <span class="whitespace-nowrap hidden xl:block">Hello World</span>
          <span class="whitespace-nowrap hidden md:block xl:hidden">A&B</span>
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
            class="w-6 h-6 text-red-500"
          >
            <path
              d="M12.586 2.586A2 2 0 0 0 11.172 2H4a2 2 0 0 0-2 2v7.172a2 2 0 0 0 .586 1.414l8.704 8.704a2.426 2.426 0 0 0 3.42 0l6.58-6.58a2.426 2.426 0 0 0 0-3.42z"
            />
            <circle cx="7.5" cy="7.5" r=".5" fill="currentColor" />
          </svg>
        </RouterLink>
      </div>

      <button
        @click="toggleMenu"
        class="md:hidden flex items-center p-2 transition-transform duration-300 transform hover:scale-110"
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
        >
          <path d="M3 12h18M3 6h18M3 18h18"></path>
        </svg>
      </button>

      <div
        :class="{
          'block': showMenu,
          'hidden': !showMenu,
          'sm:block': true
        }"
        class="md:flex flex-1 flex-wrap items-center justify-between mb-4 md:mb-0"
      >
        <div class="flex items-center gap-x-2">
          <RouterLink
            to="/"
            class="relative inline-block text-black before:absolute before:bottom-0.5 before:start-0 before:-z-[1] font-medium dark:text-white hover:text-white hover:bg-red-500 hover:rounded-lg px-3 py-2"
          >
            Home
          </RouterLink>
         
          <RouterLink
          to="/categories"
          class="inline-block text-black font-medium hover:text-white hover:bg-red-500 hover:rounded-lg px-3 py-2 dark:text-white"
        >
        Dashbaord
        </RouterLink>

       
          
         
        </div>
        
        <div class="flex-1 mx-4 md:mx-8">
          <SearchBar />
        </div>

        <div class="flex items-center gap-x-2 md:gap-x-4">
          <button
            @click="settings.toggleDarkMode()"
            type="button"
            class="flex items-center text-gray-600 hover:text-red-500 font-medium dark:hover:text-neutral-500"
          >
            <svg
              v-if="!settings.isDarkMode"
              class="flex-shrink-0 size-4"
              xmlns="http://www.w3.org/2000/svg"
              width="24"
              height="24"
              viewBox="0 0 24 24"
              fill="none"
              stroke="currentColor"
              stroke-width="2"
              stroke-linecap="round"
              stroke-linejoin="round"
            >
              <path d="M12 3a6 6 0 0 0 9 9 9 9 0 1 1-9-9Z"></path>
            </svg>
            <svg
              v-else
              class="flex-shrink-0 size-4"
              xmlns="http://www.w3.org/2000/svg"
             width="34" height="49"
              viewBox="0 0 24 24"
              fill="none"
              stroke="currentColor"
              stroke-width="2"
              stroke-linecap="round"
              stroke-linejoin="round"
            >
              <circle cx="12" cy="12" r="4"></circle>
              <path d="M12 2v2"></path>
              <path d="M12 20v2"></path>
              <path d="m4.93 4.93 1.41 1.41"></path>
              <path d="m17.66 17.66 1.41 1.41"></path>
              <path d="M2 12h2"></path>
              <path d="M20 12h2"></path>
              <path d="m6.34 17.66-1.41 1.41"></path>
              <path d="m19.07 4.93-1.41 1.41"></path>
            </svg>
          </button>

          <Cart />

          <Favorites />

          <div class="hidden sm:block border-e border-gray-200 w-px h-6 mx-1.5 dark:border-neutral-700"></div>
          <AccountDropDown class="z-40" />
        </div>
      </div>
    </nav>
  </header>
</template>
