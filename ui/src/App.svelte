<script lang="ts">
  import NotFound from "./routes/NotFound.svelte";
  import Router from "svelte-spa-router";
  import { wrap } from "svelte-spa-router/wrap";

  const routes = {
    // Exact path
    "/": wrap({
      asyncComponent: () => import("./routes/Home.svelte"),
    }),

    // Catch-all
    // This is optional, but if present it must be the last
    "*": NotFound,
  };
</script>

<div class="container mt-8 pt-2 w-[750px] h-96 uto border rounded-xl">
  <Router {routes} />
</div>
