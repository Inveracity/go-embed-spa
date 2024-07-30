export type Hello = {
  hello: string
}

export async function load(): Promise<Hello> {
  const res = await fetch(`/api`);

  if (!res.ok) {
    return { "hello": "nothing" }
  }

  const item = await res.json();
  console.log(item)
  return item;

}
