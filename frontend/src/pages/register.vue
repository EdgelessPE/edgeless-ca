<template>
  <LoginLayout>
    <div>
      <div class="flex gap-2 items-center">
        <h2 class="text-2xl font-bold">{{ t('welcomeToEdgelessCa') }}</h2>
      </div>
      <span>{{ t('configYourFirstPassword') }}</span>

      <Form
        v-slot="$form"
        :resolver="resolver"
        :initialValues="initialValues"
        @submit="onFormSubmit"
        class="mt-20px flex flex-col gap-2 w-full"
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
          :label="t('set')"
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
import { Recover } from '../api/auth';
import { router } from '../router';
import { PASSWORD_VALIDATOR } from '../utils/validator';
import { toast } from '../utils/toast';
import { t } from '../i18n';
import { useRoute } from 'vue-router';
import { useLoginInfo } from '../utils/login';

const route = useRoute();
const userInfo = useLoginInfo();

const submitLoading = ref(false);
const initialValues = ref({
  password: '',
});

const resolver = ref(
  zodResolver(
    z.object({
      password: PASSWORD_VALIDATOR,
    }),
  ),
);

const onFormSubmit = async ({ valid, values }: FormSubmitEvent) => {
  if (valid && userInfo.value?.email) {
    submitLoading.value = true;
    await Recover(
      userInfo.value?.email as string,
      route.query.opt as string,
      values.password,
    ).finally(() => {
      submitLoading.value = false;
    });
    toast('success', t('success'), t('passwordSetSuccess'), 3000);
    router.push('/');
  }
};
</script>
