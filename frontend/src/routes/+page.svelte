<script lang="ts">
  import { Events, Window } from "@wailsio/runtime";
  import { App } from "../../bindings/window-service-watcher/internal/app";

  let status = $state("Checking...");
  let isHealthy = $state(false);
  let logs = $state<string[]>([]);

  $effect(() => {
    const interval = setInterval(async () => {
      try {
        const res = await App.GetServiceStatus();
        status = res.status;
        isHealthy = res.is_healthy;
      } catch (err) {
        console.error("Failed to get status:", err);
        status = "Connection Error";
      }
    }, 2000);

    // App.WatchLogs('logs/temp.log');

    const stopLogListener = Events.On("new-log", (event) => {
      const line = event.data;
      logs = [...logs.slice(-4), ...line];
    });

    return () => {
      clearInterval(interval);
      stopLogListener();
    };
  });

  const handleMin = () => Window.Minimise();
</script>

<main class="widget-container">
  <div
    class="flex h-full w-full cursor-default flex-col gap-2 overflow-hidden border border-white/20 bg-white/10 p-4"
  >
    <div class="flex items-center justify-between">
      <h3
        class="text-sm font-semibold text-white"
        style="--wails-draggable:drag"
      >
        Blogic Report Service
      </h3>
    </div>

    <div class="flex items-center gap-1.5 mb-2">
      <div
        class="h-2 w-2 rounded-full {isHealthy
          ? 'bg-green-400'
          : 'animate-pulse bg-red-400'}"
      ></div>
      <span class="font-mono text-xs text-white/80">{status}</span>
    </div>

    <div
      class="group relative flex-1 overflow-hidden rounded-lg bg-black/20 p-2 font-mono text-[10px] text-white/70"
    >
      <div class="scrollbar-hide absolute inset-0 overflow-y-auto pl-2">
        {#if logs.length === 0}
          <div class="text-white/30 italic">Waiting for logs...</div>
        {/if}
        {#each logs as log}
          <div
            class="truncate border-b border-white/5 py-0.5 transition-colors last:border-0 hover:text-white"
          >
            {log}
          </div>
        {/each}
      </div>
    </div>

    <footer class="flex justify-end">
      <button
        class="rounded bg-white/10 px-3 py-1 text-xs text-white/70 transition-colors hover:bg-white/20 hover:text-white"
        on:click={handleMin}
        tabindex="-1"
      >
        Minimize
      </button>
    </footer>
  </div>
</main>
