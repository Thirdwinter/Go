import { createRouter, createWebHashHistory } from 'vue-router'
import LoginView from '../views/Login.vue'

const routes = [
  {
    path: '/login',
    name: 'LoginView',
    component: LoginView
  }
]

const router = createRouter({
  history: createWebHashHistory(),
  routes
})

export default router
