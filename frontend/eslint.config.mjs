import js from '@eslint/js';
import globals from 'globals';
import ts from 'typescript-eslint';
import prettier from 'eslint-plugin-prettier/recommended';

export default [
  { languageOptions: { globals: globals.browser } },
  js.configs.recommended,
  ...ts.configs.recommended,
  { ignores: ['dist/'] },
  prettier,
];
