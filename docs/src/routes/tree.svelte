<script lang="ts">
    import Self from "./tree.svelte";
    import { type KeyBinding } from "./model.svelte";
    const { binding }: { binding: KeyBinding } = $props();

    const path: string =
        binding.children.length === 0
            ? binding
                  .path()
                  .filter((i) => i.key)
                  .map((i) => i.key)
                  .join(" > ")
            : "";
</script>

<ul class="item">
    <li>
        {#if binding.key}
            <code>{binding.key}</code>
        {/if}

        {#if binding.key || binding.name}
            :
        {/if}

        {binding.name}

        {#if path.length > 0}
            : <code>{path}</code>
        {/if}

        {#if binding.children.length > 0}
            <button onclick={() => binding.toggleExpended()}>
                ({binding.children.length})
            </button>

            {#if binding.expanded}
                {#each binding.children as child}
                    <Self binding={child}></Self>
                {/each}
            {/if}
        {/if}
    </li>
</ul>

<style>
    code {
        border-color: lightgray;
        border-radius: 1px;
        border-style: solid;
        padding: 0.25em;
    }

    li {
        margin: 0.25em;
    }
</style>
