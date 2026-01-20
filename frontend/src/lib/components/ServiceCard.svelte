<script lang="ts">
	import type { Service } from '$lib/stores/services.svelte';

	let { service } = $props<{ service: Service }>();

	let isProcessing = $derived(service.status === 'starting' || service.status === 'stopping');

    let statusColor = $derived.by(() => {
        switch (service.status) {
            case 'running': return 'bg-emerald-500 shadow-emerald-500/20';
            case 'stopped': return 'bg-neutral-500 shadow-neutral-500/20';
            case 'error': return 'bg-rose-500 shadow-rose-500/20';
            default: return 'bg-amber-500 shadow-amber-500/20 animate-pulse';
        }
    });

    let statusText = $derived.by(() => {
         switch (service.status) {
            case 'running': return 'Active';
            case 'stopped': return 'Stopped';
            case 'error': return 'Error';
            default: return 'Processing...';
        }
    });

</script>

<div class="glass-card group relative flex flex-col gap-4 p-5">
    <!-- Header -->
	<div class="flex items-start justify-between">
		<div class="flex items-center gap-3">
			<div class={`h-2.5 w-2.5 rounded-full shadow-[0_0_10px] ${statusColor}`}></div>
			<div>
				<h3 class="font-medium text-white/90 group-hover:text-white transition-colors">
                    {service.name}
                </h3>
				<p class="text-xs text-white/40 max-w-60">{service.description}</p>
			</div>
		</div>

        <!-- Status Badge -->
        <span class="absolute top-2 right-3 rounded-full bg-white/5 px-2 py-0.5 text-[10px] font-medium text-white/40 border border-white/5">
            {statusText}
        </span>
	</div>

    <!-- Actions -->
	<div class="mt-auto flex items-center justify-between pt-2">
		<div class="flex items-center gap-2">
            {#if service.status === 'running'}
                <button
                    class="glass-btn h-8 w-8 p-0! rounded-full hover:bg-rose-500/20 hover:border-rose-500/30 hover:text-rose-200"
                    onclick={() => service.stop()}
                    title="Stop Service"
                    disabled={isProcessing}
                >
                    <i class="icon-[duotone--stop] size-3.5"></i>
                </button>

                <button
                    class="glass-btn h-8 w-8 p-0! rounded-full hover:bg-amber-500/20 hover:border-amber-500/30 hover:text-amber-200"
                    onclick={() => service.restart()}
                    title="Restart Service"
                    disabled={isProcessing}
                >
                    <i class="icon-[duotone--rotate] size-3.5"></i>
                </button>
            {:else}
                 <button
                    class="glass-btn h-8 w-8 !p-0 rounded-full hover:bg-emerald-500/20 hover:border-emerald-500/30 hover:text-emerald-200"
                    onclick={() => service.start()}
                    title="Start Service"
                    disabled={isProcessing}
                >
                    <i class="icon-[duotone--play] size-3.5"></i>
                </button>
            {/if}
		</div>

        {#if service.logsUrl}
            <a
                href={service.logsUrl}
                target="_blank"
                class="text-xs text-blue-400/60 hover:text-blue-400 hover:underline decoration-blue-400/30 underline-offset-4 transition-all"
            >
                View Logs &rarr;
            </a>
        {/if}
	</div>
</div>
