import BACKEND_BASE_URL from "./config"
export type Hello = {
  hello: string
}

export async function load(): Promise<Hello> {
  const res = await fetch(`${BACKEND_BASE_URL}/api`);

  if (!res.ok) {
    return { "hello": "nothing" }
  }

  return await res.json();
}
