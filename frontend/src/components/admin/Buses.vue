<script setup lang="ts">
import { computed, ref, onMounted } from 'vue';
import { Search, BusFront } from 'lucide-vue-next';
import RegisteredBusCard from './BusCard.vue';
import { API_BASE } from '../../api';

interface Bus {
  id: string;
  numberPlate: string;
  make: string;
  model: string;
  seats: number;
  company: string;
}

const searchQuery = ref('');
const currentPage = ref(1);
const itemsPerPage = 6;

const buses = ref<Bus[]>([]);

const loadBuses = async () => {
  try {
    const response = await fetch(`${API_BASE}/buses`);
    const data = await response.json();
    if (!Array.isArray(data)) {
      buses.value = [];
      return;
    }
    buses.value = data.map((bus: any) => ({
      id: String(bus.id),
      numberPlate: String(bus.numberplate || bus.numberPlate || ''),
      make: String(bus.make || ''),
      model: String(bus.model || ''),
      seats: Number(bus.seats) || 0,
      company: String(bus.company || bus.company_id || ''),
    }));
  } catch (error) {
    console.error('Failed to load buses:', error);
    buses.value = [];
  }
};

onMounted(loadBuses);

const filteredBuses = computed(() => {
  const query = searchQuery.value.trim().toLowerCase();

  if (!query) {
    return buses.value;
  }

  return buses.value.filter((bus) =>
    bus.id.toLowerCase().includes(query) ||
    bus.numberPlate.toLowerCase().includes(query) ||
    bus.make.toLowerCase().includes(query) ||
    bus.model.toLowerCase().includes(query) ||
    bus.company.toLowerCase().includes(query) ||
    bus.seats.toString().includes(query)
  );
});

const totalPages = computed(() =>
  Math.max(1, Math.ceil(filteredBuses.value.length / itemsPerPage))
);

const paginatedBuses = computed(() => {
  const start = (currentPage.value - 1) * itemsPerPage;
  return filteredBuses.value.slice(start, start + itemsPerPage);
});

const onSearchChange = () => {
  currentPage.value = 1;
};
</script>

<template>
  <div class="flex w-full flex-col gap-6">
    <!-- Search -->
    <div class="w-full rounded-2xl border border-slate-800/80 bg-[#151b29] p-5 shadow-lg">
      <div class="relative w-full md:w-1/2 lg:w-1/3">
        <Search
          class="absolute left-4 top-1/2 h-5 w-5 -translate-y-1/2 text-slate-500"
        />

        <input
          v-model="searchQuery"
          type="text"
          placeholder="Search by bus ID, number plate, make, model, company..."
          class="w-full rounded-xl border border-slate-700/50 bg-[#0a0d14] py-3 pl-12 pr-4 text-white shadow-inner outline-none transition-all placeholder:text-slate-600 focus:border-[#5eafff] focus:ring-1 focus:ring-[#5eafff]"
          @input="onSearchChange"
        />
      </div>
    </div>

    <!-- Bus list -->
<div class="flex w-full flex-col gap-4">
  <div
    v-if="paginatedBuses.length === 0"
    class="flex w-full flex-col items-center justify-center rounded-2xl border border-dashed border-slate-800/50 bg-[#151b29]/50 py-16 text-slate-500"
  >
    <BusFront class="mb-4 h-10 w-10 opacity-50" />

    <p class="text-sm font-bold uppercase tracking-widest">
      No buses match your search
    </p>
  </div>

  <RegisteredBusCard
    v-for="bus in paginatedBuses"
    :key="bus.id"
    :bus="bus"
  />
</div>

    <!-- Pagination -->
    <div
      v-if="totalPages > 1"
      class="mt-2 flex w-full items-center justify-between rounded-2xl border border-slate-800/80 bg-[#151b29] p-4 shadow-lg"
    >
      <button
        :disabled="currentPage === 1"
        class="rounded-lg border border-slate-700/50 bg-[#0a0d14] px-4 py-2 text-xs font-bold uppercase tracking-widest text-slate-300 transition-colors hover:bg-[#1e2638] disabled:cursor-not-allowed disabled:opacity-50"
        @click="currentPage--"
      >
        Previous
      </button>

      <span class="text-sm font-bold text-slate-400">
        Page
        <span class="text-white">{{ currentPage }}</span>
        of {{ totalPages }}
      </span>

      <button
        :disabled="currentPage === totalPages"
        class="rounded-lg border border-slate-700/50 bg-[#0a0d14] px-4 py-2 text-xs font-bold uppercase tracking-widest text-slate-300 transition-colors hover:bg-[#1e2638] disabled:cursor-not-allowed disabled:opacity-50"
        @click="currentPage++"
      >
        Next
      </button>
    </div>
  </div>
</template>