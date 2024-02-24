import { createRouter, createWebHashHistory } from 'vue-router'
import LoginView from '../views/login.vue'
import AdminView from '../views/admin.vue'

const routes = [
  {
    path: '/',
    name: 'LoginView',
    component: LoginView
  },
  {
    path: '/admin',
    name: 'AdminView',
    component: AdminView
  },
]

const router = createRouter({
  history: createWebHashHistory(),
  routes
})

export default router
