<script lang="ts">
  import { timeMessages } from '$lib/store';
  import { Streamer } from '$lib/stream';
  import { onDestroy, onMount } from 'svelte';

  const s = new Streamer('/stream/time', timeMessages);

  onMount(async () => {
    s.connect();
  });

  onDestroy(() => {
    s.disconnect();
  });
</script>

<p class="text-mono text-xl">{$timeMessages.msg}</p>
