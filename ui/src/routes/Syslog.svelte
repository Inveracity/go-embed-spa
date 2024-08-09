<script lang="ts">
  import { es, connected, syslogMsgs } from "$lib/events";
  import { onDestroy, onMount } from "svelte";

  onMount(() => {
    if ($connected) {
      $es.addEventListener("syslog", async function (event) {
        const data = event.data;
        syslogMsgs.set(data);
        return () => $es.close();
      });
    }
  });

  onDestroy(() => {
    $es.removeEventListener("syslog", async function (event) {
      console.log("");
    });
  });
  console.log("Unmounted syslog");
</script>

{$syslogMsgs}
