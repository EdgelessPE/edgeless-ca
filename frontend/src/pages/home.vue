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
      <div class="bg-white dark:bg-#18181b p-4 rounded-3xl w-60%">
        <InputGroup>
          <InputText placeholder="使用邮箱或用户名查询其他作者的公钥" />
          <InputGroupAddon>
            <Button icon="pi pi-search" variant="text" />
          </InputGroupAddon>
        </InputGroup>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { onMounted } from 'vue';
import { getToken } from '../utils/login';
import { useRouter } from 'vue-router';

const router = useRouter();

onMounted(() => {
  const token = getToken();
  if (!token) {
    router.push('/login');
  }
});
</script>
