<script setup lang="ts">
import { ref } from 'vue'
import { Send, MapPin, Calendar, Search, ArrowUpDown } from 'lucide-vue-next'

const form = ref({
  from: '',
  to: '',
  date: ''
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
  <!-- Mobile: Rounded rect box | Desktop: Full pill -->
  <div class="w-full mx-auto p-2 bg-white/[0.03] backdrop-blur-xl rounded-[2rem] lg:rounded-full border border-white/[0.08] shadow-[0_8px_30px_rgb(0,0,0,0.4)]">
    
    <form @submit.prevent="handleSearch" class="flex flex-col lg:flex-row items-center w-full">
      
      <!-- LOCATION GROUP (From & To) -->
      <div class="relative flex flex-col lg:flex-row w-full lg:w-auto lg:flex-[2]">
        
        <!-- Departure -->
        <div class="flex-1 w-full flex items-center px-6 py-4 lg:py-3 hover:bg-white/[0.04] rounded-t-[1.5rem] lg:rounded-full transition-all group cursor-text">
          <Send class="h-5 w-5 text-slate-500 mr-4 -rotate-45 shrink-0 group-focus-within:text-[#5fa8ff] transition-colors" />
          <div class="flex flex-col w-full text-left">
            <label class="text-[10px] font-bold uppercase tracking-[0.1em] text-slate-400 mb-0.5">Leaving from</label>
            <input 
              v-model="form.from" 
              type="text" 
              placeholder="City or Station" 
              class="bg-transparent border-none p-0 focus:ring-0 text-white font-semibold placeholder:text-white outline-none w-full text-lg" 
              required 
            />
          </div>
        </div>

        <!-- Mobile Horizontal Divider -->
        <div class="w-[calc(100%-3rem)] h-px bg-white/[0.08] mx-auto block lg:hidden"></div>
        <!-- Desktop Vertical Divider -->
        <div class="hidden lg:block w-px h-12 bg-white/[0.08] mx-2 self-center"></div>

        <!-- The Swap Button -->
        <!-- Mobile: Floats right between From/To | Desktop: Centers on vertical divider -->
        <div class="absolute right-6 top-1/2 -translate-y-1/2 lg:left-1/2 lg:top-1/2 lg:-translate-x-1/2 lg:-translate-y-1/2 z-10 flex items-center justify-center">
          <button 
            type="button"
            @click="swapLocations"
            class="w-9 h-9 flex items-center justify-center bg-[#0a1128] border border-white/[0.08] rounded-full shadow-lg text-slate-400 hover:text-[#5fa8ff] hover:border-[#5fa8ff]/50 hover:rotate-180 transition-all duration-300"
          >
            <!-- Up/Down icon looks better on mobile, but rotate handles desktop layout -->
            <ArrowUpDown class="h-4 w-4 lg:-rotate-90" />
          </button>
        </div>

        <!-- Arrival -->
        <div class="flex-1 w-full flex items-center px-6 py-4 lg:py-3 hover:bg-white/[0.04] lg:rounded-full transition-all group cursor-text">
          <MapPin class="h-5 w-5 text-slate-500 mr-4 shrink-0 group-focus-within:text-[#5fa8ff] transition-colors" />
          <div class="flex flex-col w-full text-left">
            <label class="text-[10px] font-bold uppercase tracking-[0.1em] text-slate-400 mb-0.5">Going to</label>
            <input 
              v-model="form.to" 
              type="text" 
              placeholder="Destination" 
              class="bg-transparent border-none p-0 focus:ring-0 text-white font-semibold placeholder:text-white outline-none w-full text-lg" 
              required 
            />
          </div>
        </div>
      </div>

      <!-- Mobile Horizontal Divider -->
      <div class="w-[calc(100%-3rem)] h-px bg-white/[0.08] mx-auto block lg:hidden"></div>
      <!-- Desktop Vertical Divider -->
      <div class="hidden lg:block w-px h-12 bg-white/[0.08] mx-2"></div>

      <!-- Date Picker -->
      <div class="flex-1 w-full flex items-center px-6 py-4 lg:py-3 hover:bg-white/[0.04] rounded-b-[1.5rem] lg:rounded-full transition-all group cursor-pointer">
        <Calendar class="h-5 w-5 text-slate-500 mr-4 shrink-0 group-hover:text-[#5fa8ff] transition-colors" />
        <div class="flex flex-col w-full text-left">
          <label class="text-[10px] font-bold uppercase tracking-[0.1em] text-slate-400 mb-0.5">Date of travel</label>
          <input 
            v-model="form.date" 
            type="date" 
            class="native-date-input bg-transparent border-none p-0 focus:ring-0 text-white font-semibold text-lg placeholder:text-slate-600 outline-none w-full cursor-pointer" 
            required
          />
        </div>
      </div>

      <!-- Search Button -->
      <button 
        type="submit" 
        class="w-full lg:w-auto mt-2 lg:mt-0 lg:ml-2 px-8 py-4 rounded-[1.5rem] lg:rounded-full flex items-center justify-center gap-2 bg-[#5fa8ff] hover:bg-[#4a90e2] text-[#0a1128] border-none transition-all shadow-lg shrink-0"
      >
        <Search class="h-5 w-5 font-bold text-[#0a1128]" />
        <span class="font-bold text-lg">Search</span>
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