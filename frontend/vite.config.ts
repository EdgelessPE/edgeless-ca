import { defineConfig } from 'vite';
import vue from '@vitejs/plugin-vue';
import UnoCSS from 'unocss/vite';
import { presetAttributify } from '@unocss/preset-attributify';
import { presetWind3 } from '@unocss/preset-wind3';
import Components from 'unplugin-vue-components/vite';
import { PrimeVueResolver } from '@primevue/auto-import-resolver';

// https://vite.dev/config/
export default defineConfig({
  plugins: [
    vue(),
    Components({
      resolvers: [PrimeVueResolver()],
    }),
    UnoCSS({
      presets: [
        presetWind3({
          dark: 'media',
        }),
        presetAttributify(),
      ],
    }),
  ],
});
