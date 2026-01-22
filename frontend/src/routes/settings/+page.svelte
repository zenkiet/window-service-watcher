<script lang="ts">
	let autoRefresh = $state(true);
	let notifications = $state(true);
	let runOnStartup = $state(false);
</script>

<div
	class="animate-in fade-in slide-in-from-bottom-4 mx-auto flex h-full w-full max-w-2xl flex-col p-6 duration-500"
>
	<!-- Header -->
	<header class="mb-10 flex items-center gap-4">
		<a
			href="/"
			class="glass-btn h-10 w-10 rounded-full p-0! hover:bg-white/10"
			title="Back to Dashboard"
		>
		 	<i class="icon-[duotone--arrow-left] size-4"></i>
		</a>
		<div>
			<h1 class="text-2xl font-light text-white">Settings</h1>
			<p class="text-sm text-white/40">Manage your watcher preferences</p>
		</div>
	</header>

	<!-- Settings Form -->
	<div class="glass-panel flex flex-col gap-6 rounded-2xl p-6">
		<!-- Setting Item -->
		<div class="flex items-center justify-between">
			<div class="flex flex-col">
				<span class="font-medium text-white/80">Auto-refresh Dashboard</span>
				<span class="text-xs text-white/40"
					>Automatically update service status every 5 seconds</span
				>
			</div>
			<button
				class={`h-6 w-12 rounded-full p-1 transition-colors duration-300 ${autoRefresh ? 'bg-blue-500' : 'bg-white/10'}`}
				onclick={() => (autoRefresh = !autoRefresh)}
			>
				<div
					class={`h-4 w-4 rounded-full bg-white shadow-sm transition-transform duration-300 ${autoRefresh ? 'translate-x-6' : 'translate-x-0'}`}
				></div>
			</button>
		</div>

		<div class="h-px w-full bg-white/5"></div>

		<!-- Setting Item -->
		<div class="flex items-center justify-between">
			<div class="flex flex-col">
				<span class="font-medium text-white/80">System Notifications</span>
				<span class="text-xs text-white/40">Get notified when a service stops unexpectedly</span>
			</div>
			<button
				class={`h-6 w-12 rounded-full p-1 transition-colors duration-300 ${notifications ? 'bg-blue-500' : 'bg-white/10'}`}
				onclick={() => (notifications = !notifications)}
			>
				<div
					class={`h-4 w-4 rounded-full bg-white shadow-sm transition-transform duration-300 ${notifications ? 'translate-x-6' : 'translate-x-0'}`}
				></div>
			</button>
		</div>

		<div class="h-px w-full bg-white/5"></div>

		<!-- Setting Item -->
		<div class="flex items-center justify-between">
			<div class="flex flex-col">
				<span class="font-medium text-white/80">Launch at Startup</span>
				<span class="text-xs text-white/40">Start Service Watcher when you log in</span>
			</div>
			<button
				class={`h-6 w-12 rounded-full p-1 transition-colors duration-300 ${runOnStartup ? 'bg-blue-500' : 'bg-white/10'}`}
				onclick={() => (runOnStartup = !runOnStartup)}
			>
				<div
					class={`h-4 w-4 rounded-full bg-white shadow-sm transition-transform duration-300 ${runOnStartup ? 'translate-x-6' : 'translate-x-0'}`}
				></div>
			</button>
		</div>
	</div>

	<div class="mt-8 rounded-xl border border-dashed border-white/5 bg-white/[0.02] p-4 text-center">
		<p class="text-xs text-white/30">Version 1.0.0</p>
	</div>
</div>
