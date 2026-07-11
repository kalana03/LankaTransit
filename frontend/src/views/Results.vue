<script setup lang="ts">
import Header from '../components/Header.vue';
import ResultCard from '../components/ResultCard.vue';
import SearchFormResults from '../components/SearchFormResults.vue';

// Dummy data for loop
const results = [1, 2, 3, 4];
</script>

<template>
  <div class="min-h-screen bg-dark text-white">
    <Header />

    <!-- Main wrapper -->
    <main class="max-w-7xl mx-auto px-4 pt-[8rem] pb-8">
      
      <!-- Grid container with alignment -->
      <div class="grid grid-cols-1 lg:grid-cols-12 gap-2 items-start">
        
        <!-- Sidebar (Sticky) -->
        <!-- 1. sticky top-[8rem] aligns it exactly below your header -->
        <!-- 2. z-10 ensures it stays above results if content scrolls -->
        <!-- 3. max-h-[calc(100vh-8rem)] ensures it doesn't get cut off on mobile if the form is long -->
        <aside class="sm:sticky top-[8rem] z-10 lg:col-span-4 xl:col-span-3 overflow-y-auto max-h-[calc(100vh-9rem)]">
          <SearchFormResults />
        </aside>

        <!-- Results Column -->
        <div class="lg:col-span-8 xl:col-span-9 flex flex-col gap-4">
          <ResultCard v-for="i in results" :key="i" />
        </div>

      </div>
    </main>
  </div>
</template>