<script lang="ts">
	import { serviceStore } from '$lib/stores/services.svelte';
	import ServiceCard from './ServiceCard.svelte';
</script>

<div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-4 pb-20">
	{#each serviceStore.services as service (service.id)}
		<ServiceCard {service} />
	{/each}

   <button class="glass-card group flex min-h-35 flex-col items-center justify-center gap-2 border-dashed border-white/10 bg-transparent hover:bg-white/5 hover:border-white/20 transition-all cursor-not-allowed opacity-50">
        <div class="h-8 w-8 rounded-full bg-white/5 flex items-center justify-center text-white/20 group-hover:text-white/40 group-hover:bg-white/10 transition-colors">
            <i class="icon-[duotone--plus] size-4"></i>
        </div>
        <span class="text-xs font-medium text-white/20 group-hover:text-white/40">Add Service</span>
    </button>
</div>
