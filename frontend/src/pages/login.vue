<template>
  <div class="flex flex-col items-center justify-center h-screen bg-black">
    <div
      class="flex items-center justify-between h-lg w-80% bg-#1f1f1f rounded-3xl"
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
          可信 Nep 包密钥对服务
        </span>
      </div>
      <div class="w-40% h-full flex flex-col gap-4 rounded-3xl bg-#18181b p-12">
        <div>
          <h2 class="text-2xl font-bold">登录</h2>
          <span>使用 GitHub 注册或登录</span>
        </div>
        <Button
          class="w-full"
          label="使用 GitHub 账号"
          icon="pi pi-github"
          :size="showEmailLogin ? 'small' : 'large'"
          @click="LoginWithGitHub"
        />
        <Divider align="center" type="dotted" class="!mb-0">
          <b>或</b>
        </Divider>
        <template v-if="showEmailLogin">
          <span>使用邮箱登录</span>
          <div class="flex flex-col gap-2 w-full">
            <FloatLabel variant="on">
              <InputText
                v-model="email"
                name="email"
                type="text"
                size="small"
                fluid
              />
              <label for="on_label">邮箱</label>
            </FloatLabel>
            <FloatLabel variant="on">
              <Password
                v-model="password"
                toggleMask
                size="small"
                fluid
                :feedback="false"
              />
              <label for="on_label">密码</label>
            </FloatLabel>
          </div>
          <div style="height: 32px">
            <Button class="w-full" label="登录" @click="onSubmit" />
          </div>
        </template>
        <template v-else>
          <span>使用邮箱登录</span>
          <Button
            class="w-full"
            label="使用邮箱账号"
            icon="pi pi-envelope"
            :size="showEmailLogin ? undefined : 'large'"
            @click="showEmailLogin = true"
          />
        </template>
        <div class="text-12px cursor-pointer mt-2" href="#">忘记密码？</div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue';
import { Login } from '../api/auth';
import { LoginWithGitHub } from '../api/oauth';

const showEmailLogin = ref(false);

const email = ref('');
const password = ref('');

const onSubmit = async () => {
  const res = await Login(email.value, password.value);
  console.log(res);
};
</script>
