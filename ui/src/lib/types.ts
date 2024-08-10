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

export type Hello = {
  hello: string
}
