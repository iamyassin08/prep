<script setup lang="ts">
// import { ref, onBeforeMount } from "vue";
// import { getAllUsers, updateUser, addUser } from "@/services/api"; 



</script>

<template>
    <div class="container mx-auto px-4 py-6">
      <!-- Header -->
  
    
      <div class="overflow-x-auto bg-white dark:bg-neutral-800 shadow-lg rounded-lg">
        <h2 class="text-xl font-semibold text-gray-800 dark:text-white mb-4">Your Profile</h2>
        <table class="min-w-full table-auto text-sm text-left">
          <thead>
            <tr class="bg-gray-100 dark:bg-neutral-700">
              <th class="px-4 py-2 text-gray-700 dark:text-neutral-200">Name</th>
              <th class="px-4 py-2 text-gray-700 dark:text-neutral-200">Email</th>
              <th class="px-4 py-2 text-gray-700 dark:text-neutral-200">Actions</th>
            </tr>
          </thead>
          <tbody>
        
            <tr class="border-b dark:border-neutral-600">
              <td class="px-4 py-2">John Doe</td>
              <td class="px-4 py-2">johndoe@example.com</td>
              <td class="px-4 py-2">
                <button class="text-blue-600 dark:text-blue-400">Edit</button>
              </td>
            </tr>
            <tr class="border-b dark:border-neutral-600">
              <td class="px-4 py-2">Jane Smith</td>
              <td class="px-4 py-2">janesmith@example.com</td>
              <td class="px-4 py-2">
                <button class="text-blue-600 dark:text-blue-400">Edit</button>
              </td>
            </tr>
            <tr class="border-b dark:border-neutral-600">
              <td class="px-4 py-2">Alice Johnson</td>
              <td class="px-4 py-2">alicej@example.com</td>
              <td class="px-4 py-2">
                <button class="text-blue-600 dark:text-blue-400">Edit</button>
              </td>
            </tr>
          </tbody>
        </table>
      </div>
  
  

    </div>
  </template>