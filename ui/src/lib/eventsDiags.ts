import { writable } from "svelte/store";

export const diagMessages = writable<Memory>();
export const diagEs = writable<EventSource>()
export const diagConnected = writable<boolean>(false);

export function diagSubscribe(es: EventSource) {
  diagConnected.set(true)
  console.log("Subscribed to diagnostics")
  es.addEventListener("diagnostics", async function (event) {
    const data: Memory = JSON.parse(event.data);
    diagMessages.set(data);
    return () => es.close();
  })
}

export function diagUnsubscribe(es: EventSource) {
  diagConnected.set(false)
  es.close()
  console.log("Unsubscribed from diagnostics")
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
