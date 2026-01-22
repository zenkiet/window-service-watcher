<script lang="ts">
	import ServiceList from '$lib/components/ServiceList.svelte';
	import { GetConfig } from '../../bindings/window-service-watcher/internal/app/app';

	let searchQuery = $state('');
</script>

<div class="mx-auto flex h-full w-full max-w-7xl flex-col overflow-y-auto scroll-smooth p-5">
	<!-- Header -->
	<div class="mb-8 flex flex-col gap-6">
		<div class="flex items-center justify-between">
			<h1 class="text-3xl font-bold tracking-tight text-white">Zen Service Watcher</h1>

			<a href="/settings" class="glass-btn group h-10 w-10 rounded-lg p-0!" title="Settings">
				<i class="icon-[duotone--gear] size-4.5 text-white/40 transition-all duration-500 group-hover:rotate-90 group-hover:text-white "></i>
			</a>
		</div>

		<!-- Dashboard Controls -->
		<div class="flex items-center gap-3">
			<!-- Search -->
			<div class="group relative">
				<i class="icon-[duotone--magnifying-glass] size-4 -translate-y-1/2 absolute top-1/2 left-3 text-white/20 transition-colors group-focus-within:text-white/50"></i>

				<input
					type="text"
					bind:value={searchQuery}
					placeholder="Filter services..."
					class="glass-input w-64! pl-10!"
				/>
			</div>
		</div>
	</div>

	<ServiceList />
</div>
