import { writable } from "svelte/store";
import { type Memory } from "./types";
// Connection and EventSource stores
export const es = writable<EventSource>()
export const connected = writable<boolean>(false);

// EventSource message stores
export const timeMsgs = writable<string>("");
export const memoryMsgs = writable<Memory>();
export const syslogMsgs = writable<Array<string>>([]);

// Establishes a connection to a provided EventSource but requires EventHandlers to be set elsewhere.
export function subscribe(es: EventSource) {
  es.onopen = function (_) {
    connected.set(true)
    console.log("Subscribed", es.url)
  }

  es.onerror = function (err) {
    console.log("SSE Error: ", err)
    connected.set(false)
    es.close()
  }
}

export function unsubscribe(es: EventSource) {
  connected.set(false)
  es.close()
  console.log("Unubscribed", es.url)
}
