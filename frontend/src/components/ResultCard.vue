<script setup lang="ts">
import { BusFront, Map, Armchair, ArrowRight, Clock } from 'lucide-vue-next'

interface BusResult {
  company: string
  routeNumber: string
  date: string
  startTime: string
  endTime: string
  startLocation: string
  endLocation: string
  duration: string
  seatsLeft: number
  price: string
}

withDefaults(defineProps<{ result?: BusResult }>(), {
  result: () => ({
    company: 'Superline Express',
    routeNumber: 'EX-17',
    date: 'Oct 24, 2023',
    startTime: '08:30 AM',
    endTime: '11:45 AM',
    startLocation: 'Colombo Central Bus Stand',
    endLocation: 'Kandy Goods Shed',
    duration: '3h 15m',
    seatsLeft: 4,
    price: '1,200'
  })
})
</script>

<template>
  <!-- 
    Card Container 
    Adjusted padding and border radius for mobile (p-4 rounded-[1.5rem] -> p-6 rounded-[2rem])
  -->
  <div class="w-full max-w-4xl mx-auto p-5 sm:p-6 lg:p-8 bg-white/[0.02] backdrop-blur-md rounded-[1.5rem] sm:rounded-[2rem] border border-white/[0.08] shadow-lg hover:border-[#5eafff]/30 hover:bg-white/[0.04] sm:hover:-translate-y-1 transition-all duration-300 font-sans group">
    
    <!-- TOP ROW: Bus Info & Date -->
    <div class="flex items-start sm:items-center justify-between gap-4 mb-6 sm:mb-8">
      
      <!-- Company & Route -->
      <div class="flex items-center gap-3">
        <div class="w-10 h-10 sm:w-12 sm:h-12 rounded-full bg-[#0f172a] border border-white/[0.1] flex items-center justify-center text-[#5eafff] shadow-inner shrink-0">
          <BusFront class="w-5 h-5 sm:w-6 sm:h-6" />
        </div>
        <div class="flex flex-col">
          <h3 class="text-white font-bold text-base sm:text-lg tracking-tight leading-tight sm:leading-none mb-1">
            {{ result.company }}
          </h3>
          <div class="flex items-center gap-2">
            <span class="inline-block px-2 py-0.5 bg-white/[0.05] text-[#5eafff] text-[9px] sm:text-[10px] font-bold uppercase tracking-widest rounded-md border border-[#5eafff]/20">
              Route {{ result.routeNumber }}
            </span>
          </div>
        </div>
      </div>

      <!-- Date Badge -->
      <div class="flex items-center gap-1.5 text-slate-400 text-xs sm:text-sm font-medium px-3 py-1.5 sm:px-4 sm:py-2 bg-[#0f172a]/50 rounded-full border border-white/[0.05] shrink-0">
        <Clock class="w-3.5 h-3.5 sm:w-4 sm:h-4" />
        <span class="hidden sm:inline">{{ result.date }}</span>
        <!-- Shortened date for very small screens -->
        <span class="sm:hidden">{{ result.date.split(',')[0] }}</span> 
      </div>
    </div>

    <!-- ========================================== -->
    <!-- DESKTOP TIMELINE (Hidden on mobile)        -->
    <!-- ========================================== -->
    <div class="hidden sm:flex flex-row items-center justify-between gap-2 mb-8 relative">
      <!-- Start Point -->
      <div class="flex flex-col items-start w-1/4 text-left z-10">
        <p class="text-2xl lg:text-3xl font-bold text-white tracking-tight mb-1">{{ result.startTime }}</p>
        <p class="text-slate-400 font-medium text-sm lg:text-base truncate w-full" :title="result.startLocation">{{ result.startLocation }}</p>
      </div>

      <!-- Timeline Graphic & Map Link -->
      <div class="flex-1 flex flex-col items-center w-full px-4 relative z-0">
        <p class="text-xs font-bold text-slate-500 uppercase tracking-widest mb-2 bg-[#0f172a] px-2 relative z-10 rounded-full">
          {{ result.duration }}
        </p>
        
        <div class="w-full flex items-center relative">
          <div class="w-2.5 h-2.5 rounded-full border-2 border-[#5eafff] bg-[#0f172a] z-10"></div>
          <div class="flex-1 border-t-2 border-dashed border-white/[0.1] group-hover:border-[#5eafff]/30 transition-colors"></div>
          <ArrowRight class="w-4 h-4 text-slate-600 absolute left-1/2 -translate-x-1/2 bg-[#0f172a] px-1 rounded-full" />
          <div class="w-2.5 h-2.5 rounded-full border-2 border-[#5eafff] bg-[#0f172a] z-10"></div>
        </div>

        <button class="mt-3 flex items-center gap-1.5 text-xs font-bold text-[#5eafff] hover:text-[#4a90e2] hover:underline transition-all">
          <Map class="w-3.5 h-3.5" /> View route on map
        </button>
      </div>

      <!-- End Point -->
      <div class="flex flex-col items-end w-1/4 text-right z-10">
        <p class="text-2xl lg:text-3xl font-bold text-white tracking-tight mb-1">{{ result.endTime }}</p>
        <p class="text-slate-400 font-medium text-sm lg:text-base truncate w-full" :title="result.endLocation">{{ result.endLocation }}</p>
      </div>
    </div>

    <!-- ========================================== -->
    <!-- MOBILE TIMELINE (Hidden on desktop)        -->
    <!-- ========================================== -->
    <div class="flex sm:hidden w-full mb-6">
      
      <!-- Graphic Left Column -->
      <div class="flex flex-col items-center w-6 mr-4 mt-1.5 mb-1.5 shrink-0">
        <div class="w-3 h-3 rounded-full border-[2.5px] border-[#5eafff] bg-[#0f172a] z-10 shadow-[0_0_8px_rgba(94,175,255,0.4)]"></div>
        <div class="w-px flex-1 border-l-2 border-dashed border-white/[0.15] my-1 group-hover:border-[#5eafff]/30 transition-colors"></div>
        <div class="w-3 h-3 rounded-full border-[2.5px] border-[#5eafff] bg-[#0f172a] z-10 shadow-[0_0_8px_rgba(94,175,255,0.4)]"></div>
      </div>

      <!-- Content Right Column -->
      <div class="flex-1 flex flex-col justify-between">
        
        <!-- Start -->
        <div class="mb-5">
          <p class="text-xl font-bold text-white leading-none tracking-tight mb-1">{{ result.startTime }}</p>
          <p class="text-slate-400 text-sm leading-tight">{{ result.startLocation }}</p>
        </div>

        <!-- Duration & Map Row (Middle) -->
        <div class="flex items-center gap-3 mb-5">
          <span class="text-[10px] font-bold text-slate-500 uppercase tracking-widest bg-[#0f172a]/50 px-2 py-0.5 rounded-md border border-white/[0.05]">
            {{ result.duration }}
          </span>
          <span class="w-1 h-1 bg-slate-700 rounded-full"></span>
          <button class="flex items-center gap-1.5 text-xs font-bold text-[#5eafff] hover:underline active:opacity-70">
            <Map class="w-3 h-3" /> Map
          </button>
        </div>

        <!-- End -->
        <div>
          <p class="text-xl font-bold text-white leading-none tracking-tight mb-1">{{ result.endTime }}</p>
          <p class="text-slate-400 text-sm leading-tight">{{ result.endLocation }}</p>
        </div>

      </div>
    </div>

    <!-- ========================================== -->
    <!-- BOTTOM ROW: Footer Actions                 -->
    <!-- ========================================== -->
    <div class="flex flex-col sm:flex-row items-center justify-between pt-5 sm:pt-6 border-t border-white/[0.08] gap-5 sm:gap-4">
      
      <!-- Seats Left & Mobile Price Group -->
      <div class="flex items-center justify-between w-full sm:w-auto">
        <!-- Seats Indicator -->
        <div class="flex items-center gap-3">
          <div 
            class="flex items-center justify-center w-8 h-8 sm:w-9 sm:h-9 rounded-full shrink-0"
            :class="result.seatsLeft <= 5 ? 'bg-amber-500/10 text-amber-500' : 'bg-emerald-500/10 text-emerald-500'"
          >
            <Armchair class="w-4 h-4 sm:w-4 sm:h-4" />
          </div>
          <div>
            <p 
              class="text-sm sm:text-base font-bold leading-none mb-1"
              :class="result.seatsLeft <= 5 ? 'text-amber-500' : 'text-emerald-500'"
            >
              {{ result.seatsLeft }} seats left
            </p>
            <p class="text-[10px] sm:text-xs font-medium text-slate-300 leading-none">Available to book</p>
          </div>
        </div>

        <!-- Price (Visible here ONLY on mobile to save vertical space) -->
        <div class="text-right sm:hidden">
          <p class="text-[9px] font-bold uppercase tracking-widest text-slate-500 mb-0.5">Per seat</p>
          <p class="text-lg font-bold text-white leading-none">LKR {{ result.price }}</p>
        </div>
      </div>

      <!-- Desktop Price & Global Book Button -->
      <div class="flex items-center justify-end w-full sm:w-auto gap-6">
        
        <!-- Price (Visible here ONLY on Desktop) -->
        <div class="text-right hidden sm:block">
          <p class="text-[10px] font-bold uppercase tracking-widest text-slate-500 mb-1">Per seat</p>
          <p class="text-xl lg:text-2xl font-bold text-white leading-none">LKR {{ result.price }}</p>
        </div>
        
        <!-- Action Button (Full width on mobile, auto on desktop) -->
        <button class="w-full sm:w-auto px-6 py-3.5 sm:py-3 bg-[#5eafff] hover:bg-[#4a90e2] text-[#0f172a] font-bold rounded-full transition-all shadow-[0_0_15px_rgba(94,175,255,0.2)] hover:shadow-[0_0_20px_rgba(94,175,255,0.4)] active:scale-[0.98] flex items-center justify-center gap-2 text-base sm:text-md">
          Book Seats
          <ArrowRight class="w-4 h-4 sm:w-4 sm:h-4" />
        </button>

      </div>

    </div>
  </div>
</template>