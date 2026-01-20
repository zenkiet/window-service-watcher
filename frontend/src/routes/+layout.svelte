<script lang="ts">
	import Footer from '$lib/components/Footer.svelte';
	import { serviceStore } from '$lib/stores/services.svelte';
	import './layout.css';
	let { children } = $props();

	$effect(() => {
		serviceStore.init();
	})

</script>

<div class="flex h-screen w-screen flex-col overflow-hidden selection:bg-blue-500/30 pt-10">
	<main class="relative flex-1 overflow-hidden">
		{@render children()}
	</main>

	<Footer />
</div>
