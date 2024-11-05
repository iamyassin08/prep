<script setup lang="ts">
import { RouterView, useRoute } from "vue-router";
import { type IStaticMethods } from "preline/preline";
import { computed } from 'vue';

declare global {
  interface Window {
    HSStaticMethods: IStaticMethods;
  }
}

const route = useRoute();
const layout = computed(() => route.meta.layout || 'div');
</script>

<template>
  <component :is="layout">
    <RouterView />
  </component>
</template>

<style>
body {
  @apply bg-slate-50 text-slate-800 dark:bg-slate-800 dark:text-slate-50;
}
</style>