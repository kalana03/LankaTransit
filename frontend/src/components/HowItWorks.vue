<script setup lang="ts">
import { Search, TicketCheck, BusFront } from 'lucide-vue-next'
</script>

<template>
  <section class="w-full max-w-5xl mx-auto px-4 font-sans relative">
    
    <!-- Section Header -->
    <!-- <div class="text-center mb-16 lg:mb-16">
      <p class="text-slate-300 text-lg lg:text-xl font-medium">
        Your journey across Sri Lanka starts with three simple steps.
      </p>
    </div> -->

    <!-- Process Steps Container -->
    <div class="relative flex flex-col md:flex-row items-center md:items-start justify-between gap-12 md:gap-4">
      
      <!-- DESKTOP Connecting Line -->
      <div class="hidden md:block absolute top-8 left-[15%] right-[15%] h-[2px] bg-white/[0.08] z-0"></div>
      
      <!-- MOBILE Connecting Line -->
      <div class="block md:hidden absolute top-[10%] bottom-[10%] left-[50%] w-[2px] bg-white/[0.08] -translate-x-1/2 z-0"></div>

      <!-- Step 1 -->
      <div class="relative z-10 flex flex-col items-center text-center w-full md:flex-1 group">
        <div class="w-16 h-16 rounded-2xl bg-[#0a1128] border border-white/[0.15] flex items-center justify-center mb-4 transition-transform duration-300 group-hover:-translate-y-1.5 group-hover:border-[#5fa8ff] shadow-lg shadow-black/50">
          <Search class="w-7 h-7 text-[#5fa8ff]" />
        </div>
        <!-- FIXED TEXT COLORS FOR DARK THEME -->
        <h3 class="text-xl font-bold text-white mb-2">1. Find your route</h3>
        <p class="text-sm text-slate-400 max-w-[250px] md:max-w-[300px] leading-relaxed">
          Enter your departure, destination, and dates to compare available buses instantly.
        </p>
      </div>

      <!-- Step 2 -->
      <div class="relative z-10 flex flex-col items-center text-center w-full md:flex-1 group">
        <div class="w-16 h-16 rounded-2xl bg-[#0a1128] border border-white/[0.15] flex items-center justify-center mb-4 transition-transform duration-300 group-hover:-translate-y-1.5 group-hover:border-[#5fa8ff] shadow-lg shadow-black/50">
          <TicketCheck class="w-7 h-7 text-[#5fa8ff]" />
        </div>
        <h3 class="text-xl font-bold text-white mb-2">2. Choose & Book</h3>
        <p class="text-sm text-slate-400 max-w-[250px] md:max-w-[300px] leading-relaxed">
          Select your preferred seats, verify details, and check out securely.
        </p>
      </div>

      <!-- Step 3 -->
      <div class="relative z-10 flex flex-col items-center text-center w-full md:flex-1 group">
        <div class="w-16 h-16 rounded-2xl bg-[#0a1128] border border-white/[0.15] flex items-center justify-center mb-4 transition-transform duration-300 group-hover:-translate-y-1.5 group-hover:border-[#5fa8ff] shadow-lg shadow-black/50">
          <BusFront class="w-7 h-7 text-[#5fa8ff]" />
        </div>
        <h3 class="text-xl font-bold text-white mb-2">3. Enjoy the ride</h3>
        <p class="text-sm text-slate-400 max-w-[250px] md:max-w-[300px] leading-relaxed">
          Receive your digital ticket immediately. Just show it on your phone and board.
        </p>
      </div>

    </div>
  </section>
</template>