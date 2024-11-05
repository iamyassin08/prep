<script setup lang="ts">
import { ref } from "vue";
import usersData from "@/assets/users.json"; 

const users = ref(usersData); 
const userName = ref("");
const userEmail = ref("");
const successMessage = ref("");
const errorMessage = ref("");

const createUser = (event: Event) => {
  event.preventDefault(); 
  
  if (!userName.value || !userEmail.value) {
    errorMessage.value = "Please fill in both fields.";
    return;
  }

  const newUser = {
    ID: users.value.length + 1, 
    FirstName: userName.value,
    Email: userEmail.value,
  };

  users.value.push(newUser); 
  successMessage.value = "User created successfully!";
  

  userName.value = "";
  userEmail.value = "";
};
</script>

<template>
  <div class="container mx-auto px-4 py-6">
    <!-- Header -->
    <h1 class="text-2xl font-semibold text-gray-800 dark:text-white mb-4">User Management</h1>

    <div class="mt-6 bg-white dark:bg-neutral-800 p-6 shadow-lg rounded-lg">
      <h2 class="text-xl font-medium text-gray-800 dark:text-white mb-4">Add New User</h2>
      
      <!-- Success and Error Messages -->
      <div v-if="successMessage" class="mb-4 text-green-600">{{ successMessage }}</div>
      <div v-if="errorMessage" class="mb-4 text-red-600">{{ errorMessage }}</div>

      <form @submit="createUser">
        <div class="space-y-4">
          <div>
            <label for="userName" class="block text-gray-700 dark:text-neutral-300">Name</label>
            <input
              type="text"
              id="userName"
              v-model="userName"
              class="mt-1 block w-full px-4 py-2 border rounded-lg focus:outline-none focus:ring-2 focus:ring-green-600 dark:bg-neutral-900 dark:text-white dark:border-neutral-700"
              placeholder="Enter user name"
              required
            />
          </div>
          <div>
            <label for="userEmail" class="block text-gray-700 dark:text-neutral-300">Email</label>
            <input
              type="email"
              id="userEmail"
              v-model="userEmail"
              class="mt-1 block w-full px-4 py-2 border rounded-lg focus:outline-none focus:ring-2 focus:ring-green-600 dark:bg-neutral-900 dark:text-white dark:border-neutral-700"
              placeholder="Enter user email"
              required
            />
          </div>
        </div>
        <div class="mt-6 flex gap-4">
          <button type="submit" class="px-6 py-2 bg-green-600 text-white rounded-lg hover:bg-green-700">
            Create User
          </button>
          <button type="button" class="px-6 py-2 bg-gray-400 text-white rounded-lg hover:bg-gray-500">Cancel</button>
        </div>
      </form>
    </div>

    <!-- Display Users -->
    <div class="mt-6">
      <h2 class="text-xl font-medium text-gray-800 dark:text-white mb-4">Users List</h2>
      <ul>
        <li v-for="user in users" :key="user.ID" class="border-b py-2">
          {{ user.FirstName }} ({{ user.Email }})
        </li>
      </ul>
    </div>
  </div>
</template>
