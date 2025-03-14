import { defineConfig } from '@rsbuild/core';
import { pluginVue } from '@rsbuild/plugin-vue';
import Components from 'unplugin-vue-components/rspack';
import {PrimeVueResolver} from '@primevue/auto-import-resolver';
import { UnoCSSRspackPlugin } from '@unocss/webpack/rspack';
import { presetAttributify } from '@unocss/preset-attributify';
import { presetUno } from '@unocss/preset-uno';


export default defineConfig({
  plugins: [pluginVue(),],
  tools:{
    rspack:{
      plugins:[
        UnoCSSRspackPlugin({
          presets: [presetUno(), presetAttributify()],
        }),
        Components({
          resolvers: [
            PrimeVueResolver()
          ]
        })
      ]
    }
  }
});
