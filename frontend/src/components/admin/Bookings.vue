<script setup lang="ts">
import { ref, computed, onMounted } from 'vue';
import { Search, Filter } from 'lucide-vue-next';
import BookingCard from './BookingCard.vue';
import { API_BASE } from '../../api';

interface Booking {
  id: string;
  tourId: string;
  start: string;
  destination: string;
  date: string;
  time: string;
  name: string;
  status: string;
  seats: string[];
  totalPrice: number | string;
  email: string;
  phone: string;
  address: string;
  busCompany: string;
  busId: string;
  busModel: string;
  driverId: string;
  assistantId: string;
  nic: string;
}

const bookings = ref<Booking[]>([]);

const loadBookings = async () => {
  try {
    const response = await fetch(`${API_BASE}/bookings`);
    const data = await response.json();
    if (!Array.isArray(data)) {
      bookings.value = [];
      return;
    }
    bookings.value = data.map((booking: any) => ({
      id: String(booking.id),
      tourId: String(booking.tour_id || booking.tourId || ''),
      start: String(booking.start || ''),
      destination: String(booking.destination || ''),
      date: String(booking.date || ''),
      time: String(booking.time || ''),
      name: String(booking.name || ''),
      status: String(booking.status || 'Active'),
      seats: Array.isArray(booking.seats) ? booking.seats : [],
      totalPrice: booking.total_price ?? booking.totalPrice ?? 0,
      email: String(booking.email || ''),
      phone: String(booking.phone || ''),
      address: String(booking.address || ''),
      busCompany: String(booking.bus_company || booking.busCompany || ''),
      busId: String(booking.bus_id || booking.busId || ''),
      busModel: String(booking.bus_model || booking.busModel || ''),
      driverId: String(booking.driver_id || booking.driverId || ''),
      assistantId: String(booking.assistant_id || booking.assistantId || ''),
      nic: String(booking.nic || ''),
    }));
  } catch (error) {
    console.error('Failed to load bookings:', error);
    bookings.value = [];
  }
};

const searchQuery = ref('');
const statusFilter = ref('All');
const currentPage = ref(1);
const itemsPerPage = 10;

const handleCancelBooking = async (id: string) => {
  const booking = bookings.value.find((booking) => booking.id === id);
  if (!booking) return;

  if (!confirm(`Are you sure you want to cancel booking ${id}?`)) {
    return;
  }

  try {
    const response = await fetch(`${API_BASE}/bookings/update`, {
      method: 'PUT',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({
        ...booking,
        tour_id: Number(booking.tourId),
        id: Number(booking.id),
        status: 'Cancelled',
      }),
    });

    if (!response.ok) {
      throw new Error('Failed to update booking');
    }

    booking.status = 'Cancelled';
  } catch (error) {
    console.error('Cancellation failed:', error);
  }
};

onMounted(loadBookings);

const filteredBookings = computed(() => {
  return bookings.value.filter(booking => {
    if (statusFilter.value !== 'All' && booking.status !== statusFilter.value) return false;

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

const totalPages = computed(() => Math.ceil(filteredBookings.value.length / itemsPerPage) || 1);

const paginatedBookings = computed(() => {
  const start = (currentPage.value - 1) * itemsPerPage;
  const end = start + itemsPerPage;
  return filteredBookings.value.slice(start, end);
});

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