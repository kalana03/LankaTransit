<script setup lang="ts">
// Navigation logic here if needed
</script>

<template>
  <header class="fixed top-0 z-50 w-full bg-transparent backdrop-blur-md border-b border-white/5 font-sans">
    <div class="max-w-7xl mx-auto flex items-center justify-between px-6 py-4 lg:py-5">
      
      <!-- Brand -->
      <a href="/" class="text-xl lg:text-2xl font-bold tracking-tight text-white outline-none">
        Lanka<span class="text-[#5fa8ff]">Transit</span>
      </a>

      <!-- Navigation -->
      <nav>
        <ul class="flex items-center gap-4 lg:gap-10 text-sm lg:text-base font-semibold text-white">
          
          <!-- Hidden on mobile, visible on medium+ screens -->
          <li class="hidden md:block">
            <a href="#" class="text-slate-300 hover:text-white transition-colors outline-none">
              Help
            </a>
          </li>
          <li class="hidden md:block">
            <a href="#" class="text-slate-300 hover:text-white transition-colors outline-none">
              Sign In
            </a>
          </li>

          <!-- CTA Button (Visible on all screens, slightly smaller padding on mobile) -->
          <li>
            <button class="bg-[#5fa8ff] text-[#0a1128] px-5 py-2 lg:px-7 lg:py-2.5 rounded-full font-bold transition-all hover:bg-[#4a90e2] active:scale-95 shadow-md">
              Book a Ride
            </button>
          </li>
        </ul>
      </nav>
      
    </div>
  </header>
</template>