import axios, { AxiosError } from 'axios';
import type { BaseResponse } from './types';
import { toast } from '../utils/toast';
import { router } from '../router';
import { getToken } from '../utils/login';

export const BASE_URL = import.meta.env.VITE_BASE_URL;

export const instance = axios.create({
  baseURL: BASE_URL,
});

instance.interceptors.request.use(
  (config) => {
    // 携带认证信息
    const token = getToken();
    if (token) {
      config.headers.Authorization = token;
    }
    return config;
  },
  (error) => {
    return Promise.reject(error);
  },
);

instance.interceptors.response.use(
  (response) => {
    return response;
  },
  (error: AxiosError<BaseResponse<unknown>>) => {
    console.error(error);
    if (error.response) {
      toast('error', '错误', error.response.data.msg, 3000);
    } else {
      toast('error', '错误', error.message, 3000);
    }
    // 如果是401，则跳转到登录页
    if (error.response?.status === 401) {
      router.push('/login');
    }
    return Promise.reject(error);
  },
);
