import { writable } from "svelte/store";

export const state = writable<string>("");

export function subscribe() {
  const es = new EventSource("http://localhost:3000/stream");

  // Create a new websocket
  es.onmessage = async function (event) {
    const data = event.data;
    state.set(data);
    return () => es.close();
  };
}
