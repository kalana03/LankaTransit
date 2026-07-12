<script setup lang="ts">
import { ref, computed } from 'vue';
import { Search, Filter } from 'lucide-vue-next';
import BookingCard from './BookingCard.vue';

// --------------------------------------------------------
// MOCK DATA
// --------------------------------------------------------
const generateMockBookings = () => {
  const locations = ['Colombo', 'Kandy', 'Galle', 'Matara', 'Jaffna', 'Trincomalee'];
  const data = [];
  // Increased loop to 45 to better demonstrate pagination functionality
  for (let i = 1; i <= 45; i++) {
    const isCancelled = i % 7 === 0;
    const numSeats = Math.floor(Math.random() * 3) + 1;
    
    data.push({
      id: `LTX-${8000 + i}`,
      tourId: `TRP-24${i.toString().padStart(3, '0')}`, 
      name: `Passenger ${i}`,
      nic: `199${i % 9}123456${i % 2 === 0 ? 'V' : 'X'}`,
      start: locations[i % locations.length],
      destination: locations[(i + 2) % locations.length],
      date: `2026-07-${10 + (i % 15)}`,
      time: '08:30 AM',
      seats: Array.from({ length: numSeats }, (_, idx) => `A${idx + 1}`),
      status: isCancelled ? 'Cancelled' : 'Active',
      totalPrice: (2450 * numSeats).toFixed(2),
      email: `passenger${i}@example.com`,
      phone: `+94 77 123 45${i.toString().padStart(2, '0')}`,
      address: `No. ${i}, High Level Road, ${locations[i % locations.length]}`, 
      busCompany: 'LankaTransit Express',
      busId: `WP-ND-${1000 + i}`,
      busModel: 'Volvo B11R',
      driverId: `DRV-${100 + i}`,
      assistantId: `AST-${500 + i}`
    });
  }
  return data;
};

const bookings = ref(generateMockBookings());

// --------------------------------------------------------
// STATE & LOGIC
// --------------------------------------------------------
const searchQuery = ref('');
const statusFilter = ref('All');
const currentPage = ref(1);
const itemsPerPage = 10; // Max items per page

// Handle Cancellation emitted from the Modal
const handleCancelBooking = (id: string) => {
  if (confirm(`Are you sure you want to cancel booking ${id}?`)) {
    const index = bookings.value.findIndex(b => b.id === id);
    if (index !== -1) bookings.value[index].status = 'Cancelled';
  }
};

// --------------------------------------------------------
// SEARCH, FILTER, & PAGINATION COMPUTATIONS
// --------------------------------------------------------
const filteredBookings = computed(() => {
  return bookings.value.filter(booking => {
    // Status Filter
    if (statusFilter.value !== 'All' && booking.status !== statusFilter.value) return false;
    
    // Search Query (ID, Tour ID, Name, or NIC)
    const q = searchQuery.value.toLowerCase();
    if (q) {
      return booking.id.toLowerCase().includes(q) || 
             booking.tourId.toLowerCase().includes(q) || 
             booking.name.toLowerCase().includes(q) || 
             booking.nic.toLowerCase().includes(q);
    }
    
    return true;
  });
});

// Pagination Math
const totalPages = computed(() => Math.ceil(filteredBookings.value.length / itemsPerPage) || 1);

const paginatedBookings = computed(() => {
  const start = (currentPage.value - 1) * itemsPerPage;
  const end = start + itemsPerPage;
  return filteredBookings.value.slice(start, end);
});

// Reset pagination when search or filter values change
const resetPagination = () => {
  currentPage.value = 1;
};
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
          @input="resetPagination"
          type="text" 
          placeholder="Search by Booking ID, Tour ID, Name..." 
          class="w-full bg-[#0a0d14] text-white pl-12 pr-4 py-3 rounded-xl border border-slate-700/50 focus:border-[#5eafff] focus:ring-1 focus:ring-[#5eafff] outline-none transition-all placeholder:text-slate-600 shadow-inner"
        >
      </div>

      <!-- Filters -->
      <div class="flex items-center gap-3 w-full md:w-auto">
        <div class="flex items-center gap-2 bg-[#0a0d14] px-4 py-3 rounded-xl border border-slate-700/50">
          <Filter class="w-4 h-4 text-slate-500" />
          <select 
            v-model="statusFilter"
            @change="resetPagination"
            class="bg-transparent text-slate-300 text-sm font-bold outline-none cursor-pointer"
          >
            <option value="All" class="bg-[#151b29]">All Statuses</option>
            <option value="Active" class="bg-[#151b29]">Active</option>
            <option value="Cancelled" class="bg-[#151b29]">Cancelled</option>
          </select>
        </div>
      </div>

    </div>

    <!-- BOOKINGS LIST -->
    <div class="w-full flex flex-col gap-3 min-h-[500px]">
      
      <div v-if="paginatedBookings.length === 0" class="w-full py-16 flex flex-col items-center justify-center text-slate-500 bg-[#151b29]/50 rounded-2xl border border-slate-800/50 border-dashed">
        <Search class="w-10 h-10 mb-4 opacity-50" />
        <p class="font-bold tracking-widest uppercase text-sm">No bookings found</p>
      </div>

      <!-- Render Paginated Booking Cards -->
      <BookingCard 
        v-for="booking in paginatedBookings" 
        :key="booking.id"
        :booking="booking"
        @cancel="handleCancelBooking"
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