<script setup lang="ts">
import { computed, onBeforeUnmount, ref } from 'vue';
import {
  ArrowRight,
  Banknote,
  Ban,
  BusFront,
  CalendarDays,
  Clock3,
  Fingerprint,
  Mail,
  MapPin,
  Phone,
  ShieldCheck,
  User,
  X
} from 'lucide-vue-next';

interface Booking {
  id: string;
  tourId: string; // <-- Added Tour ID
  start: string;
  destination: string;
  date: string;
  time: string;
  name: string;
  status: string;
  seats: string[];
  busCompany: string;
  busId: string;
  busModel: string;
  driverId: string;
  assistantId: string;
  nic: string;
  phone: string;
  email: string;
  address?: string;
  totalPrice: number | string;
}

const props = defineProps<{
  booking: Booking;
}>();

const emit = defineEmits<{
  cancel: [bookingId: string];
}>();

const isModalOpen = ref(false);

const booking = computed(() => props.booking);
const isActive = computed(() => booking.value.status === 'Active');

const formattedTotal = computed(() => {
  return new Intl.NumberFormat('en-LK', {
    style: 'currency',
    currency: 'LKR',
    minimumFractionDigits: 2
  }).format(Number(booking.value.totalPrice));
});

const initials = computed(() => {
  return booking.value.name
    .split(' ')
    .filter(Boolean)
    .slice(0, 2)
    .map((part) => part[0])
    .join('')
    .toUpperCase();
});

const openModal = () => {
  isModalOpen.value = true;
  document.body.style.overflow = 'hidden';
};

const closeModal = () => {
  isModalOpen.value = false;
  document.body.style.overflow = '';
};

const cancelBooking = () => {
  emit('cancel', booking.value.id);
  closeModal();
};

const handleKeydown = (event: KeyboardEvent) => {
  if (event.key === 'Escape' && isModalOpen.value) {
    closeModal();
  }
};

window.addEventListener('keydown', handleKeydown);

onBeforeUnmount(() => {
  window.removeEventListener('keydown', handleKeydown);
  document.body.style.overflow = '';
});
</script>

<template>
  <!-- Booking summary card -->
  <article
    role="button"
    tabindex="0"
    :aria-label="`View booking ${booking.id}`"
    @click="openModal"
    @keydown.enter="openModal"
    @keydown.space.prevent="openModal"
    class="w-full bg-[#151b29] rounded-2xl border border-slate-800/80
           hover:border-[#5eafff]/40 hover:shadow-[0_10px_30px_rgba(0,0,0,0.45)]
           focus:outline-none focus:ring-2 focus:ring-[#5eafff]/60
           transition-all duration-300 cursor-pointer group"
  >
    <div class="px-5 sm:px-6 py-5 flex items-center justify-between gap-4">
      <div class="grid grid-cols-2 md:grid-cols-3 lg:grid-cols-8 gap-5 w-full items-center">
        <!-- Booking ID -->
        <div class="flex flex-col">
          <span class="text-[10px] text-slate-500 uppercase tracking-widest font-bold mb-1">
            Booking ID
          </span>
          <span class="text-white font-mono font-bold">{{ booking.id }}</span>
        </div>

        <!-- Tour ID (Replaced 'Route No') -->
        <div class="flex flex-col">
          <span class="text-[10px] text-slate-500 uppercase tracking-widest font-bold mb-1">
            Tour ID
          </span>
          <div class="flex items-center gap-2 text-slate-200 font-semibold">
            <span class="truncate font-mono">{{ booking.tourId }}</span>
          </div>
        </div>

        <!-- Route -->
        <div class="flex flex-col xl:col-span-2">
          <span class="text-[10px] text-slate-500 uppercase tracking-widest font-bold mb-1">
            Route
          </span>
          <div class="flex items-center gap-2 text-slate-200 font-semibold">
            <span class="truncate">{{ booking.start }}</span>
            <ArrowRight class="w-3.5 h-3.5 text-[#5eafff] shrink-0" />
            <span class="truncate">{{ booking.destination }}</span>
          </div>
        </div>

        <!-- Date -->
        <div class="flex flex-col">
          <span class="text-[10px] text-slate-500 uppercase tracking-widest font-bold mb-1">
            Date
          </span>
          <span class="text-slate-200 font-semibold text-sm">{{ booking.date }}</span>
        </div>

        <!-- Time -->
        <div class="flex flex-col">
          <span class="text-[10px] text-slate-500 uppercase tracking-widest font-bold mb-1">
            Departure
          </span>
          <span class="text-[#5eafff] text-sm font-bold">{{ booking.time }}</span>
        </div>

        <!-- Seats (Fixed typo here) -->
        <div class="flex flex-col">
          <span class="text-[10px] text-slate-500 uppercase tracking-widest font-bold mb-1">
            Seats
          </span>
          <span class="text-[#5eafff] font-bold">{{ booking.seats.length }}</span>
        </div>
        
        <!-- Status -->
        <div class="flex flex-col items-start justify-center">
          <span
              :class="[
              'w-fit px-2.5 py-1 rounded-md text-[10px] font-black uppercase tracking-widest border',
              isActive
                  ? 'bg-emerald-500/10 text-emerald-400 border-emerald-500/20'
                  : 'bg-rose-500/10 text-rose-400 border-rose-500/20'
              ]"
          >
              {{ booking.status }}
          </span>
        </div>
        
      </div>

      <!-- View details action -->
      <div
        class="w-10 h-10 rounded-full bg-[#0a0d14] border border-slate-700/50
               flex items-center justify-center shrink-0 transition-all duration-300
               group-hover:bg-[#5eafff] group-hover:border-[#5eafff]
               group-hover:shadow-[0_0_18px_rgba(94,175,255,0.35)]"
      >
        <ArrowRight class="w-4 h-4 text-slate-400 group-hover:text-[#0f172a] transition-colors" />
      </div>
    </div>
  </article>

  <!-- Modal -->
  <Teleport to="body">
    <Transition name="modal">
      <div
        v-if="isModalOpen"
        class="fixed inset-0 z-[100] flex items-center justify-center p-4 sm:p-6"
      >
        <!-- Backdrop -->
        <div
          class="modal-backdrop absolute inset-0 bg-[#070b12]/85 backdrop-blur-md"
          @click="closeModal"
        />

        <!-- Large ticket details panel -->
        <section
          role="dialog"
          aria-modal="true"
          aria-labelledby="booking-modal-title"
          class="modal-panel relative w-full max-w-6xl max-h-[92vh] overflow-hidden
                 bg-[#101827] border border-slate-700/60 rounded-[2rem] sm:rounded-[2.5rem]
                 shadow-[0_30px_80px_-20px_rgba(0,0,0,0.9)] flex flex-col"
          @click.stop
        >
          <!-- Ambient highlight -->
          <div
            class="absolute -top-28 left-1/3 w-[34rem] h-56 bg-[#5eafff]
                   opacity-[0.07] blur-[90px] pointer-events-none"
          />

          <!-- Header -->
          <header
            class="relative z-10 shrink-0 px-6 sm:px-8 py-5 sm:py-6
                   border-b border-slate-800/90 flex items-center justify-between gap-4"
          >
            <div>
              <div class="flex items-center gap-3">
                <div class="w-9 h-9 rounded-xl bg-[#5eafff]/10 border border-[#5eafff]/20 flex items-center justify-center">
                  <BusFront class="w-4 h-4 text-[#5eafff]" />
                </div>

                <div>
                  <h2 id="booking-modal-title" class="text-xl sm:text-2xl font-bold text-white">
                    Booking Details
                  </h2>
                  <p class="text-[10px] sm:text-xs text-slate-500 uppercase tracking-widest font-bold mt-1">
                    Reference:
                    <span class="text-[#5eafff] font-mono">{{ booking.id }}</span>
                  </p>
                </div>
              </div>
            </div>

            <div class="flex items-center gap-3">
              <span
                :class="[
                  'hidden sm:inline-flex px-3 py-1.5 rounded-full text-[10px] font-black uppercase tracking-widest border',
                  isActive
                    ? 'bg-emerald-500/10 text-emerald-400 border-emerald-500/20'
                    : 'bg-rose-500/10 text-rose-400 border-rose-500/20'
                ]"
              >
                {{ booking.status }}
              </span>

              <button
                type="button"
                aria-label="Close booking details"
                @click="closeModal"
                class="w-10 h-10 rounded-full bg-slate-800/80 border border-slate-700
                       text-slate-400 hover:bg-slate-700 hover:text-white
                       focus:outline-none focus:ring-2 focus:ring-[#5eafff]/60 transition-colors"
              >
                <X class="w-5 h-5 mx-auto" />
              </button>
            </div>
          </header>

          <!-- Scrollable body -->
          <div class="relative z-10 overflow-y-auto p-5 sm:p-8 custom-scrollbar">
            <!-- Route ticket banner -->
            <section
              class="mb-7 sm:mb-8 rounded-3xl border border-slate-800
                     bg-[#0a0f18] shadow-[inset_0_2px_16px_rgba(0,0,0,0.45)]
                     p-5 sm:p-7"
            >
              <div class="grid grid-cols-1 md:grid-cols-[1fr_auto_1fr] items-center gap-5">
                <!-- Origin -->
                <div>
                  <p class="text-[10px] text-slate-500 uppercase tracking-widest font-bold mb-2">
                    Departure
                  </p>
                  <p class="text-2xl sm:text-3xl font-black text-white">{{ booking.start }}</p>

                  <div class="flex flex-wrap items-center gap-3 mt-3">
                    <span class="inline-flex items-center gap-1.5 text-sm font-bold text-[#5eafff]">
                      <Clock3 class="w-4 h-4" />
                      {{ booking.time }}
                    </span>
                    <span class="inline-flex items-center gap-1.5 text-sm text-slate-400">
                      <CalendarDays class="w-4 h-4" />
                      {{ booking.date }}
                    </span>
                  </div>
                </div>

                <!-- Travel path -->
                <div class="hidden md:flex items-center justify-center w-36 relative">
                  <div class="absolute left-0 right-0 border-t border-dashed border-slate-700" />
                  <div class="relative w-11 h-11 rounded-full bg-[#172236] border border-[#5eafff]/30 flex items-center justify-center">
                    <BusFront class="w-5 h-5 text-[#5eafff]" />
                  </div>
                </div>

                <!-- Destination -->
                <div class="md:text-right">
                  <p class="text-[10px] text-slate-500 uppercase tracking-widest font-bold mb-2">
                    Destination
                  </p>
                  <p class="text-2xl sm:text-3xl font-black text-white">{{ booking.destination }}</p>
                  <p class="text-sm text-slate-400 mt-3">
                    {{ booking.busCompany }}
                  </p>
                </div>
              </div>
            </section>

            <!-- Main details layout -->
            <div class="grid grid-cols-1 lg:grid-cols-12 gap-6">
              <!-- Customer details: deliberately larger -->
              <section
                class="lg:col-span-8 relative overflow-hidden rounded-3xl
                       bg-gradient-to-br from-[#1a2940] via-[#162134] to-[#121b2a]
                       border border-[#5eafff]/25 p-6 sm:p-7
                       shadow-[0_12px_35px_rgba(0,0,0,0.28)]"
              >
                <div
                  class="absolute -right-20 -top-20 w-56 h-56 rounded-full
                         bg-[#5eafff] opacity-[0.08] blur-[70px] pointer-events-none"
                />

                <div class="relative z-10">
                  <div class="flex flex-col sm:flex-row sm:items-center justify-between gap-5 mb-7">
                    <div class="flex items-center gap-4">
                      <div
                        class="w-16 h-16 rounded-2xl bg-[#5eafff] text-[#0f172a]
                               flex items-center justify-center font-black text-xl
                               shadow-[0_0_25px_rgba(94,175,255,0.35)]"
                      >
                        {{ initials }}
                      </div>

                      <div>
                        <p class="text-[10px] text-[#5eafff] uppercase tracking-[0.18em] font-bold mb-1">
                          Primary Passenger
                        </p>
                        <h3 class="text-2xl sm:text-3xl font-black text-white">
                          {{ booking.name }}
                        </h3>
                        <p class="text-sm text-slate-400 mt-1">
                          Passenger and booking contact information
                        </p>
                      </div>
                    </div>

                    <div class="inline-flex items-center gap-2 self-start sm:self-auto px-3 py-2 rounded-full bg-emerald-500/10 border border-emerald-500/20">
                      <ShieldCheck class="w-4 h-4 text-emerald-400" />
                      <span class="text-[10px] text-emerald-400 uppercase tracking-widest font-bold">
                        Verified Booking
                      </span>
                    </div>
                  </div>

                  <!-- Customer detail tiles -->
                  <div class="grid grid-cols-1 sm:grid-cols-2 gap-3">
                    <div class="sm:col-span-2 detail-tile">
                      <Mail class="detail-icon" />
                      <div class="min-w-0">
                        <p class="detail-label">Email Address</p>
                        <p class="detail-value truncate">{{ booking.email }}</p>
                      </div>
                    </div>

                    <div class="detail-tile">
                      <Phone class="detail-icon" />
                      <div class="min-w-0">
                        <p class="detail-label">Contact Number</p>
                        <p class="detail-value">{{ booking.phone }}</p>
                      </div>
                    </div>

                    <div class="detail-tile">
                      <Fingerprint class="detail-icon" />
                      <div class="min-w-0">
                        <p class="detail-label">NIC Number</p>
                        <p class="detail-value font-mono">{{ booking.nic }}</p>
                      </div>
                    </div>

                    <div class="sm:col-span-2 detail-tile">
                      <MapPin class="detail-icon shrink-0" />
                      <div class="min-w-0">
                        <p class="detail-label">Address</p>
                        <p class="detail-value">
                          {{ booking.address || 'Address not provided' }}
                        </p>
                      </div>
                    </div>
                  </div>

                  <!-- Seats -->
                  <div class="mt-5 pt-5 border-t border-[#5eafff]/15 flex flex-col sm:flex-row sm:items-center justify-between gap-4">
                    <div>
                      <p class="text-[10px] text-slate-400 uppercase tracking-widest font-bold">
                        Reserved Seats
                      </p>
                      <p class="text-sm text-slate-400 mt-1">
                        {{ booking.seats.length }} seat{{ booking.seats.length === 1 ? '' : 's' }} reserved for this journey
                      </p>
                    </div>

                    <div class="flex flex-wrap gap-2">
                      <span
                        v-for="seat in booking.seats"
                        :key="seat"
                        class="px-3 py-1.5 rounded-lg bg-[#5eafff]/10 border border-[#5eafff]/30
                               text-[#5eafff] font-mono font-bold text-sm"
                      >
                        {{ seat }}
                      </span>
                    </div>
                  </div>
                </div>
              </section>

              <!-- Right-side ticket and payment information -->
              <aside class="lg:col-span-4 flex flex-col gap-6">
                <!-- Bus information -->
                <section class="rounded-3xl bg-[#151b29] border border-slate-800 p-6">
                  <div class="flex items-center gap-2 mb-5">
                    <BusFront class="w-4 h-4 text-[#5eafff]" />
                    <h3 class="text-[10px] text-[#5eafff] uppercase tracking-widest font-bold">
                      Fleet Information
                    </h3>
                  </div>

                  <dl class="space-y-4 text-sm">
                    <!-- Added TOUR ID Here -->
                    <div class="flex justify-between gap-4 pb-4 border-b border-slate-800/80">
                      <dt class="text-slate-400 font-bold uppercase tracking-widest text-[10px] self-center">Tour ID</dt>
                      <dd class="text-[#5eafff] font-mono font-black text-right">{{ booking.tourId }}</dd>
                    </div>

                    <div class="flex justify-between gap-4">
                      <dt class="text-slate-500">Operator</dt>
                      <dd class="text-white font-semibold text-right">{{ booking.busCompany }}</dd>
                    </div>

                    <div class="flex justify-between gap-4">
                      <dt class="text-slate-500">Bus Number</dt>
                      <dd class="text-white font-mono font-semibold uppercase text-right">{{ booking.busId }}</dd>
                    </div>

                    <div class="flex justify-between gap-4">
                      <dt class="text-slate-500">Model</dt>
                      <dd class="text-slate-300 text-right">{{ booking.busModel }}</dd>
                    </div>

                    <div class="pt-4 border-t border-slate-800 flex justify-between gap-4">
                      <dt class="text-slate-500">Driver ID</dt>
                      <dd class="text-slate-300 font-mono text-xs">{{ booking.driverId }}</dd>
                    </div>

                    <div class="flex justify-between gap-4">
                      <dt class="text-slate-500">Assistant ID</dt>
                      <dd class="text-slate-300 font-mono text-xs">{{ booking.assistantId }}</dd>
                    </div>
                  </dl>
                </section>

                <!-- Payment information -->
                <section class="rounded-3xl bg-[#151b29] border border-slate-800 p-6">
                  <div class="flex items-center gap-2 mb-5">
                    <Banknote class="w-4 h-4 text-[#5eafff]" />
                    <h3 class="text-[10px] text-[#5eafff] uppercase tracking-widest font-bold">
                      Payment Summary
                    </h3>
                  </div>

                  <div class="rounded-2xl bg-[#0a0d14] border border-slate-800 p-4 mb-5">
                    <p class="text-[10px] text-slate-500 uppercase tracking-widest font-bold">
                      Total Paid
                    </p>
                    <p class="text-3xl font-black text-white mt-2">
                      {{ formattedTotal }}
                    </p>
                  </div>

                  <button
                    v-if="isActive"
                    type="button"
                    @click="cancelBooking"
                    class="w-full py-3.5 bg-rose-500/10 text-rose-400 hover:bg-rose-500 hover:text-white
                           border border-rose-500/20 hover:border-rose-500 rounded-xl
                           font-bold text-xs uppercase tracking-widest transition-all
                           flex items-center justify-center gap-2"
                  >
                    <Ban class="w-4 h-4" />
                    Cancel Booking
                  </button>

                  <div
                    v-else
                    class="w-full py-3.5 bg-[#0a0d14] border border-slate-800 text-slate-600
                           rounded-xl font-bold text-xs uppercase tracking-widest text-center"
                  >
                    Booking Cancelled
                  </div>
                </section>
              </aside>
            </div>
          </div>
        </section>
      </div>
    </Transition>
  </Teleport>
</template>

<style scoped>
.detail-tile {
  @apply flex items-center gap-3 rounded-2xl bg-[#0b1320]/70 border border-slate-700/60 p-4;
}

.detail-icon {
  @apply w-4 h-4 text-[#5eafff] shrink-0;
}

.detail-label {
  @apply text-[10px] text-slate-500 uppercase tracking-widest font-bold mb-1;
}

.detail-value {
  @apply text-sm text-white font-semibold;
}

/* Modal fade + panel movement */
.modal-enter-active .modal-backdrop,
.modal-leave-active .modal-backdrop {
  transition: opacity 0.28s ease;
}

.modal-enter-active .modal-panel,
.modal-leave-active .modal-panel {
  transition: opacity 0.28s ease, transform 0.28s cubic-bezier(0.16, 1, 0.3, 1);
}

.modal-enter-from .modal-backdrop,
.modal-leave-to .modal-backdrop {
  opacity: 0;
}

.modal-enter-from .modal-panel,
.modal-leave-to .modal-panel {
  opacity: 0;
  transform: translateY(18px) scale(0.97);
}

/* Modal scrollbar */
.custom-scrollbar::-webkit-scrollbar {
  width: 6px;
}

.custom-scrollbar::-webkit-scrollbar-track {
  background: transparent;
}

.custom-scrollbar::-webkit-scrollbar-thumb {
  background: #334155;
  border-radius: 999px;
}

.custom-scrollbar::-webkit-scrollbar-thumb:hover {
  background: #475569;
}
</style>