import { defineConfig } from '@rsbuild/core';
import { pluginVue } from '@rsbuild/plugin-vue';
import Components from 'unplugin-vue-components/rspack';
import {PrimeVueResolver} from '@primevue/auto-import-resolver';

export default defineConfig({
  plugins: [pluginVue(),],
  tools:{
    rspack:{
      plugins:[
        Components({
          resolvers: [
            PrimeVueResolver()
          ]
        })
      ]
    }
  }
});
