<template>
  <div>
    <code
      class="px-3 py-1.5 rounded-lg bg-gray-100 dark:bg-#18181b text-gray-700 dark:text-gray-300 hover:bg-gray-200 dark:hover:bg-#27272a transition-colors duration-200"
      aria-haspopup="true"
      aria-controls="overlay_menu"
      @click="toggle"
    >
      {{ text }}
    </code>
    <Menu ref="op" id="overlay_menu" :model="items" :popup="true" />
  </div>
</template>

<script setup lang="ts">
import { useLocalStorage } from '@vueuse/core';
import { computed, onMounted, ref } from 'vue';

const op = ref<any>(null);
const userInfo = useLocalStorage<{ email: string; name: string } | undefined>(
  'userInfo',
  undefined,
);

const text = computed(() => {
  if (!userInfo.value) {
    return '--';
  }
  return `${userInfo.value.name} <${userInfo.value.email}>`;
});

const items = [{ label: '注销', icon: 'pi pi-sign-out' }];

onMounted(async () => {
  userInfo.value = {
    email: 'cno@edgeless.top',
    name: 'Cno',
  };
});

const toggle = (event: MouseEvent) => {
  op.value?.toggle(event);
};
</script>
