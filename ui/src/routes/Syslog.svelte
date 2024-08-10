<script lang="ts">
  import BACKEND_BASE_URL from '$lib/config';
  import { es, subscribe, syslogMsgs, unsubscribe } from '$lib/events';
  import LoaderCircle from 'lucide-svelte/icons/loader-circle';
  import { onDestroy, onMount, tick } from 'svelte';
  let scrollAreaElement: Element;

  async function syslogHandler(event: MessageEvent<any>) {
    const data = event.data;
    if (data !== undefined) {
      $syslogMsgs = [...$syslogMsgs, data];
      scrollToBottom(scrollAreaElement);
      await tick();
    }
    return () => $es.close();
  }

  onMount(() => {
    es.set(new EventSource(`${BACKEND_BASE_URL}/stream/syslog`));
    $es.addEventListener('syslog', syslogHandler, false);
    subscribe($es);
  });

  onDestroy(() => {
    $es.removeEventListener('syslog', syslogHandler, false);
    unsubscribe($es);
    $syslogMsgs.length = 0; // zero out the array after unsubscribing
  });

  const scrollToBottom = async (node: Element) => {
    node.scroll({
      top: node.scrollHeight,
      behavior: 'instant',
    });
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
      {#each $syslogMsgs as msg}
        <p class="text-nowrap font-mono text-sm">{msg}</p>
      {/each}
    </div>
  </div>
</main>
