<template>
  <div class="flex flex-col items-center justify-center h-screen bg-black">
    <div
      class="flex items-center justify-between h-lg w-80% bg-#1f1f1f rounded-3xl"
    >
      <slot />
    </div>
  </div>
</template>

<script setup lang="ts">
defineSlots<{
  default: () => any;
}>();
</script>
