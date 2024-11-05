<script setup lang="ts">
import { ref, onBeforeMount } from 'vue';
import { getAllUsers } from '@/services/api';

const users = ref([]); // Step 1: Define a ref to hold user data

// Step 2: Create a method to fetch users
const fetchUsers = async () => {
  try {
    users.value = await getAllUsers(); // Fetch and assign the user data
  } catch (error) {
    console.error('Failed to fetch users:', error); // Handle any errors
  }
};

// Step 3: Use the onBeforeMount lifecycle hook to fetch users
onBeforeMount(fetchUsers);
</script>

<template>
  <!-- Hero -->
  <div class="overflow-hidden">
    <div class="max-w-[85rem] mx-auto px-4 sm:px-6 lg:px-8 py-20">
      <div class="relative mx-auto max-w-4xl grid space-y-5 sm:space-y-10">
        <!-- Title -->
        <div class="text-center">
          <p class="text-xs font-semibold text-gray-500 tracking-wide uppercase mb-3 dark:text-neutral-200">
            Mini Hello World Project
          </p>
          <h1 class="text-3xl text-gray-800 font-bold sm:text-5xl lg:text-6xl lg:leading-tight dark:text-neutral-200">
            Turn Your <span class="text-green-500">User to Life</span>
          </h1>
        </div>
        <!-- End Title -->
        
        <!-- Table to display users -->
        <div class="overflow-x-auto bg-white shadow-lg rounded-lg border dark:bg-neutral-900 dark:border-neutral-700">
          <table class="min-w-full text-sm text-left text-gray-800 dark:text-white">
            <thead class="text-xs text-gray-700 uppercase bg-gray-100 dark:bg-neutral-700 dark:text-neutral-200">
              <tr>
                <th scope="col" class="py-3 px-6 font-semibold tracking-wider">ID</th>
                <th scope="col" class="py-3 px-6 font-semibold tracking-wider">Name</th>
                <th scope="col" class="py-3 px-6 font-semibold tracking-wider">Email</th>
              </tr>
            </thead>
            <tbody>
              <!-- Step 4: Dynamically render users data -->
              <tr v-for="user in users" :key="user.id" class="border-b hover:bg-gray-50 dark:hover:bg-neutral-800">
                <td class="py-3 px-6">{{ user.id }}</td>
                <td class="py-3 px-6">{{ user.name }}</td>
                <td class="py-3 px-6">{{ user.email }}</td>
              </tr>
            </tbody>
          </table>
        
          <!-- Button Section -->
          <div class="w-full flex justify-center pt-6">
            <button @click="fetchUsers" class="py-3 px-5 inline-flex justify-center items-center gap-x-2 text-sm font-semibold rounded-lg bg-green-600 text-white hover:bg-green-700 focus:outline-none focus:ring-4 focus:ring-green-500 focus:ring-opacity-50 transition duration-200 ease-in-out">
              Get Users
            </button>
          </div>
        </div>
        
        <!-- End Table -->
      </div>
    </div>
  </div>
  <!-- End Hero -->
</template>
