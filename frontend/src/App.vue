<template>
  <div
    class="flex flex-col items-center justify-center h-screen bg-#eee dark:bg-#000000"
  >
    <Toast />
    <User class="absolute top-4 right-4" />
    <RouterView />
  </div>
</template>

<script setup lang="ts">
import Toast from 'primevue/toast';
import { useToast } from 'primevue/usetoast';
import { emitter } from './utils/emit';
import User from './components/User.vue';
const toast = useToast();

emitter.on(
  'toast',
  (
    severity: 'success' | 'info' | 'warn' | 'error',
    summary: string,
    detail: string,
    life: number,
  ) => {
    toast.add({ severity, summary, detail, life });
  },
);
</script>
