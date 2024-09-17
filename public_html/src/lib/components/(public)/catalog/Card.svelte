<script>
    import Image from "$lib/components/Image.svelte";
    import Rating from "$lib/components/(public)/catalog/ratings/VerticalList.svelte";
    import LatestSeriesUpdate from "$lib/components/(public)/catalog/tv-series/LatestSeriesUpdate.svelte";

    export let data;

    let highestRating = 0

    const GetHighestRating = () => {
        if (data.external_ids === null) {
            highestRating = 0
            return
        }

        for (let i = 0; i < data.external_ids.length; i++) {
            if (data.external_ids[i].rating > 0 && data.external_ids[i].rating > highestRating) {
                highestRating = data.external_ids[i].rating
            }
        }

        return 0
    }
</script>

<div class="group hover:scale-115 transition duration-150">
    <div class="rounded-xl space-y-2 relative">
        <div class="relative rounded-2xl overflow-hidden group-hover:shadow-xl duration-150 transition">
            <a href="/catalog/{data.url}" class="block overflow-hidden h-[22.5rem] scale-105 group-hover:scale-100 transition duration-150">
                <Image src="{data.poster}" />
            </a>
            <div class="absolute top-0 right-0 space-y-2 opacity-25 group-hover:opacity-100 transition duration-150">
                <div class="flex justify-end overflow-hidden rounded-tr-2xl">
                    <button class="py-2 px-4 bg-slate-800 text-white w-max hover:text-blue-500 translate duration-300 rounded-bl-xl">
                        <svg xmlns="http://www.w3.org/2000/svg" fill="currentColor" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" class="w-5 h-5">
                            <path stroke-linecap="round" stroke-linejoin="round" d="M17.593 3.322c1.1.128 1.907 1.077 1.907 2.185V21L12 17.25 4.5 21V5.507c0-1.108.806-2.057 1.907-2.185a48.507 48.507 0 0 1 11.186 0Z" />
                        </svg>
                    </button>
                </div>
                <Rating data={data.external_ids} />
            </div>
        </div>
        <div class="absolute bottom-2 w-[90%] inset-x-0 mx-auto px-4 py-4 bg-gradient-to-t from-black to-black/30 space-y-2 backdrop-blur-sm rounded-xl transition duration-150">
            {#if data.IsTvSeriesType()}
                <div class="absolute top-0 -translate-y-1/2 inset-x-0 mx-auto w-max">
                    <LatestSeriesUpdate data={data} />
                </div>
            {/if}
            <ul class="flex justify-center gap-x-4 text-xs dark:text-white/50">
                {#if GetHighestRating()}
                    <li class="flex gap-x-1 text-yellow-600">
                        <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="2" stroke="currentColor" class="w-4 h-4">
                            <path stroke-linecap="round" stroke-linejoin="round" d="M11.48 3.499a.562.562 0 0 1 1.04 0l2.125 5.111a.563.563 0 0 0 .475.345l5.518.442c.499.04.701.663.321.988l-4.204 3.602a.563.563 0 0 0-.182.557l1.285 5.385a.562.562 0 0 1-.84.61l-4.725-2.885a.562.562 0 0 0-.586 0L6.982 20.54a.562.562 0 0 1-.84-.61l1.285-5.386a.562.562 0 0 0-.182-.557l-4.204-3.602a.562.562 0 0 1 .321-.988l5.518-.442a.563.563 0 0 0 .475-.345L11.48 3.5Z" />
                        </svg>
                        <span>{ GetHighestRating() }</span>
                    </li>
                {/if}
                {#if data.genres !== null}
                    <li class="capitalize">
                        <a href="" class="hover:text-white transition duration-150">{ data.genres[0].title }</a>
                    </li>
                {/if}
                {#if data.countries !== null}
                    <li>
                        <a href="" class="hover:text-white transition duration-150">{ data.countries[0].title }</a>
                    </li>
                {/if}
                <li>
                    <a href="" class="hover:text-white transition duration-150">{ data.year }</a>
                </li>
            </ul>
        </div>
    </div>
    <a href="/catalog/{data.url}" class="flex transition dark:text-white/80 group-hover:dark:text-white-500/80 transition text-md tracking-wide font-semibold h-12 items-center justify-center text-center leading-4 px-2">
        { data.ru_title.slice(0, 48) }{ data.ru_title.length > 48 ? '...' : '' }
    </a>
</div>