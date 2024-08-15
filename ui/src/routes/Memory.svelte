<script lang="ts">
  import { memoryMessages } from '$lib/store';
  import { Streamer } from '$lib/stream';
  import Square from 'lucide-svelte/icons/square';
  import { onDestroy, onMount } from 'svelte';

  const s = new Streamer('/stream/memory', memoryMessages);
  let squarray = [0, 0, 0, 0, 0, 0, 0, 0, 0, 0];
  let mem;

  function flipArrayBasedOnPercentage(arr: Array<number>, percentage: number) {
    // Determine the number of elements to flip based on the percentage
    const elementsToFlip = Math.floor(percentage / 10);

    // Loop through the array and flip the first 'elementsToFlip' elements
    for (let i = 0; i < elementsToFlip; i++) {
      arr[i] = arr[i] === 0 ? 1 : arr[i];
    }

    return arr;
  }

  $: mem = Number(
    (Math.round((Number($memoryMessages.msg) || 0) * 100) / 100).toFixed(2),
  );

  $: squarray = flipArrayBasedOnPercentage(squarray, mem);

  onMount(() => {
    s.connect();
  });

  onDestroy(() => {
    s.disconnect();
  });
</script>

<main class="flex flex-1 flex-col gap-4 p-4 md:gap-8 md:p-10">
  <div class="mx-auto grid w-full max-w-6xl justify-center gap-2 text-center">
    <h1 class="text-3xl font-semibold">Memory</h1>
    <p class="text-mono text-xl">{mem}%</p>
    <div class="flex flex-col content-center">
      {#each squarray.slice().reverse() as sq}
        <div class="self-center">
          <Square class={sq === 1 ? 'fill-green-500' : 'opacity-15'} />
        </div>
      {/each}
    </div>
  </div>
</main>
