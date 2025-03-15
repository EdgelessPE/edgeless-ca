import { instance } from './index';
import { blake3 } from '../utils/bake3';
import type { BaseResponse } from './types';

export async function Login(email: string, password: string) {
  return instance.post<
    BaseResponse<{ name: string; email: string; token: string }>
  >('/api/auth/login', {
    email,
    pwdHash: await blake3(password),
  });
}
