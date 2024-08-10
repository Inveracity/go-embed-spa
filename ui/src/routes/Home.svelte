<script lang="ts">
  import type { Hello } from '$lib/api';
  import BACKEND_BASE_URL from '$lib/config';
  import { apiData } from '$lib/store';
  import { onMount } from 'svelte';

  onMount(async () => {
    const response = await fetch(`${BACKEND_BASE_URL}/api`);
    if (response.ok) {
      const data: Hello = await response.json();
      console.log(data);
      apiData.set(data);
    }
  });
</script>

<div class="p-[10vh] text-center">
  <h1>Hello {$apiData.hello}</h1>
</div>
