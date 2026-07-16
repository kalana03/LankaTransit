<script setup lang="ts">
import { 
  LayoutDashboard,  
  Navigation, 
  Ticket, 
  CalendarDays, 
  Route, 
  BusFront, 
  ClipboardCheck, 
  UserCog, 
  Users, 
  LogOut 
} from 'lucide-vue-next';

// 1. Accept the activeTab from the parent Layout
defineProps({
  activeTab: {
    type: String,
    required: true
  }
});

// 2. Emit the event when a tab is clicked
const emit = defineEmits(['update:activeTab']);

// 3. Map navigation items directly to Lucide Vue components
const navItems = [
  { name: 'Dashboard', icon: LayoutDashboard },  
  { name: 'Bookings', icon: Ticket },
  { name: 'Upcoming Tours', icon: CalendarDays },
  { name: 'Routes', icon: Route },
  { name: 'Registered Buses', icon: BusFront },
  { name: 'Companies', icon: UserCog },
];

const handleSignOut = () => {
  console.log('Signing out...');
};
</script>

<template>
  <!-- Sidebar Container -->
  <aside class="w-72 h-screen bg-gradient-to-b from-[#162032] to-[#0f172a] border-r border-slate-700/50 flex flex-col font-sans overflow-hidden shadow-2xl relative z-40">
    
    <!-- Top Branding Section -->
    <div class="px-8 pt-10 pb-8 border-b border-slate-800/80">
      <div class="flex items-center gap-3">
        <!-- Logo Mark -->
        <div class="w-10 h-10 rounded-xl bg-[#0f172a] border border-slate-700/60 flex items-center justify-center shadow-inner relative overflow-hidden">
          <div class="absolute inset-0 bg-[#5eafff] opacity-10 blur-sm"></div>
          <Navigation class="w-5 h-5 text-[#5eafff]" :stroke-width="2.5" />
        </div>
        <!-- Logo Text -->
        <div class="flex flex-col">
          <h1 class="text-xl font-black tracking-wide text-white uppercase leading-none">
            Lanka<span class="text-[#5eafff]">Transit</span>
          </h1>
          <span class="text-[10px] text-slate-500 font-bold tracking-[0.2em] uppercase mt-1">Admin Portal</span>
        </div>
      </div>
    </div>

    <!-- Navigation Menu -->
    <nav class="flex-grow px-4 py-6 overflow-y-auto custom-scrollbar space-y-1.5">
      <button
        v-for="item in navItems"
        :key="item.name"
        @click="emit('update:activeTab', item.name)"
        :class="[
          'w-full flex items-center gap-4 px-4 py-3.5 rounded-2xl transition-all duration-300 ease-out group relative overflow-hidden',
          activeTab === item.name 
            ? 'bg-[#1e2638] border border-slate-700/50' 
            : 'bg-transparent border border-transparent hover:bg-slate-800/40 hover:border-slate-700/30'
        ]"
      >
        <!-- Active Tab Indicator Glow -->
        <div v-if="activeTab === item.name" class="absolute left-0 top-1/2 -translate-y-1/2 w-1.5 h-1/2 bg-[#5eafff] rounded-r-full shadow-[0_0_12px_rgba(94,175,255,0.8)]"></div>

        <!-- Lucide Icon Component -->
        <component 
          :is="item.icon" 
          :class="[
            'w-5 h-5 transition-colors duration-300 relative z-10',
            activeTab === item.name ? 'text-[#5eafff]' : 'text-slate-500 group-hover:text-slate-300'
          ]" 
          :stroke-width="1.75"
        />

        <!-- Label -->
        <span 
          :class="[
            'text-sm font-bold tracking-wide transition-colors duration-300 relative z-10',
            activeTab === item.name ? 'text-white' : 'text-slate-400 group-hover:text-slate-200'
          ]"
        >
          {{ item.name }}
        </span>
      </button>
    </nav>

    <!-- Bottom Profile & Sign Out Section -->
    <div class="p-6 border-t border-slate-800/80 bg-[#0f172a]/50 backdrop-blur-md">
      <div class="flex items-center justify-between">
        
        <!-- Admin Info -->
        <div class="flex items-center gap-3">
          <div class="w-10 h-10 rounded-full bg-[#1e2638] border border-slate-700 flex items-center justify-center text-[#5eafff] font-bold text-sm shadow-inner">
            K
          </div>
          <div class="flex flex-col text-left">
            <span class="text-xs text-slate-500 font-bold tracking-widest uppercase mb-0.5">Logged In</span>
            <span class="text-sm font-bold text-slate-200">Kalana</span>
          </div>
        </div>

        <!-- Sign Out Button -->
        <button 
          @click="handleSignOut"
          class="p-2.5 rounded-xl text-slate-500 hover:text-rose-400 hover:bg-rose-500/10 transition-colors duration-200 group"
          title="Sign Out"
        >
          <LogOut class="w-5 h-5 group-hover:-translate-x-0.5 transition-transform" :stroke-width="2" />
        </button>
      </div>
    </div>

  </aside>
</template>

<style scoped>
/* Subtle scrollbar for the navigation area if items overflow */
.custom-scrollbar::-webkit-scrollbar {
  width: 4px;
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