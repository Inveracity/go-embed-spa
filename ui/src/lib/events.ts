import { writable } from "svelte/store";

export const messages = writable<string>("");
export const es = writable<EventSource>()
export const connected = writable<boolean>(false);
export const memoryMsgs = writable<Memory>();
export const syslogMsgs = writable<string>();

export function subscribe(es: EventSource) {
  es.onopen = function (_) {
    connected.set(true)
    console.log("Subscribed")
  }
  es.onmessage = async function (event) {
    const data = event.data;
    messages.set(data);
    return () => es.close();
  };
  es.addEventListener("memory", async function (event) {
    const data: Memory = JSON.parse(event.data);
    memoryMsgs.set(data);
    return () => es.close();
  })

  es.onerror = function (_) {
    connected.set(false)
    messages.set("")
    es.close()
  }
}

export function unsubscribe(es: EventSource) {
  connected.set(false)
  es.close()
}

export interface Memory {
  total: number
  available: number
  used: number
  usedPercent: number
  free: number
  active: number
  inactive: number
  wired: number
  laundry: number
  buffers: number
  cached: number
  writeBack: number
  dirty: number
  writeBackTmp: number
  shared: number
  slab: number
  sreclaimable: number
  sunreclaim: number
  pageTables: number
  swapCached: number
  commitLimit: number
  committedAS: number
  highTotal: number
  highFree: number
  lowTotal: number
  lowFree: number
  swapTotal: number
  swapFree: number
  mapped: number
  vmallocTotal: number
  vmallocUsed: number
  vmallocChunk: number
  hugePagesTotal: number
  hugePagesFree: number
  hugePagesRsvd: number
  hugePagesSurp: number
  hugePageSize: number
  anonHugePages: number
}
