<script lang="ts">
    import type { PageProps } from './$types';
    import { init, KeyBinding } from "./model.svelte";
    import Tree from "./tree.svelte";

    let { data }: PageProps = $props();
    const global = init(data.config);
    let expanded = $state(false);
    // $inspect(global);

    function fn() {
        expanded = !expanded;
        global.walkDown(function (binding: KeyBinding): boolean {
            binding.expanded = expanded;
            return true;
        });
    }
</script>

<button onclick={fn}>Expand All</button>
<Tree binding={global} />
