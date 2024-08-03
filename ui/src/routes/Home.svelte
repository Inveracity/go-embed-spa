<script lang="ts">
  import "../app.css";
  import { onDestroy, onMount } from "svelte";
  import {
    subscribe,
    unsubscribe,
    es,
    messages,
    memoryMsgs,
    connected,
    syslogMsgs,
  } from "$lib/events";
  import Button from "$lib/components/ui/button/button.svelte";
  import Separator from "$lib/components/ui/separator/separator.svelte";
  import BACKEND_BASE_URL from "$lib/config";
  import Wifi from "lucide-svelte/icons/wifi";
  import Wifioff from "lucide-svelte/icons/wifi-off";
  const connect = () => {
    if (!$connected) {
      es.set(new EventSource(`${BACKEND_BASE_URL}/stream`));
      subscribe($es);
    }
  };
  onMount(() => {
    connect();
  });
  // Automatically unsubscribe when leaving the page
  onDestroy(() => {
    unsubscribe($es);
  });

  // Handle user manually unsubscribing from server side events
  function disconnect() {
    unsubscribe($es);
  }
</script>

<div class="flex flex-col">
  <div class="flex flex-col space-y-2 justify-center items-center">
    {#if $connected}
      <Wifi />
      <Button class="w-36" on:click={disconnect}>Disconnect</Button>
    {:else}
      <Wifioff />
      <Button class="w-36" on:click={connect}>Connect</Button>
    {/if}
  </div>
  <div class="my-4">
    <Separator />
  </div>
  <div class="items-center">
    <div class="flex items-center justify-center space-x-4 text-lg h-[100px]">
      <Separator orientation="vertical" />
      <div class="flex flex-col items-center w-96">
        <p>Time</p>
        <p>{$messages}</p>
      </div>
      <Separator orientation="vertical" />
      <div class="flex flex-col items-center w-96">
        <p>Memory</p>
        <p>{$memoryMsgs?.usedPercent || ""}%</p>
      </div>
      <Separator orientation="vertical" />
    </div>
    <div>
      {$syslogMsgs}
    </div>
  </div>
</div>
