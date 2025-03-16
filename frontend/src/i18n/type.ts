import type { zh } from './dictionaries/zh';

export type LangKeys = 'en' | 'zh';
export type Dictionary = Record<keyof typeof zh, string>;
