<template>
  <div class="absolute bottom-4 text-12px text-#808080">
    <span>Copyright © {{ year() }}</span>
    <Button
      class="!text-12px !text-#808080 !p-0 ml-2"
      label="Edgeless Project"
      variant="link"
      @click="openEdgelessCA"
    />
  </div>
</template>

<script setup lang="ts">
import Button from 'primevue/button';

const year = () => new Date().getFullYear();

const openEdgelessCA = () => {
  window.open('https://github.com/EdgelessPE/edgeless-ca', '_blank');
};
</script>
