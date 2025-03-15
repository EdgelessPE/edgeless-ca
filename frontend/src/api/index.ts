import axios, { AxiosError } from 'axios';
import { emitter } from '../utils/emit';
import type { BaseResponse } from './types';

export const BASE_URL = import.meta.env.VITE_BASE_URL;

export const instance = axios.create({
  baseURL: BASE_URL,
});

instance.interceptors.request.use(
  (config) => {
    // 携带认证信息
    const token = sessionStorage.getItem('token');
    if (token) {
      config.headers.Authorization = `Bearer ${token}`;
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
      emitter.emit('toast', 'error', '错误', error.response.data.msg, 3000);
    } else {
      emitter.emit('toast', 'error', '错误', error.message, 3000);
    }
    return Promise.reject(error);
  },
);
