<script setup lang="ts">
import { ref } from 'vue';
import { API_BASE } from '../../api';
import { Plus } from 'lucide-vue-next';

const emit = defineEmits<{
  (e: 'created', route: any): void;
}>();

const form = ref({
  start: '',
  destination: '',
  distanceKm: 0,
  durationMinutes: 0,
  fare: 0,
  status: 'Active'
});

const submitting = ref(false);

const createRoute = async () => {
  if (!form.value.start || !form.value.destination || !form.value.distanceKm || !form.value.durationMinutes || !form.value.fare) {
    return;
  }

  submitting.value = true;

  try {
    const response = await fetch(`${API_BASE}/routes/create`, {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({
        start_location: form.value.start,
        destination: form.value.destination,
        distance_km: form.value.distanceKm,
        duration_minutes: form.value.durationMinutes,
        fare: form.value.fare,
        status: form.value.status,
      }),
    });
    if (!response.ok) {
      throw new Error('Failed to create route');
    }
    const route = await response.json();
    emit('created', route);
    form.value = {
      start: '',
      destination: '',
      distanceKm: 0,
      durationMinutes: 0,
      fare: 0,
      status: 'Active',
    };
  } catch (error) {
    console.error('Failed to create route:', error);
  } finally {
    submitting.value = false;
  }
};
</script>

<template>
  <div class="rounded-3xl border border-slate-800/80 bg-[#151b29] p-6 shadow-lg">
    <div class="mb-5 flex items-center justify-between gap-4">
      <div>
        <h3 class="text-lg font-bold text-white">New Route</h3>
        <p class="text-sm text-slate-400">Add a route to the backend route catalog.</p>
      </div>
      <Plus class="h-5 w-5 text-[#5eafff]" />
    </div>

    <div class="grid gap-4 md:grid-cols-2">
      <label class="flex flex-col gap-2 text-sm text-slate-300">
        Start Location
        <input v-model="form.start" type="text" class="rounded-2xl border border-slate-700/80 bg-[#0a0d14] p-3 text-white outline-none focus:border-[#5eafff]" />
      </label>
      <label class="flex flex-col gap-2 text-sm text-slate-300">
        Destination
        <input v-model="form.destination" type="text" class="rounded-2xl border border-slate-700/80 bg-[#0a0d14] p-3 text-white outline-none focus:border-[#5eafff]" />
      </label>
      <label class="flex flex-col gap-2 text-sm text-slate-300">
        Distance (km)
        <input v-model.number="form.distanceKm" type="number" min="0" class="rounded-2xl border border-slate-700/80 bg-[#0a0d14] p-3 text-white outline-none focus:border-[#5eafff]" />
      </label>
      <label class="flex flex-col gap-2 text-sm text-slate-300">
        Duration (minutes)
        <input v-model.number="form.durationMinutes" type="number" min="0" class="rounded-2xl border border-slate-700/80 bg-[#0a0d14] p-3 text-white outline-none focus:border-[#5eafff]" />
      </label>
      <label class="flex flex-col gap-2 text-sm text-slate-300 md:col-span-2">
        Fare (LKR)
        <input v-model.number="form.fare" type="number" min="0" step="0.01" class="rounded-2xl border border-slate-700/80 bg-[#0a0d14] p-3 text-white outline-none focus:border-[#5eafff]" />
      </label>
      <label class="flex flex-col gap-2 text-sm text-slate-300 md:col-span-2">
        Status
        <select v-model="form.status" class="rounded-2xl border border-slate-700/80 bg-[#0a0d14] p-3 text-white outline-none focus:border-[#5eafff]">
          <option value="Active">Active</option>
          <option value="Inactive">Inactive</option>
        </select>
      </label>
    </div>

    <button
      type="button"
      @click="createRoute"
      :disabled="submitting"
      class="mt-6 inline-flex items-center justify-center rounded-2xl bg-[#5eafff] px-5 py-3 text-sm font-bold uppercase tracking-widest text-slate-900 transition hover:bg-[#74b9ff] disabled:cursor-not-allowed disabled:opacity-50"
    >
      {{ submitting ? 'Saving…' : 'Create Route' }}
    </button>
  </div>
</template>
