<template>
  <LoginLayout>
    <div>
      <h2 class="text-2xl font-bold">{{ t('login') }}</h2>
      <span>{{ t('useGitHubToRegisterOrLogin') }}</span>
    </div>
    <Button
      class="w-full min-h-33px"
      :label="t('useGitHubAccount')"
      icon="pi pi-github"
      :size="showEmailLogin ? 'small' : 'large'"
      :loading="githubLoading"
      @click="
        githubLoading = true;
        LoginWithGitHub();
      "
    />
    <Divider align="center" type="dotted" class="!mb-0">
      <b>或</b>
    </Divider>
    <template v-if="showEmailLogin">
      <span>{{ t('useEmailToLogin') }}</span>
      <Form
        v-slot="$form"
        :resolver="resolver"
        :initialValues="initialValues"
        @submit="onFormSubmit"
        class="flex flex-col gap-2 w-full"
      >
        <FloatLabel variant="on">
          <InputText name="email" type="text" size="small" fluid autofocus />
          <label for="on_label">{{ t('email') }}</label>
        </FloatLabel>
        <Message
          v-if="$form.email?.invalid"
          severity="error"
          size="small"
          variant="simple"
          >{{ $form.email.error?.message }}</Message
        >
        <FloatLabel variant="on">
          <Password
            name="password"
            toggleMask
            size="small"
            fluid
            :feedback="false"
          />
          <label for="on_label">{{ t('password') }}</label>
        </FloatLabel>
        <Message
          v-if="$form.password?.invalid"
          severity="error"
          size="small"
          variant="simple"
          >{{ $form.password.error?.message }}</Message
        >
        <div style="height: 32px">
          <Button
            class="w-full"
            type="submit"
            :label="t('login')"
            :loading="emailLoading"
            :disabled="
              !$form.email?.value ||
              !$form.password?.value ||
              $form.email?.invalid ||
              $form.password?.invalid
            "
          />
        </div>
      </Form>
    </template>
    <template v-else>
      <span>{{ t('useEmailToLogin') }}</span>
      <Button
        class="w-full"
        :label="t('useEmailAccount')"
        icon="pi pi-envelope"
        :size="showEmailLogin ? undefined : 'large'"
        @click="showEmailLogin = true"
      />
    </template>
    <div class="text-12px cursor-pointer mt-2" @click="router.push('/recover')">
      {{ t('forgotPassword') }}
    </div>
  </LoginLayout>
</template>

<script setup lang="ts">
import { ref } from 'vue';
import { Login } from '../api/auth';
import { LoginWithGitHub } from '../api/oauth';
import { z } from 'zod';
import { zodResolver } from '@primevue/forms/resolvers/zod';
import { EMAIL_VALIDATOR, PASSWORD_SIMPLE_VALIDATOR } from '../utils/validator';
import type { FormSubmitEvent } from '@primevue/forms';
import LoginLayout from '../layouts/LoginLayout.vue';
import { setLoginInfo } from '../utils/login';
import { useRouter } from 'vue-router';
import { t } from '../i18n';

const router = useRouter();

const githubLoading = ref(false);
const emailLoading = ref(false);
const showEmailLogin = ref(false);
const initialValues = ref({
  email: '',
  password: '',
});

const resolver = ref(
  zodResolver(
    z.object({
      email: EMAIL_VALIDATOR,
      password: PASSWORD_SIMPLE_VALIDATOR,
    }),
  ),
);

const onFormSubmit = async ({ valid, values }: FormSubmitEvent) => {
  if (valid) {
    emailLoading.value = true;
    const {
      data: { data },
    } = await Login(values.email, values.password).finally(() => {
      emailLoading.value = false;
    });
    setLoginInfo(data);
    router.push('/');
  }
};
</script>
