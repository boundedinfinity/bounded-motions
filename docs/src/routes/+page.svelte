<script lang="ts">
    import { init } from "./model";
    import Tree from "./tree.svelte";

    const global = init();
    // $inspect(global);
</script>

<Tree binding={global} />
