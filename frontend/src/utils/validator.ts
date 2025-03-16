import { z } from 'zod';
import { t } from '../i18n';
export const EMAIL_VALIDATOR = z
  .string()
  .min(1, { message: t('emailRequired') })
  .email({ message: t('emailInvalid') });

export const PASSWORD_SIMPLE_VALIDATOR = z
  .string()
  .min(1, { message: t('passwordRequired') })
  .min(8, { message: t('passwordLengthAtLeast8') });

export const PASSWORD_VALIDATOR = z
  .string()
  .min(1, { message: t('passwordRequired') })
  .min(8, { message: t('passwordLengthAtLeast8') })
  .regex(/[A-Z]/, { message: t('passwordMustContainUppercase') })
  .regex(/[a-z]/, { message: t('passwordMustContainLowercase') })
  .regex(/[0-9]/, { message: t('passwordMustContainNumber') })
  .regex(/[!@#$%^&*(),.?":{}|<>]/, {
    message: t('passwordMustContainSpecialCharacter'),
  });
