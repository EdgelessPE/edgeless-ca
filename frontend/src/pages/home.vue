<template>
  <MyKeypair v-model:visible="myKeypairVisible" />
  <div class="h-full w-full">
    <User
      class="absolute top-5 right-4"
      @view-my-key="myKeypairVisible = true"
    />
    <div class="flex flex-col gap-4 items-center w-full mt-40vh">
      <div
        class="text-48px font-bold bg-gradient-to-b from-#134597 to-#338aca bg-clip-text text-transparent"
      >
        Edgeless CA
      </div>
      <span class="text-16px line-height-24px text-#808080">
        {{ t('trustedNepPackageKeyPairService') }}
      </span>
      <div class="bg-$p-surface-600 p-3 rounded-xl w-60%">
        <InputGroup>
          <InputText v-model="q" :placeholder="t('queryPlaceholder')" />
          <InputGroupAddon>
            <Button
              icon="pi pi-search"
              variant="text"
              :loading="queryLoading"
            />
          </InputGroupAddon>
        </InputGroup>
        <KeyViewer v-if="publicKey" :view-key="publicKey" />
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue';
import { computedAsync, refDebounced } from '@vueuse/core';
import { GetPublicKey } from '../api/token';
import { z } from 'zod';
import User from '../components/User.vue';
import KeyViewer from '../components/KeyViewer.vue';
import { t } from '../i18n';

const myKeypairVisible = ref(false);

const q = ref('');
const debouncedQ = refDebounced(q, 500);
const queryLoading = ref(false);

const publicKey = computedAsync(
  async () => {
    if (
      !debouncedQ.value ||
      !z.string().email().safeParse(debouncedQ.value).success
    )
      return undefined;
    queryLoading.value = true;
    try {
      const {
        data: { data },
      } = await GetPublicKey(debouncedQ.value);
      return data;
    } catch (error) {
      return undefined;
    } finally {
      queryLoading.value = false;
    }
  },
  undefined,
  { lazy: true },
);
</script>
