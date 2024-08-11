<script lang="ts">
  import Time from '$lib/components/Time.svelte';
  import { connected } from '$lib/events';
  import { routeLoadingState } from '$lib/store';
  import Moon from 'lucide-svelte/icons/moon';
  import Squareterminal from 'lucide-svelte/icons/square-terminal';
  import Sun from 'lucide-svelte/icons/sun';
  import Wifi from 'lucide-svelte/icons/wifi';
  import Wifioff from 'lucide-svelte/icons/wifi-off';
  import { ModeWatcher } from 'mode-watcher';
  import Router, { link } from 'svelte-spa-router';
  import { wrap } from 'svelte-spa-router/wrap';
  import NotFound from './routes/NotFound.svelte';

  import { Button } from '$lib/components/ui/button/index.js';
  import { toggleMode } from 'mode-watcher';
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

  function routeLoading(event: any) {
    routeLoadingState.set(true);
  }

  function routeLoaded(event: any) {
    routeLoadingState.set(false);
  }
</script>

<ModeWatcher />
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
          <Button on:click={toggleMode} variant="outline" size="icon">
            <Sun
              class="h-[1.2rem] w-[1.2rem] rotate-0 scale-100 transition-all dark:-rotate-90 dark:scale-0" />
            <Moon
              class="absolute h-[1.2rem] w-[1.2rem] rotate-90 scale-0 transition-all dark:rotate-0 dark:scale-100" />
            <span class="sr-only">Toggle theme</span>
          </Button>
        </div>
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
  <div class="container">
    <Router
      {routes}
      on:routeLoading={routeLoading}
      on:routeLoaded={routeLoaded} />
    {#if $routeLoadingState}
      <p>Loading...</p>
    {/if}
  </div>
</div>
