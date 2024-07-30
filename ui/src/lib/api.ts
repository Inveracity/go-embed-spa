export type Hello = {
  hello: string
}

export async function load(): Promise<Hello> {
  const res = await fetch(`${import.meta.env.THING_BACKEND_URL}/api`);

  if (!res.ok) {
    return { "hello": "nothing" }
  }

  const item = await res.json();
  console.log(item)
  return item;

}
