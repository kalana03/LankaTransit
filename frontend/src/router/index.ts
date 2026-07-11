import { createRouter, createWebHistory, type RouteRecordRaw } from 'vue-router'

const routes: Array<RouteRecordRaw> = [
  {
    path: '/',
    name: 'Home',
    // We will create this view in the next step
    component: () => import('../views/Landing.vue')
  }
]

const router = createRouter({
  history: createWebHistory(),
  routes
})

export default router