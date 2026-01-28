<script lang="ts">
	import { theme } from '$lib/stores/theme.svelte';
	import Icon from '../ui/Icon.svelte';
</script>

<div class="fixed bottom-6 right-6 z-50">
  <button
    class="justify-center flex items-center w-10 h-10 rounded-full bg-surface shadow-xl text-text-muted hover:text-primary transition-all duration-300"
    onclick={() => theme.toggle()}
    aria-label="Toggle Theme"
  >
    {#if theme.isDark}
      <Icon name="sun" />
    {:else}
      <Icon name="moon" />
    {/if}
  </button>
</div>