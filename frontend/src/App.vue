<template>
  <div
    class="flex flex-col items-center justify-center h-screen bg-#eee dark:bg-#000000"
  >
    <Toast />
    <RouterView />
    <Footer />
  </div>
</template>

<script setup lang="ts">
import Toast from 'primevue/toast';
import { useToast } from 'primevue/usetoast';
import { emitter } from './utils/emit';
import Footer from './layouts/Footer.vue';

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
