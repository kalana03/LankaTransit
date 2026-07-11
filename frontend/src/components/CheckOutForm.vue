<script setup lang="ts">
import { ref } from 'vue';

// ------------------------------------------------------------------------
// FORM STATE & LOGIC
// ------------------------------------------------------------------------
const currentStep = ref<1 | 2>(1); // 1 = Passenger Details, 2 = Payment

// Form Data Models
const passenger = ref({
  firstName: '',
  lastName: '',
  email: '',
  phone: '',
  nic: '',
  address: ''
});

const payment = ref({
  cardNumber: '',
  expiry: '',
  cvc: '',
  nameOnCard: ''
});

// Actions
const proceedToPayment = () => {
  // In a real app, validate form fields here
  currentStep.value = 2;
  window.scrollTo({ top: 0, behavior: 'smooth' }); // Scroll to top for mobile
};

const goBack = () => {
  currentStep.value = 1;
};

const submitPayment = () => {
  // Handle final payment logic
  alert('Payment Successful! Processing ticket...');
};

// Define props if you want to pass the total price down to the Pay button
defineProps({
  totalAmount: {
    type: Number,
    default: 0
  }
});
</script>

<template>
  <!-- Premium Dark Shell: Matching the deep navy travel dashboard theme -->
  <div class="relative w-full p-8 lg:p-12 bg-gradient-to-b from-[#162032] to-[#0f172a] rounded-[2.5rem] shadow-2xl font-sans border border-white/5 overflow-hidden">
    
    <!-- Ambient Blue Glow -->
    <div class="absolute -top-40 -left-40 w-96 h-96 bg-[#5eafff] opacity-[0.03] blur-[100px] pointer-events-none"></div>

    <!-- Internal Form Stepper Header -->
    <div class="flex items-center justify-between mb-10 border-b border-slate-700/50 pb-6 relative z-10">
      
      <!-- Step 1 Indicator -->
      <div class="flex items-center gap-4 transition-opacity duration-300" :class="currentStep === 1 ? 'opacity-100' : 'opacity-40'">
        <div class="w-8 h-8 rounded-full flex items-center justify-center font-bold text-sm transition-colors duration-300"
             :class="currentStep === 1 ? 'bg-[#5eafff] text-[#0f172a] shadow-[0_0_15px_rgba(94,175,255,0.4)]' : 'bg-transparent border border-slate-600 text-slate-500'">
          1
        </div>
        <span class="text-white font-bold tracking-wide">Passenger Details</span>
      </div>
      
      <!-- Connector Line -->
      <div class="flex-grow h-px bg-slate-700/50 mx-6 hidden sm:block"></div>

      <!-- Step 2 Indicator -->
      <div class="flex items-center gap-4 transition-opacity duration-300" :class="currentStep === 2 ? 'opacity-100' : 'opacity-40'">
        <div class="w-8 h-8 rounded-full flex items-center justify-center font-bold text-sm transition-colors duration-300"
             :class="currentStep === 2 ? 'bg-[#5eafff] text-[#0f172a] shadow-[0_0_15px_rgba(94,175,255,0.4)]' : 'bg-transparent border border-slate-600 text-slate-500'">
          2
        </div>
        <span class="text-white font-bold tracking-wide">Payment</span>
      </div>
    </div>

    <!-- Transition Wrapper for smooth sliding between steps -->
    <div class="relative min-h-[400px]">
      <Transition name="slide" mode="out-in">
        
        <!-- ============================================== -->
        <!-- STEP 1: PASSENGER DETAILS FORM -->
        <!-- ============================================== -->
        <div v-if="currentStep === 1" key="step1" class="relative z-10 space-y-7">
          
          <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
            <!-- First Name -->
            <div class="space-y-2">
              <label class="text-[10px] text-slate-400 uppercase tracking-widest font-bold ml-1">First Name</label>
              <input v-model="passenger.firstName" type="text" placeholder="John" 
                     class="w-full bg-transparent text-white px-5 py-4 rounded-2xl border border-slate-600/60 hover:border-slate-500 focus:border-[#5eafff] focus:ring-1 focus:ring-[#5eafff] outline-none transition-all placeholder:text-slate-600">
            </div>
            <!-- Last Name -->
            <div class="space-y-2">
              <label class="text-[10px] text-slate-400 uppercase tracking-widest font-bold ml-1">Last Name</label>
              <input v-model="passenger.lastName" type="text" placeholder="Doe" 
                     class="w-full bg-transparent text-white px-5 py-4 rounded-2xl border border-slate-600/60 hover:border-slate-500 focus:border-[#5eafff] focus:ring-1 focus:ring-[#5eafff] outline-none transition-all placeholder:text-slate-600">
            </div>
          </div>

          <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
            <!-- Email -->
            <div class="space-y-2">
              <label class="text-[10px] text-slate-400 uppercase tracking-widest font-bold ml-1">Email Address</label>
              <input v-model="passenger.email" type="email" placeholder="john@example.com" 
                     class="w-full bg-transparent text-white px-5 py-4 rounded-2xl border border-slate-600/60 hover:border-slate-500 focus:border-[#5eafff] focus:ring-1 focus:ring-[#5eafff] outline-none transition-all placeholder:text-slate-600">
            </div>
            <!-- Contact Number -->
            <div class="space-y-2">
              <label class="text-[10px] text-slate-400 uppercase tracking-widest font-bold ml-1">Contact Number</label>
              <input v-model="passenger.phone" type="tel" placeholder="+94 77 123 4567" 
                     class="w-full bg-transparent text-white px-5 py-4 rounded-2xl border border-slate-600/60 hover:border-slate-500 focus:border-[#5eafff] focus:ring-1 focus:ring-[#5eafff] outline-none transition-all placeholder:text-slate-600">
            </div>
          </div>

          <!-- NIC -->
          <div class="space-y-2">
            <label class="text-[10px] text-slate-400 uppercase tracking-widest font-bold ml-1">NIC / Passport Number</label>
            <input v-model="passenger.nic" type="text" placeholder="Enter your ID number" 
                   class="w-full bg-transparent text-white px-5 py-4 rounded-2xl border border-slate-600/60 hover:border-slate-500 focus:border-[#5eafff] focus:ring-1 focus:ring-[#5eafff] outline-none transition-all placeholder:text-slate-600">
          </div>

          <!-- Step 1 Actions -->
          <div class="pt-6 flex justify-end">
            <button @click="proceedToPayment" 
                    class="px-8 py-4 bg-[#5eafff] text-[#0f172a] rounded-full font-bold text-sm uppercase tracking-widest shadow-[0_0_20px_rgba(94,175,255,0.2)] hover:shadow-[0_4px_25px_rgba(94,175,255,0.4)] hover:-translate-y-1 hover:bg-[#4ca0f0] transition-all flex items-center gap-3">
              <span>Proceed to Payment</span>
              <svg class="w-4 h-4" fill="none" stroke="currentColor" stroke-width="2.5" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" d="M14 5l7 7m0 0l-7 7m7-7H3"></path></svg>
            </button>
          </div>
        </div>

        <!-- ============================================== -->
        <!-- STEP 2: PAYMENT DETAILS FORM -->
        <!-- ============================================== -->
        <div v-else key="step2" class="relative z-10 space-y-7">
          
          <!-- Secure Payment Badge -->
          <div class="flex items-center justify-center gap-2 bg-emerald-500/10 border border-emerald-500/20 text-emerald-400 p-3.5 rounded-2xl mb-8 backdrop-blur-sm">
            <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 15v2m-6 4h12a2 2 0 002-2v-6a2 2 0 00-2-2H6a2 2 0 00-2 2v6a2 2 0 002 2zm10-10V7a4 4 0 00-8 0v4h8z"></path></svg>
            <span class="text-xs font-bold tracking-widest uppercase">Secure 256-bit Encryption</span>
          </div>

          <!-- Card Number -->
          <div class="space-y-2">
            <label class="text-[10px] text-slate-400 uppercase tracking-widest font-bold ml-1">Card Number</label>
            <div class="relative">
              <input v-model="payment.cardNumber" type="text" placeholder="0000 0000 0000 0000" 
                     class="w-full bg-transparent text-white px-5 py-4 pl-12 rounded-2xl border border-slate-600/60 hover:border-slate-500 focus:border-[#5eafff] focus:ring-1 focus:ring-[#5eafff] outline-none transition-all placeholder:text-slate-600 font-mono tracking-widest text-lg">
              <svg class="w-6 h-6 text-slate-500 absolute left-4 top-1/2 -translate-y-1/2" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M3 10h18M7 15h1m4 0h1m-7 4h12a3 3 0 003-3V8a3 3 0 00-3-3H6a3 3 0 00-3 3v8a3 3 0 003 3z"></path></svg>
            </div>
          </div>

          <div class="grid grid-cols-2 gap-6">
            <!-- Expiry -->
            <div class="space-y-2">
              <label class="text-[10px] text-slate-400 uppercase tracking-widest font-bold ml-1">Expiry Date</label>
              <input v-model="payment.expiry" type="text" placeholder="MM/YY" 
                     class="w-full bg-transparent text-white px-5 py-4 rounded-2xl border border-slate-600/60 hover:border-slate-500 focus:border-[#5eafff] focus:ring-1 focus:ring-[#5eafff] outline-none transition-all placeholder:text-slate-600 font-mono text-center tracking-widest">
            </div>
            <!-- CVC -->
            <div class="space-y-2">
              <label class="text-[10px] text-slate-400 uppercase tracking-widest font-bold ml-1">CVC / CVV</label>
              <input v-model="payment.cvc" type="password" placeholder="•••" 
                     class="w-full bg-transparent text-white px-5 py-4 rounded-2xl border border-slate-600/60 hover:border-slate-500 focus:border-[#5eafff] focus:ring-1 focus:ring-[#5eafff] outline-none transition-all placeholder:text-slate-600 font-mono text-center tracking-widest">
            </div>
          </div>

          <!-- Name on Card -->
          <div class="space-y-2">
            <label class="text-[10px] text-slate-400 uppercase tracking-widest font-bold ml-1">Name on Card</label>
            <input v-model="payment.nameOnCard" type="text" placeholder="JOHN DOE" 
                   class="w-full bg-transparent text-white px-5 py-4 rounded-2xl border border-slate-600/60 hover:border-slate-500 focus:border-[#5eafff] focus:ring-1 focus:ring-[#5eafff] outline-none transition-all placeholder:text-slate-600 uppercase">
          </div>

          <!-- Step 2 Actions -->
          <div class="pt-8 flex items-center justify-between">
            <button @click="goBack" class="text-slate-400 hover:text-white text-xs font-bold uppercase tracking-widest flex items-center gap-2 transition-colors">
              <svg class="w-4 h-4" fill="none" stroke="currentColor" stroke-width="2" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" d="M10 19l-7-7m0 0l7-7m-7 7h18"></path></svg>
              Back
            </button>
            
            <!-- In a real flow, this button would trigger `submitPayment`. Since your ticket handles it, I left it as per your layout. -->
          </div>
        </div>

      </Transition>
    </div>
  </div>
</template>

<style scoped>
/* Vue <Transition> Classes for sliding between Form Steps */
.slide-enter-active,
.slide-leave-active {
  transition: all 0.4s cubic-bezier(0.4, 0, 0.2, 1);
}

.slide-enter-from {
  opacity: 0;
  transform: translateX(20px);
}

.slide-leave-to {
  opacity: 0;
  transform: translateX(-20px);
}
</style>