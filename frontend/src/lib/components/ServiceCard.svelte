<script lang="ts">
  import { formatRelativeTime, readFileAsBytes } from '$lib/helpers';
  import { Status, type Service } from '$lib/stores/services.svelte';
	import InstallModal from './modals/InstallModal.svelte';
	import Tooltip from './Tooltip.svelte';

  let { service } = $props<{ service: Service }>();

  const isRunning = $derived(service.status === Status.RUNNING);

  let showInstallModal = $state(false);

  async function handleInstall(files: File[]) {
    if (!files || files.length === 0) return;

    const payload = files.map(async (file) => {
      const data = await readFileAsBytes(file);
      return {
        name: file.name,
        data
      };
    });
    await service.install(await Promise.all(payload));
  }
</script>

<div
  class="group relative flex flex-col gap-4 rounded-xl border border-white/5 bg-white/5 p-4 backdrop-blur-md transition-all hover:border-white/10 hover:bg-white/[0.07] hover:shadow-2xl hover:shadow-black/20"
>
  <div class="flex items-start justify-between gap-3">
    <div class="min-w-0 flex-1">
      <div class="mb-1 flex items-center gap-2">
        <span class="font-mono text-xs font-semibold text-white/30"
          >#{service.metrics.pid ? service.metrics.pid : '--'}</span
        >
      </div>

      <div class="flex flex-col">
        <h3
          class="truncate font-semibold text-white/90 transition-colors group-hover:text-white"
          title={service.name}
        >
          {service.name}
        </h3>
        <p class="line-clamp-2 min-h-[2.5em] text-xs leading-relaxed text-white/40">
          {service.description}
        </p>
      </div>
    </div>

    <div class="flex gap-2">
      <Tooltip content={!service.installable ? 'Coming Soon' : ''}>
        <button
            title="install service"
            onclick={() => showInstallModal = true}
            class={`
                flex h-8 w-8 items-center justify-center rounded-lg border transition-all
                ${!service.installable
                    ? 'cursor-not-allowed border-white/5 bg-white/5 text-white/20 opacity-50'
                    : 'cursor-pointer border-blue-500/20 bg-blue-500/10 text-blue-400 hover:border-blue-500 hover:bg-blue-500/40 hover:text-white'
                }
            `}
            disabled={!service.installable}
        >
            <span class="size-4.5 icon-[duotone--wrench]"></span>
        </button>
      </Tooltip>

      <button
        onclick={() => service.openExplorer()}
        class="flex h-8 w-8 cursor-pointer items-center justify-center rounded-lg border border-white/5 bg-white/5 text-white/40 transition-colors hover:border-white/20 hover:bg-white/10 hover:text-white"
        title="Open in Explorer"
      >
        <span class="size-5 icon-[duotone--folder-open]"></span>
      </button>

      <button
        onclick={isRunning ? () => service.stop() : () => service.start()}
        disabled={service.loading}
        class={`
       flex h-8 w-8 cursor-pointer items-center justify-center rounded-lg active:scale-[0.98] transition-all
      disabled:cursor-not-allowed disabled:opacity-50
      ${
        isRunning
          ? 'bg-rose-500/10 text-rose-400 ring-1 ring-rose-500/20 ring-inset hover:bg-rose-500/20'
          : 'bg-emerald-500 text-emerald-950 shadow-[0_0_15px_-3px_rgba(16,185,129,0.4)] hover:bg-emerald-400'
      }
    `}
      >
        {#if service.loading}
          <span class="size-4.5 animate-spin icon-[duotone--spinner]"></span>
        {:else if isRunning}
          <span class="size-4.5 icon-[duotone--stop]"></span>
        {:else}
          <span class="size-4.5 icon-[duotone--play]"></span>
        {/if}
      </button>
    </div>
  </div>

  <div class="grid grid-cols-3 gap-px overflow-hidden rounded-lg border border-white/5 bg-white/5">
    {@render metric(
      'CPU',
      service.metrics.cpu ? `${service.metrics.cpu.toFixed(1)}%` : '--',
      'text-blue-300'
    )}
    {@render metric(
      'RAM',
      service.metrics.mem ? `${(service.metrics.mem / 1024 / 1024).toFixed(0)} MB` : '--',
      'text-purple-300'
    )}
    {@render metric(
      'TIME',
      service.metrics.createTime ? formatRelativeTime(service.metrics.createTime) : '--',
      'text-emerald-300'
    )}
  </div>
</div>

<InstallModal
  isOpen={showInstallModal}
  serviceName={service.name}
  onClose={() => showInstallModal = false}
  onInstall={handleInstall}
/>

{#snippet metric(label: string, value: string, colorClass: string)}
  <div class="flex flex-col items-center justify-center gap-1 bg-white/2 py-2 transition-colors hover:bg-white/4">
    <span class="text-[9px] font-bold tracking-wider text-white/30 uppercase">{label}</span>
    <span class={`font-mono text-xs font-medium ${value === '--' ? 'text-white/20' : colorClass}`}>
      {value}
    </span>
  </div>
{/snippet}