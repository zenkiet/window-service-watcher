<script lang="ts">
  import Icon from '../ui/Icon.svelte';

  const projects = [
    { name: 'Blogic Service', icon: 'dashboard', active: true },
    { name: 'Legacy Systems', icon: 'dns', active: false },
  ];

  const config = [
    { name: 'Settings', icon: 'settings' },
    // { name: 'Logs', icon: 'history' },
  ];
</script>

<aside class="w-64 flex flex-col border-r border-border bg-surface h-full shrink-0 transition-colors duration-300 pt-5 overflow-auto scrollable">
  <div class="p-6">
    <!-- Brand -->
    <div class="flex items-center gap-3 mb-8">
      <h1 class="font-bold text-lg tracking-tight text-text-main">Service Watcher</h1>
    </div>

    <!-- Project Nav -->
    <nav class="space-y-1 mb-8">
      <p class="px-3 text-xs font-semibold text-text-muted uppercase tracking-wider mb-2">Projects</p>
      {#each projects as item}
        <a
          href="#"
          class={`flex items-center gap-3 px-3 py-2 text-sm font-medium rounded-md transition-all duration-200 group
            ${item.active
              ? 'bg-blue-50 dark:bg-blue-500/10 text-primary'
              : 'text-text-muted hover:bg-surface-highlight hover:text-text-main'}`}
        >
          <Icon name={item.icon} class={item.active ? 'text-primary' : 'group-hover:text-text-main transition-colors'} />
          {item.name}
        </a>
      {/each}
    </nav>

    <!-- Config Nav -->
    <nav class="space-y-1">
      <p class="px-3 text-xs font-semibold text-text-muted uppercase tracking-wider mb-2">Configuration</p>
      {#each config as item}
        <a
          href="#"
          class="flex items-center gap-3 px-3 py-2 text-sm font-medium rounded-md text-text-muted hover:bg-surface-highlight hover:text-text-main transition-colors group"
        >
          <Icon name={item.icon} class="group-hover:text-text-main transition-colors" />
          {item.name}
        </a>
      {/each}
    </nav>
  </div>
</aside>