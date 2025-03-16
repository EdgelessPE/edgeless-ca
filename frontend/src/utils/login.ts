import { useLocalStorage } from '@vueuse/core';

const TOKEN_KEY = 'token';
const USER_INFO_KEY = 'userInfo';

const userInfo = useLocalStorage<{
  name: string;
  email: string;
} | null>(USER_INFO_KEY, null, {
  serializer: {
    read: (v: string) => (v ? JSON.parse(v) : null),
    write: (v: unknown) => JSON.stringify(v),
  },
});

export function getToken() {
  return localStorage.getItem(TOKEN_KEY);
}

export function useLoginInfo() {
  return userInfo;
}

export function setLoginInfo({
  name,
  email,
  token,
}: {
  name: string;
  email: string;
  token: string;
}) {
  localStorage.setItem(TOKEN_KEY, token);
  userInfo.value = {
    name,
    email,
  };
}

export function removeLoginInfo() {
  localStorage.removeItem(TOKEN_KEY);
  userInfo.value = null;
}
