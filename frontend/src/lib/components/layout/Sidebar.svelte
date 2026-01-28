<script lang="ts">
	import { page } from '$app/state';
	import Footer from '../Footer.svelte';

	let { isOpen = false, onClose } = $props<{
		isOpen?: boolean;
		onClose?: () => void;
	}>();

	const projects = [
		{ name: 'Services', class: 'icon-[regular--server] size-4', href: '/' },
		{ name: 'Projects', class: 'icon-[regular--diagram-project] size-4', href: '/folders' },
		{ name: 'Run Script', class: 'icon-[regular--code] size-4', href: '/script' }
	];

	const config = [
		{ name: 'Settings', class: 'icon-[regular--gear-complex] size-4', href: '/settings' }
		// { name: 'Logs', class: 'icon-[regular--rectangle-history] size-5', href: '/logs' }
	];

	function isActive(path: string): boolean {
		return (
			page.url.pathname.replaceAll('/', '').toLowerCase() === path.replaceAll('/', '').toLowerCase()
		);
	}
</script>

<aside
	class={`
        fixed inset-y-0 left-0 z-50 flex h-full w-64 shrink-0 flex-col overflow-y-auto
        border-r border-border bg-surface pt-5 shadow-2xl transition-transform duration-300 ease-in-out md:shadow-none
        ${isOpen ? 'translate-x-0' : '-translate-x-full'}
        md:static md:translate-x-0
    `}
>
	<div class="flex h-full flex-col px-4 pb-4">
		<div class="mb-6 flex items-center justify-between gap-3">
			<h1 class="text-lg font-bold tracking-tight text-main">ZenB Tool</h1>
		</div>

		<nav class="mb-8 space-y-1">
			<p class="tracking mb-2 px-3 text-xs font-semibold text-muted/70 uppercase">Features</p>
			{#each projects as item}
				<a
					href={item.href}
					onclick={() => window.innerWidth < 768 && onClose?.()}
					class={`group relative flex items-center gap-3 rounded-md px-3 py-2 text-sm font-medium transition-all duration-200
            ${
							isActive(item.href)
								? 'bg-blue-50 text-primary dark:bg-blue-500/10'
								: 'hover:bg-surface-highlight text-muted hover:text-main'
						}`}
				>
					<div
						class="absolute top-0 left-0 h-full w-1 rounded-l-md bg-blue-500 opacity-0 transition-opacity group-hover:opacity-100"
					></div>
					<i class={item.class}></i>
					{item.name}
				</a>
			{/each}
		</nav>

		<nav class="space-y-1">
			<p class="tracking mb-2 px-3 text-xs font-semibold text-muted/70 uppercase">Configuration</p>
			{#each config as item}
				<a
					href={item.href}
					onclick={() => window.innerWidth < 768 && onClose?.()}
					class={`group relative flex items-center gap-3 rounded-md px-3 py-2 text-sm font-medium transition-all duration-200
            ${
							isActive(item.href)
								? 'bg-blue-50 text-primary dark:bg-blue-500/10'
								: 'hover:bg-surface-highlight text-muted hover:text-main'
						}`}
				>
					<div
						class="absolute top-0 left-0 h-full w-1 rounded-l-md bg-blue-500 opacity-0 transition-opacity group-hover:opacity-100"
					></div>
					<i class={item.class}></i>
					{item.name}
				</a>
			{/each}
		</nav>
	</div>

	<div class="mt-auto">
		<Footer />
	</div>
</aside>
