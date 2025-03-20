import { instance, BASE_URL } from './index';
import type { BaseResponse } from './types';

export async function LoginWithGitHub() {
  window.location.href = `${BASE_URL}/api/oauth/login`;
}

export async function LoginWithGitHubCallback(code: string, state: string) {
  return instance.get<
    BaseResponse<{
      name: string;
      email: string;
      token: string;
      tmpOpt?: string;
    }>
  >(`/api/oauth/callback?code=${code}&state=${state}`);
}
