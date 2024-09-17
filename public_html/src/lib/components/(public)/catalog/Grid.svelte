<script>
    import Card from "$lib/components/(public)/catalog/Card.svelte";

    export let data;

    export let grid = 5;
</script>

<div class="grid grid-cols-{ grid } gap-x-4 gap-y-8">
    {#each data as item}
        <Card data={item} />
    {/each}
</div>