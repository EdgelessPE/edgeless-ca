import { defineConfig } from '@rsbuild/core';
import { pluginVue } from '@rsbuild/plugin-vue';
import Components from 'unplugin-vue-components/rspack';
import { PrimeVueResolver } from '@primevue/auto-import-resolver';
import { UnoCSSRspackPlugin } from '@unocss/webpack/rspack';
import { presetAttributify } from '@unocss/preset-attributify';
import { presetWind3 } from '@unocss/preset-wind3';

export default defineConfig({
  html: {
    title: 'Edgeless CA',
    favicon: './src/assets/nep.ico',
  },
  plugins: [pluginVue()],
  tools: {
    rspack: {
      plugins: [
        UnoCSSRspackPlugin({
          presets: [presetWind3(), presetAttributify()],
        }),
        Components({
          resolvers: [PrimeVueResolver()],
        }),
      ],
    },
  },
});
