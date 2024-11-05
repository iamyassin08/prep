<script setup lang="ts">
// import { ref, onBeforeMount } from "vue";
// import { getAllUsers, updateUser, addUser } from "@/services/api"; 



</script>

<template>
    <div class="container mx-auto px-4 py-6">
      <!-- Header -->
      <h1 class="text-2xl font-semibold text-gray-800 dark:text-white mb-4">User Management</h1>
  
    
      <div class="overflow-x-auto bg-white dark:bg-neutral-800 shadow-lg rounded-lg">
        <h2 class="text-xl font-semibold text-gray-800 dark:text-white mb-4">Users List</h2>
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
  
  
      <div class="mt-6 bg-white dark:bg-neutral-800 p-6 shadow-lg rounded-lg">
        <h2 class="text-xl font-medium text-gray-800 dark:text-white mb-4">Add New User</h2>
        <form>
          <div class="space-y-4">
            <div>
              <label for="userName" class="block text-gray-700 dark:text-neutral-300">Name</label>
              <input
                type="text"
                id="userName"
                value="New User"
                class="mt-1 block w-full px-4 py-2 border rounded-lg focus:outline-none focus:ring-2 focus:ring-green-600 dark:bg-neutral-900 dark:text-white dark:border-neutral-700"
                placeholder="Enter user name"
              />
            </div>
            <div>
              <label for="userEmail" class="block text-gray-700 dark:text-neutral-300">Email</label>
              <input
                type="email"
                id="userEmail"
                value="newuser@example.com"
                class="mt-1 block w-full px-4 py-2 border rounded-lg focus:outline-none focus:ring-2 focus:ring-green-600 dark:bg-neutral-900 dark:text-white dark:border-neutral-700"
                placeholder="Enter user email"
              />
            </div>
          </div>
          <div class="mt-6 flex gap-4">
            <button type="submit" class="px-6 py-2 bg-green-600 text-white rounded-lg hover:bg-green-700">Create User</button>
            <button type="button" class="px-6 py-2 bg-gray-400 text-white rounded-lg hover:bg-gray-500">Cancel</button>
          </div>
        </form>
      </div>
    </div>
  </template>