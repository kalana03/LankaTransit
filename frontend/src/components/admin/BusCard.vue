<script setup lang="ts">
import { Building2, Armchair, BadgeInfo } from 'lucide-vue-next';

defineProps<{
  bus: {
    id: string;
    numberPlate: string;
    make: string;
    model: string;
    seats: number;
    company: string;
  };
}>();
</script>

<template>
  <article
    class="w-full overflow-hidden rounded-3xl border border-slate-800/80 bg-[#151b29] shadow-lg transition-all duration-300 hover:-translate-y-0.5 hover:shadow-[0_20px_70px_rgba(0,0,0,0.35)]"
  >
    <div
      class="grid w-full gap-6 px-6 py-5 sm:px-8 md:grid-cols-2 lg:grid-cols-[0.9fr_1.2fr_1fr_1fr_0.7fr_1.4fr] lg:items-center"
    >
      <!-- Bus ID -->
      <div class="min-w-0">
        <p class="mb-2 text-[11px] font-bold uppercase tracking-[0.2em] text-slate-500">
          Bus ID
        </p>
        <span
          class="inline-flex rounded-full border border-slate-700/60 bg-slate-900/50 px-4 py-2 text-sm font-bold text-white"
        >
          {{ bus.id }}
        </span>
      </div>

      <!-- Number Plate -->
      <div class="min-w-0">

        <span
          class="inline-flex rounded-xl border border-blue-500/20 bg-blue-500/10 px-5 py-2 text-2xl font-extrabold tracking-wider text-[#5eafff]"
        >
          {{ bus.numberPlate }}
        </span>
      </div>

      <!-- Make -->
      <div class="min-w-0">
        <div class="mb-2 flex items-center gap-2 text-slate-400">
          <BadgeInfo class="h-4 w-4" />
          <span class="text-[11px] font-bold uppercase tracking-[0.2em]">Make</span>
        </div>
        <p class="truncate text-lg font-semibold text-white">
          {{ bus.make }}
        </p>
      </div>

      <!-- Model -->
      <div class="min-w-0">
        <div class="mb-2 flex items-center gap-2 text-slate-400">
          <BadgeInfo class="h-4 w-4" />
          <span class="text-[11px] font-bold uppercase tracking-[0.2em]">Model</span>
        </div>
        <p class="truncate text-lg font-semibold text-white">
          {{ bus.model }}
        </p>
      </div>

      <!-- Seats -->
      <div class="min-w-0">
        <div class="mb-2 flex items-center gap-2 text-slate-400">
          <Armchair class="h-4 w-4" />
          <span class="text-[11px] font-bold uppercase tracking-[0.2em]">Seats</span>
        </div>
        <p class="text-lg font-semibold text-white">
          {{ bus.seats }}
        </p>
      </div>

      <!-- Company -->
      <div class="min-w-0">
        <div class="mb-2 flex items-center gap-2 text-slate-400 lg:justify-start">
          <Building2 class="h-4 w-4" />
          <span class="text-[11px] font-bold uppercase tracking-[0.2em]">Company</span>
        </div>
        <p class="truncate text-lg font-semibold text-white">
          {{ bus.company }}
        </p>
      </div>
    </div>
  </article>
</template>