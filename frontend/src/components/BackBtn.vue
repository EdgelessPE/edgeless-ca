<template>
  <Button
    icon="pi pi-arrow-left"
    class="h-32px w-32px"
    severity="secondary"
    size="small"
    @click="router.back()"
  />
</template>

<script setup lang="ts">
import { router } from '../router';
import { Button } from 'primevue';
</script>
