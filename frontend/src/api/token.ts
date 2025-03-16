import { instance } from './index';
import type { BaseResponse } from './types';

export async function GetPublicKey(name: string) {
  return instance.get<BaseResponse<string>>('/api/token/public', {
    params: {
      name,
    },
  });
}

export async function GetMyKeypair() {
  return instance.get<
    BaseResponse<{ public_key: string; private_key: string }>
  >('/api/token/keypair');
}
