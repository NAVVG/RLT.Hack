import { createRouter, createWebHistory } from "vue-router"

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: "/",
      name: "home",
      component: () => import("@/views/HomeView.vue")
    },
    {
      path: "/profile",
      name: "profile",
      component: () => import("@/views/ProfileView.vue")
    },
    {
      path: "/purchase",
      name: "purchase",
      component: () => import("@/views/PurchaseView.vue")
    }
  ]
})

export default router
