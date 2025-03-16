<template>
  <div>
    <Button
      icon="pi pi-language"
      severity="secondary"
      size="small"
      aria-haspopup="true"
      aria-controls="overlay_menu"
      @click="op.toggle($event)"
    />
    <Menu ref="op" id="overlay_menu" :model="items" :popup="true" />
  </div>
</template>

<script setup lang="ts">
import Button from 'primevue/button';
import Menu from 'primevue/menu';
import type { MenuItem } from 'primevue/menuitem';
import { ref } from 'vue';
import { changeLanguage } from '../i18n';

const items: MenuItem[] = [
  { label: 'English', value: 'en', command: () => changeLanguage('en') },
  { label: '简体中文', value: 'zh', command: () => changeLanguage('zh') },
];

const op = ref<any>(null);
</script>
