<script lang="ts">
  import BACKEND_BASE_URL from '$lib/config';
  import { connected, es, subscribe, timeMsgs, unsubscribe } from '$lib/events';
  import { onDestroy, onMount } from 'svelte';

  const timeHandler = async (event: MessageEvent<any>) => {
    if (event.data) {
      const data = event.data;
      timeMsgs.set(data);
      return () => $es.close();
    }
  };

  const connect = async () => {
    if (!$connected) {
      es.set(new EventSource(`${BACKEND_BASE_URL}/stream/time`));
      subscribe($es);
    }
  };

  onMount(() => {
    connect();
    $es.addEventListener('time', timeHandler, true);
  });

  onDestroy(() => {
    $es.removeEventListener('time', timeHandler, true);
    unsubscribe($es);
  });
</script>

<p class="text-mono text-xl">{$timeMsgs}</p>
