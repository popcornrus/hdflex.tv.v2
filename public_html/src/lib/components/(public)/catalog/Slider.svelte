<script>
    import {register} from 'swiper/element/bundle';
    import {onMount} from "svelte";
    import Rating from "$lib/components/(public)/catalog/ratings/VerticalList.svelte";
    import Image from "$lib/components/Image.svelte";
    import {browser} from "$app/environment";
    import LatestSeriesUpdate from "$lib/components/(public)/catalog/tv-series/LatestSeriesUpdate.svelte";

    /** @type {import('./$types').PageData} */
    export let data;

    const nextSlide = () => {
        const swiperEl = document.querySelector('swiper-container').swiper;
        swiperEl.slideNext();
    }

    const prevSlide = () => {
        const swiperEl = document.querySelector('swiper-container').swiper;
        swiperEl.slidePrev();
    }

    onMount(async () => {
        if (browser) {
            register();

            const swiperEl = document.querySelector('swiper-container');

            const swiperParams = {
                slidesPerView: 8,
                slidesPerGroup: 6,
                loop: true,
                spaceBetween: 10,
                loopedSlides: 8,
                updateOnWindowResize: true,
            };

            Object.assign(swiperEl, swiperParams);
            swiperEl.initialize();
        }
    })
</script>

<div class="overflow-hidden relative">
    <div class="text-white">
        <div class="absolute top-0 left-0 bg-gradient-to-r from-slate-950 w-28 h-full z-10 flex items-center justify-start">
            <button on:click={prevSlide} class="hover:scale-125 transition duration-300 h-full px-4 w-full">
                <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" class="size-14">
                    <path stroke-linecap="round" stroke-linejoin="round" d="M15.75 19.5 8.25 12l7.5-7.5" />
                </svg>
            </button>
        </div>
        <div class="absolute top-0 right-0 bg-gradient-to-l from-slate-950 w-28 h-full z-10 flex items-center justify-end">
            <button on:click={nextSlide} class="hover:scale-125 transition duration-300 h-full w-full px-4">
                <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" class="size-14">
                    <path stroke-linecap="round" stroke-linejoin="round" d="m8.25 4.5 7.5 7.5-7.5 7.5" />
                </svg>
            </button>
        </div>
    </div>
    <div class="py-8 w-[calc(100%_+_20rem)] -ml-[10rem]">
        <swiper-container init="false">
            {#each data as item}
                <swiper-slide class="space-y-2">
                    <div class="shadow shadow-slate-900/80 rounded-2xl overflow-hidden hover:shadow-slate-900/90 hover:shadow-lg overflow-hidden transition duration-150">
                        <div class="overflow-hidden relative group">
                            <a href="/catalog/{item.url}"
                               class="rounded-2xl overflow-hidden block h-[25rem] scale-105 group-hover:scale-100 transition duration-150">
                                <Image width={295} height={450} src="{item.poster}"/>
                            </a>
                            <div class="absolute top-0 right-0 space-y-2 opacity-25 group-hover:opacity-100 transition duration-150">
                                <div class="flex justify-end overflow-hidden rounded-tr-2xl">
                                    <button class="py-2 px-4 bg-slate-800 text-white w-max hover:text-blue-500 translate duration-300 rounded-bl-xl">
                                        <svg xmlns="http://www.w3.org/2000/svg" fill="currentColor" viewBox="0 0 24 24"
                                             stroke-width="1.5" stroke="currentColor" class="w-5 h-5">
                                            <path stroke-linecap="round" stroke-linejoin="round"
                                                  d="M17.593 3.322c1.1.128 1.907 1.077 1.907 2.185V21L12 17.25 4.5 21V5.507c0-1.108.806-2.057 1.907-2.185a48.507 48.507 0 0 1 11.186 0Z"/>
                                        </svg>
                                    </button>
                                </div>
                                <Rating data={item.external_ids}/>
                            </div>
                            <div class="absolute bottom-2 w-[90%] inset-x-0 mx-auto px-4 py-4 bg-gradient-to-t from-black to-black/30 space-y-2 backdrop-blur-sm rounded-xl">
                                {#if item.IsTvSeriesType()}
                                    <div class="absolute top-0 -translate-y-1/2 inset-x-0 mx-auto w-max">
                                        <LatestSeriesUpdate data={item}/>
                                    </div>
                                {/if}
                                <ul class="flex justify-center gap-x-4 text-xs dark:text-white/50 items-center">
                                    <li>
                                        <a href="" class="hover:text-white transition duration-150">{ item.year }</a>
                                    </li>
                                    {#if item.genres !== null}
                                        <li class="capitalize">
                                            <a href=""
                                               class="hover:text-white transition duration-150">{ item.genres[0].title }</a>
                                        </li>
                                    {/if}
                                    {#if item.countries !== null}
                                        <li>
                                            <a href=""
                                               class="hover:text-white transition duration-150">{ item.countries[0].title }</a>
                                        </li>
                                    {/if}
                                </ul>
                            </div>
                        </div>
                    </div>
                    <div class="flex gap-x-2 justify-center text-white h-12 items-center px-4">
                        <h3 class="text-balance hyphens-auto text-md font-bold text-shadow text-center"
                            title="{item.ru_title}" lang="ru">
                            <a href="/catalog/{item.url}">{ item.ru_title.slice(0, 48) }{ item.ru_title.length > 48 ? '...' : '' }</a>
                        </h3>
                    </div>
                </swiper-slide>
            {/each}
        </swiper-container>
    </div>
</div>