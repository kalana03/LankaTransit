<script setup lang="ts">
import { computed, ref } from 'vue';
import {
  Activity,
  ArrowUpRight,
  ArrowDownRight,
  Banknote,
  Building2,
  BusFront,
  CircleAlert,
  Clock3,
  MapPinned,
  Route,
  ShieldCheck,
  Ticket,
  Users
} from 'lucide-vue-next';

const selectedRange = ref('Last 7 Days');

const overview = ref({
  totalBookings: 1842,
  activeRoutes: 42,
  registeredBuses: 128,
  companies: 18,
  revenue: 2487500,
  occupancyRate: 78
});

const weeklyBookings = ref([
  { day: 'Mon', bookings: 180 },
  { day: 'Tue', bookings: 220 },
  { day: 'Wed', bookings: 260 },
  { day: 'Thu', bookings: 240 },
  { day: 'Fri', bookings: 310 },
  { day: 'Sat', bookings: 390 },
  { day: 'Sun', bookings: 340 }
]);

const topRoutes = ref([
  {
    id: 'RT-001',
    routeNumber: 'R-101',
    from: 'Colombo',
    to: 'Kandy',
    bookings: 324,
    occupancy: 92,
    revenue: 486000,
    status: 'High Demand'
  },
  {
    id: 'RT-002',
    routeNumber: 'R-102',
    from: 'Colombo',
    to: 'Nuwara Eliya',
    bookings: 281,
    occupancy: 87,
    revenue: 421500,
    status: 'Stable'
  },
  {
    id: 'RT-003',
    routeNumber: 'R-103',
    from: 'Galle',
    to: 'Colombo',
    bookings: 248,
    occupancy: 81,
    revenue: 372000,
    status: 'Stable'
  },
  {
    id: 'RT-004',
    routeNumber: 'R-104',
    from: 'Jaffna',
    to: 'Anuradhapura',
    bookings: 198,
    occupancy: 73,
    revenue: 297000,
    status: 'Monitor'
  }
]);

const companyPerformance = ref([
  { name: 'Lanka Express', trips: 412, contribution: 26, revenue: 645000 },
  { name: 'Island Travels', trips: 351, contribution: 21, revenue: 521000 },
  { name: 'Southern Lines', trips: 298, contribution: 18, revenue: 438000 },
  { name: 'Northern Route', trips: 233, contribution: 14, revenue: 326000 },
  { name: 'Coastal Express', trips: 191, contribution: 11, revenue: 281000 },
  { name: 'Hilltop Coaches', trips: 147, contribution: 10, revenue: 236500 }
]);

const fleetStats = ref({
  activeBuses: 104,
  maintenance: 12,
  inactive: 12,
  seatCapacity: 6120,
  seatsBookedToday: 4638
});

const alerts = ref([
  {
    title: 'Route RT-004 needs attention',
    description: 'Occupancy dropped below 75% over the last 3 days.',
    level: 'warning'
  },
  {
    title: 'Weekend bookings increased',
    description: 'Saturday bookings increased by 14.6% compared to last week.',
    level: 'success'
  },
  {
    title: '3 buses scheduled for maintenance',
    description: 'Fleet availability may reduce slightly tomorrow morning.',
    level: 'info'
  }
]);

const maxWeeklyBookings = computed(() =>
  Math.max(...weeklyBookings.value.map((item) => item.bookings))
);

const seatUtilization = computed(() =>
  Math.round((fleetStats.value.seatsBookedToday / fleetStats.value.seatCapacity) * 100)
);

const formattedRevenue = computed(() =>
  new Intl.NumberFormat('en-LK', {
    style: 'currency',
    currency: 'LKR',
    maximumFractionDigits: 0
  }).format(overview.value.revenue)
);

const totalTrips = computed(() =>
  companyPerformance.value.reduce((sum, company) => sum + company.trips, 0)
);

const bookingGrowth = ref(12.4);
const revenueGrowth = ref(8.7);
const routeGrowth = ref(4.2);
</script>

<template>
  <section class="w-full flex flex-col gap-6">
    <!-- Header -->


    <!-- KPI cards -->
    <div class="grid grid-cols-1 gap-4 md:grid-cols-2 xl:grid-cols-3 2xl:grid-cols-6">
      <div class="metric-card">
        <div class="metric-icon bg-sky-500/10 text-sky-400 border-sky-500/20">
          <Ticket class="h-5 w-5" />
        </div>
        <div>
          <p class="metric-label">Total Bookings</p>
          <p class="metric-value">{{ overview.totalBookings.toLocaleString() }}</p>
          <div class="metric-change text-emerald-400">
            <ArrowUpRight class="h-3.5 w-3.5" />
            +{{ bookingGrowth }}%
          </div>
        </div>
      </div>

      <div class="metric-card">
        <div class="metric-icon bg-indigo-500/10 text-indigo-400 border-indigo-500/20">
          <Route class="h-5 w-5" />
        </div>
        <div>
          <p class="metric-label">Active Routes</p>
          <p class="metric-value">{{ overview.activeRoutes }}</p>
          <div class="metric-change text-emerald-400">
            <ArrowUpRight class="h-3.5 w-3.5" />
            +{{ routeGrowth }}%
          </div>
        </div>
      </div>

      <div class="metric-card">
        <div class="metric-icon bg-cyan-500/10 text-cyan-400 border-cyan-500/20">
          <BusFront class="h-5 w-5" />
        </div>
        <div>
          <p class="metric-label">Registered Buses</p>
          <p class="metric-value">{{ overview.registeredBuses }}</p>
          <div class="metric-change text-slate-400">
            <Activity class="h-3.5 w-3.5" />
            Fleet stable
          </div>
        </div>
      </div>

      <div class="metric-card">
        <div class="metric-icon bg-fuchsia-500/10 text-fuchsia-400 border-fuchsia-500/20">
          <Building2 class="h-5 w-5" />
        </div>
        <div>
          <p class="metric-label">Companies</p>
          <p class="metric-value">{{ overview.companies }}</p>
          <div class="metric-change text-slate-400">
            <Users class="h-3.5 w-3.5" />
            Active
          </div>
        </div>
      </div>

      <div class="metric-card col-span-2">
        <div class="metric-icon bg-emerald-500/10 text-emerald-400 border-emerald-500/20">
          <Banknote class="h-5 w-5" />
        </div>
        <div>
          <p class="metric-label">Revenue</p>
          <p class="metric-value">{{ formattedRevenue }}</p>
          <div class="metric-change text-emerald-400">
            <ArrowUpRight class="h-3.5 w-3.5" />
            +{{ revenueGrowth }}%
          </div>
        </div>
      </div>

      
    </div>

    <!-- Main analytics area -->
    <div class="grid grid-cols-1 gap-6 xl:grid-cols-12">
      <!-- Weekly bookings -->
      <section
        class="rounded-3xl border border-slate-800/80 bg-[#151b29] p-6 shadow-lg xl:col-span-7"
      >
        <div class="mb-6 flex items-center justify-between gap-4">
          <div>
            <p class="text-[11px] font-bold uppercase tracking-[0.25em] text-slate-500">
              Booking Activity
            </p>
            <h3 class="mt-2 text-xl font-black text-white">
              Weekly Booking Trend
            </h3>
          </div>

          <div class="rounded-xl border border-slate-700/60 bg-[#0a0d14] px-4 py-2">
            <span class="text-xs font-bold uppercase tracking-widest text-slate-400">
              {{ selectedRange }}
            </span>
          </div>
        </div>

        <div class="grid h-[280px] grid-cols-7 items-end gap-3">
          <div
            v-for="item in weeklyBookings"
            :key="item.day"
            class="flex h-full flex-col items-center justify-end gap-3"
          >
            <span class="text-xs font-bold text-slate-400">
              {{ item.bookings }}
            </span>

            <div
              class="relative flex w-full items-end justify-center rounded-2xl bg-[#0a0d14] px-1"
              style="height: 220px"
            >
              <div
                class="w-full rounded-xl bg-gradient-to-t from-[#3b82f6] to-[#60a5fa] shadow-[0_0_20px_rgba(96,165,250,0.25)] transition-all duration-300"
                :style="{ height: `${(item.bookings / maxWeeklyBookings) * 100}%` }"
              />
            </div>

            <span class="text-[11px] font-bold uppercase tracking-wider text-slate-500">
              {{ item.day }}
            </span>
          </div>
        </div>
      </section>

      <!-- Fleet utilization -->
      <section
        class="rounded-3xl border border-slate-800/80 bg-[#151b29] p-6 shadow-lg xl:col-span-5"
      >
        <div class="mb-6">
          <p class="text-[11px] font-bold uppercase tracking-[0.25em] text-slate-500">
            Fleet Health
          </p>
          <h3 class="mt-2 text-xl font-black text-white">
            Utilization Summary
          </h3>
        </div>

        <div class="space-y-5">
          <div class="rounded-2xl border border-slate-800/80 bg-[#0a0d14] p-5">
            <div class="mb-3 flex items-center justify-between">
              <span class="text-sm font-semibold text-slate-300">Seat Utilization Today</span>
              <span class="text-lg font-black text-[#5eafff]">{{ seatUtilization }}%</span>
            </div>
            <div class="h-3 overflow-hidden rounded-full bg-slate-800">
              <div
                class="h-full rounded-full bg-gradient-to-r from-[#3b82f6] to-[#60a5fa]"
                :style="{ width: `${seatUtilization}%` }"
              />
            </div>
            <p class="mt-3 text-xs text-slate-500">
              {{ fleetStats.seatsBookedToday.toLocaleString() }} / {{ fleetStats.seatCapacity.toLocaleString() }} seats booked
            </p>
          </div>

          <div class="grid grid-cols-1 gap-4 sm:grid-cols-3">
            <div class="rounded-2xl border border-emerald-500/15 bg-emerald-500/5 p-4">
              <p class="text-[10px] font-bold uppercase tracking-widest text-emerald-400">
                Active Buses
              </p>
              <p class="mt-2 text-2xl font-black text-white">
                {{ fleetStats.activeBuses }}
              </p>
            </div>

            <div class="rounded-2xl border border-amber-500/15 bg-amber-500/5 p-4">
              <p class="text-[10px] font-bold uppercase tracking-widest text-amber-400">
                Maintenance
              </p>
              <p class="mt-2 text-2xl font-black text-white">
                {{ fleetStats.maintenance }}
              </p>
            </div>

            <div class="rounded-2xl border border-rose-500/15 bg-rose-500/5 p-4">
              <p class="text-[10px] font-bold uppercase tracking-widest text-rose-400">
                Inactive
              </p>
              <p class="mt-2 text-2xl font-black text-white">
                {{ fleetStats.inactive }}
              </p>
            </div>
          </div>

          <div class="rounded-2xl border border-slate-800/80 bg-[#0f172a] p-5">
            <div class="mb-4 flex items-center gap-2">
              <MapPinned class="h-4 w-4 text-[#5eafff]" />
              <p class="text-[11px] font-bold uppercase tracking-widest text-[#5eafff]">
                Operational Snapshot
              </p>
            </div>

            <div class="space-y-3 text-sm">
              <div class="flex items-center justify-between">
                <span class="text-slate-400">Fleet Availability</span>
                <span class="font-semibold text-white">81.2%</span>
              </div>
              <div class="flex items-center justify-between">
                <span class="text-slate-400">Avg. Route Occupancy</span>
                <span class="font-semibold text-white">78%</span>
              </div>
              <div class="flex items-center justify-between">
                <span class="text-slate-400">Buses Ready for Dispatch</span>
                <span class="font-semibold text-white">96</span>
              </div>
            </div>
          </div>
        </div>
      </section>

      <!-- Top routes -->
      <section
        class="rounded-3xl border border-slate-800/80 bg-[#151b29] p-6 shadow-lg xl:col-span-7"
      >
        <div class="mb-6">
          <p class="text-[11px] font-bold uppercase tracking-[0.25em] text-slate-500">
            Route Insights
          </p>
          <h3 class="mt-2 text-xl font-black text-white">
            Top Performing Routes
          </h3>
        </div>

        <div class="space-y-4">
          <div
            v-for="route in topRoutes"
            :key="route.id"
            class="rounded-2xl border border-slate-800/80 bg-[#0a0d14] p-5"
          >
            <div class="flex flex-col gap-4 lg:flex-row lg:items-center lg:justify-between">
              <div class="min-w-0">
                <div class="flex flex-wrap items-center gap-3">
                  <span class="rounded-full border border-slate-700/60 bg-slate-900/50 px-3 py-1 text-xs font-bold text-white">
                    {{ route.id }}
                  </span>
                  <span class="rounded-lg border border-blue-500/20 bg-blue-500/10 px-3 py-1 text-xs font-bold text-[#5eafff]">
                    {{ route.routeNumber }}
                  </span>
                </div>

                <div class="mt-3 flex flex-wrap items-center gap-2 text-white">
                  <span class="font-semibold">{{ route.from }}</span>
                  <span class="text-slate-500">→</span>
                  <span class="font-semibold">{{ route.to }}</span>
                </div>
              </div>

              <div class="grid grid-cols-2 gap-4 sm:grid-cols-4 lg:min-w-[420px]">
                <div>
                  <p class="text-[10px] font-bold uppercase tracking-widest text-slate-500">
                    Bookings
                  </p>
                  <p class="mt-1 text-base font-bold text-white">
                    {{ route.bookings }}
                  </p>
                </div>

                <div>
                  <p class="text-[10px] font-bold uppercase tracking-widest text-slate-500">
                    Occupancy
                  </p>
                  <p class="mt-1 text-base font-bold text-[#5eafff]">
                    {{ route.occupancy }}%
                  </p>
                </div>

                <div>
                  <p class="text-[10px] font-bold uppercase tracking-widest text-slate-500">
                    Revenue
                  </p>
                  <p class="mt-1 text-base font-bold text-white">
                    LKR {{ route.revenue.toLocaleString() }}
                  </p>
                </div>

                <div>
                  <p class="text-[10px] font-bold uppercase tracking-widest text-slate-500">
                    Status
                  </p>
                  <span
                    :class="[
                      'mt-1 inline-flex rounded-full px-2.5 py-1 text-[10px] font-black uppercase tracking-widest border',
                      route.status === 'High Demand'
                        ? 'bg-emerald-500/10 text-emerald-400 border-emerald-500/20'
                        : route.status === 'Monitor'
                        ? 'bg-amber-500/10 text-amber-400 border-amber-500/20'
                        : 'bg-sky-500/10 text-sky-400 border-sky-500/20'
                    ]"
                  >
                    {{ route.status }}
                  </span>
                </div>
              </div>
            </div>

            <div class="mt-4">
              <div class="mb-2 flex items-center justify-between text-xs">
                <span class="font-bold uppercase tracking-widest text-slate-500">
                  Occupancy Progress
                </span>
                <span class="font-semibold text-slate-300">
                  {{ route.occupancy }}%
                </span>
              </div>
              <div class="h-2 overflow-hidden rounded-full bg-slate-800">
                <div
                  class="h-full rounded-full bg-gradient-to-r from-[#3b82f6] to-[#60a5fa]"
                  :style="{ width: `${route.occupancy}%` }"
                />
              </div>
            </div>
          </div>
        </div>
      </section>

      <!-- Company contribution -->
      <section
        class="rounded-3xl border border-slate-800/80 bg-[#151b29] p-6 shadow-lg xl:col-span-5"
      >
        <div class="mb-6">
          <p class="text-[11px] font-bold uppercase tracking-[0.25em] text-slate-500">
            Operator Performance
          </p>
          <h3 class="mt-2 text-xl font-black text-white">
            Company Contribution
          </h3>
        </div>

        <div class="space-y-4">
          <div
            v-for="company in companyPerformance"
            :key="company.name"
            class="rounded-2xl border border-slate-800/80 bg-[#0a0d14] p-4"
          >
            <div class="mb-3 flex items-center justify-between gap-3">
              <div class="min-w-0">
                <p class="truncate font-semibold text-white">
                  {{ company.name }}
                </p>
                <p class="text-xs text-slate-500">
                  {{ company.trips }} trips · {{ company.contribution }}% share
                </p>
              </div>

              <div class="text-right">
                <p class="text-sm font-bold text-[#5eafff]">
                  LKR {{ company.revenue.toLocaleString() }}
                </p>
              </div>
            </div>

            <div class="h-2 overflow-hidden rounded-full bg-slate-800">
              <div
                class="h-full rounded-full bg-gradient-to-r from-cyan-500 to-sky-400"
                :style="{ width: `${company.contribution}%` }"
              />
            </div>
          </div>
        </div>

        <div class="mt-5 rounded-2xl border border-slate-800/80 bg-[#0f172a] p-4">
          <div class="flex items-center justify-between">
            <span class="text-sm text-slate-400">Total Trips Recorded</span>
            <span class="text-lg font-black text-white">{{ totalTrips }}</span>
          </div>
        </div>
      </section>
    </div>

    <!-- Bottom insights -->
    <div class="grid grid-cols-1 gap-6 xl:grid-cols-12">
      <section
        class="rounded-3xl border border-slate-800/80 bg-[#151b29] p-6 shadow-lg xl:col-span-8"
      >
        <div class="mb-6">
          <p class="text-[11px] font-bold uppercase tracking-[0.25em] text-slate-500">
            Strategic Insights
          </p>
          <h3 class="mt-2 text-xl font-black text-white">
            Performance Highlights
          </h3>
        </div>

        <div class="grid gap-4 md:grid-cols-2 xl:grid-cols-3">
          <div class="rounded-2xl border border-emerald-500/15 bg-emerald-500/5 p-5">
            <p class="text-[10px] font-bold uppercase tracking-widest text-emerald-400">
              Best Route
            </p>
            <p class="mt-2 text-lg font-black text-white">
              Colombo → Kandy
            </p>
            <p class="mt-1 text-sm text-slate-400">
              Highest occupancy and strongest revenue performance this week.
            </p>
          </div>

          <div class="rounded-2xl border border-sky-500/15 bg-sky-500/5 p-5">
            <p class="text-[10px] font-bold uppercase tracking-widest text-sky-400">
              Growth Signal
            </p>
            <p class="mt-2 text-lg font-black text-white">
              Weekend Demand Surge
            </p>
            <p class="mt-1 text-sm text-slate-400">
              Friday to Sunday bookings continue to outperform weekday traffic.
            </p>
          </div>

          <div class="rounded-2xl border border-amber-500/15 bg-amber-500/5 p-5">
            <p class="text-[10px] font-bold uppercase tracking-widest text-amber-400">
              Focus Area
            </p>
            <p class="mt-2 text-lg font-black text-white">
              Northern Route Efficiency
            </p>
            <p class="mt-1 text-sm text-slate-400">
              Underperforming occupancy suggests route timing or pricing review.
            </p>
          </div>
        </div>
      </section>

      <section
        class="rounded-3xl border border-slate-800/80 bg-[#151b29] p-6 shadow-lg xl:col-span-4"
      >
        <div class="mb-6">
          <p class="text-[11px] font-bold uppercase tracking-[0.25em] text-slate-500">
            Live Alerts
          </p>
          <h3 class="mt-2 text-xl font-black text-white">
            Operational Notices
          </h3>
        </div>

        <div class="space-y-4">
          <div
            v-for="alert in alerts"
            :key="alert.title"
            class="rounded-2xl border bg-[#0a0d14] p-4"
            :class="[
              alert.level === 'warning'
                ? 'border-amber-500/20'
                : alert.level === 'success'
                ? 'border-emerald-500/20'
                : 'border-sky-500/20'
            ]"
          >
            <div class="flex items-start gap-3">
              <div
                class="mt-0.5 flex h-9 w-9 shrink-0 items-center justify-center rounded-full"
                :class="[
                  alert.level === 'warning'
                    ? 'bg-amber-500/10 text-amber-400'
                    : alert.level === 'success'
                    ? 'bg-emerald-500/10 text-emerald-400'
                    : 'bg-sky-500/10 text-sky-400'
                ]"
              >
                <CircleAlert class="h-4 w-4" />
              </div>

              <div>
                <p class="font-bold text-white">
                  {{ alert.title }}
                </p>
                <p class="mt-1 text-sm leading-relaxed text-slate-400">
                  {{ alert.description }}
                </p>
              </div>
            </div>
          </div>
        </div>
      </section>
    </div>
  </section>
</template>

<style scoped>
.metric-card {
  @apply rounded-3xl border border-slate-800/80 bg-[#151b29] p-5 shadow-lg flex items-start gap-4;
}

.metric-icon {
  @apply flex h-12 w-12 shrink-0 items-center justify-center rounded-2xl border;
}

.metric-label {
  @apply text-[10px] font-bold uppercase tracking-widest text-slate-500;
}

.metric-value {
  @apply mt-2 text-2xl font-black text-white;
}

.metric-change {
  @apply mt-2 inline-flex items-center gap-1 text-xs font-bold uppercase tracking-widest;
}
</style>