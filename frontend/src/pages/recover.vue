<template>
  <LoginLayout>
    <div>
      <div class="flex gap-2 items-center">
        <BackBtn />
        <h2 class="text-2xl font-bold">{{ t('resetPassword') }}</h2>
      </div>
      <span>{{ t('verifyEmailToContinue') }}</span>

      <Form
        v-slot="$form"
        :resolver="resolver"
        :initialValues="initialValues"
        @submit="onFormSubmit"
        class="mt-20px flex flex-col gap-2 w-full"
      >
        <InputGroup>
          <FloatLabel variant="on">
            <InputText
              name="email"
              type="text"
              fluid
              autofocus
              v-model="email"
            />
            <label for="on_label">{{ t('email') }}</label>
          </FloatLabel>
          <Button v-if="sendLoading" icon="pi pi-spin pi-spinner" disabled />
          <Button
            v-else
            :label="countdown ? `${countdown}s` : t('verify')"
            class="whitespace-nowrap"
            :disabled="
              Boolean(countdown) || !$form.email?.value || $form.email?.invalid
            "
            @click="onSendEmail"
          />
        </InputGroup>
        <Message
          v-if="$form.email?.invalid"
          severity="error"
          variant="simple"
          size="small"
          >{{ $form.email.error?.message }}</Message
        >
        <FloatLabel variant="on">
          <InputText name="code" type="text" fluid :feedback="false" />
          <label for="on_label">{{ t('verificationCode') }}</label>
        </FloatLabel>
        <Message
          v-if="$form.code?.invalid"
          severity="error"
          variant="simple"
          size="small"
          >{{ $form.code.error?.message }}</Message
        >
        <FloatLabel variant="on" class="mt-16px">
          <Password name="password" toggleMask fluid :feedback="false" />
          <label for="on_label">{{ t('newPassword') }}</label>
        </FloatLabel>
        <Message
          v-if="$form.password?.invalid"
          severity="error"
          variant="simple"
          size="small"
          >{{ $form.password.error?.message }}</Message
        >
        <Button
          class="w-full mt-12px"
          type="submit"
          :label="t('reset')"
          :loading="submitLoading"
        />
      </Form>
    </div>
  </LoginLayout>
</template>

<script setup lang="ts">
import { zodResolver } from '@primevue/forms/resolvers/zod';
import LoginLayout from '../layouts/LoginLayout.vue';
import { ref } from 'vue';
import type { FormSubmitEvent } from '@primevue/forms';
import { z } from 'zod';
import { Recover, SendEmail } from '../api/auth';
import { router } from '../router';
import { EMAIL_VALIDATOR, PASSWORD_VALIDATOR } from '../utils/validator';
import { useIntervalFn, useLocalStorage } from '@vueuse/core';
import { toast } from '../utils/toast';
import { t } from '../i18n';
const lastSendEmailTime = useLocalStorage('lastSendEmailTime', 0);

const sendLoading = ref(false);
const submitLoading = ref(false);
const email = ref('');

const getCountdown = () => {
  const passed = Date.now() - lastSendEmailTime.value;
  const oneMinute = 1 * 60 * 1000;
  if (passed < oneMinute) {
    return Math.ceil((oneMinute - passed) / 1000);
  }
  return undefined;
};

const countdown = ref<number | undefined>(getCountdown());

const { pause, resume } = useIntervalFn(() => {
  countdown.value = getCountdown();
  if (!countdown.value) {
    pause();
  }
}, 1000);

const onSendEmail = async () => {
  sendLoading.value = true;
  await SendEmail(email.value).finally(() => {
    sendLoading.value = false;
  });
  lastSendEmailTime.value = Date.now();
  resume();
  countdown.value = 60;
};

const initialValues = ref({
  email: '',
  password: '',
});

const resolver = ref(
  zodResolver(
    z.object({
      email: EMAIL_VALIDATOR,
      code: z
        .string()
        .min(6, t('verificationCodeLength'))
        .max(6, t('verificationCodeLength')),
      password: PASSWORD_VALIDATOR,
    }),
  ),
);

const onFormSubmit = async ({ valid, values }: FormSubmitEvent) => {
  if (valid) {
    submitLoading.value = true;
    await Recover(values.email, values.code, values.password).finally(() => {
      submitLoading.value = false;
    });
    toast('success', t('success'), t('passwordResetSuccess'), 3000);
    router.push('/login');
  }
};
</script>
