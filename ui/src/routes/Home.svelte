<script lang="ts">
  import "../app.css";
  import { onMount } from "svelte";
  import { load } from "../api";
  import type { Hello } from "../api";
  import Counter from "$lib/Counter.svelte";
  import * as Card from "$lib/components/ui/card";
  import { ArrowUp } from "lucide-svelte";
  import { apiData, hellodata } from "../store";
  let data: Hello;
  onMount(async () => {
    data = await load();
    apiData.set(data);
  });
</script>

<main>
  <div class="md:container md:mx-auto">
    <Card.Root>
      <Card.Header>
        <Card.Title>{$hellodata.hello}</Card.Title>
        <Card.Description
          >Click the button to increase the value</Card.Description
        >
      </Card.Header>
      <Card.Content>
        <Counter></Counter>
      </Card.Content>
      <Card.Footer>You'll be OK</Card.Footer>
    </Card.Root>

    <ArrowUp></ArrowUp>
  </div>

  <script lang="ts">
  </script>
</main>
