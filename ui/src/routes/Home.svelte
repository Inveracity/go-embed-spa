<script lang="ts">
  import BACKEND_BASE_URL from '$lib/config';
  import { apiData } from '$lib/store';
  import { messages, Streamer } from '$lib/stream';
  import type { Hello } from '$lib/types';
  import { onDestroy, onMount } from 'svelte';

  const s = new Streamer('/stream/time', messages);

  onMount(async () => {
    const response = await fetch(`${BACKEND_BASE_URL}/api`);
    if (response.ok) {
      const data: Hello = await response.json();
      apiData.set(data);
    }

    s.connect();
  });

  onDestroy(() => s.disconnect());
</script>

<div class="p-[10vh] text-center">
  <h1>Hello {$apiData.hello}</h1>
  {$messages.msg}
</div>
