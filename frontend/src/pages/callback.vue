<template>
  <LoginLayout>
    <div
      class="flex flex-col items-center justify-center gap-36px h-full w-full"
    >
      <div class="pi pi-github" style="font-size: 2.5rem" />
      <div class="flex items-center justify-center gap-4">
        <div class="text-2xl">
          <template v-if="!failed"> 正在认证中 </template>
          <template v-else>
            认证失败，请尝试
            <RouterLink to="/login">重新登录</RouterLink>
          </template>
        </div>
        <div v-if="!failed" class="pi pi-spin pi-spinner" />
      </div>
    </div>
  </LoginLayout>
</template>

<script setup lang="ts">
import { onMounted, ref } from 'vue';
import { useRoute } from 'vue-router';
import { LoginWithGitHubCallback } from '../api/oauth';
import LoginLayout from '../layouts/LoginLayout.vue';
import { setToken } from '../utils/token';
import { useRouter } from 'vue-router';

const route = useRoute();
const router = useRouter();

const failed = ref(false);

onMounted(async () => {
  const code = route.query.code;
  const state = route.query.state;
  if (code && state) {
    try {
      const {
        data: {
          data: { token },
        },
      } = await LoginWithGitHubCallback(code as string, state as string);
      setToken(token);
      router.push('/');
    } catch (error) {
      failed.value = true;
    }
  }
});
</script>
