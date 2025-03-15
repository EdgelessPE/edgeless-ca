import { instance, BASE_URL } from './index';

export async function LoginWithGitHub() {
  window.open(`${BASE_URL}/api/oauth/login`);
}

export async function LoginWithGitHubCallback(code: string) {
  return instance.get(`/api/oauth/callback?code=${code}&state=github-login`);
}
