<template>
  <div
    class="flex items-center justify-between h-lg w-80% bg-#1f1f1f rounded-3xl rounded-rt-32px rounded-rb-32px"
  >
    <div
      class="flex-1 items-center justify-center flex flex-col items-center gap-4px"
    >
      <div
        class="h-full text-36px line-height-48px font-bold bg-gradient-to-b from-#134597 to-#338aca bg-clip-text text-transparent"
      >
        Edgeless CA
      </div>
      <span class="text-16px line-height-24px text-#808080">
        {{ t('trustedNepPackageKeyPairService') }}
      </span>
    </div>
    <div
      class="w-40% h-full flex flex-col gap-4 rounded-3xl bg-white dark:bg-#18181b p-12"
    >
      <slot />
    </div>
  </div>
</template>

<script setup lang="ts">
import { t } from '../i18n';

defineSlots<{
  default: () => any;
}>();
</script>
