<script lang="ts">
	import { serviceStore } from '$lib/stores/services.svelte';
	import ServiceCard from './ServiceCard.svelte';
	import { fade } from 'svelte/transition';

	let { searchQuery } = $props<{
		searchQuery: string;
	}>();

	const filteredServices = $derived.by(() =>
		serviceStore.services.filter(
			(service) =>
				service.name.toLowerCase().includes(searchQuery.toLowerCase()) ||
				service.description.toLowerCase().includes(searchQuery.toLowerCase()) ||
				service.metrics?.pid?.toString().includes(searchQuery)
		)
	);
</script>

<div class="pb-20">
	{#if filteredServices.length === 0 && searchQuery}
		<div class="flex flex-col items-center justify-center py-20 text-white/30" in:fade>
			<i class="mb-4 text-4xl icon-[duotone--magnifying-glass]"></i>
			<p class="text-sm font-medium">No services found for "{searchQuery}"</p>
			<button
				class="mt-2 text-xs text-blue-400 hover:text-blue-300 hover:underline"
				onclick={() => (searchQuery = '')}
			>
				Clear filter
			</button>
		</div>
	{:else}
		<div class="grid grid-cols-1 gap-4 md:grid-cols-2 lg:grid-cols-3">
			{#each filteredServices as service (service.id)}
				<div in:fade={{ duration: 200 }}>
					<ServiceCard {service} />
				</div>
			{/each}

			<button
				class="
                    glass-card flex min-h-45 cursor-not-allowed flex-col items-center justify-center gap-3 border-dashed border-white/10 bg-transparent opacity-50 transition-all hover:border-white/20 hover:bg-white/5"
			>
				<div
					class="flex h-10 w-10 items-center justify-center rounded-full bg-white/5 text-white/20 transition-colors group-hover:bg-white/10 group-hover:text-white/50"
				>
					<i class="text-lg icon-[duotone--plus]"></i>
				</div>
				<span class="text-xs font-medium text-white/20 group-hover:text-white/40">
					Add Service
				</span>
			</button>
		</div>
	{/if}
</div>
