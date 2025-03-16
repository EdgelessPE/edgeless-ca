import axios from 'axios';
import type { BaseResponse } from './types';

export async function getToken(name: string) {
  return axios.get<BaseResponse<string>>('/api/token/public', {
    params: {
      name,
    },
  });
}

export async function getMyKeypair() {
  return axios.get<BaseResponse<{ public_key: string; private_key: string }>>(
    '/api/token/keypair',
  );
}
