import BACKEND_BASE_URL from "./config"
import type { Hello } from "./types";


export async function load(): Promise<Hello> {
  const res = await fetch(`${BACKEND_BASE_URL}/api`);

  if (!res.ok) {
    return { "hello": "nothing" }
  }

  return await res.json();
}
