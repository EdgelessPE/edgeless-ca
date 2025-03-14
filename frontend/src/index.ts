import { createApp } from 'vue';
import App from './App.vue';
import PrimeVue from "primevue/config";
import './index.css';
import Aura from '@primeuix/themes/aura';


createApp(App).use(PrimeVue, {
  ripple: true,
  theme: {
      preset: Aura,
      options: {
        darkModeSelector: 'system',
      }
  },

}).mount('#root');
