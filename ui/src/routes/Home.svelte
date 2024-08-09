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
  import * as Tabs from "$lib/components/ui/tabs/index";
  import { Progress } from "$lib/components/ui/progress/index";
  import Syslog from "src/routes/Syslog.svelte";
  import Time from "$lib/components/Time.svelte";
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
  <Time>{$messages}</Time>
  <div class="my-4">
    <Separator />
  </div>
  <Tabs.Root value="syslog" class="w-[400]">
    <Tabs.List class="grid w-full grid-cols-2">
      <Tabs.Trigger value="syslog">syslog</Tabs.Trigger>
      <Tabs.Trigger value="memory">memory</Tabs.Trigger>
    </Tabs.List>
    <Tabs.Content value="syslog">
      <Syslog />
    </Tabs.Content>
    <Tabs.Content value="memory">
      {$memoryMsgs?.usedPercent || ""}%
      <Progress
        value={$memoryMsgs?.usedPercent || 0}
        max={100}
        class="w-[60%]"
      />
    </Tabs.Content>
  </Tabs.Root>
</div>

<div class="grid h-screen w-full pl-[53px]">
  <aside class="inset-y fixed left-0 z-20 flex h-full flex-col border-r">
    <div class="border-b p-2">
      <Button variant="outline" size="icon" aria-label="Home">
        <p>home</p>
      </Button>
    </div>
    <nav class="grid gap-1 p-2">
      <Button variant="ghost" size="icon" class="rounded-lg" aria-label="Models"
        ><p>syslog</p></Button
      >
      <Button variant="ghost" size="icon" class="rounded-lg" aria-label="API"
      ></Button>
    </nav>
    <nav class="mt-auto grid gap-1 p-2">
      <Button
        variant="ghost"
        size="icon"
        class="mt-auto rounded-lg"
        aria-label="Help"><p>memory</p></Button
      >
      <Button
        variant="ghost"
        size="icon"
        class="mt-auto rounded-lg"
        aria-label="Account"
      ></Button>
    </nav>
  </aside>
  <div class="flex flex-col">
    <header
      class="bg-background sticky top-0 z-10 flex h-[57px] items-center gap-1 border-b px-4"
    >
      <h1 class="text-xl font-semibold">Playground</h1>
    </header>
    <p>content</p>
  </div>
</div>
