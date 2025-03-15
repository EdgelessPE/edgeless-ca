import { createBLAKE3 } from 'hash-wasm';

export async function blake3(password: string) {
  const hash = await createBLAKE3();
  return hash.update(password).digest('hex');
}
