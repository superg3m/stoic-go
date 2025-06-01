import { createRouter, createWebHistory } from 'vue-router'
import LoginView from '../views/LoginView.vue'
import SettingsView from "@/views/SettingsView.vue";
import HomeView from '@/views/HomeView.vue';
import ResourcesView from '@/views/ResourcesView.vue';
import {isAuthorized} from "@/UserStore.js";

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'login',
      component: LoginView,
      meta: { requiresAuth: false }
    },
    {
      path: '/Settings',
      name: 'settings',
      component: SettingsView,
      meta: { requiresAuth: true }
    },
    {
      path: '/home',
      name: 'Home',
      component: HomeView,
      meta: { requiresAuth: true }
    },
    {
      path: '/resources',
      name: 'Resources',
      component: ResourcesView,
      meta: { requiresAuth: true }
    },
    {
      path: '/resources',
      name: 'Resources',
      component: ResourcesView,
    }
  ],
})

router.beforeEach(async (to, from, next) => {
  const authorized = await isAuthorized();

  if (authorized && to.path === "/") {

    return next('/home');
  }

  if (!authorized && to.meta.requiresAuth) {
    return next('/');
  }

  next();
});

export default router
