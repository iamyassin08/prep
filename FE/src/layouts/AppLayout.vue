<script setup lang="ts">
import AppFooter from "@/components/layout/Footer.vue";
import AppHeader from "@/components/layout/TheHeader.vue";


</script>
<template>
    <div class="dark:bg-neutral-800">

    <AppHeader />
    <main id="content" class="dark:bg-neutral-800 dark:text-white" role="main">

        <slot></slot>
        
    </main>

    <AppFooter />


    
</div>

</template>