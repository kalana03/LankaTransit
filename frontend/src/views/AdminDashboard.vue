<script setup lang="ts">
import { ref, computed } from 'vue';
import AdminSidebar from '../components/admin/Sidebar.vue';

// Import your section components here
import Bookings from '../components/admin/Bookings.vue';
import UpcomingTours from '../components/admin/UpcomingTours.vue';
import Routes from '../components/admin/Routes.vue';
// import Buses from './sections/Buses.vue';
// import RegisteredBuses from './sections/RegisteredBuses.vue';
// import BusOwners from './sections/BusOwners.vue';
// import RegisteredCustomers from './sections/RegisteredCustomers.vue';

// State to track the active view
const currentView = ref('Bookings');

// 1. Map the string names from the sidebar to the imported Vue components
const viewComponents: Record<string, any> = {
  'Bookings': Bookings,
  'Upcoming Tours': UpcomingTours, // Uncomment these as you build them
  'Routes': Routes,
  // 'Buses': Buses,
};

// 2. Computed property to return the correct component based on currentView
const activeComponent = computed(() => {
  return viewComponents[currentView.value] || Bookings; // Fallback to Bookings if not found
});
</script>

<template>
  <!-- Full Screen Wrapper: Prevents body scrolling, handles layout -->
  <div class="flex h-screen w-full bg-[#0a0d14] overflow-hidden font-sans text-slate-200">
    
    <!-- Sidebar: Fixed on the left -->
    <AdminSidebar 
      :activeTab="currentView" 
      @update:activeTab="currentView = $event" 
      class="flex-shrink-0"
    />

    <!-- 3. Main Content Area: Takes up remaining space, scrolls internally -->
    <main class="flex-1 h-full overflow-y-auto relative custom-scrollbar">
      
      <!-- Top Decorative Glow -->
      <div class="absolute top-0 left-1/2 -translate-x-1/2 w-[800px] h-[300px] bg-[#5eafff] opacity-[0.02] blur-[120px] pointer-events-none"></div>

      <!-- Header Area -->
      <header class="w-full px-10 py-8 sticky top-0 bg-[#0a0d14]/80 backdrop-blur-md z-20 border-b border-slate-800/50">
        <h2 class="text-2xl font-black text-white tracking-tight">{{ currentView }}</h2>
        <p class="text-sm text-slate-500 mt-1 font-medium">Manage and view system data</p>
      </header>

      <!-- 4. Dynamic Section Rendering -->
      <div class="p-10 relative z-10">
        <Transition name="fade" mode="out-in">
          <!-- Vue will automatically swap the component here based on activeComponent -->
          <component :is="activeComponent" />
        </Transition>
      </div>

    </main>
  </div>
</template>

<style scoped>
/* Smooth fade between sections */
.fade-enter-active,
.fade-leave-active {
  transition: opacity 0.2s ease, transform 0.2s ease;
}
.fade-enter-from,
.fade-leave-to {
  opacity: 0;
  transform: translateY(10px);
}

/* Scrollbar for the main content area */
.custom-scrollbar::-webkit-scrollbar {
  width: 8px;
}
.custom-scrollbar::-webkit-scrollbar-track {
  background: transparent;
}
.custom-scrollbar::-webkit-scrollbar-thumb {
  background: #1e293b;
  border-radius: 10px;
}
.custom-scrollbar:hover::-webkit-scrollbar-thumb {
  background: #334155;
}
</style>