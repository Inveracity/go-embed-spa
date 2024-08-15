<script lang="ts">
  import { syslogMsgs } from '$lib/events';
  import { syslogMessages } from '$lib/store';
  import { Streamer } from '$lib/stream';
  import LoaderCircle from 'lucide-svelte/icons/loader-circle';
  import { onDestroy, onMount } from 'svelte';
  let scrollAreaElement: Element;

  const s = new Streamer('/stream/syslog', syslogMessages);

  onMount(async () => {
    s.connect();
  });

  onDestroy(() => {
    s.disconnect();
    $syslogMessages.length = 0;
  });

  $: if ($syslogMessages) {
    scrollToBottom(scrollAreaElement);
  }

  const scrollToBottom = async (node: Element) => {
    if (node) {
      node.scroll({
        top: node.scrollHeight,
        behavior: 'smooth',
      });
    }
  };
</script>

<main class="flex flex-col items-center gap-4">
  <h1 class="text-3xl font-semibold">Syslog</h1>
  {#if $syslogMsgs.length <= 1}
    <LoaderCircle class="animate-spin" />
  {/if}
  <div
    bind:this={scrollAreaElement}
    class=" bg-muted/100 h-[70vh] w-[80%] overflow-auto rounded-md border">
    <div class="p-5">
      {#each $syslogMessages as msg}
        <p class="text-nowrap font-mono text-sm">{msg.msg}</p>
      {/each}
    </div>
  </div>
</main>
