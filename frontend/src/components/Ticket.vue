<script setup lang="ts">
import { computed } from 'vue';

const props = defineProps({
  busDetails: {
    type: Object,
    required: true
  },
  selectedSeats: {
    type: Array,
    required: true
  }
});

const totalPrice = computed(() => {
  return props.selectedSeats.length * props.busDetails.pricePerSeat;
});

// Helper to generate 3-letter city codes (e.g., Colombo -> COL)
const getCityCode = (city: string) => city.substring(0, 3).toUpperCase();
</script>

<template>
  <!-- Premium Dark-Matte Ticket Shell -->
  <!-- Uses a very subtle gradient and inner border for a high-end physical card feel -->
  <div class="relative w-full max-w-[400px] mx-auto rounded-[2.5rem] bg-gradient-to-b from-[#162032] to-[#0f172a] border border-slate-700/60 shadow-[0_20px_50px_-12px_rgba(0,0,0,0.8)] overflow-hidden font-sans text-slate-200 flex flex-col group">
    
    <!-- Ambient Glow Effect behind the card content -->
    <div class="absolute top-0 left-1/2 -translate-x-1/2 w-3/4 h-32 bg-[#5eafff] opacity-[0.07] blur-[50px] pointer-events-none"></div>

    <!-- Header: Brand & Class -->
    <div class="px-8 pt-8 pb-6 flex justify-between items-center relative z-10">
      <h2 class="text-lg font-bold text-white tracking-wide flex items-center gap-2">
        <!-- Minimalist Logo Mark -->
        <svg class="w-5 h-5 text-[#5eafff]" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2.5" d="M13 5l7 7-7 7M5 5l7 7-7 7"></path></svg>
        Lanka<span class="text-acc">Transit</span>
      </h2>
      
    </div>

    <!-- Hero Route Section (Airport Style) -->
    <div class="px-8 pb-8 relative z-10">
      <div class="flex items-center justify-between">
        
        <!-- Origin -->
        <div class="flex flex-col">
          <span class="text-[40px] font-black text-white leading-none tracking-tighter">{{ getCityCode(busDetails.start) }}</span>
          <span class="text-xs text-slate-400 font-medium mt-1 truncate max-w-[100px]">{{ busDetails.start }}</span>
          <span class="text-[#5eafff] font-bold mt-2">{{ busDetails.departureTime }}</span>
        </div>

        <!-- Animated Connection Graphic -->
        <div class="flex-1 mx-4 relative flex items-center justify-center h-12">
          <!-- Dashed Track -->
          <div class="absolute w-full border-t-[1.5px] border-dashed border-slate-600/60 top-1/2 -translate-y-1/2"></div>
          <!-- Center Icon -->
          <div class="relative bg-[#162032] p-2 rounded-full border border-slate-700/80 shadow-lg">
            <svg class="w-5 h-5 text-[#5eafff]" fill="currentColor" viewBox="0 0 24 24">
              <path d="M4 16c0 .88.39 1.67 1 2.22V20c0 .55.45 1 1 1h1c.55 0 1-.45 1-1v-1h8v1c0 .55.45 1 1 1h1c.55 0 1-.45 1-1v-1.78c.61-.55 1-1.34 1-2.22V6c0-3.5-3.58-4-8-4s-8 .5-8 4v10zm3.5 1c-.83 0-1.5-.67-1.5-1.5S6.67 14 7.5 14s1.5.67 1.5 1.5S8.33 17 7.5 17zm9 0c-.83 0-1.5-.67-1.5-1.5s.67-1.5 1.5-1.5 1.5.67 1.5 1.5-.67 1.5-1.5 1.5zm1.5-6H6V6h12v5z"/>
            </svg>
          </div>
        </div>

        <!-- Destination -->
        <div class="flex flex-col text-right">
          <span class="text-[40px] font-black text-white leading-none tracking-tighter">{{ getCityCode(busDetails.destination) }}</span>
          <span class="text-xs text-slate-400 font-medium mt-1 truncate max-w-[100px]">{{ busDetails.destination }}</span>
          <span class="text-[#5eafff] font-bold mt-2">{{ busDetails.arrivalTime }}</span>
        </div>
      </div>
    </div>

    <!-- Sleek Data Grid -->
    <div class="px-8 py-5 border-y border-slate-700/50 bg-slate-900/20 backdrop-blur-sm grid grid-cols-3 gap-4 relative z-10">
      <div>
        <p class="text-[9px] text-slate-500 uppercase tracking-widest font-bold mb-1">Date</p>
        <p class="text-sm font-semibold text-slate-200">{{ busDetails.date }}</p>
      </div>
      <div>
        <p class="text-[9px] text-slate-500 uppercase tracking-widest font-bold mb-1">Bus Plate</p>
        <p class="text-sm font-semibold text-slate-200">{{ busDetails.plate }}</p>
      </div>
      <div class="text-right">
        <p class="text-[9px] text-slate-500 uppercase tracking-widest font-bold mb-1">Operator</p>
        <p class="text-sm font-semibold text-slate-200 truncate">{{ busDetails.company.split(' ')[0] }}</p>
      </div>
    </div>

    <!-- Perforated Cutout (Modern Style) -->
    <div class="relative flex items-center h-2 opacity-90 my-2 z-10">
      <div class="absolute -left-3 w-6 h-6 bg-[#0f172a] rounded-full border-r border-slate-700/60 shadow-inner"></div>
      <div class="w-full border-t-2 border-dashed border-slate-700/70"></div>
      <div class="absolute -right-3 w-6 h-6 bg-[#0f172a] rounded-full border-l border-slate-700/60 shadow-inner"></div>
    </div>

    <!-- Bottom Section: Seats, Price, Barcode -->
    <div class="px-8 pt-5 pb-8 flex-grow flex flex-col relative z-10">
      
      <div class="flex flex-col gap-8 items-start mb-8">
        <!-- Seats Selection -->
        <div class="flex-1">
          <p class="text-[9px] text-slate-500 uppercase tracking-widest font-bold mb-2">
            Selected Seats ({{ selectedSeats.length }})
          </p>
          <div class="flex flex-wrap gap-2 min-h-[2rem]">
            <span v-if="selectedSeats.length === 0" class="text-slate-600 italic text-sm font-medium">Select seats to continue</span>
            <span 
              v-else 
              v-for="seat in selectedSeats" 
              :key="seat.id"
              class="px-3 py-1 bg-[#5eafff]/10 border border-[#5eafff]/30 text-[#5eafff] text-sm font-bold rounded-lg shadow-sm"
            >
              {{ seat.id }}
            </span>
          </div>
        </div>

        <!-- Total Price Element -->
        <div class="text-left  border-slate-700/50">
          <p class="text-[9px] text-slate-500 uppercase tracking-widest font-bold mb-1">Total</p>
          <p class="text-3xl font-black text-white tracking-tight leading-none">
            <span class="text-lg text-slate-400 font-bold">LKR</span>  {{ totalPrice.toFixed(2) }}
          </p>
        </div>
      </div>

     

      <!-- Premium CTA Button -->
      <button 
        :disabled="selectedSeats.length === 0"
        :class="[
          'w-full py-4 rounded-2xl font-bold text-sm transition-all duration-300 tracking-widest uppercase flex items-center justify-center gap-3',
          selectedSeats.length > 0 
            ? 'bg-acc text-[#0f172a] shadow-[0_0_30px_rgba(255,255,255,0.15)] hover:shadow-[0_0_40px_rgba(255,255,255,0.25)] hover:scale-[1.02] active:scale-[0.98]' 
            : 'bg-slate-800/50 text-slate-500 cursor-not-allowed border border-slate-700'
        ]"
      >
        <span>Confirm & Pay</span>
        <svg v-if="selectedSeats.length > 0" class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M14 5l7 7m0 0l-7 7m7-7H3"></path></svg>
      </button>
      
    </div>
  </div>
</template>