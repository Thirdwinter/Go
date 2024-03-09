import { createRouter, createWebHashHistory } from 'vue-router'


let routes = [
  { path: "/", component: () => import("../views/index.vue") },
  { path: "/test", component: () => import("../views/test.vue") },
  { path: "/login", component: () => import("../views/login.vue") },
  { path: "/admin", component: () => import("../views/admin.vue"),
    children: [
      { path: "/admin/article", component: () => import("../views/back/article.vue") },
      { path: "/admin/category", component: () => import("../views/back/category.vue") },
      { path: "/admin/user", component: () => import("../views/back/user.vue") },

      //{ path: "users", component: () => import("../views/users.vue") },
    ]
  }
]

const router = createRouter({
  history: createWebHashHistory(),
  routes
})


//路由守卫
router.beforeEach((to, from, next) => {
  const token = localStorage.getItem('atoken');

  if ((to.path === '/admin'||to.path==="/admin/article"||to.path==="/admin/user"||to.path==="/admin/category") && !token) {
    // 如果不是在登录页面且 token 不存在，则强制返回登录页面
    next('/login');
  } else {
    next(); // 继续路由导航
  }
});
export { router,routes }