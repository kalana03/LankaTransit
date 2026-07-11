<script setup lang="ts">
import { ref } from 'vue'
import { Mail, Lock, ArrowRight, Eye, EyeOff } from 'lucide-vue-next'

const form = ref({
  email: '',
  password: ''
})

const showPassword = ref(false)

const handleLogin = () => {
  console.log('Logging in with:', form.value)
  // Add authentication logic here
}

const handleGoogleLogin = () => {
  console.log('Initiating Google Login...')
  // Add Google OAuth logic here
}
</script>

<template>
  <div class="flex min-h-screen bg-[#0f172a] text-white font-sans overflow-hidden">
    
    <!-- LEFT COLUMN: Form Area -->
    <div class="w-full lg:w-[45%] xl:w-[40%] flex flex-col px-6 sm:px-12 md:px-20 py-8 relative z-10">
      
      <!-- Subtle Background Glow -->
      <div class="absolute top-[-10%] left-[-20%] w-[70%] h-[50%] bg-[#5eafff]/10 blur-[120px] rounded-full pointer-events-none"></div>

      <!-- Top Nav / Logo -->
      <header class="flex items-center justify-between w-full">
        <a href="/" class="text-2xl font-bold tracking-tight text-white outline-none">
          Lanka<span class="text-[#5eafff]">Transit</span>
        </a>
      </header>

      <!-- Form Container (Vertically Centered) -->
      <main class="flex-1 flex flex-col justify-center max-w-md w-full mx-auto my-12">
        
        <div class="mb-8">
          <h1 class="text-3xl sm:text-4xl font-bold mb-3 tracking-tight">Welcome back</h1>
          <p class="text-slate-400 text-sm sm:text-base font-medium">
            Enter your details to access your dashboard and manage your bookings.
          </p>
        </div>

        <!-- Google Sign In Button -->
        <button 
          @click="handleGoogleLogin"
          type="button" 
          class="flex items-center justify-center gap-3 w-full px-6 py-3.5 bg-white/[0.02] hover:bg-white/[0.06] border border-white/[0.08] rounded-full transition-all text-sm font-semibold mb-6 outline-none focus-visible:border-[#5eafff]"
        >
          <!-- Google SVG Icon -->
          <svg class="w-5 h-5" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
            <path d="M22.56 12.25c0-.78-.07-1.53-.2-2.25H12v4.26h5.92c-.26 1.37-1.04 2.53-2.21 3.31v2.77h3.57c2.08-1.92 3.28-4.74 3.28-8.09z" fill="#4285F4"/>
            <path d="M12 23c2.97 0 5.46-.98 7.28-2.66l-3.57-2.77c-.98.66-2.23 1.06-3.71 1.06-2.86 0-5.29-1.93-6.16-4.53H2.18v2.84C3.99 20.53 7.7 23 12 23z" fill="#34A853"/>
            <path d="M5.84 14.09c-.22-.66-.35-1.36-.35-2.09s.13-1.43.35-2.09V7.07H2.18C1.43 8.55 1 10.22 1 12s.43 3.45 1.18 4.93l2.85-2.22.81-.62z" fill="#FBBC05"/>
            <path d="M12 5.38c1.62 0 3.06.56 4.21 1.64l3.15-3.15C17.45 2.09 14.97 1 12 1 7.7 1 3.99 3.47 2.18 7.07l3.66 2.84c.87-2.6 3.3-4.53 6.16-4.53z" fill="#EA4335"/>
          </svg>
          Continue with Google
        </button>

        <!-- Divider -->
        <div class="flex items-center gap-4 mb-6">
          <div class="flex-1 h-px bg-white/[0.08]"></div>
          <span class="text-[10px] font-bold uppercase tracking-widest text-slate-500">Or sign in with email</span>
          <div class="flex-1 h-px bg-white/[0.08]"></div>
        </div>

        <!-- Email & Password Form -->
        <form @submit.prevent="handleLogin" class="flex flex-col gap-4">
          
          <!-- Email Input -->
          <div class="group flex items-center px-6 py-3.5 bg-white/[0.02] hover:bg-white/[0.04] rounded-full border border-white/[0.08] focus-within:border-[#5eafff]/50 focus-within:bg-white/[0.04] transition-all cursor-text">
            <Mail class="h-5 w-5 text-slate-500 mr-4 group-focus-within:text-[#5eafff] transition-colors shrink-0" />
            <div class="flex flex-col w-full text-left">
              <label class="text-[10px] font-bold uppercase tracking-[0.1em] text-slate-400 mb-0.5 cursor-text">Email Address</label>
              <input 
                v-model="form.email" 
                type="email" 
                placeholder="hello@example.com" 
                class="bg-transparent border-none p-0 focus:ring-0 text-white font-semibold placeholder:text-slate-600 outline-none w-full text-base" 
                required 
              />
            </div>
          </div>

          <!-- Password Input -->
          <div class="group flex items-center px-6 py-3.5 bg-white/[0.02] hover:bg-white/[0.04] rounded-full border border-white/[0.08] focus-within:border-[#5eafff]/50 focus-within:bg-white/[0.04] transition-all cursor-text">
            <Lock class="h-5 w-5 text-slate-500 mr-4 group-focus-within:text-[#5eafff] transition-colors shrink-0" />
            <div class="flex flex-col w-full text-left">
              <label class="text-[10px] font-bold uppercase tracking-[0.1em] text-slate-400 mb-0.5 cursor-text">Password</label>
              <input 
                v-model="form.password" 
                :type="showPassword ? 'text' : 'password'" 
                placeholder="••••••••" 
                class="bg-transparent border-none p-0 focus:ring-0 text-white font-semibold placeholder:text-slate-600 outline-none w-full text-base tracking-widest font-mono" 
                required 
              />
            </div>
            <!-- Toggle Password Visibility -->
            <button 
              type="button" 
              @click="showPassword = !showPassword"
              class="ml-2 p-2 text-slate-500 hover:text-white transition-colors rounded-full outline-none focus-visible:text-[#5eafff]"
            >
              <Eye v-if="!showPassword" class="h-4 w-4" />
              <EyeOff v-else class="h-4 w-4" />
            </button>
          </div>

          <!-- Secondary Form Actions -->
          <div class="flex items-center justify-between px-2 mt-2 mb-4">
            <label class="flex items-center gap-2 cursor-pointer group">
              <div class="relative flex items-center justify-center w-4 h-4 rounded border border-white/[0.2] group-hover:border-[#5eafff] transition-colors">
                <input type="checkbox" class="peer absolute opacity-0 w-full h-full cursor-pointer" />
                <div class="absolute inset-0 bg-[#5eafff] scale-0 peer-checked:scale-100 transition-transform rounded-[3px]"></div>
              </div>
              <span class="text-xs font-medium text-slate-400 group-hover:text-slate-300 transition-colors">Remember me</span>
            </label>
            
            <a href="#" class="text-xs font-bold text-[#5eafff] hover:text-[#4a90e2] hover:underline transition-all">
              Forgot password?
            </a>
          </div>

          <!-- Primary Submit Button -->
          <button 
            type="submit" 
            class="w-full px-8 py-4 rounded-full flex items-center justify-center gap-2 bg-[#5eafff] hover:bg-[#4a90e2] text-[#0f172a] font-bold text-lg transition-all shadow-[0_0_20px_rgba(94,175,255,0.2)] hover:shadow-[0_0_25px_rgba(94,175,255,0.4)] active:scale-[0.98]"
          >
            Sign In
            <ArrowRight class="h-5 w-5" />
          </button>

        </form>

      </main>

      <!-- Footer / Switch to Sign Up -->
      <footer class="text-center md:text-left mt-auto pt-8 border-t border-white/[0.05]">
        <p class="text-slate-400 text-sm font-medium">
          Don't have an account? 
          <a href="#" class="text-[#5eafff] font-bold hover:underline transition-all ml-1">
            Sign up now
          </a>
        </p>
      </footer>

    </div>

    <!-- RIGHT COLUMN: Visual / Image Area (Hidden on mobile/tablet) -->
    <div class="hidden lg:flex lg:flex-1 relative bg-[url('/login.jpg')] bg-cover bg-center">
      
      <!-- Overlay: Deep navy with gradient fade -->
      <div class="absolute inset-0 bg-gradient-to-br from-[#0f172a]/90 via-[#0f172a]/60 to-[#0f172a]/40 z-0"></div>
      
      <!-- Decorative UI element in the image -->
      <div class="relative z-10 flex flex-col justify-end p-16 w-full h-full">
        <!-- Floating Glassmorphism Quote/Highlight Card -->
        <div class="max-w-md p-8 bg-white/[0.03] backdrop-blur-md border border-white/[0.08] rounded-[2rem] shadow-2xl">
          <div class="flex gap-1 mb-4">
            <!-- 5 Stars -->
            <svg v-for="i in 5" :key="i" class="w-5 h-5 text-[#5eafff]" fill="currentColor" viewBox="0 0 20 20">
              <path d="M9.049 2.927c.3-.921 1.603-.921 1.902 0l1.07 3.292a1 1 0 00.95.69h3.462c.969 0 1.371 1.24.588 1.81l-2.8 2.034a1 1 0 00-.364 1.118l1.07 3.292c.3.921-.755 1.688-1.54 1.118l-2.8-2.034a1 1 0 00-1.175 0l-2.8 2.034c-.784.57-1.838-.197-1.539-1.118l1.07-3.292a1 1 0 00-.364-1.118L2.98 8.72c-.783-.57-.38-1.81.588-1.81h3.461a1 1 0 00.951-.69l1.07-3.292z" />
            </svg>
          </div>
          <p class="text-white text-lg font-medium leading-relaxed mb-6">
            "Booking a bus from Colombo to Kandy has never been easier. The platform is sleek, fast, and completely reliable. Highly recommended for daily commuters."
          </p>
          <div class="flex items-center gap-4">
            <div class="w-10 h-10 rounded-full bg-gradient-to-tr from-[#5eafff] to-[#4a90e2] flex items-center justify-center text-[#0f172a] font-bold text-sm">
              SD
            </div>
            <div>
              <p class="text-white font-bold text-sm">Sahan Dissanayake</p>
              <p class="text-slate-400 text-xs font-medium">Frequent Traveler</p>
            </div>
          </div>
        </div>
      </div>

    </div>

  </div>
</template>

<style scoped>
/* Autofill stylings to prevent white background override in dark themes */
input:-webkit-autofill,
input:-webkit-autofill:hover, 
input:-webkit-autofill:focus, 
input:-webkit-autofill:active {
  -webkit-box-shadow: 0 0 0 30px #0f172a inset !important;
  -webkit-text-fill-color: white !important;
  transition: background-color 5000s ease-in-out 0s;
}
</style>