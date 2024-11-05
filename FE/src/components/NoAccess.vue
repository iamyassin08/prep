<template>
<div class="p-2 sm:p-5 sm:pb-0 space-y-5 text-center">
    <!-- Content -->
    <div class="max-w-md mx-auto space-y-3">
      <h1 class="leading-none text-8xl font-semibold text-gray-800 dark:text-neutral-200">
        500
      </h1>
      <h2 class="text-3xl font-medium tracking-tight text-gray-800 dark:text-neutral-200">
        Page not found
      </h2>
      <p class="text-gray-500 dark:text-neutral-500">
        The server encountered an internal error or misconfiguration and was unable to complete your request.
      </p>
      <div>
        <RouterLink :to="'/'" type="button" class="py-2 px-3 inline-flex justify-center items-center gap-x-2 text-start bg-blue-600 border border-blue-600 text-white text-xs sm:text-sm font-medium rounded-lg shadow-sm align-middle hover:bg-blue-700 focus:outline-none focus:ring-1 focus:ring-blue-300 dark:focus:ring-blue-500">
          Go Home
        </RouterLink>
      </div>
    </div>
    <!-- End Content -->
  </div>
</template>
  