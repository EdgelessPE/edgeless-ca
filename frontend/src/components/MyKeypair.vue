<template>
  <Dialog v-model:visible="visible" header="我的密钥对" modal>
    <template v-if="loading">
      <div class="pi pi-spin pi-spinner" />
    </template>
    <template v-else-if="keys">
      <div>公钥</div>
      <KeyViewer :view-key="keys.publicKey" />
      <div class="mt-4">私钥</div>
      <KeyViewer :view-key="keys.privateKey" />
    </template>
  </Dialog>
</template>

<script setup lang="ts">
import { ref, watch } from 'vue';
import { GetMyKeypair } from '../api/token';

const visible = defineModel<boolean>('visible');

const keys = ref<{
  publicKey: string;
  privateKey: string;
} | null>(null);
const loading = ref(false);

watch(visible, async (visible) => {
  if (!visible || loading.value || keys.value) return;
  loading.value = true;
  const {
    data: { data },
  } = await GetMyKeypair().finally(() => {
    loading.value = false;
  });
  keys.value = {
    publicKey: data.public_key,
    privateKey: data.private_key,
  };
});
</script>
