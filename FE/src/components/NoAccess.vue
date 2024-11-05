<template>
<div class="p-2 sm:p-5 sm:pb-0 space-y-5 text-center">
    <!-- Content -->
    <div class="max-w-md mx-auto space-y-3">
      <h1 class="leading-none text-8xl font-semibold text-neutral-800 dark:text-neutral-200">
        500
      </h1>
      <h2 class="text-3xl font-medium tracking-tight text-neutral-800 dark:text-neutral-200">
        Page not found
      </h2>
      <p class="text-neutral-500 dark:text-neutral-500">
        The server encountegreen an internal error or misconfiguration and was unable to complete your request.
      </p>
      <div>
        <RouterLink :to="'/'" type="button" class="py-2 px-3 inline-flex justify-center items-center gap-x-2 text-start bg-green-600 border border-green-600 text-white text-xs sm:text-sm font-medium rounded-lg shadow-sm align-middle hover:bg-green-700 focus:outline-none focus:ring-1 focus:ring-green-300 dark:focus:ring-green-500">
          Go Home
        </RouterLink>
      </div>
    </div>
    <!-- End Content -->
  </div>
</template>
  