import { writable } from "svelte/store";

export const messages = writable<string>("");
export const es = writable<EventSource>()
export const connected = writable<boolean>(false);

export function subscribe(es: EventSource) {
  connected.set(true)
  console.log("Subscribed")
  es.onmessage = async function (event) {
    const data = event.data;
    messages.set(data);
    return () => es.close();
  };
}

export function unsubscribe(es: EventSource) {
  connected.set(false)
  es.close()
  console.log("Unsubscribed")
}
