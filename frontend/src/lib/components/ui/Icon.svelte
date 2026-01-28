<script lang="ts">
  let { name, size = '5', class: className = '' } = $props<{
    name: string;
    size?: string;
    class?: string;
  }>();
</script>

<i class={`icon-[duotone--${name}] size-${size} ${className}`}></i>