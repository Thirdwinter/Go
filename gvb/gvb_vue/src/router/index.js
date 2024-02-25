import { createRouter, createWebHashHistory } from 'vue-router'


let routes = [
  { path: "/login", component: () => import("../views/login.vue") },
  { path: "/admin", component: () => import("../views/AdminView.vue") }
]

const router = createRouter({
  history: createWebHashHistory(),
  routes
})

export { router,routes }