<script lang="ts">
  import { CheckUpdate, Update } from '../../../../wailsjs/go/app/App';
  import { Quit } from '../../../../wailsjs/runtime/runtime';
  import { onMount } from 'svelte';
  import { fade, scale } from 'svelte/transition';

  let show = $state(false);
  let loading = $state(false);
  let updateInfo = $state<any>(null);
  let errorMsg = $state('');

  onMount(async () => {
    try {
      const info = await CheckUpdate();
      if (info.available) {
        updateInfo = info;
        show = true;
      }
    } catch (e) {
      console.error("Check update failed:", e);
    }
  });

  async function handleUpdate() {
    if (!updateInfo) return;
    loading = true;
    errorMsg = '';
    try {
      await Update(updateInfo.downloadUrl);
      Quit();
    } catch (e: any) {
      errorMsg = e.message;
      loading = false;
    }
  }
</script>

{#if show}
  <div
    class="fixed inset-0 z-50 flex items-center justify-center bg-black/60 p-4 backdrop-blur-sm"
    transition:fade={{ duration: 200 }}
  >
    <div
      class="glass-panel relative flex min-h-75 w-full max-w-lg flex-col overflow-hidden rounded-xl border border-white/10 bg-[#1a1b26]/95 ring-1 ring-white/5"
      transition:scale={{ start: 0.98, duration: 200 }}
    >
      <div class="flex items-start justify-between border-b border-neutral-800 px-6 py-5">
        <div class="flex flex-col gap-1">
          <h3 class="text-base font-semibold text-white">Update Available</h3>
          <p class="text-sm text-neutral-400">
            A new version of the application is ready to install.
          </p>
        </div>

        <div class="rounded border border-neutral-700 bg-neutral-800 px-2.5 py-1 text-xs font-mono font-medium text-neutral-300">
          {updateInfo?.latestVersion}
        </div>
      </div>

      <div class="px-6 py-5">
        <label for="changelog" class="mb-2 block text-xs font-medium uppercase tracking-wider text-neutral-500">
          Release Notes
        </label>

        <div class="h-48 overflow-y-auto rounded border border-neutral-800 bg-black/30 p-3 scrollable">
          <pre class="whitespace-pre-wrap font-mono text-sm leading-relaxed text-neutral-300">{updateInfo?.releaseNotes || "No details provided."}</pre>
        </div>

        {#if errorMsg}
          <div class="mt-4 flex items-center gap-2 rounded border border-red-900/50 bg-red-900/20 px-3 py-2 text-sm text-red-200">
            <span class="icon-[heroicons--exclamation-circle] text-red-400"></span>
            <span>{errorMsg}</span>
          </div>
        {/if}
      </div>

      <div class="flex items-center justify-end gap-3 rounded-b-lg border-t border-neutral-800 bg-neutral-900/50 px-6 py-3">
        <button
          onclick={() => show = false}
          class="rounded-lg px-4 py-2 text-xs font-medium text-neutral-400 transition hover:bg-neutral-800 hover:text-white disabled:opacity-50"
          disabled={loading}
        >
          Skip
        </button>

        <button
          onclick={handleUpdate}
          class="flex cursor-pointer items-center gap-2 rounded-lg bg-blue-600 px-4 py-2 text-xs font-medium text-white transition hover:bg-blue-700 disabled:bg-blue-600/50"
          disabled={loading}
        >
          {#if loading}
            <span class="icon-[duotone--arrows-rotate] animate-spin"></span>
            <span>Updating...</span>
          {:else}
            <span class="icon-[duotone--download]"></span>
            <span>Update</span>
          {/if}
        </button>
      </div>
    </div>
  </div>
{/if}