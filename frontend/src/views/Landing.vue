<script setup lang="ts">
import Header from '../components/Header.vue';
import HowItWorks from '../components/HowItWorks.vue';
import SearchForm from '../components/SearchForm.vue';
</script>

<template>
  <!-- Added 'relative' and 'bg-center' to ensure the image centers nicely -->
  <div class="relative bg-[url('/bg.jpeg')] bg-cover bg-center min-h-screen font-sans">
    
    <!-- THE OVERLAY: Dark blue color with 85% opacity -->
    <div class="absolute inset-0 bg-[#0a1128]/70 z-0"></div>

    <!-- Wrapped content in 'relative z-10' so it sits above the background overlay -->
    <div class="relative z-10">
      <Header />

      <!-- 
        Adjusted padding and minimum height for mobile.
        Removed the fixed 'h-[100dvh]' so mobile scrolling works perfectly.
      -->
      <main class="flex flex-col items-center justify-center w-full pt-32 lg:pt-[12rem] pb-16 lg:pb-[8rem] min-h-screen text-white px-4">
        
        <div class="text-center w-full max-w-5xl">
          <!-- Responsive typography (smaller on mobile, 7xl on desktop) -->
          <h1 class="text-5xl lg:text-7xl font-bold mb-4 lg:mb-6 tracking-tight">
            Book your tickets <span class="text-[#5fa8ff]">with ease.</span>
          </h1>
          
          <p class="text-lg lg:text-2xl mb-12 lg:mb-20 font-medium text-slate-300 px-4">
            Your reliable partner for seamless transportation across Sri Lanka.
          </p>

          <SearchForm />
        </div>  
          
        <!-- Margin top on mobile to separate form from how it works -->
        <div class="mt-32 w-full">
          <HowItWorks />
        </div>

      </main>
    </div>
  </div>
</template>