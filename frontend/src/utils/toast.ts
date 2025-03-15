import { emitter } from './emit';

export function toast(
  severity: 'success' | 'info' | 'warn' | 'error',
  summary: string,
  detail: string,
  life: number,
) {
  emitter.emit('toast', severity, summary, detail, life);
}
