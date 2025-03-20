<template>
  <LoginLayout>
    <div
      class="flex flex-col items-center justify-center gap-36px h-full w-full"
    >
      <div class="pi pi-github" style="font-size: 2.5rem" />
      <div class="flex items-center justify-center gap-4">
        <div class="text-2xl">
          <template v-if="!failed"> {{ t('authenticating') }} </template>
          <template v-else>
            {{ t('authenticationFailed') }}
            <RouterLink to="/login"> {{ t('reLogin') }} </RouterLink>
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
import { setLoginInfo } from '../utils/login';
import { useRouter } from 'vue-router';
import { t } from '../i18n';

const route = useRoute();
const router = useRouter();

const failed = ref(false);

onMounted(async () => {
  const code = route.query.code;
  const state = route.query.state;
  if (code && state) {
    try {
      const {
        data: { data },
      } = await LoginWithGitHubCallback(code as string, state as string);
      setLoginInfo(data);
      const opt = data.tmpOpt;
      if (opt) {
        router.push(`/register?opt=${opt}`);
      } else {
        router.push('/');
      }
    } catch (error) {
      failed.value = true;
    }
  }
});
</script>
