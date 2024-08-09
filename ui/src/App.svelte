<script lang="ts">
  import BACKEND_BASE_URL from "$lib/config";
  import NotFound from "./routes/NotFound.svelte";
  import Router from "svelte-spa-router";
  import Squareterminal from "lucide-svelte/icons/square-terminal";
  import { wrap } from "svelte-spa-router/wrap";
  import { link } from "svelte-spa-router";
  import { onDestroy, onMount } from "svelte";
  import { subscribe, unsubscribe, es, connected } from "$lib/events";

  const routes = {
    // Exact path
    "/": wrap({
      asyncComponent: () => import("./routes/Home.svelte"),
    }),

    "/memory": wrap({
      asyncComponent: () => import("./routes/Homev2.svelte"),
    }),

    "/syslog": wrap({
      asyncComponent: () => import("./routes/Syslog.svelte"),
    }),
    // Catch-all
    // This is optional, but if present it must be the last
    "*": NotFound,
  };

  let menuStyle =
    "text-muted-foreground hover:text-foreground transition-colors";

  const connect = () => {
    if (!$connected) {
      es.set(new EventSource(`${BACKEND_BASE_URL}/stream`));
      subscribe($es);
    }
  };

  // Automatically subscribe when going to the page
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

<div class="flex min-h-screen w-full flex-col">
  <header
    class="bg-background sticky top-0 flex h-16 items-center gap-4 border-b px-4"
  >
    <nav class="flex flex-row gap-6 text-lg font-medium items-center">
      <a
        href="/"
        class="flex items-center gap-2 text-lg font-semibold"
        use:link
      >
        <Squareterminal class="h-6 w-6" />
        <span class="sr-only">SPA</span>
      </a>
      <a href="/memory" use:link class={menuStyle}> Memory </a>
      <a href="/syslog" use:link class={menuStyle}> Syslog </a>
    </nav>
    <div class="flex items-end gap-4">
      <p>hello</p>
    </div>
  </header>
  <Router {routes} />
</div>
