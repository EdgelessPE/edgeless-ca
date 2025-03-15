<template>
  <div>
    <Toast />
    <RouterView />
  </div>
</template>

<script setup lang="ts">
import Toast from 'primevue/toast';
import { useToast } from 'primevue/usetoast';
import { emitter } from './utils/emit';

const toast = useToast();

emitter.on(
  'toast',
  (
    severity: 'success' | 'info' | 'warn' | 'error',
    summary: string,
    detail: string,
    life: number,
  ) => {
    console.log(severity, summary, detail, life);
    toast.add({ severity, summary, detail, life });
  },
);
</script>
