<script>
    import OptionsBar from "$lib/components/(public)/OptionsBar.svelte";
    import CatalogGrid from "$lib/components/(public)/catalog/Grid.svelte";
    import Slider from "$lib/components/(public)/catalog/Slider.svelte";
    import Pagination from "$lib/components/Pagination.svelte";

    export let data;

    let catalogLimit = 40;
    let category = 0;

    let catalog = new Promise(() => {
        }),
        slider = new Promise(() => {
        }),
        count = new Promise(() => {
        });

    let sorting = 'created_at';

    const GetCatalog = async (offset, sort) => {
        sorting = sort

        count = await GetCount(sort)

        const request = data.content.applyParams({
            order: sort,
            direction: 'desc',
            offset: offset,
            limit: catalogLimit,
            category: category,
        })

        return request.get();
    }

    const GetCount = (sort) => {
        const request = data.content.applyParams({
            order: sort,
            direction: 'desc',
            category: category,
        })

        return request.count();
    }

    const GetSlider = () => {
        return data.content.applyParams({
            order: 'world_premiere',
            direction: 'desc',
            limit: 24,
        }).get();
    }

    catalog = GetCatalog(0, sorting)
    slider = GetSlider()
</script>

{#await slider}
    <div class="py-1">
        <div class="grid grid-cols-8 gap-x-3 animate-pulse w-[calc(100%_+_20rem)] ml-[-10rem] py-8">
            {#each Array(8) as _}
                <div class="bg-gray-200 rounded-2xl dark:bg-neutral-700 rounded-2xl overflow-hidden opacity-50">
                    <div class="h-[25rem] bg-gray-200 rounded-2xl"></div>
                </div>
            {/each}
        </div>
    </div>
{:then slider}
    <Slider data="{ slider }"/>
{/await}

<div class="max-w-7xl mx-auto space-y-4">
    <OptionsBar
        bind:category={category}
        onChange="{ (o, s) => catalog = GetCatalog(0, s) }"/>

    {#await catalog}
        <div class="grid grid-cols-5 gap-x-4 gap-y-8 animate-pulse">
            {#each Array(catalogLimit) as _}
                <div class="bg-gray-200 rounded-2xl dark:bg-neutral-700 rounded-2xl overflow-hidden opacity-50">
                    <div class="h-[22.5rem] bg-gray-200 rounded-2xl"></div>
                </div>
            {/each}
        </div>
    {:then catalog}
        <CatalogGrid data="{catalog}"/>
    {/await}

    <div class="flex justify-center pt-4">
        <Pagination
                total={count}
                limit={catalogLimit}
                onChange={ o => catalog = GetCatalog(o, sorting) }
        />
    </div>
</div>