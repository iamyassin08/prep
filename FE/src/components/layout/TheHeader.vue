<script setup lang="ts">
import { ref } from 'vue';

const showMenu = ref(false);

const toggleMenu = () => {
  showMenu.value = !showMenu.value;
};
</script>

<template>
  <header class="w-full py-4 px-6 dark:bg-neutral-800 z-40">
    <nav class="relative max-w-7xl mx-auto flex items-center justify-between">

      <!-- Mobile Menu Toggle -->
      <div class="sm:hidden flex items-center">
        <div @click="toggleMenu" class="cursor-pointer text-white">
          <svg v-if="!showMenu" xmlns="http://www.w3.org/2000/svg" class="h-6 w-6" fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 6h16M4 12h16m-7 6h7" />
          </svg>
          <svg v-else xmlns="http://www.w3.org/2000/svg" class="h-6 w-6" fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12" />
          </svg>
        </div>
      </div>

      <!-- Logo -->
      <div class="flex items-center flex-shrink-0 mr-4">
        <RouterLink
          to="/"
          class="flex items-center gap-x-1 text-xl font-bold text-green-500 transition-transform duration-300 transform hover:scale-105"
          aria-label="Brand"
        >
          <span class="whitespace-nowrap hidden xl:block">Hello World</span>
          <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="lucide lucide-earth"><path d="M21.54 15H17a2 2 0 0 0-2 2v4.54"/><path d="M7 3.34V5a3 3 0 0 0 3 3a2 2 0 0 1 2 2c0 1.1.9 2 2 2a2 2 0 0 0 2-2c0-1.1.9-2 2-2h3.17"/><path d="M11 21.95V18a2 2 0 0 0-2-2a2 2 0 0 1-2-2v-1a2 2 0 0 0-2-2H2.05"/><circle cx="12" cy="12" r="10"/></svg>
        </RouterLink>
      </div>

      <!-- Desktop and Mobile Navigation Links -->
      <div
        :class="{
          'block': showMenu,
          'hidden': !showMenu,
          'sm:flex': true
        }"
        class="flex-1 flex items-center justify-center sm:justify-between space-x-6 md:space-x-8"
      >
        <!-- Desktop Navigation Links -->
        <div class="flex items-center space-x-6 text-gray-800 dark:text-white">
          <RouterLink
            to="/"
            class="relative inline-block text-black font-medium hover:text-white hover:bg-green-500 hover:rounded-lg px-4 py-2 transition"
          >
            Home
          </RouterLink>
          
          <RouterLink
            to="/prep"
            class="inline-block text-black font-medium hover:text-white hover:bg-green-500 hover:rounded-lg px-4 py-2 transition"
          >
            Dashboard
          </RouterLink>
        </div>

        <!-- Theme Switcher and Other Buttons -->
        <div class="flex items-center space-x-4">
          <!-- Dark Mode Toggle -->
          <button
            type="button"
            class="hs-dark-mode-active:hidden block hs-dark-mode font-medium text-gray-800 rounded-full hover:bg-gray-200 focus:outline-none dark:text-neutral-200 dark:hover:bg-neutral-800"
            data-hs-theme-click-value="dark"
          >
            <span class="group inline-flex justify-center items-center">
              <svg xmlns="http://www.w3.org/2000/svg" class="w-6 h-6" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                <path d="M12 3a6 6 0 0 0 9 9 9 9 0 1 1-9-9Z"></path>
              </svg>
            </span>
          </button>
          <!-- Light Mode Toggle -->
          <button
            type="button"
            class="hs-dark-mode-active:block hidden hs-dark-mode font-medium text-gray-800 rounded-full hover:bg-gray-200 focus:outline-none dark:text-neutral-200 dark:hover:bg-neutral-800"
            data-hs-theme-click-value="light"
          >
            <span class="group inline-flex justify-center items-center">
              <svg xmlns="http://www.w3.org/2000/svg" class="w-6 h-6" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
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
            </span>
          </button>
        </div>
      </div>
    </nav>
  </header>
</template>