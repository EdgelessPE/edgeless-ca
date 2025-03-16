import { zh } from './dictionaries/zh';
import { en } from './dictionaries/en';
import type { Dictionary, LangKeys } from './type';

const LANG_KEY = 'i18n-key';
let cachedKey: LangKeys | null = null;

export function changeLanguage(key: LangKeys) {
  localStorage.setItem(LANG_KEY, key);
  window.location.reload();
}

export function getLanguage(): LangKeys {
  const browserLang = navigator.language.toLowerCase();
  const isChinese = browserLang.includes('zh');
  const key =
    (localStorage.getItem(LANG_KEY) as LangKeys) || (isChinese ? 'zh' : 'en');
  cachedKey = key;
  return key;
}

export function t(key: keyof Dictionary) {
  if (!cachedKey) {
    cachedKey = getLanguage();
  }
  const dictionary: Dictionary = cachedKey === 'zh' ? zh : en;
  return dictionary[key];
}
