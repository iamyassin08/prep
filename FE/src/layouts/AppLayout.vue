<script setup lang="ts">
import AppFooter from "@/components/layout/Footer.vue";
import AppHeader from "@/components/layout/TheHeader.vue";


</script>
<template>
    <div class="dark:bg-neutral-800">

    <!-- <header -->
    <!-- class="fixed top-0 inset-x-0 flex flex-wrap md:justify-start md:flex-nowrap z-50 bg-white border-b border-gray-200 dark:bg-neutral-800 dark:border-neutral-700"> -->
    <AppHeader />
    <main id="content" class="dark:bg-neutral-800 dark:text-white" role="main">
    <!-- <main id="content" role="main" class="pt-[59px]"> -->
        <slot></slot>
    </main>
    <!-- </header> -->
    <AppFooter />


    
</div>

</template>