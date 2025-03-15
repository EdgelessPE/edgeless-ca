import { createApp } from 'vue';
import App from './App.vue';
import PrimeVue from 'primevue/config';
import Aura from '@primeuix/themes/aura';
import 'primeicons/primeicons.css';
import 'virtual:uno.css';
import './index.css';
import { router } from './router';

createApp(App)
  .use(PrimeVue, {
    ripple: true,
    theme: {
      preset: Aura,
      options: {
        // darkModeSelector: 'system',
      },
    },
  })
  .use(router)
  .mount('#root');
