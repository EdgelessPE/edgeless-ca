<template>
  <div>
    <h1>Home</h1>
  </div>
</template>

<script setup lang="ts">
import { onMounted } from 'vue';
import { getToken } from '../utils/token';
import { useRouter } from 'vue-router';

const router = useRouter();

onMounted(() => {
  const token = getToken();
  if (!token) {
    router.push('/login');
  }
});
</script>
