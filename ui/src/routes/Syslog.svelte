<script lang="ts">
  import BACKEND_BASE_URL from "$lib/config";
  import {
    es,
    syslogMsgs,
    connected,
    subscribe,
    unsubscribe,
  } from "$lib/events";

  import { onDestroy, onMount, tick } from "svelte";
  let scrollAreaElement: any;

  async function syslogHandler(event: MessageEvent<any>) {
    const data = event.data;
    if (data !== undefined) {
      $syslogMsgs = [...$syslogMsgs, data];
      scrollToBottom(scrollAreaElement);
      await tick();
    }
    return () => $es.close();
  }

  const connect = () => {
    if (!$connected) {
      es.set(new EventSource(`${BACKEND_BASE_URL}/stream/syslog`));
      subscribe($es);
    }
  };

  onMount(() => {
    connect();
    $es.addEventListener("syslog", syslogHandler, false);
  });

  onDestroy(() => {
    $es.removeEventListener("syslog", syslogHandler, false);
    unsubscribe($es);
    $syslogMsgs.length = 0; // zero out the array after unsubscribing
  });

  const scrollToBottom = async (node: Element) => {
    node.scroll({
      top: node.scrollHeight,
      behavior: "instant",
    });
  };
</script>

<main class="flex flex-col gap-4 items-center">
  <h1 class="text-3xl font-semibold">Syslog</h1>
  <div
    bind:this={scrollAreaElement}
    class=" h-[70vh] w-[80%] rounded-md border bg-muted/100 overflow-auto"
  >
    <div class="p-5">
      {#each $syslogMsgs as msg}
        <p class="font-mono text-nowrap text-sm">{msg}</p>
      {/each}
    </div>
  </div>
</main>
