<script setup lang="ts">
import { ref, computed } from 'vue';
import Header from '../components/Header.vue';
import Ticket from '../components/Ticket.vue';
import CheckoutForm from '../components/CheckOutForm.vue';

// ------------------------------------------------------------------------
// MOCKED STATE (In a real app, this comes from Pinia/Vuex or Vue Router)
// ------------------------------------------------------------------------
const currentSelection = ref([
  { id: 'B2', row: 'B', number: 2, status: 'selected' },
  { id: 'B3', row: 'B', number: 3, status: 'selected' }
]);

const busData = ref({
  company: 'LankaTransit Express',
  plate: 'WP-ND-8921',
  start: 'Colombo',
  destination: 'Kandy',
  departureTime: '06:30 AM',
  arrivalTime: '09:45 AM',
  date: 'Oct 24, 2024',
  pricePerSeat: 15.00
});

// Calculate total to pass into the form's Pay button
const totalAmount = computed(() => {
  return currentSelection.value.length * busData.value.pricePerSeat;
});
</script>

<template>
  <!-- Main Background: Deep Navy (#0f172a) -->
  <div class="min-h-screen pt-[8rem] pb-16 bg-[#0f172a] flex flex-col items-center font-sans selection:bg-[#5eafff] selection:text-[#0f172a]">
    
    <Header />

    <div class="w-full max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
        
      <!-- Page Title Area -->
      <div class="mb-10 text-center lg:text-left">
          <h1 class="text-3xl md:text-4xl font-extrabold text-white tracking-tight">
              Complete your <span class="text-[#5eafff]">Booking</span>
          </h1>
          <p class="text-slate-400 mt-3 text-sm md:text-base mx-auto lg:mx-0">
              Enter your details to finalize the reservation.
          </p>
      </div>

      <!-- Grid Layout -->
      <div class="grid grid-cols-1 lg:grid-cols-12 gap-8 lg:gap-12 items-start">
          
        <!-- LEFT SIDE: The new Checkout Form Component -->
        <main class="lg:col-span-7 xl:col-span-8 animate-fade-in-up">
            <CheckoutForm :totalAmount="totalAmount" />
        </main>

        <!-- RIGHT SIDE: Sticky Ticket (Same as Bookings page) -->
        <aside class="lg:col-span-5 xl:col-span-4 sticky top-[8rem] z-30">
            <Ticket 
                :busDetails="busData" 
                :selectedSeats="currentSelection" 
            />
        </aside>

      </div>  
    </div>
  </div>
</template>

<style scoped>
/* Page Load Animation */
.animate-fade-in-up {
  animation: fadeInUp 0.6s ease-out forwards;
}

@keyframes fadeInUp {
  from { opacity: 0; transform: translateY(20px); }
  to { opacity: 1; transform: translateY(0); }
}
</style>