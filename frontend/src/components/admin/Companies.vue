<script setup lang="ts">
import { computed, ref, onBeforeUnmount } from 'vue';
import { Search, Building2, Plus, X } from 'lucide-vue-next';
import CompanyCard from './CompanyCard.vue';

interface Company {
  id: string;
  name: string;
  ownerName: string;
  contactNumber: string;
  email: string;
  address: string;
}

const searchQuery = ref('');
const currentPage = ref(1);
const itemsPerPage = 6;

const companies = ref<Company[]>([
  {
    id: 'CMP-001',
    name: 'Lanka Express',
    ownerName: 'Saman Perera',
    contactNumber: '+94 77 123 4567',
    email: 'info@lankaexpress.com',
    address: 'No. 25, Galle Road, Colombo 03, Sri Lanka'
  },
  {
    id: 'CMP-002',
    name: 'Island Travels',
    ownerName: 'Nimal Fernando',
    contactNumber: '+94 71 234 5678',
    email: 'contact@islandtravels.lk',
    address: '45 Main Street, Kandy, Sri Lanka'
  },
  {
    id: 'CMP-003',
    name: 'Southern Lines',
    ownerName: 'Kasun Jayawardena',
    contactNumber: '+94 76 345 6789',
    email: 'support@southernlines.com',
    address: '12 Beach Road, Galle, Sri Lanka'
  },
  {
    id: 'CMP-004',
    name: 'Northern Route',
    ownerName: 'Malaka Silva',
    contactNumber: '+94 75 456 7890',
    email: 'admin@northernroute.lk',
    address: '88 Hospital Road, Jaffna, Sri Lanka'
  },
  {
    id: 'CMP-005',
    name: 'Coastal Express',
    ownerName: 'Tharindu Wickramasinghe',
    contactNumber: '+94 78 567 8901',
    email: 'hello@coastalexpress.com',
    address: '102 Station Road, Matara, Sri Lanka'
  },
  {
    id: 'CMP-006',
    name: 'Hilltop Coaches',
    ownerName: 'Amal Weerasinghe',
    contactNumber: '+94 70 678 9012',
    email: 'service@hilltopcoaches.lk',
    address: '17 Lake View, Kurunegala, Sri Lanka'
  }
]);

const showNewCompanyModal = ref(false);

const newCompanyForm = ref({
  name: '',
  ownerName: '',
  contactNumber: '',
  email: '',
  address: ''
});

const filteredCompanies = computed(() => {
  const query = searchQuery.value.trim().toLowerCase();

  if (!query) {
    return companies.value;
  }

  return companies.value.filter((company) =>
    company.id.toLowerCase().includes(query) ||
    company.name.toLowerCase().includes(query) ||
    company.ownerName.toLowerCase().includes(query) ||
    company.contactNumber.toLowerCase().includes(query) ||
    company.email.toLowerCase().includes(query) ||
    company.address.toLowerCase().includes(query)
  );
});

const totalPages = computed(() =>
  Math.max(1, Math.ceil(filteredCompanies.value.length / itemsPerPage))
);

const paginatedCompanies = computed(() => {
  const start = (currentPage.value - 1) * itemsPerPage;
  return filteredCompanies.value.slice(start, start + itemsPerPage);
});

const nextCompanyId = computed(() => {
  const maxNumber = companies.value.reduce((max, company) => {
    const match = company.id.match(/CMP-(\d+)/);
    const num = match ? Number(match[1]) : 0;
    return Math.max(max, num);
  }, 0);

  return `CMP-${String(maxNumber + 1).padStart(3, '0')}`;
});

const onSearchChange = () => {
  currentPage.value = 1;
};

const resetNewCompanyForm = () => {
  newCompanyForm.value = {
    name: '',
    ownerName: '',
    contactNumber: '',
    email: '',
    address: ''
  };
};

const openNewCompanyModal = () => {
  resetNewCompanyForm();
  showNewCompanyModal.value = true;
  document.body.style.overflow = 'hidden';
};

const closeNewCompanyModal = () => {
  showNewCompanyModal.value = false;
  document.body.style.overflow = '';
};

const addCompany = () => {
  const company: Company = {
    id: nextCompanyId.value,
    name: newCompanyForm.value.name.trim(),
    ownerName: newCompanyForm.value.ownerName.trim(),
    contactNumber: newCompanyForm.value.contactNumber.trim(),
    email: newCompanyForm.value.email.trim(),
    address: newCompanyForm.value.address.trim()
  };

  if (
    !company.name ||
    !company.ownerName ||
    !company.contactNumber ||
    !company.email ||
    !company.address
  ) {
    return;
  }

  companies.value.unshift(company);
  currentPage.value = 1;
  closeNewCompanyModal();
  resetNewCompanyForm();
};

onBeforeUnmount(() => {
  document.body.style.overflow = '';
});
</script>

<template>
  <div class="flex w-full flex-col gap-6">
    <!-- Search + New Company -->
    <div
      class="flex w-full flex-col items-stretch justify-between gap-4 rounded-2xl border border-slate-800/80 bg-[#151b29] p-5 shadow-lg md:flex-row md:items-center"
    >
      <div class="relative w-full md:w-1/2 lg:w-1/3">
        <Search
          class="absolute left-4 top-1/2 h-5 w-5 -translate-y-1/2 text-slate-500"
        />

        <input
          v-model="searchQuery"
          type="text"
          placeholder="Search by company ID, name, owner, contact, email..."
          class="w-full rounded-xl border border-slate-700/50 bg-[#0a0d14] py-3 pl-12 pr-4 text-white shadow-inner outline-none transition-all placeholder:text-slate-600 focus:border-[#5eafff] focus:ring-1 focus:ring-[#5eafff]"
          @input="onSearchChange"
        />
      </div>

      <button
        type="button"
        @click="openNewCompanyModal"
        class="inline-flex items-center justify-center gap-2 rounded-xl border border-[#5eafff]/20 bg-[#5eafff]/10 px-5 py-3 text-sm font-bold uppercase tracking-widest text-[#5eafff] transition-colors hover:bg-[#5eafff] hover:text-[#0f172a]"
      >
        <Plus class="h-4 w-4" />
        New Company
      </button>
    </div>

    <!-- Company list -->
    <div class="flex w-full flex-col gap-4">
      <div
        v-if="paginatedCompanies.length === 0"
        class="flex w-full flex-col items-center justify-center rounded-2xl border border-dashed border-slate-800/50 bg-[#151b29]/50 py-16 text-slate-500"
      >
        <Building2 class="mb-4 h-10 w-10 opacity-50" />

        <p class="text-sm font-bold uppercase tracking-widest">
          No companies match your search
        </p>
      </div>

      <CompanyCard
        v-for="company in paginatedCompanies"
        :key="company.id"
        :company="company"
      />
    </div>

    <!-- Pagination -->
    <div
      v-if="totalPages > 1"
      class="mt-2 flex w-full items-center justify-between rounded-2xl border border-slate-800/80 bg-[#151b29] p-4 shadow-lg"
    >
      <button
        :disabled="currentPage === 1"
        class="rounded-lg border border-slate-700/50 bg-[#0a0d14] px-4 py-2 text-xs font-bold uppercase tracking-widest text-slate-300 transition-colors hover:bg-[#1e2638] disabled:cursor-not-allowed disabled:opacity-50"
        @click="currentPage--"
      >
        Previous
      </button>

      <span class="text-sm font-bold text-slate-400">
        Page
        <span class="text-white">{{ currentPage }}</span>
        of {{ totalPages }}
      </span>

      <button
        :disabled="currentPage === totalPages"
        class="rounded-lg border border-slate-700/50 bg-[#0a0d14] px-4 py-2 text-xs font-bold uppercase tracking-widest text-slate-300 transition-colors hover:bg-[#1e2638] disabled:cursor-not-allowed disabled:opacity-50"
        @click="currentPage++"
      >
        Next
      </button>
    </div>
  </div>

  <!-- New Company Modal -->
  <Teleport to="body">
    <Transition name="modal">
      <div
        v-if="showNewCompanyModal"
        class="fixed inset-0 z-[100] flex items-center justify-center p-4 sm:p-6"
      >
        <div
          class="modal-backdrop absolute inset-0 bg-[#070b12]/85 backdrop-blur-md"
          @click="closeNewCompanyModal"
        />

        <section
          class="modal-panel relative w-full max-w-2xl overflow-hidden rounded-[2rem] border border-slate-700/60 bg-[#101827] shadow-[0_30px_80px_-20px_rgba(0,0,0,0.9)]"
          @click.stop
        >
          <div
            class="pointer-events-none absolute -top-20 left-1/3 h-40 w-80 bg-[#5eafff] opacity-[0.07] blur-[80px]"
          />

          <header
            class="relative z-10 flex items-center justify-between border-b border-slate-800/90 px-6 py-5 sm:px-8 sm:py-6"
          >
            <div class="flex items-center gap-3">
              <div class="flex h-10 w-10 items-center justify-center rounded-xl border border-[#5eafff]/20 bg-[#5eafff]/10">
                <Building2 class="h-5 w-5 text-[#5eafff]" />
              </div>

              <div>
                <h2 class="text-xl font-bold text-white sm:text-2xl">
                  Add New Company
                </h2>
                <p class="mt-1 text-xs font-bold uppercase tracking-widest text-slate-500">
                  Register a new transport company
                </p>
              </div>
            </div>

            <button
              type="button"
              @click="closeNewCompanyModal"
              class="flex h-10 w-10 items-center justify-center rounded-full border border-slate-700 bg-slate-800/80 text-slate-400 transition-colors hover:bg-slate-700 hover:text-white"
            >
              <X class="h-5 w-5" />
            </button>
          </header>

          <form @submit.prevent="addCompany" class="relative z-10 p-6 sm:p-8">
            <div class="mb-6 rounded-2xl border border-slate-800 bg-[#0a0d14] p-4">
              <p class="mb-1 text-[10px] font-bold uppercase tracking-widest text-slate-500">
                Company ID
              </p>
              <p class="font-mono text-lg font-black text-[#5eafff]">
                {{ nextCompanyId }}
              </p>
            </div>

            <div class="grid grid-cols-1 gap-4 sm:grid-cols-2">
              <div>
                <label class="mb-2 block text-[11px] font-bold uppercase tracking-[0.2em] text-slate-500">
                  Company Name
                </label>
                <input
                  v-model="newCompanyForm.name"
                  type="text"
                  required
                  class="w-full rounded-xl border border-slate-700/50 bg-[#0a0d14] px-4 py-3 text-white outline-none transition-all placeholder:text-slate-600 focus:border-[#5eafff] focus:ring-1 focus:ring-[#5eafff]"
                  placeholder="Enter company name"
                />
              </div>

              <div>
                <label class="mb-2 block text-[11px] font-bold uppercase tracking-[0.2em] text-slate-500">
                  Owner Name
                </label>
                <input
                  v-model="newCompanyForm.ownerName"
                  type="text"
                  required
                  class="w-full rounded-xl border border-slate-700/50 bg-[#0a0d14] px-4 py-3 text-white outline-none transition-all placeholder:text-slate-600 focus:border-[#5eafff] focus:ring-1 focus:ring-[#5eafff]"
                  placeholder="Enter owner name"
                />
              </div>

              <div>
                <label class="mb-2 block text-[11px] font-bold uppercase tracking-[0.2em] text-slate-500">
                  Contact Number
                </label>
                <input
                  v-model="newCompanyForm.contactNumber"
                  type="text"
                  required
                  class="w-full rounded-xl border border-slate-700/50 bg-[#0a0d14] px-4 py-3 text-white outline-none transition-all placeholder:text-slate-600 focus:border-[#5eafff] focus:ring-1 focus:ring-[#5eafff]"
                  placeholder="Enter contact number"
                />
              </div>

              <div>
                <label class="mb-2 block text-[11px] font-bold uppercase tracking-[0.2em] text-slate-500">
                  Email
                </label>
                <input
                  v-model="newCompanyForm.email"
                  type="email"
                  required
                  class="w-full rounded-xl border border-slate-700/50 bg-[#0a0d14] px-4 py-3 text-white outline-none transition-all placeholder:text-slate-600 focus:border-[#5eafff] focus:ring-1 focus:ring-[#5eafff]"
                  placeholder="Enter email address"
                />
              </div>

              <div class="sm:col-span-2">
                <label class="mb-2 block text-[11px] font-bold uppercase tracking-[0.2em] text-slate-500">
                  Address
                </label>
                <textarea
                  v-model="newCompanyForm.address"
                  required
                  rows="4"
                  class="w-full resize-none rounded-xl border border-slate-700/50 bg-[#0a0d14] px-4 py-3 text-white outline-none transition-all placeholder:text-slate-600 focus:border-[#5eafff] focus:ring-1 focus:ring-[#5eafff]"
                  placeholder="Enter company address"
                />
              </div>
            </div>

            <div class="mt-6 flex justify-end gap-3 border-t border-slate-800 pt-6">
              <button
                type="button"
                @click="closeNewCompanyModal"
                class="rounded-xl border border-slate-700/60 bg-[#0a0d14] px-5 py-3 text-sm font-bold uppercase tracking-widest text-slate-300 transition-colors hover:bg-[#1e2638]"
              >
                Cancel
              </button>

              <button
                type="submit"
                class="rounded-xl border border-[#5eafff]/20 bg-[#5eafff]/10 px-5 py-3 text-sm font-bold uppercase tracking-widest text-[#5eafff] transition-colors hover:bg-[#5eafff] hover:text-[#0f172a]"
              >
                Add Company
              </button>
            </div>
          </form>
        </section>
      </div>
    </Transition>
  </Teleport>
</template>

<style scoped>
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
</style>