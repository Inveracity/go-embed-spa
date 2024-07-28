<script lang="ts">
  import Home from "./routes/Home.svelte";
  import NotFound from "./routes/NotFound.svelte";
  import Router from "svelte-spa-router";

  const routes = {
    // Exact path
    "/": Home,

    // Catch-all
    // This is optional, but if present it must be the last
    "*": NotFound,
  };
</script>

<Router {routes} />
