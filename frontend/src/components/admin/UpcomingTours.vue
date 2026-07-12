<script setup lang="ts">
import { ref, computed } from 'vue';
import { Search, Filter } from 'lucide-vue-next';
import TourCard from './TourCard.vue'; 

// --------------------------------------------------------
// MOCK DATA GENERATOR
// --------------------------------------------------------
const generateMockTours = () => {
  const locations = ['Colombo', 'Kandy', 'Galle', 'Matara', 'Jaffna', 'Trincomalee', 'Anuradhapura', 'Kurunegala'];
  const companies = ['ExpressLine NCG', 'Lanka Super', 'Southern Transit', 'Northern Star', 'Highway Express'];
  const driverNames = ['Saman Kumara', 'Nimal Perera', 'Ruwan Silva', 'Sunil Jayasinghe', 'Kamal Fernando', 'Jagath Kumara'];
  
  const data = [];
  for (let i = 1; i <= 45; i++) {
    const day = (i % 28) + 1;
    const dateStr = `2026-08-${day.toString().padStart(2, '0')}`;
    
    data.push({
      id: `TR-260${i.toString().padStart(3, '0')}`,
      start: locations[i % locations.length],
      destination: locations[(i + 3) % locations.length], 
      date: dateStr,
      time: ['06:00 AM', '08:30 AM', '11:15 AM', '02:00 PM', '05:30 PM', '09:45 PM'][i % 6],
      busCompany: companies[i % companies.length],
      driverName: driverNames[i % driverNames.length], // Replaced driverId with Driver Name
      driverPhone: `+94 77 ${Math.floor(100 + Math.random() * 900)} ${Math.floor(1000 + Math.random() * 9000)}`
    });
  }
  
  // Sort by date ascending to make it realistic for "Upcoming" tours
  return data.sort((a, b) => new Date(a.date).getTime() - new Date(b.date).getTime());
};

const tours = ref(generateMockTours());

// --------------------------------------------------------
// STATE & LOGIC
// --------------------------------------------------------
const searchQuery = ref('');
const operatorFilter = ref('All');
const currentPage = ref(1);
const itemsPerPage = 20;

// --------------------------------------------------------
// SEARCH & FILTER COMPUTATIONS
// --------------------------------------------------------
const filteredTours = computed(() => {
  return tours.value.filter(tour => {
    // 1. Operator Filter
    if (operatorFilter.value !== 'All' && tour.busCompany !== operatorFilter.value) {
      return false;
    }
    
    // 2. Search Query (ID, Route, or Driver Name)
    const q = searchQuery.value.toLowerCase();
    if (q) {
      return tour.id.toLowerCase().includes(q) || 
             tour.start.toLowerCase().includes(q) || 
             tour.destination.toLowerCase().includes(q) ||
             tour.driverName.toLowerCase().includes(q); // Now searching by name
    }
    
    return true;
  });
});

// Pagination Computations
const totalPages = computed(() => Math.ceil(filteredTours.value.length / itemsPerPage) || 1);

const paginatedTours = computed(() => {
  const start = (currentPage.value - 1) * itemsPerPage;
  const end = start + itemsPerPage;
  return filteredTours.value.slice(start, end);
});

// Reset pagination when search/filter changes
const onSearchOrFilterChange = () => {
  currentPage.value = 1;
};

// Unique operators list for the dropdown filter
const uniqueOperators = computed(() => {
  const operators = new Set(tours.value.map(t => t.busCompany));
  return Array.from(operators).sort();
});
</script>

<template>
  <div class="w-full flex flex-col gap-6">
    
    <!-- TOP BAR: Search & Filters -->
    <div class="w-full bg-[#151b29] p-5 rounded-2xl border border-slate-800/80 shadow-lg flex flex-col md:flex-row items-center gap-4 justify-between">
      
      <!-- Search Bar -->
      <div class="relative w-full md:w-1/2 lg:w-1/3">
        <Search class="absolute left-4 top-1/2 -translate-y-1/2 w-5 h-5 text-slate-500" />
        <input 
          v-model="searchQuery" 
          @input="onSearchOrFilterChange"
          type="text" 
          placeholder="Search by ID, Route, or Driver Name..." 
          class="w-full bg-[#0a0d14] text-white pl-12 pr-4 py-3 rounded-xl border border-slate-700/50 focus:border-[#5eafff] focus:ring-1 focus:ring-[#5eafff] outline-none transition-all placeholder:text-slate-600 shadow-inner"
        >
      </div>

      <!-- Filters -->
      <div class="flex items-center gap-3 w-full md:w-auto">
        <div class="flex items-center gap-2 bg-[#0a0d14] px-4 py-3 rounded-xl border border-slate-700/50">
          <Filter class="w-4 h-4 text-slate-500" />
          <select 
            v-model="operatorFilter" 
            @change="onSearchOrFilterChange"
            class="bg-transparent text-slate-300 text-sm font-bold outline-none cursor-pointer max-w-[200px] truncate"
          >
            <option value="All" class="bg-[#151b29]">All Operators</option>
            <option v-for="operator in uniqueOperators" :key="operator" :value="operator" class="bg-[#151b29]">
              {{ operator }}
            </option>
          </select>
        </div>
      </div>

    </div>

    <!-- TOURS LIST -->
    <div class="w-full flex flex-col gap-3 min-h-[500px]">
      
      <!-- Empty State -->
      <div v-if="paginatedTours.length === 0" class="w-full py-16 flex flex-col items-center justify-center text-slate-500 bg-[#151b29]/50 rounded-2xl border border-slate-800/50 border-dashed">
        <Search class="w-10 h-10 mb-4 opacity-50" />
        <p class="font-bold tracking-widest uppercase text-sm">No tours match your search</p>
      </div>

      <!-- Render Flat Tour Cards -->
      <TourCard 
        v-for="tour in paginatedTours" 
        :key="tour.id"
        :tour="tour"
      />

    </div>

    <!-- PAGINATION CONTROLS -->
    <div v-if="totalPages > 1" class="w-full flex items-center justify-between bg-[#151b29] p-4 rounded-2xl border border-slate-800/80 shadow-lg mt-2">
      <button 
        @click="currentPage--" 
        :disabled="currentPage === 1"
        class="px-4 py-2 bg-[#0a0d14] border border-slate-700/50 rounded-lg text-slate-300 text-xs font-bold uppercase tracking-widest hover:bg-[#1e2638] disabled:opacity-50 disabled:cursor-not-allowed transition-colors"
      >
        Previous
      </button>

      <span class="text-slate-400 text-sm font-bold">
        Page <span class="text-white">{{ currentPage }}</span> of {{ totalPages }}
      </span>

      <button 
        @click="currentPage++" 
        :disabled="currentPage === totalPages"
        class="px-4 py-2 bg-[#0a0d14] border border-slate-700/50 rounded-lg text-slate-300 text-xs font-bold uppercase tracking-widest hover:bg-[#1e2638] disabled:opacity-50 disabled:cursor-not-allowed transition-colors"
      >
        Next
      </button>
    </div>

  </div>
</template>