import { createRouter, createWebHistory, type RouteRecordRaw } from 'vue-router'

const routes: Array<RouteRecordRaw> = [
  {
    path: '/',
    name: 'Home',
    // We will create this view in the next step
    component: () => import('../views/Landing.vue')
  },
  {
    path: '/login',
    name: 'Login',
    // We will create this view in the next step
    component: () => import('../views/Login.vue')
  },
  {
    path: '/results',
    name: 'Results',
    // We will create this view in the next step
    component: () => import('../views/Results.vue')
  },
  {
    path: '/booking',
    name: 'Booking',
    // We will create this view in the next step
    component: () => import('../views/Booking.vue')
  }
  ,
  {
    path: '/checkout',
    name: 'CheckOut',
    // We will create this view in the next step
    component: () => import('../views/CheckOut.vue')
  },
  {
    path: '/admin',
    name: 'Admin',
    // We will create this view in the next step
    component: () => import('../views/AdminDashboard.vue')
  }
]

const router = createRouter({
  history: createWebHistory(),
  routes
})

export default router