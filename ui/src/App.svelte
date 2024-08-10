<script lang="ts">
  import Time from '$lib/components/Time.svelte';
  import { connected } from '$lib/events';
  import Squareterminal from 'lucide-svelte/icons/square-terminal';
  import Wifi from 'lucide-svelte/icons/wifi';
  import Wifioff from 'lucide-svelte/icons/wifi-off';
  import Router, { link } from 'svelte-spa-router';
  import { wrap } from 'svelte-spa-router/wrap';
  import NotFound from './routes/NotFound.svelte';

  const routes = {
    // Exact path
    '/': wrap({
      asyncComponent: () => import('./routes/Home.svelte'),
    }),

    '/memory': wrap({
      asyncComponent: () => import('./routes/Memory.svelte'),
    }),

    '/syslog': wrap({
      asyncComponent: () => import('./routes/Syslog.svelte'),
    }),
    // Catch-all
    // This is optional, but if present it must be the last
    '*': NotFound,
  };

  let menuStyle =
    'text-muted-foreground hover:text-foreground transition-colors';
</script>

<div class="flex min-h-screen w-full flex-col">
  <header
    class="bg-background sticky top-0 flow-root h-12 w-[100%] items-center gap-4 border-b px-4 pt-2">
    <nav class="flex flex-row items-center gap-6 text-lg font-medium">
      <a
        href="/"
        class="flex items-center gap-2 text-lg font-semibold"
        use:link>
        <Squareterminal class="h-6 w-6" />
        <span class="sr-only">SPA</span>
      </a>
      <a href="/memory" use:link class={menuStyle}>Memory</a>
      <a href="/syslog" use:link class={menuStyle}>Syslog</a>
      <div
        class="absolute right-0 flex flex-row items-center justify-center gap-5">
        <div>
          {#if $connected}
            <Wifi />
          {:else}
            <Wifioff />
          {/if}
        </div>
        <div>
          <Time />
        </div>
      </div>
    </nav>
  </header>
  <Router {routes} />
</div>
