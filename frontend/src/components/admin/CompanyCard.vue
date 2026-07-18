<script setup lang="ts">
import { Building2, User, Phone, Mail, MapPin } from 'lucide-vue-next';

defineProps<{
  company: {
    id: number | string;
    name: string;
    owner: string;
    contactno: string;
    email: string;
    address: string;
  };
}>();
</script>

<template>
  <article
    class="w-full overflow-hidden rounded-3xl border border-slate-800/80 bg-[#151b29] shadow-lg transition-all duration-300 hover:-translate-y-0.5 hover:shadow-[0_20px_70px_rgba(0,0,0,0.35)]"
  >
    <div
      class="grid w-full gap-6 px-6 py-5 sm:px-8 md:grid-cols-2 lg:grid-cols-[0.8fr_1.3fr_1.2fr_1.1fr_1.5fr_1.8fr] lg:items-center"
    >
      <!-- Company ID -->
      <div class="min-w-0">
        <p class="mb-2 text-[11px] font-bold uppercase tracking-[0.2em] text-slate-500">
          Company ID
        </p>
        <span
          class="inline-flex rounded-full border border-slate-700/60 bg-slate-900/50 px-4 py-2 text-sm font-bold text-white"
        >
          {{ company.id }}
        </span>
      </div>

      <!-- Company Name -->
      <div class="min-w-0">
        <div class="mb-2 flex items-center gap-2 text-slate-400">
          <Building2 class="h-4 w-4" />
          <span class="text-[11px] font-bold uppercase tracking-[0.2em]">Company Name</span>
        </div>
        <p class="truncate text-base font-semibold text-[#5eafff]">
          {{ company.name }}
        </p>
      </div>

      <!-- Owner Name -->
      <div class="min-w-0">
        <div class="mb-2 flex items-center gap-2 text-slate-400">
          <User class="h-4 w-4" />
          <span class="text-[11px] font-bold uppercase tracking-[0.2em]">Owner Name</span>
        </div>
        <!-- Owner Name -->
        <p class="truncate text-base font-semibold text-white">
          {{ company.owner }}
        </p>
      </div>

      <!-- Contact Number -->
      <div class="min-w-0">
        <div class="mb-2 flex items-center gap-2 text-slate-400">
          <Phone class="h-4 w-4" />
          <span class="text-[11px] font-bold uppercase tracking-[0.2em]">Contact</span>
        </div>
        <!-- Contact Number -->
        <p class="truncate text-base font-semibold text-white">
          {{ company.contactno }}
        </p>
      </div>

      <!-- Email -->
      <div class="min-w-0">
        <div class="mb-2 flex items-center gap-2 text-slate-400">
          <Mail class="h-4 w-4" />
          <span class="text-[11px] font-bold uppercase tracking-[0.2em]">Email</span>
        </div>
        <p class="truncate text-base font-semibold text-white">
          {{ company.email }}
        </p>
      </div>

      <!-- Address -->
      <div class="min-w-0">
        <div class="mb-2 flex items-center gap-2 text-slate-400">
          <MapPin class="h-4 w-4" />
          <span class="text-[11px] font-bold uppercase tracking-[0.2em]">Address</span>
        </div>
        <p
          class="truncate text-base font-semibold text-white"
          :title="company.address"
        >
          {{ company.address }}
        </p>
      </div>
    </div>
  </article>
</template>