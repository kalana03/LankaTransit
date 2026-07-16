Here's the updated `RouteCard.vue` with a **custom modal** instead of the `alert()` confirmation:

```vue
<script setup lang="ts">
import { ref, computed } from 'vue';
import { ArrowRight, MapPin, ChevronDown, Trash2, AlertTriangle, X } from 'lucide-vue-next';

const props = defineProps<{
  route: {
    id: string;
    routeNumber: string;
    from: string;
    to: string;
    status: 'Active' | 'Inactive';
  };
}>();

const isActive = computed(() => props.route.status === 'Active');

const emit = defineEmits<{
  delete: [routeId: string];
}>();

const isMapOpen = ref(false);
const showDeleteModal = ref(false);

const toggleMap = () => {
  isMapOpen.value = !isMapOpen.value;
};

const closeMap = () => {
  isMapOpen.value = false;
};

const openDeleteModal = () => {
  showDeleteModal.value = true;
};

const closeDeleteModal = () => {
  showDeleteModal.value = false;
};

const confirmDelete = () => {
  emit('delete', props.route.id);
  closeDeleteModal();
  closeMap();
};

// Google Maps embed URL for the route directions
const mapSrc = computed(() => {
  const origin = encodeURIComponent(props.route.from);
  const destination = encodeURIComponent(props.route.to);

  return `https://www.google.com/maps?saddr=${origin}&daddr=${destination}&output=embed`;
});
</script>

<template>
  <article
    class="overflow-hidden rounded-3xl border border-slate-800/80 bg-[#151b29] shadow-lg transition-all duration-300 hover:-translate-y-0.5 hover:shadow-[0_20px_70px_rgba(0,0,0,0.35)]"
  >
    <div
      class="grid items-center gap-6 px-6 py-2 sm:px-8 md:grid-cols-[auto_auto_1fr]"
    >
      <!-- Status dot + Route ID -->
<div>


  <div class="flex flex-wrap items-center gap-3">
    <span
      :class="[
        'inline-block h-3 w-3 rounded-full',
        isActive
          ? 'bg-emerald-400 shadow-[0_0_12px_rgba(52,211,153,0.9)]'
          : 'bg-rose-400 shadow-[0_0_12px_rgba(251,113,133,0.9)]'
      ]"
    />

    <span
      class="inline-flex rounded-full border border-slate-700/60 bg-slate-900/50 px-4 py-2 text-sm font-bold text-white"
    >
      {{ route.id }}
    </span>
  </div>
</div>

      <!-- Route Number -->
      <div>
        <span
          class="inline-flex rounded-lg border border-blue-500/20 bg-blue-500/10 px-4 py-2 text-3xl font-bold text-[#5eafff]"
        >
          {{ route.routeNumber }}
        </span>
      </div>

      <!-- Cities with dropdown arrow -->
      <div class="flex items-center gap-3 md:justify-self-end">
        <div class="flex flex-wrap items-center gap-3">
          <div class="flex items-center gap-2">
            <MapPin class="h-4 w-4 text-[#5eafff]" />
            <span class="font-semibold text-white">
              {{ route.from }}
            </span>
          </div>

          <ArrowRight class="h-5 w-5 text-slate-500" />

          <div class="flex items-center gap-2">
            <MapPin class="h-4 w-4 text-[#5eafff]" />
            <span class="font-semibold text-white">
              {{ route.to }}
            </span>
          </div>
        </div>

        <!-- Down arrow button -->
        <button
          class="ml-2 flex h-8 w-8 items-center justify-center rounded-full border border-slate-700/60 bg-[#0a0d14] transition-colors hover:bg-[#1e2638]"
          @click="toggleMap"
        >
          <ChevronDown
            class="h-4 w-4 text-slate-300 transition-transform duration-300"
            :class="{ 'rotate-180': isMapOpen }"
          />
        </button>
      </div>
    </div>

    <!-- Map dropdown content -->
    <Transition
      enter-active-class="transition-all duration-300 ease-out"
      enter-from-class="max-h-0 opacity-0"
      enter-to-class="max-h-[500px] opacity-100"
      leave-active-class="transition-all duration-300 ease-in"
      leave-from-class="max-h-[500px] opacity-100"
      leave-to-class="max-h-0 opacity-0"
    >
      <div v-if="isMapOpen" class="overflow-hidden">
        <div class="border-t border-slate-800/80 p-6">
          <!-- Map -->
          <div
            class="overflow-hidden rounded-2xl border border-slate-800/80"
          >
            <iframe
              :src="mapSrc"
              class="h-[350px] w-full"
              style="border: 0"
              loading="lazy"
              referrerpolicy="no-referrer-when-downgrade"
              allowfullscreen
            />
          </div>

          <!-- Delete button below map -->
          <div class="mt-4 flex justify-end">
            <button
              class="inline-flex items-center gap-2 rounded-lg border border-rose-500/20 bg-rose-500/10 px-5 py-3 text-xs font-bold uppercase tracking-widest text-rose-300 transition-colors hover:bg-rose-500/20"
              @click="openDeleteModal"
            >
              <Trash2 class="h-4 w-4" />
              Delete Route
            </button>
          </div>
        </div>
      </div>
    </Transition>

    <!-- Delete Confirmation Modal -->
    <Transition
      enter-active-class="transition-all duration-300 ease-out"
      enter-from-class="opacity-0 scale-95"
      enter-to-class="opacity-100 scale-100"
      leave-active-class="transition-all duration-300 ease-in"
      leave-from-class="opacity-100 scale-100"
      leave-to-class="opacity-0 scale-95"
    >
      <div
        v-if="showDeleteModal"
        class="fixed inset-0 z-50 flex items-center justify-center bg-black/60 backdrop-blur-sm"
        @click.self="closeDeleteModal"
      >
        <div
          class="mx-4 w-full max-w-md rounded-3xl border border-slate-800/80 bg-[#151b29] p-8 shadow-2xl"
        >
          <!-- Modal header -->
          <div class="mb-6 flex items-center justify-between">
            <div class="flex items-center gap-3">
              <div
                class="flex h-12 w-12 items-center justify-center rounded-full bg-rose-500/10"
              >
                <AlertTriangle class="h-6 w-6 text-rose-400" />
              </div>
              <div>
                <h3 class="text-lg font-bold text-white">
                  Delete Route
                </h3>
                <p class="text-sm text-slate-400">
                  This action cannot be undone
                </p>
              </div>
            </div>

            <button
              class="flex h-8 w-8 items-center justify-center rounded-full border border-slate-700/60 bg-[#0a0d14] transition-colors hover:bg-[#1e2638]"
              @click="closeDeleteModal"
            >
              <X class="h-4 w-4 text-slate-400" />
            </button>
          </div>

          <!-- Route details -->
          <div
            class="mb-6 rounded-2xl border border-slate-800/80 bg-[#0a0d14] p-4"
          >
            <div class="mb-3 flex items-center justify-between">
              <span class="text-xs font-bold uppercase tracking-widest text-slate-500">
                Route ID
              </span>
              <span class="font-bold text-white">
                {{ route.id }}
              </span>
            </div>
            <div class="flex items-center justify-between">
              <span class="text-xs font-bold uppercase tracking-widest text-slate-500">
                Route
              </span>
              <span class="font-semibold text-[#5eafff]">
                {{ route.from }} → {{ route.to }}
              </span>
            </div>
          </div>

          <!-- Modal actions -->
          <div class="flex gap-3">
            <button
              class="flex-1 rounded-xl border border-slate-700/60 bg-[#0a0d14] px-5 py-3 text-sm font-bold uppercase tracking-widest text-slate-300 transition-colors hover:bg-[#1e2638]"
              @click="closeDeleteModal"
            >
              Cancel
            </button>

            <button
              class="flex-1 rounded-xl bg-rose-500/10 px-5 py-3 text-sm font-bold uppercase tracking-widest text-rose-300 transition-colors hover:bg-rose-500/20"
              @click="confirmDelete"
            >
              Delete
            </button>
          </div>
        </div>
      </div>
    </Transition>
  </article>
</template>
```

### What changed

1. **Added modal state** — `showDeleteModal` ref to control the modal visibility.

2. **Replaced `confirm()` with modal** — instead of calling `alert()`, clicking the delete button now opens a custom modal.

3. **Modal design**:
   - **Backdrop** — semi-transparent black with blur effect, clicking outside closes it.
   - **Warning icon** — an `AlertTriangle` icon in a red-tinted circle.
   - **Route details** — shows the route ID and cities being deleted.
   - **Two buttons** — "Cancel" and "Delete" (styled in red).
   - **Close button** — an `X` button in the top-right corner.

4. **Smooth transitions** — the modal fades in/out with a scale effect.

5. **Auto-close map on delete** — `confirmDelete` also calls `closeMap()` to close the map dropdown after deletion.