<script setup lang="ts">
import { ref } from 'vue';
import Header from '../components/Header.vue';
import SeatSelector from '../components/SeatSelector.vue';
import Ticket from '../components/Ticket.vue';

// Reactive state to hold selected seats from the child component
const currentSelection = ref([]);

// Receive updates from SeatSelector.vue
const handleSelectionUpdate = (seats: any) => {
  currentSelection.value = seats;
};

// Dynamic Bus Data
const busData = ref({
  company: 'LankaTransit Express',
  plate: 'WP-ND-8921',
  start: 'Colombo',
  destination: 'Kandy',
  departureTime: '06:30 AM',
  arrivalTime: '09:45 AM',
  date: 'Oct 24, 2024',
  pricePerSeat: 1500.00
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
                    Select your <span class="text-[#5eafff]">seats</span>
                </h1>
                <p class="text-slate-400 mt-3 text-sm md:text-base mx-auto lg:mx-0">
                    Choose your preferred seats for the journey. The front of the bus is at the top of the layout.
                </p>
            </div>

            <!-- Grid Layout -->
            <div class="grid grid-cols-1 lg:grid-cols-12 gap-8 lg:gap-12 items-start">
                
                <!-- LEFT SIDE: Seat Selector (Takes 7 or 8 columns) -->
                <main class="lg:col-span-7 xl:col-span-8 animate-fade-in-up">
                    <SeatSelector @update:selection="handleSelectionUpdate" />
                </main>

                <!-- RIGHT SIDE: Ticket (Takes 5 or 4 columns, Sticky) -->
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
/* Optional: A smooth entrance animation for the page load */
.animate-fade-in-up {
  animation: fadeInUp 0.6s ease-out forwards;
}

@keyframes fadeInUp {
  from {
    opacity: 0;
    transform: translateY(20px);
  }
  to {
    opacity: 1;
    transform: translateY(0);
  }
}
</style>