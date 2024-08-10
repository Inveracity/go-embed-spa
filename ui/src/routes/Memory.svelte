<script lang="ts">
  import BACKEND_BASE_URL from '$lib/config';
  import {
    connected,
    es,
    memoryMsgs,
    subscribe,
    unsubscribe,
  } from '$lib/events';
  import { type Memory } from '$lib/types';
  import { onDestroy, onMount } from 'svelte';

  const memHandler = async (event: MessageEvent<any>) => {
    if (event.data) {
      const data: Memory = JSON.parse(event.data);
      memoryMsgs.set(data);
      return () => $es.close();
    }
  };

  const connect = async () => {
    if (!$connected) {
      es.set(new EventSource(`${BACKEND_BASE_URL}/stream/memory`));
      subscribe($es);
    }
  };

  onMount(() => {
    connect();
    $es.addEventListener('memory', memHandler, true);
  });

  onDestroy(() => {
    $es.removeEventListener('memory', memHandler, true);
    unsubscribe($es);
  });

  let mem;
  $: mem = (Math.round(($memoryMsgs?.usedPercent || 0) * 100) / 100).toFixed(2);
</script>

<main class="flex flex-1 flex-col gap-4 p-4 md:gap-8 md:p-10">
  <div class="mx-auto grid w-full max-w-6xl justify-center gap-2 text-center">
    <h1 class="text-3xl font-semibold">Memory</h1>
    <p class="text-mono text-xl">{mem}%</p>
  </div>
</main>
