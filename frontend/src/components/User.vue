<template>
  <div>
    <template v-if="userInfo">
      <code
        class="px-3 py-1.5 rounded-lg bg-gray-100 dark:bg-#18181b text-gray-700 dark:text-gray-300 hover:bg-gray-200 dark:hover:bg-#27272a transition-colors duration-200"
        aria-haspopup="true"
        aria-controls="overlay_menu"
        @click="toggle"
      >
        {{ text }}
      </code>
      <Menu ref="op" id="overlay_menu" :model="items" :popup="true" />
    </template>
    <template v-else>
      <Button
        icon="pi pi-user"
        label="登录"
        severity="secondary"
        size="small"
        @click="router.push('/login')"
      />
    </template>
  </div>
</template>

<script setup lang="ts">
import { computed, ref } from 'vue';
import { useLoginInfo, removeLoginInfo } from '../utils/login';
import { useRouter } from 'vue-router';

const emits = defineEmits(['view-my-key']);

const op = ref<any>(null);
const userInfo = useLoginInfo();
const router = useRouter();

const text = computed(() => {
  if (!userInfo.value) {
    return '--';
  }
  return `${userInfo.value?.name} <${userInfo.value?.email}>`;
});

const items = [
  {
    label: '查看我的密钥对',
    icon: 'pi pi-key',
    command: () => {
      emits('view-my-key');
    },
  },
  {
    label: '注销',
    icon: 'pi pi-sign-out',
    command: () => {
      removeLoginInfo();
      router.push('/login');
    },
  },
];

const toggle = (event: MouseEvent) => {
  op.value?.toggle(event);
};
</script>
