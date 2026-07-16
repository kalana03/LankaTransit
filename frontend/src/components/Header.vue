<script setup lang="ts">
import { ref } from 'vue';


const steps = [
  { num: 1, name: 'Select Route', status: 'completed' },
  { num: 2, name: 'Choose Seats', status: 'current' },
  { num: 3, name: 'Payment', status: 'upcoming' },
];
</script>

<template>
  <!-- Fixed Header: Frosted Glass over #0f172a background -->
  <header class="fixed top-0 left-0 w-full z-50 bg-[#0f172a]/0 backdrop-blur-xl border-b border-slate-800 transition-all font-sans">
    
    <!-- Top Row: Logo & Navigation -->
    <div class="h-16 w-full max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 flex items-center justify-between border-b border-slate-800/50">
      
      <!-- Logo: LankaTransit -->
      <a href="/" class="flex items-center gap-2 group">
        <svg class="w-6 h-6 text-[#5eafff] group-hover:scale-110 transition-transform" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2.5" d="M13 5l7 7-7 7M5 5l7 7-7 7"></path>
        </svg>
        <span class="text-xl font-bold text-white tracking-wide">
          Lanka<span class="text-[#5eafff]">Transit</span>
        </span>
      </a>

      

      <!-- Right Side: User Profile / Account Pill -->
      <div class="flex items-center gap-4">
        
        
        <!-- User Avatar Pill -->
        <button class="flex items-center gap-2 pl-2 pr-4 py-1.5 bg-slate-800 rounded-full border border-slate-700 hover:border-slate-500 transition-colors cursor-pointer group">
          <div class="w-6 h-6 rounded-full bg-gradient-to-tr from-[#5eafff] to-indigo-500 flex items-center justify-center text-[10px] text-white font-bold shadow-sm">
            JD
          </div>
          <span class="text-sm font-medium text-slate-200 group-hover:text-white">John Doe</span>
        </button>
      </div>

    </div>

    
  </header>
</template>