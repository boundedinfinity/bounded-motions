<script lang="ts">
    import Self from "./tree.svelte";
    import { type KeyBinding } from "./model";
    const props: { binding: KeyBinding } = $props();
    const path = props.binding
        .path()
        .filter((i) => i.key)
        .map((i) => i.key)
        .join(" > ");
</script>

<ul class="item">
    <li>
        {props.binding.key}
        {#if props.binding.key || props.binding.name}
            :
        {/if}
        {props.binding.name}

        {#if path.length > 0}
        : <b>{path}</b>
        {/if}

        {#each props.binding.children as child}
            <Self binding={child}></Self>
        {/each}
    </li>
</ul>

<style>
</style>
