<script setup lang="ts">

</script>

<template>
  <footer class="mt-auto w-full max-w-[85rem] py-10 px-4 sm:px-6 lg:px-8 mx-auto">
    <!-- Grid Layout -->
    <div class="grid grid-cols-1 md:grid-cols-3 items-center gap-5">
      <!-- Brand Section -->
      <div class="flex items-center flex-shrink-0 w-full md:w-auto mb-4 md:mb-0 mr-4">
        <RouterLink
          to="/"
          class="flex items-center gap-x-1 text-xl font-bold text-green-500 transition-transform duration-300 transform-gpu hover:scale-105"
          aria-label="Brand"
        >
          <span class="whitespace-nowrap hidden xl:block">Hello World</span>
          <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="lucide lucide-earth"><path d="M21.54 15H17a2 2 0 0 0-2 2v4.54"/><path d="M7 3.34V5a3 3 0 0 0 3 3a2 2 0 0 1 2 2c0 1.1.9 2 2 2a2 2 0 0 0 2-2c0-1.1.9-2 2-2h3.17"/><path d="M11 21.95V18a2 2 0 0 0-2-2a2 2 0 0 1-2-2v-1a2 2 0 0 0-2-2H2.05"/><circle cx="12" cy="12" r="10"/></svg>
        </RouterLink>
      </div>

      <!-- Navigation Links -->
      <nav>
        <ul class="text-center">
          <li class="inline-block relative pe-8 last:pe-0 last-of-type:before:hidden before:absolute before:top-1/2 before:end-3 before:-translate-y-1/2 before:content-['/'] before:text-neutral-300 dark:before:text-neutral-600">
            <RouterLink to="/" class="inline-flex gap-x-2 text-sm text-neutral-500 hover:text-neutral-800 focus:outline-none focus:text-neutral-800 dark:text-neutral-500 dark:hover:text-neutral-200 dark:focus:text-neutral-200">
              Home
            </RouterLink>
          </li>
          <li class="inline-block relative pe-8 last:pe-0 last-of-type:before:hidden before:absolute before:top-1/2 before:end-3 before:-translate-y-1/2 before:content-['/'] before:text-neutral-300 dark:before:text-neutral-600">
            <RouterLink to="/prep" class="inline-flex gap-x-2 text-sm text-neutral-500 hover:text-neutral-800 focus:outline-none focus:text-neutral-800 dark:text-neutral-500 dark:hover:text-neutral-200 dark:focus:text-neutral-200">
              Dashboard
            </RouterLink>
          </li>
        </ul>
      </nav>

      <!-- Social Media Section -->
      <div class="md:text-end space-x-2">
        <a  href="/https://github.com/iamyassin08"

         class="size-8 inline-flex justify-center items-center gap-x-2 text-sm font-semibold rounded-full border border-transparent text-neutral-500 hover:bg-neutral-50 focus:outline-none focus:bg-neutral-50 disabled:opacity-50 disabled:pointer-events-none dark:text-neutral-400 dark:hover:bg-neutral-700 dark:focus:bg-neutral-700">
          <svg class="shrink-0 size-3.5" xmlns="http://www.w3.org/2000/svg" width="16" height="16" fill="currentColor" viewBox="0 0 16 16">
            <path d="M8 0C3.58 0 0 3.58 0 8c0 3.54 2.29 6.53 5.47 7.59.4.07.55-.17.55-.38 0-.19-.01-.82-.01-1.49-2.01.37-2.53-.49-2.69-.94-.09-.23-.48-.94-.82-1.13-.28-.15-.68-.52-.01-.53.63-.01 1.08.58 1.23.82.72 1.21 1.87.87 2.33.66.07-.52.28-.87.51-1.07-1.78-.2-3.64-.89-3.64-3.95 0-.87.31-1.59.82-2.15-.08-.2-.36-1.02.08-2.12 0 0 .67-.21 2.2.82.64-.18 1.32-.27 2-.27.68 0 1.36.09 2 .27 1.53-1.04 2.2-.82 2.2-.82.44 1.1.16 1.92.08 2.12.51.56.82 1.27.82 2.15 0 3.07-1.87 3.75-3.65 3.95.29.25.54.73.54 1.48 0 1.07-.01 1.93-.01 2.2 0 .21.15.46.55.38A8.012 8.012 0 0 0 16 8c0-4.42-3.58-8-8-8z" />
          </svg>
        </a>
      </div>
    </div>
    <!-- End Grid Layout -->
  </footer>
</template>


