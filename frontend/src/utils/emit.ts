/* eslint-disable @typescript-eslint/no-explicit-any */

class Emitter<Keys extends string> {
  private callMap: Map<Keys, ((...args: any[]) => void)[]> = new Map();

  on(event: Keys, callback: (...args: any[]) => void) {
    const callbacks = this.callMap.get(event) || [];
    callbacks.push(callback);
    this.callMap.set(event, callbacks);
  }

  emit(event: Keys, ...args: any[]) {
    const callbacks = this.callMap.get(event) || [];
    callbacks.forEach((callback) => callback(...args));
  }

  off(event: Keys, callback: (...args: any[]) => void) {
    const callbacks = this.callMap.get(event) || [];
    callbacks.filter((cb) => cb !== callback);
    this.callMap.set(event, callbacks);
  }
}

export const emitter = new Emitter<'toast'>();
