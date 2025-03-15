import { z } from 'zod';

export const EMAIL_VALIDATOR = z
  .string()
  .min(1, { message: '请输入邮箱' })
  .email({ message: '邮箱格式错误' });

export const PASSWORD_SIMPLE_VALIDATOR = z
  .string()
  .min(1, { message: '请输入密码' })
  .min(8, { message: '密码长度至少为8位' });

export const PASSWORD_VALIDATOR = z
  .string()
  .min(1, { message: '请输入密码' })
  .min(8, { message: '密码长度至少为8位' })
  .regex(/[A-Z]/, { message: '密码必须包含大写字母' })
  .regex(/[a-z]/, { message: '密码必须包含小写字母' })
  .regex(/[0-9]/, { message: '密码必须包含数字' })
  .regex(/[!@#$%^&*(),.?":{}|<>]/, { message: '密码必须包含特殊字符' });
