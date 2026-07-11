<script setup lang="ts">
import { ref } from 'vue';

const emit = defineEmits(['update:selection']);

type SeatStatus = 'available' | 'selected' | 'occupied' | 'temp-unavailable';
interface Seat {
  id: string;
  row: string;
  number: number;
  status: SeatStatus;
}

const rows = ['A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I'];
const seatsPerRow = 4;

const generateSeats = () => {
  const layout: Seat[][] = [];
  rows.forEach((rowLabel) => {
    const row: Seat[] = [];
    for (let i = 1; i <= seatsPerRow; i++) {
      const random = Math.random();
      let status: SeatStatus = 'available';
      if (random < 0.2) status = 'occupied';
      else if (random < 0.3) status = 'temp-unavailable';
      row.push({
        id: `${rowLabel}${i}`,
        row: rowLabel,
        number: i,
        status
      });
    }
    layout.push(row);
  });
  return layout;
};

const seatingLayout = ref<Seat[][]>(generateSeats());

const toggleSeat = (seat: Seat) => {
  if (seat.status === 'occupied' || seat.status === 'temp-unavailable') return;
  seat.status = seat.status === 'selected' ? 'available' : 'selected';
  const selectedSeats = seatingLayout.value.flat().filter(s => s.status === 'selected');
  emit('update:selection', selectedSeats);
};
</script>

<template>
  <!-- Premium Dark Shell -->
  <div class="w-full p-8 lg:p-12 bg-gradient-to-b from-[#162032] to-[#0f172a] rounded-[2.5rem] font-sans border border-white/5 shadow-2xl relative">
    
    <!-- Header & Legend -->
    <div class="flex flex-col xl:flex-row justify-between items-start xl:items-center gap-6 mb-12 relative z-20">
      <div>
        <h3 class="text-white font-black text-[28px] tracking-tight">Deck Plan</h3>
        <p class="text-slate-400 text-sm mt-1">Select your preferred seating</p>
      </div>

      <!-- High-End Pill Legend -->
      <div class="flex flex-wrap items-center gap-1 bg-transparent border border-slate-700/50 rounded-full px-2 py-1 backdrop-blur-sm">
        <!-- Available -->
        <div class="flex items-center gap-2 px-3 py-1.5 rounded-full">
          <div class="w-3 h-3 rounded-full border-2 border-slate-500"></div>
          <span class="text-slate-300 uppercase tracking-widest text-[10px] font-bold">Avail</span>
        </div>
        <!-- Selected -->
        <div class="flex items-center gap-2 px-4 py-1.5 bg-[#1d3354] rounded-full border border-[#2a4d7a]">
          <div class="w-3 h-3 rounded-full bg-[#5eafff] shadow-[0_0_8px_rgba(94,175,255,0.5)]"></div>
          <span class="text-[#5eafff] uppercase tracking-widest text-[10px] font-bold">Selected</span>
        </div>
        <!-- Taken -->
        <div class="flex items-center gap-2 px-3 py-1.5 rounded-full">
          <div class="w-3 h-3 rounded-full border-2 border-[#0a0d14] bg-[#1a2130]"></div>
          <span class="text-slate-500 uppercase tracking-widest text-[10px] font-bold">Taken</span>
        </div>
        <!-- Held -->
        <div class="flex items-center gap-2 px-3 py-1.5 rounded-full">
          <div class="w-3 h-3 rounded-full border-2 border-[#d97706]"></div>
          <span class="text-[#d97706] uppercase tracking-widest text-[10px] font-bold">Held</span>
        </div>
      </div>
    </div>

    <!-- Bus Cavity (Transparent layout, holding the grid and seats) -->
    <div class="w-full max-w-[800px] mx-auto bg-transparent rounded-[2.5rem] relative overflow-hidden py-4">
      
      <!-- Technical Grid Overlay (Faint dotted pattern masked to fade at edges) -->
      <div class="absolute inset-0 opacity-[0.15] pointer-events-none [mask-image:radial-gradient(ellipse_at_center,white,transparent_80%)]" 
           style="background-image: radial-gradient(#94a3b8 1px, transparent 1px); background-size: 24px 24px;">
      </div>

      <!-- Seating Grid -->
      <div class="w-full relative z-10 flex flex-col items-center pb-4">
        <div class="flex flex-col gap-5 sm:gap-6">
          
          <div v-for="(row, rowIndex) in seatingLayout" :key="rowIndex" class="flex items-center gap-6 sm:gap-8 group/row">
            
            <!-- Left Row Identifier -->
            <div class="w-4 text-right text-slate-600 font-bold text-xs transition-colors duration-300 group-hover/row:text-slate-400">{{ row[0].row }}</div>

            <!-- Seat Buttons -->
            <div class="flex items-center gap-4 sm:gap-5">
              <button
                v-for="(seat, seatIndex) in row"
                :key="seat.id"
                @click="toggleSeat(seat)"
                :disabled="seat.status === 'occupied' || seat.status === 'temp-unavailable'"
                :aria-label="`Seat ${seat.id}, ${seat.status}`"
                :aria-pressed="seat.status === 'selected'"
                :class="[
                  'w-14 h-14 sm:w-16 sm:h-16 rounded-[1rem] transition-all duration-300 ease-out flex items-center justify-center text-base font-bold select-none border relative overflow-hidden',
                  
                  // Aisle Spacer
                  seatIndex === 1 ? 'mr-12 sm:mr-20 lg:mr-24' : '',
                  
                  // AVAILABLE: Slate blue, off-white text, subtle lift on hover
                  seat.status === 'available' ? 'bg-[#1e2638] border-transparent text-white hover:bg-[#252f45] hover:-translate-y-0.5 hover:shadow-lg hover:border-slate-600/50 cursor-pointer' : '',
                  
                  // SELECTED: Darker base, blue border, blue text, subtle outer glow
                  seat.status === 'selected' ? 'bg-acc text-dark border-[#5eafff] text-[#5eafff] shadow-[0_0_12px_rgba(94,175,255,0.25)] scale-105 cursor-pointer z-10' : '',
                  
                  // OCCUPIED: Very dark, almost blends with bg, subtle X
                  seat.status === 'occupied' ? 'bg-red-500 border-[#d97706]/70 text-white cursor-not-allowed' : '',
                  
                  // TEMP UNAVAILABLE: Very dark, amber border, amber icon
                  seat.status === 'temp-unavailable' ? 'bg-[#1e2638] border-amber-500 text-amber-500 cursor-not-allowed ' : ''
                ]"
              >
                <!-- Seat Number -->
                <span v-if="seat.status === 'available' || seat.status === 'selected'" class="relative z-10">{{ seat.number }}</span>
                
                <!-- Occupied (X) -->
                <svg v-if="seat.status === 'occupied'" class="w-5 h-5 relative z-10" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12"></path></svg>
                
                <!-- Temp Held (Lock) -->
                <svg v-if="seat.status === 'temp-unavailable'" class="w-5 h-5 relative z-10" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M12 15v2m-6 4h12a2 2 0 002-2v-6a2 2 0 00-2-2H6a2 2 0 00-2 2v6a2 2 0 002 2zm10-10V7a4 4 0 00-8 0v4h8z"></path></svg>
              </button>
            </div>

            <!-- Right Row Identifier -->
            <div class="w-4 text-left text-slate-600 font-bold text-xs transition-colors duration-300 group-hover/row:text-slate-400">{{ row[0].row }}</div>
          
          </div>
        </div>
      </div>
    </div>
  </div>
</template>