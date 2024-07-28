<script lang="ts">
  import "./app.css";
  import Counter from "./lib/Counter.svelte";
  import * as Card from "$lib/components/ui/card";
  import { AArrowUp } from "lucide-svelte";
</script>

<main>
  <div class="md:container md:mx-auto">
    <Card.Root>
      <Card.Header>
        <Card.Title>Hello</Card.Title>
        <Card.Description
          >Click the button to increase the value</Card.Description
        >
      </Card.Header>
      <Card.Content>
        <Counter></Counter>
      </Card.Content>
      <Card.Footer>You'll be OK</Card.Footer>
    </Card.Root>

    <AArrowUp></AArrowUp>
  </div>

  <script lang="ts">
  </script>
</main>
