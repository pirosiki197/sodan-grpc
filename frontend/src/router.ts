import { createRouter, createWebHistory } from 'vue-router'
import HomePage from './pages/HomePage.vue'
import SodanChatPage from './pages/SodanChatPage.vue'
import CreateSodanPage from './pages/CreateSodanPage.vue'

const routes = [
  { path: '/', name: 'home', component: HomePage },
  { path: '/sodan/:id', name: 'sodan', component: SodanChatPage, props: true },
  { path: '/sodan/create', name: 'create', component: CreateSodanPage }
]

const router = createRouter({
  history: createWebHistory(),
  routes
})

export default router