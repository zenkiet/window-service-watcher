<script lang="ts">
	import { slide } from 'svelte/transition';
	import { BrowserOpenURL } from '../../../wailsjs/runtime';

	// --- State ---
	let setting = $state({
		runBackground: false,
		notifications: false,
		runOnStartup: false,
		autoCheckUpdate: true
	});

	let appInfo = $state({
		version: '1.0.0',
		build: '2026.01.24',
		checking: false
	});

	async function handleCheckUpdate() {
		appInfo.checking = true;
		try {
			await new Promise((r) => setTimeout(r, 1000));
			alert('You are on the latest version!');
		} finally {
			appInfo.checking = false;
		}
	}
</script>

<div
	class="scrollable mx-auto flex h-full w-full max-w-4xl flex-col overflow-y-auto p-8"
	in:slide={{ axis: 'y', duration: 400 }}
>
	<header class="mb-10 flex items-center justify-between">
		<div class="flex items-center gap-4">
			<a href="/" title="Back">
				<i class="size-5 icon-[duotone--arrow-left]"></i>
			</a>
			<div>
				<h1 class="text-3xl font-bold tracking-tight text-white">Settings</h1>
				<p class="text-sm text-white/40">Manage application preferences and system behavior</p>
			</div>
		</div>
	</header>

	<div class="grid grid-cols-1 gap-8 md:grid-cols-12">
		<!-- Settings -->
		<div class="flex flex-col gap-6 md:col-span-8">
			<section class="flex flex-col gap-4">
				<h2 class="pl-1 text-xs font-bold tracking-wider text-white/30 uppercase">General</h2>
				<div class="glass-panel overflow-hidden rounded-xl border border-white/5 bg-white/2">
					{@render toggleItem({
						icon: 'icon-[duotone--rocket-launch]',
						label: 'Launch at Startup',
						desc: 'Automatically start service when you log in',
						bindValue: setting.runOnStartup
					})}

					<div class="h-px w-full bg-white/5"></div>

					{@render toggleItem({
						icon: 'icon-[duotone--command]',
						label: 'Run in Background',
						desc: 'Keep application running when closed',
						bindValue: setting.runBackground
					})}
				</div>
			</section>

			<section class="flex flex-col gap-4">
				<h2 class="pl-1 text-xs font-bold tracking-wider text-white/30 uppercase">Notifications</h2>
				<div class="glass-panel overflow-hidden rounded-xl border border-white/5 bg-white/2">
					{@render toggleItem({
						icon: 'icon-[duotone--bell]',
						label: 'System Notifications',
						desc: 'Get slack notified when a service status changes',
						bindValue: setting.notifications
					})}
				</div>
			</section>
		</div>

		<!-- About -->
		<div class="flex flex-col gap-4 md:col-span-4">
			<h2 class="pl-1 text-xs font-bold tracking-wider text-white/30 uppercase">About</h2>

			<div
				class="glass-panel flex flex-col items-center gap-4 rounded-xl border border-white/5 bg-white/2 p-6 text-center"
			>
				<div
					class="flex h-16 w-16 items-center justify-center rounded-2xl bg-linear-to-br from-blue-500/20 to-purple-500/20 shadow-inner ring-1 ring-white/10"
				>
					<i class="text-3xl text-blue-400 icon-[duotone--cube]"></i>
				</div>

				<div>
					<h3 class="font-semibold text-white">Zen Service Watcher</h3>
					<p class="text-xs text-white/40">Build {appInfo.build}</p>
				</div>

				<div
					class="inline-flex items-center gap-2 rounded-full border border-white/10 bg-white/5 px-3 py-1 font-mono text-xs text-white/60"
				>
					v{appInfo.version}
				</div>

				<button
					disabled={appInfo.checking}
					onclick={handleCheckUpdate}
					class="flex w-full items-center justify-center gap-2 rounded-lg bg-white/5 py-2 text-sm font-medium text-white transition-all hover:bg-white/10 disabled:cursor-not-allowed disabled:opacity-50"
				>
					{#if appInfo.checking}
						<i class="animate-spin icon-[duotone--spinner]"></i>
						<span>Checking...</span>
					{:else}
						<span>Check for Updates</span>
					{/if}
				</button>

				<div class="flex w-full justify-between border-t border-white/5 pt-4 text-xs text-white/30">
					<a
						href="https://github.com/zenkiet/window-service-watcher/blob/main/LICENSE"
						target="_blank"
						onclick={() =>
							BrowserOpenURL('https://github.com/zenkiet/window-service-watcher/blob/main/LICENSE')}
						class="transition-colors hover:text-white">License</a
					>
					<a
						href="https://buymeacoffee.com/zenkiet"
						target="_blank"
						onclick={() => BrowserOpenURL('https://buymeacoffee.com/zenkiet')}
						class="transition-colors hover:text-white">Sponsor</a
					>
					<a
						href="https://github.com/zenkiet/window-service-watcher"
						target="_blank"
						onclick={() => BrowserOpenURL('https://github.com/zenkiet/window-service-watcher')}
						class="transition-colors hover:text-white">GitHub</a
					>
				</div>
			</div>
		</div>
	</div>
</div>

{#snippet toggleItem({
	icon,
	label,
	desc,
	bindValue
}: {
	icon: string;
	label: string;
	desc: string;
	bindValue: boolean;
})}
	<button
		class="group flex w-full items-center justify-between p-4 text-left transition-colors outline-none hover:bg-white/4"
		onclick={() => {
			if (label === 'Launch at Startup') setting.runOnStartup = !setting.runOnStartup;
			if (label === 'Run in Background') setting.runBackground = !setting.runBackground;
			if (label === 'System Notifications') setting.notifications = !setting.notifications;
		}}
	>
		<div class="flex items-center gap-4">
			<div
				class="flex h-10 w-10 items-center justify-center rounded-lg bg-white/5 text-white/40 ring-1 ring-white/5 transition-colors group-hover:text-white group-hover:ring-white/10"
			>
				<span class={`${icon} text-lg`}></span>
			</div>
			<div class="flex flex-col gap-0.5">
				<span class="text-sm font-medium text-white/90">{label}</span>
				<span class="text-xs text-white/40">{desc}</span>
			</div>
		</div>

		<div
			class={`relative h-6 w-11 rounded-full border border-transparent transition-colors duration-200 ease-in-out focus:ring-2 focus:ring-blue-500 focus:ring-offset-2 focus:ring-offset-[#1a1b26] focus:outline-none ${bindValue ? 'bg-blue-600' : 'bg-white/10'}`}
		>
			<span
				aria-hidden="true"
				class={`pointer-events-none inline-block h-5.5 w-5.5 transform rounded-full bg-white shadow ring-0 transition duration-200 ease-in-out ${bindValue ? 'translate-x-5.5' : 'translate-x-0.5'}`}
			></span>
		</div>
	</button>
{/snippet}
