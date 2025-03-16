<template>
  <div class="h-full w-full">
    <User class="absolute top-5 right-4" />
    <div class="flex flex-col gap-4 items-center w-full mt-40%">
      <div
        class="text-48px font-bold bg-gradient-to-b from-#134597 to-#338aca bg-clip-text text-transparent"
      >
        Edgeless CA
      </div>
      <span class="text-16px line-height-24px text-#808080">
        可信 Nep 包密钥对服务
      </span>
      <div class="bg-$p-surface-800 p-3 rounded-2xl w-60%">
        <InputGroup>
          <InputText
            v-model="q"
            placeholder="使用邮箱或用户名查询其他作者的公钥"
          />
          <InputGroupAddon>
            <Button
              icon="pi pi-search"
              variant="text"
              :loading="queryLoading"
            />
          </InputGroupAddon>
        </InputGroup>
        <div
          v-if="publicKey"
          class="flex items-center justify-between mt-2 rounded-xl bg-#1f1f1f p-4"
        >
          <code>{{ publicKey }}</code>
          <Button
            icon="pi pi-copy"
            variant="text"
            size="small"
            @click="copyPublicKey"
          />
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue';
import { computedAsync, refDebounced } from '@vueuse/core';
import { GetPublicKey } from '../api/token';
import { z } from 'zod';
import { toast } from '../utils/toast';

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

const copyPublicKey = () => {
  if (!publicKey.value) return;
  navigator.clipboard.writeText(publicKey.value);
  toast('success', '复制成功', '', 3000);
};
</script>
