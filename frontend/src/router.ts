import { createWebHashHistory, createRouter } from 'vue-router';
import LoginView from './pages/login.vue';
import HomeView from './pages/home.vue';

const routes = [
  { path: '/', component: HomeView },
  { path: '/login', component: LoginView },
];

export const router = createRouter({
  history: createWebHashHistory(),
  routes,
});
