<script lang="ts">
  import "../app.css";
  import { onDestroy } from "svelte";
  import {
    subscribe,
    unsubscribe,
    es,
    messages,
    memoryMsgs,
    connected,
  } from "$lib/events";
  import Button from "$lib/components/ui/button/button.svelte";
  import Separator from "$lib/components/ui/separator/separator.svelte";
  import BACKEND_BASE_URL from "$lib/config";
  import Wifi from "lucide-svelte/icons/wifi";
  import Wifioff from "lucide-svelte/icons/wifi-off";

  // Automatically unsubscribe when leaving the page
  onDestroy(() => {
    unsubscribe($es);
  });

  // Handle user manually unsubscribing from server side events
  function closeHandler() {
    unsubscribe($es);
  }

  // Handle user subscribing to server side events
  function openHandler() {
    if (!$connected) {
      es.set(new EventSource(`${BACKEND_BASE_URL}/stream`));
      subscribe($es);
    }
  }
</script>

<div class="flex flex-col">
  <div class="flex flex-col space-y-2 justify-center items-center">
    {#if $connected}
      <Wifi />
      <Button class="w-36" on:click={closeHandler}>Disconnect</Button>
    {:else}
      <Wifioff />
      <Button class="w-36" on:click={openHandler}>Connect</Button>
    {/if}
  </div>
  <div class="my-4">
    <Separator />
  </div>
  <div class="items-center">
    <div class="flex h-5 items-center justify-center space-x-4 text-lg">
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
  </div>
</div>
