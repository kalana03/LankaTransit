<script setup lang="ts">
import { computed, ref, onMounted } from 'vue';
import { Search, Plus } from 'lucide-vue-next';
import RouteCard from './RouteCard.vue';
import NewRouteForm from './NewRouteForm.vue';
import { API_BASE } from '../../api';

interface RouteItem {
  id: string;
  routeNumber: string;
  from: string;
  to: string;
  status: 'Active' | 'Inactive';
}

const searchQuery = ref('');
const currentPage = ref(1);
const itemsPerPage = 6;
const showNewRouteForm = ref(false);

const routes = ref<RouteItem[]>([]);

const normalizeRoute = (route: any): RouteItem => ({
  id: String(route.id),
  routeNumber: route.routeNumber || `R-${String(route.id).padStart(3, '0')}`,
  from: String(route.start || route.start_location || ''),
  to: String(route.destination || route.destination || ''),
  status: String(route.status || 'Active').toLowerCase() === 'active' ? 'Active' : 'Inactive',
});

const loadRoutes = async () => {
  try {
    const response = await fetch(`${API_BASE}/routes`);
    const data = await response.json();
    if (!Array.isArray(data)) {
      routes.value = [];
      return;
    }
    routes.value = data.map(normalizeRoute);
  } catch (error) {
    console.error('Failed to load routes:', error);
    routes.value = [];
  }
};

const deleteRoute = async (routeId: string) => {
  try {
    const response = await fetch(`${API_BASE}/routes/delete?id=${routeId}`, {
      method: 'DELETE',
    });
    if (!response.ok) {
      throw new Error('Failed to delete route');
    }
    routes.value = routes.value.filter((route) => route.id !== routeId);
  } catch (error) {
    console.error('Delete route failed:', error);
  }
};

const onRouteCreated = (route: any) => {
  routes.value.unshift(normalizeRoute(route));
  showNewRouteForm.value = false;
};

onMounted(loadRoutes);

const filteredRoutes = computed(() => {
  const query = searchQuery.value.trim().toLowerCase();

  if (!query) {
    return routes.value;
  }

  return routes.value.filter((route) =>
    route.id.toLowerCase().includes(query) ||
    route.routeNumber.toLowerCase().includes(query) ||
    route.from.toLowerCase().includes(query) ||
    route.to.toLowerCase().includes(query)
  );
});

const totalPages = computed(() =>
  Math.max(1, Math.ceil(filteredRoutes.value.length / itemsPerPage))
);

const paginatedRoutes = computed(() => {
  const start = (currentPage.value - 1) * itemsPerPage;
  return filteredRoutes.value.slice(start, start + itemsPerPage);
});

const onSearchChange = () => {
  currentPage.value = 1;
};
</script>

<template>
  <div class="flex w-full flex-col gap-6">
    <!-- Search -->
    <div
      class="w-full rounded-2xl border border-slate-800/80 bg-[#151b29] p-5 shadow-lg"
    >
      <div class="relative w-full md:w-1/2 lg:w-1/3">
        <Search
          class="absolute left-4 top-1/2 h-5 w-5 -translate-y-1/2 text-slate-500"
        />

        <input
          v-model="searchQuery"
          type="text"
          placeholder="Search by route ID, number, or city..."
          class="w-full rounded-xl border border-slate-700/50 bg-[#0a0d14] py-3 pl-12 pr-4 text-white shadow-inner outline-none transition-all placeholder:text-slate-600 focus:border-[#5eafff] focus:ring-1 focus:ring-[#5eafff]"
          @input="onSearchChange"
        />
      </div>
    </div>

    <div class="flex w-full flex-col gap-4">
      <button
        type="button"
        @click="showNewRouteForm = !showNewRouteForm"
        class="inline-flex items-center justify-center gap-2 rounded-2xl border border-[#5eafff]/20 bg-[#5eafff]/10 px-5 py-3 text-sm font-bold uppercase tracking-widest text-[#5eafff] transition hover:bg-[#5eafff] hover:text-[#0f172a]"
      >
        <Plus class="h-4 w-4" />
        {{ showNewRouteForm ? 'Hide Route Form' : 'Add Route' }}
      </button>

      <NewRouteForm v-if="showNewRouteForm" @created="onRouteCreated" />
    </div>

    <!-- Route list -->
<div class="grid items-start w-full grid-cols-1 gap-3 md:grid-cols-2">
  <div
    v-if="paginatedRoutes.length === 0"
    class="col-span-full flex w-full flex-col items-center justify-center rounded-2xl border border-dashed border-slate-800/50 bg-[#151b29]/50 py-16 text-slate-500"
  >
    <Search class="mb-4 h-10 w-10 opacity-50" />

    <p class="text-sm font-bold uppercase tracking-widest">
      No routes match your search
    </p>
  </div>

  <RouteCard
    v-for="route in paginatedRoutes"
    :key="route.id"
    :route="route"
    @delete="deleteRoute"
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