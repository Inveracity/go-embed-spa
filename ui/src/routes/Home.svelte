<script lang="ts">
  import "../app.css";
  import { onMount } from "svelte";
  import Counter from "$lib/Counter.svelte";
  import * as Card from "$lib/components/ui/card/";
  import ArrowUp from "lucide-svelte/icons/arrow-up";
  import { apiData } from "../store";
  import { subscribe, state } from "$lib/events";

  onMount(subscribe);
</script>

<main>
  <div class="md:container md:mx-auto">
    <Card.Root>
      <Card.Header>
        <Card.Title>{$apiData.hello}</Card.Title>
        <Card.Description
          >Click the button to increase the value</Card.Description
        >
      </Card.Header>
      <Card.Content>
        <Counter></Counter>
      </Card.Content>
      <Card.Footer>
        {$state}
      </Card.Footer>
    </Card.Root>

    <ArrowUp></ArrowUp>
  </div>

  <p>{import.meta.env.THING_BACKEND_URL}</p>
</main>
