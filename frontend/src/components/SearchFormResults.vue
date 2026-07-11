<script setup lang="ts">
import { ref } from 'vue'
import { Send, MapPin, Calendar, Search, ArrowUpDown } from 'lucide-vue-next'

const form = ref({
  from: 'Colombo',
  to: 'Kandy',
  date: '2023-10-24'
})

const emit = defineEmits(['search'])

const handleSearch = () => {
  emit('search', form.value)
}

const swapLocations = () => {
  const temp = form.value.from
  form.value.from = form.value.to
  form.value.to = temp
}
</script>

<template>
  <!-- 
    Sidebar Card Container 
    Designed to be constrained by a parent column (e.g., w-80 or w-1/3 on the results page)
  -->
  <div class="sticky w-full bg-white/[0.02] backdrop-blur-xl rounded-[2rem] border border-white/[0.08] shadow-[0_8px_30px_rgb(0,0,0,0.4)] p-5 sm:p-6 font-sans">
    
    <!-- Optional: Sidebar Header -->
    <div class="mb-6">
      <h2 class="text-white text-lg font-bold tracking-tight">Modify Search</h2>
      <p class="text-slate-400 text-xs font-medium mt-1">Update your route or date</p>
    </div>

    <form @submit.prevent="handleSearch" class="flex flex-col w-full">
      
      <!-- 
        Input Group Wrapper 
        Contains From, Swap, To, and Date with internal dividers
      -->
      <div class="flex flex-col w-full bg-white/[0.01] rounded-[1.5rem] border border-white/[0.05]">
        
        <!-- Departure -->
        <div class="w-full flex items-center px-5 py-4 hover:bg-white/[0.04] rounded-t-[1.5rem] transition-all group cursor-text">
          <Send class="h-5 w-5 text-slate-500 mr-4 -rotate-45 shrink-0 group-focus-within:text-[#5fa8ff] transition-colors" />
          <div class="flex flex-col w-full text-left">
            <label class="text-[10px] font-bold uppercase tracking-[0.1em] text-slate-400 mb-0.5">Leaving from</label>
            <input 
              v-model="form.from" 
              type="text" 
              placeholder="City or Station" 
              class="bg-transparent border-none p-0 focus:ring-0 text-white font-semibold placeholder:text-slate-600 outline-none w-full text-base" 
              required 
            />
          </div>
        </div>

        <!-- Divider with Embedded Swap Button -->
        <div class="relative w-full h-px bg-white/[0.08]">
          <button 
            type="button"
            @click="swapLocations"
            class="absolute right-5 top-1/2 -translate-y-1/2 z-10 w-8 h-8 flex items-center justify-center bg-[#0a1128] border border-white/[0.08] rounded-full shadow-lg text-slate-400 hover:text-[#5fa8ff] hover:border-[#5fa8ff]/50 hover:rotate-180 transition-all duration-300"
          >
            <ArrowUpDown class="h-4 w-4" />
          </button>
        </div>

        <!-- Arrival -->
        <div class="w-full flex items-center px-5 py-4 hover:bg-white/[0.04] transition-all group cursor-text">
          <MapPin class="h-5 w-5 text-slate-500 mr-4 shrink-0 group-focus-within:text-[#5fa8ff] transition-colors" />
          <div class="flex flex-col w-full text-left">
            <label class="text-[10px] font-bold uppercase tracking-[0.1em] text-slate-400 mb-0.5">Going to</label>
            <input 
              v-model="form.to" 
              type="text" 
              placeholder="Destination" 
              class="bg-transparent border-none p-0 focus:ring-0 text-white font-semibold placeholder:text-slate-600 outline-none w-full text-base" 
              required 
            />
          </div>
        </div>

        <!-- Divider -->
        <div class="w-full h-px bg-white/[0.08]"></div>

        <!-- Date Picker -->
        <div class="w-full flex items-center px-5 py-4 hover:bg-white/[0.04] rounded-b-[1.5rem] transition-all group cursor-pointer">
          <Calendar class="h-5 w-5 text-slate-500 mr-4 shrink-0 group-hover:text-[#5fa8ff] transition-colors" />
          <div class="flex flex-col w-full text-left relative">
            <label class="text-[10px] font-bold uppercase tracking-[0.1em] text-slate-400 mb-0.5">Date of travel</label>
            <input 
              v-model="form.date" 
              type="date" 
              class="native-date-input bg-transparent border-none p-0 focus:ring-0 text-white font-semibold text-base placeholder:text-slate-600 outline-none w-full cursor-pointer" 
              required
            />
          </div>
        </div>

      </div>

      <!-- Update Search Button -->
      <button 
        type="submit" 
        class="w-full mt-6 px-6 py-4 rounded-[1.25rem] flex items-center justify-center gap-2 bg-[#5fa8ff] hover:bg-[#4a90e2] text-[#0a1128] border-none transition-all shadow-[0_0_15px_rgba(95,168,255,0.2)] hover:shadow-[0_0_20px_rgba(95,168,255,0.4)] active:scale-95 shrink-0"
      >
        <Search class="h-5 w-5 font-bold" />
        <span class="font-bold text-base">Update Search</span>
      </button>

    </form>
  </div>
</template>

<style scoped>
.native-date-input {
  color-scheme: dark;
}
.native-date-input::-webkit-calendar-picker-indicator {
  display: none;
  -webkit-appearance: none;
}
.native-date-input::-webkit-datetime-edit-year-field,
.native-date-input::-webkit-datetime-edit-month-field,
.native-date-input::-webkit-datetime-edit-day-field,
.native-date-input::-webkit-datetime-edit-text {
  color: rgba(255, 255, 255, 0.9);
}
.native-date-input:invalid::-webkit-datetime-edit {
  color: #475569;
}
</style>