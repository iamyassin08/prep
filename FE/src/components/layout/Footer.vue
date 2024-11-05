<script setup lang="ts">
import { getAllCategories } from "@/services/api";
import { onBeforeMount, ref } from "vue";

const categoryLoading = ref(true);
const categoryList: any = ref([]);

onBeforeMount(async () => {
  try {
    const response = await getAllCategories();
    if (response && response.status === 200) {
      categoryList.value = response.data;
    }
  } catch (error) {
    console.error(error);
  } finally {
    if (categoryList.value.length < 1) {
      console.error("Could not complete request");
    } else {
      categoryLoading.value = false;
    }
  }
});

const getTopLevelListWithChildren = () => {
  let topLevelCategories: any[] = [];
  for (let i = 0; i < categoryList.value.length; i++) {
    if (categoryList.value[i].ParentID == null && categoryHasChildren(categoryList.value[i].ID)) {
      topLevelCategories.push(categoryList.value[i]);
    }
  }
  return topLevelCategories;
};

const categoryHasChildren = (categoryID: number) => {
  for (let i = 0; i < categoryList.value.length; i++) {
    if (categoryList.value[i].ParentID == categoryID) {
      return true;
    }
  }
  return false;
};

const getChildCategoryList = (parentID: number) => {
  let childCategories: any[] = [];
  for (let i = 0; i < categoryList.value.length; i++) {
    if (categoryList.value[i].ParentID == parentID) {
      childCategories.push(categoryList.value[i]);
    }
  }
  return childCategories;
};

const getCategoriesWithoutChildren = () => {
  let categoriesWithoutChildren: any[] = [];
  for (let i = 0; i < categoryList.value.length; i++) {
    if (!categoryHasChildren(categoryList.value[i].ID) && categoryList.value[i].ParentID === null) {
      categoriesWithoutChildren.push(categoryList.value[i]);
    }
  }
  return categoriesWithoutChildren;
};
</script>

<template>
  <!-- ========== FOOTER ========== -->
  <footer class="max-w-[100rem] mx-auto py-10 px-4 sm:px-6 lg:px-8 dark:bg-neutral-900 dark:border-neutral-700 dark:shadow-neutral-700/70 ">
    <!-- Grid -->
    <div class="grid gap-8 lg:grid-cols-5">
      <div class="col-span-1">
        <RouterLink to="/" class="flex items-center text-xl font-bold text-red-500 transition-transform duration-300 transform-gpu hover:scale-105" aria-label="Brand">
          <span>A&B Harmony Haul</span>
          <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="w-6 h-6 text-red-500">
            <path d="M12.586 2.586A2 2 0 0 0 11.172 2H4a2 2 0 0 0-2 2v7.172a2 2 0 0 0 .586 1.414l8.704 8.704a2.426 2.426 0 0 0 3.42 0l6.58-6.58a2.426 2.426 0 0 0 0-3.42z"/>
            <circle cx="7.5" cy="7.5" r=".5" fill="currentColor"/>
          </svg>
        </RouterLink>
        <!-- Static Footer Links -->
        <div class="mt-3 text-base transition-transform duration-300 transform-gpu hover:scale-105">
          <p><RouterLink class="text-gray-900 hover:text-gray-800 dark:text-white dark:hover:text-red-500 font-bold transition-transform duration-300 transform-gpu hover:scale-105" to="/">Home</RouterLink></p>
          <p><RouterLink class="text-gray-900 hover:text-gray-800 dark:text-white dark:hover:text-red-500 transition-transform duration-300 transform-gpu hover:scale-105" to="/productlist">Products</RouterLink></p>
          <p><RouterLink class="text-gray-900 hover:text-gray-800 dark:text-white dark:hover:text-red-500 transition-transform duration-300 transform-gpu hover:scale-105" to="/about">About Us</RouterLink></p>
          <p><RouterLink class="text-gray-900 hover:text-gray-800 dark:text-white dark:hover:text-red-500 transition-transform duration-300 transform-gpu hover:scale-105" to="/contact">Contact</RouterLink></p>
        </div>
      </div>

      <!-- Dynamic Categories -->
      <div v-for="category in getTopLevelListWithChildren()" :key="category.ID" class="col-span-1">
        <div class="mt-8 text-base transition-transform duration-300 transform-gpu hover:scale-105">
          <p class="font-bold">{{ category.Name }}</p>
          <div v-for="child in getChildCategoryList(category.ID)" :key="child.ID">
            <p><RouterLink class="inline-flex gap-x-2 text-gray-900 hover:text-gray-800 dark:text-white dark:hover:text-red-500" :to="'/category/' + child.ID">{{ child.Name }}</RouterLink></p>
          </div>
        </div>
      </div>

      <!-- Categories Without Children -->
      <div class="col-span-1">
        <div class="mt-8 text-base transition-transform duration-300 transform-gpu hover:scale-105">
          <p class="font-bold">Other </p>
          <div v-for="category in getCategoriesWithoutChildren()" :key="category.ID">
            <RouterLink :to="'/category/' + category.ID" class="text-gray-900 hover:text-gray-800 dark:text-white dark:hover:text-red-500">{{ category.Name }}</RouterLink>
          </div>
        </div>
      </div>
    </div>
    <!-- End Grid -->
  </footer>
  <!-- ========== END FOOTER ========== -->
</template>
