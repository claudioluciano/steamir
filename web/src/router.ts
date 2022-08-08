import { createRouter, createWebHistory } from 'vue-router'

import homeView from '@/views/HomeView.vue'

const routes = [
  {
    path: '/',
    component: homeView,
    name: 'Home'
  }
]
export const router = createRouter({
  history: createWebHistory(),
  routes
})
