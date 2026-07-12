<script setup lang="ts">
import { MapPin, Calendar, Clock, BusFront, User, Phone } from 'lucide-vue-next';

defineProps({
  tour: {
    type: Object,
    required: true
  }
});
</script>

<template>
  <div class="w-full bg-[#151b29] rounded-2xl border border-slate-800/80 hover:border-slate-700 hover:shadow-[0_4px_20px_rgba(0,0,0,0.4)] transition-all duration-300 p-5 lg:py-4 px-6 group">
    
    <!-- Expanded 8-Column Grid Layout -->
    <div class="grid grid-cols-2 lg:grid-cols-8 gap-4 lg:gap-6 w-full items-center">
      
      <!-- 1. Tour ID -->
      <div class="flex flex-col col-span-1">
        <span class="text-[10px] text-slate-500 uppercase tracking-widest font-bold mb-1.5">Tour ID</span>
        <span class="text-white font-mono font-bold text-sm bg-[#0a0d14] px-3 py-1.5 rounded-lg border border-slate-800/80 inline-block w-max">
          {{ tour.id }}
        </span>
      </div>

      <!-- 2. Route (Takes up 2 columns) -->
      <div class="flex flex-col col-span-2">
        <span class="text-[10px] text-slate-500 uppercase tracking-widest font-bold mb-1.5 flex items-center gap-1.5">
          <MapPin class="w-3 h-3 text-[#5eafff]" /> Route
        </span>
        <div class="flex items-center gap-2.5 text-slate-200 font-bold text-sm">
          <span class="truncate">{{ tour.start }}</span>
          <svg class="w-4 h-4 text-slate-600 shrink-0" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M14 5l7 7m0 0l-7 7m7-7H3"></path></svg>
          <span class="truncate">{{ tour.destination }}</span>
        </div>
      </div>

      <!-- 3. Date -->
      <div class="flex flex-col col-span-1">
        <span class="text-[10px] text-slate-500 uppercase tracking-widest font-bold mb-1.5 flex items-center gap-1.5">
          <Calendar class="w-3 h-3 text-[#5eafff]" /> Date
        </span>
        <span class="text-slate-200 font-bold text-sm">{{ tour.date }}</span>
      </div>

      <!-- 4. Time -->
      <div class="flex flex-col col-span-1">
        <span class="text-[10px] text-slate-500 uppercase tracking-widest font-bold mb-1.5 flex items-center gap-1.5">
          <Clock class="w-3 h-3 text-[#5eafff]" /> Time
        </span>
        <span class="text-slate-200 font-bold text-sm">{{ tour.time }}</span>
      </div>

      <!-- 5. Bus Company -->
      <div class="flex flex-col col-span-1 lg:pl-4 lg:border-l border-slate-800/80">
        <span class="text-[10px] text-slate-500 uppercase tracking-widest font-bold mb-1.5 flex items-center gap-1.5">
          <BusFront class="w-3 h-3 text-[#5eafff]" /> Operator
        </span>
        <span class="text-slate-200 font-bold text-sm truncate">{{ tour.busCompany }}</span>
      </div>

      <!-- 6. Driver Name -->
      <div class="flex flex-col col-span-1">
        <span class="text-[10px] text-slate-500 uppercase tracking-widest font-bold mb-1.5 flex items-center gap-1.5">
          <User class="w-3 h-3 text-[#5eafff]" /> Driver
        </span>
        <span class="text-slate-200 font-bold text-sm truncate">{{ tour.driverName }}</span>
      </div>

      <!-- 7. Driver Contact -->
      <div class="flex flex-col col-span-1">
        <span class="text-[10px] text-slate-500 uppercase tracking-widest font-bold mb-1.5 flex items-center gap-1.5">
          <Phone class="w-3 h-3 text-[#5eafff]" /> Contact
        </span>
        <span class="text-slate-400 font-mono text-xs font-bold">{{ tour.driverPhone }}</span>
      </div>

    </div>
  </div>
</template>