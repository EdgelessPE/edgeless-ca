import { createWebHistory, createRouter } from 'vue-router';
import LoginView from './pages/login.vue';
import HomeView from './pages/home.vue';
import CallbackView from './pages/callback.vue';
import RecoverView from './pages/recover.vue';

const routes = [
  { path: '/', component: HomeView },
  { path: '/login', component: LoginView },
  { path: '/callback', component: CallbackView },
  { path: '/recover', component: RecoverView },
];

export const router = createRouter({
  history: createWebHistory(),
  routes,
});
