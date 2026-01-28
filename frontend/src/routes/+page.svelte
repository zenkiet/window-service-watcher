<script lang="ts">
	import Sidebar from '$lib/components/layout/Sidebar.svelte';
	import ServiceList from '$lib/components/ServiceList.svelte';
	import Badge from '$lib/components/ui/Badge.svelte';
	import { serviceStore } from '$lib/stores/services.svelte';
	import { fade } from 'svelte/transition';

	// Local State
	let searchQuery = $state('');
</script>

<div class="flex h-screen overflow-hidden" in:fade={{ duration: 400 }}>
	<Sidebar />

	<main class="scrollable flex-1 overflow-y-auto scroll-smooth px-6 py-4">
		<header class="flex shrink-0 flex-col gap-5 px-6 pt-6 pb-2">
			<div class="flex items-center justify-between">
				<div class="flex items-center gap-4">
					<h1 class="text-2xl font-bold tracking-tight">Service Watcher</h1>
					<div class="flex items-center gap-2 border-l border-white/10 pl-4">
						<Badge label="Running" type="success" count={serviceStore.runningCount} />
						<Badge label="Stopped" type="danger" count={serviceStore.stoppedCount} />
					</div>
				</div>
			</div>

			<div class="flex items-center justify-between backdrop-blur-md">
				<div class="group relative max-w-md flex-1">
					<i
						class="absolute top-1/2 left-3 -translate-y-1/2 text-lg text-white/20 transition-colors icon-[duotone--magnifying-glass] group-focus-within:text-white/50"
					></i>
					<input
						type="text"
						bind:value={searchQuery}
						placeholder="Filter services by name, PID or description..."
						class="h-10 w-full rounded-lg border border-transparent bg-transparent pr-4 pl-10 text-sm text-white placeholder-white/30 transition-all focus:border-white/10 focus:bg-white/5 focus:ring-0 focus:outline-none"
					/>
				</div>
			</div>
		</header>

		<ServiceList {searchQuery} />

		<div class="h-10"></div>
	</main>
</div>
