<script setup lang="ts">
import HelloWorld from "@/components/HelloWorld.vue";


</script>

<template>
 
 <MotionGroup preset="slideVisibleOnceLeft" :delay="340" :duration="1200">


  <HelloWorld />


</MotionGroup>

</template>
