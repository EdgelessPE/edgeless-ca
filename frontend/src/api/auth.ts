import { instance } from './index';
import { blake3 } from '../utils/bake3';

export async function Login(email: string, password: string) {
  return instance.post('/api/auth/login', {
    email,
    pwdHash: await blake3(password),
  });
}
