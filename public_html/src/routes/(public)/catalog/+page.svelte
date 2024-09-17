<script>
    import GenresGrid from "$lib/components/(public)/catalog/genres/Grid.svelte"
    import ItemsGrid from "$lib/components/(public)/catalog/Grid.svelte";
    import Pagination from "$lib/components/Pagination.svelte";
    import OptionsBar from "$lib/components/(public)/OptionsBar.svelte";
</script>

<div class="bg-slate-950/60 py-12">
    <GenresGrid />
</div>

<div class="max-w-7xl mx-auto text-white">
    <div class="py-12">
        <OptionsBar genresShow="{false}" />
    </div>
    <ItemsGrid />

    <div class="py-8 flex justify-center">
        <Pagination
            total="{93}"
            limit="{10}"
            offset="{1}"
        />
    </div>
</div>