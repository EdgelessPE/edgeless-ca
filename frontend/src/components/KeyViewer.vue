<template>
  <div
    class="flex items-center justify-between mt-2 rounded-md bg-white dark:bg-black p-2"
  >
    <code>{{ viewKey }}</code>
    <Button
      icon="pi pi-copy"
      variant="text"
      size="small"
      @click="copyPublicKey"
    />
  </div>
</template>

<script setup lang="ts">
import { defineProps } from 'vue';
import { toast } from '../utils/toast';
import { t } from '../i18n';
const props = defineProps<{
  viewKey: string;
}>();

const copyPublicKey = () => {
  if (!props.viewKey) return;
  navigator.clipboard.writeText(props.viewKey);
  toast('success', t('copySuccess'), '', 3000);
};
</script>
