<script>
    import Form from "$lib/components/(public)/reviews/Form.svelte";
    import ReviewsSection from "$lib/components/(public)/reviews/Section.svelte";
    import CatalogGrid from "$lib/components/(public)/catalog/Grid.svelte";
    import Image from "$lib/components/Image.svelte";
    import HorizontalList from "$lib/components/(public)/catalog/ratings/HorizontalList.svelte"
    import {ru} from "date-fns/locale";
    import {format} from "date-fns";
    import {Content} from "$lib/model/Content.js";
    import {onMount} from "svelte";
    import {page} from "$app/stores";

    let data = {
        content: new Promise(() => {
        })
    };

    let highestRating = 0

    onMount(async () => {
        const instance = new Content()
        data.content = await instance.findBySlug($page.params.slug)
    })
</script>

{#await data.content}
{:then content}
    <section class="absolute top-0 left-0 -z-50 overflow-hidden w-full min-h-[40rem]">
        <div class="brightness-[.15] -translate-y-1/4">
            <Image height={1280} src="{content.backdrop}"/>
        </div>
        <div class="absolute h-full w-full top-0 left-0 bg-gradient-to-b from-black/0 to-slate-900"></div>
    </section>

    <section class="max-w-5xl mx-auto text-white relative">
        <div class="py-32 items-center relative z-10 flex items-center gap-x-12">
            <div class="rounded-xl overflow-hidden h-[22.5rem] min-w-[15.313rem]">
                <Image src="{content.poster}"/>
            </div>
            <div class="space-y-6 w-full">
                <div class="space-y-3">
                    <h1 class="text-5xl font-semibold">{content.ru_title}</h1>
                    {#if content.slogan.length > 0}
                        <p class="italic">"{content.slogan}"</p>
                    {/if}
                </div>
                <div class="space-y-1 w-2/3 py-4">
                    <div class="flex justify-between items-center">
                        <span class="font-bold">Дата выхода</span>
                        <span class="text-gray-400">
                        <time datetime="{content.world_premiere}">
                            {format(content.world_premiere, "dd MMMM", {locale: ru})}
                            <a href=""
                               class="underline text-blue-600 hover:text-white transition duration-150">{format(content.world_premiere, "yyyy", {locale: ru})}</a>
                        </time>
                    </span>
                    </div>
                    {#if content.duration > 0}
                        <div class="flex justify-between items-center">
                            <span class="font-bold">Время</span>
                            <span class="text-gray-400">~{content.duration} мин.</span>
                        </div>
                    {/if}
                </div>

                <div class="flex">
                    <HorizontalList data={content.external_ids}/>
                </div>
            </div>
        </div>
    </section>

    <section
            class="relative bg-slate-950 before:absolute before:bottom-full before:bg-gradient-to-t before:from-slate-950 before:left-0 before:h-36 before:w-full before:content-[''] space-y-8">
        {#if content.description}
            <div class="border-b border-slate-600/60 py-16">
                <div class="container space-y-4 mx-auto text-center">
                    <h4 class="text-3xl text-white font-bold">О чём «{ content.ru_title }»</h4>
                    <p class="text-lg text-white w-2/3 mx-auto">{@html content.description}</p>
                </div>
            </div>
        {/if}
        <section class="overflow-hidden relative z-10" id="player-container">
            <div class="rounded-xl overflow-hidden transition duration-150 container mx-auto">
                <iframe src="{content.iframe_url}"
                        class="w-full min-h-[40rem]" frameborder="0" allowfullscreen></iframe>
            </div>
        </section>
        <div class="container mx-auto grid grid-cols-4 gap-x-8 gap-y-8">
            <div class="col-span-3 space-y-8" class:col-span-4={content.similar === undefined}>
                <section class="relative">
                    <div class="text-white mx-auto">
                        <div class="relative z-20 space-y-8">
                            <div class="grid grid-cols-5 gap-x-12">
                                {#if content.casts !== null}
                                    <div class="w-full relative col-span-4 space-y-4">
                                        <div class="space-y-4">
                                            <div class="flex items-end justify-between">
                                                <h4 class="text-2xl">В главных ролях</h4>
                                                <a href=""
                                                   class="text-md w-max hover:text-blue-600 transition duration-150 flex items-center gap-x-2">
                                                    <span>Полный актёрский и съёмочный состав</span>
                                                    <svg xmlns="http://www.w3.org/2000/svg" fill="none"
                                                         viewBox="0 0 24 24"
                                                         stroke-width="1.5" stroke="currentColor" class="w-4 h-4">
                                                        <path stroke-linecap="round" stroke-linejoin="round"
                                                              d="M13.19 8.688a4.5 4.5 0 0 1 1.242 7.244l-4.5 4.5a4.5 4.5 0 0 1-6.364-6.364l1.757-1.757m13.35-.622 1.757-1.757a4.5 4.5 0 0 0-6.364-6.364l-4.5 4.5a4.5 4.5 0 0 0 1.242 7.244"/>
                                                    </svg>
                                                </a>
                                            </div>
                                            <div class="overflow-y-hidden overflow-x-auto pb-4 w-full rounded-xl"
                                                 style="">
                                                <div class="w-max flex gap-x-4">
                                                    {#each content.casts.slice(0, 8) as cast}
                                                        <div class="w-44 flex relative flex-col bg-white border border-t-4 border-t-blue-600 border-x-0 shadow-sm rounded-xl dark:bg-neutral-900 dark:border-neutral-700 dark:border-t-blue-500 dark:shadow-neutral-700/70">
                                                            <div class="overflow-hidden h-full w-full rounded-t">
                                                                <Image height={265} src="{cast.person.image}"/>
                                                            </div>
                                                            <div class="p-2 absolute -bottom-0.5 bg-neutral-900/80 w-[calc(100%_+_0.2rem)] -left-0.5 backdrop-blur rounded-b">
                                                                <p class="text-center">{cast.person.name}</p>
                                                                <p class="text-gray-500 dark:text-neutral-400 text-center text-sm">
                                                                    {cast.character}
                                                                </p>
                                                            </div>
                                                        </div>
                                                    {/each}
                                                    <div class="w-44 flex items-center group cursor-pointer hover:bg-slate-950 rounded-xl transition duration-300">
                                                        <div class="flex flex-wrap justify-center w-full group-hover:text-blue-500 transition duration-300">
                                                            <svg xmlns="http://www.w3.org/2000/svg" fill="none"
                                                                 viewBox="0 0 24 24"
                                                                 stroke-width="1.5" stroke="currentColor"
                                                                 class="w-12 h-12">
                                                                <path stroke-linecap="round" stroke-linejoin="round"
                                                                      d="M6.75 12a.75.75 0 1 1-1.5 0 .75.75 0 0 1 1.5 0ZM12.75 12a.75.75 0 1 1-1.5 0 .75.75 0 0 1 1.5 0ZM18.75 12a.75.75 0 1 1-1.5 0 .75.75 0 0 1 1.5 0Z"/>
                                                            </svg>
                                                            <p class="w-full text-center">Смотреть ещё</p>
                                                        </div>
                                                    </div>
                                                </div>
                                            </div>
                                        </div>
                                    </div>
                                {/if}
                                <div class="space-y-4 justify-between items-end max-w-[95%] w-full mx-auto"
                                     class:col-span-5={content.casts === null}
                                     class:flex={content.casts === null}>
                                    <div class="space-y-2">
                                        <p class="font-bold">Исходное название</p>
                                        <h2>{ content.orig_title }</h2>
                                    </div>
                                    {#if content.crew !== null}
                                        {#if content.crew.filter(crew => crew.job === 'Director').length > 0}
                                            <div class="space-y-2">
                                                <p class="font-bold">Режиссёр</p>
                                                <ul>
                                                    {#each content.crew.filter(crew => crew.job === 'Director') as director}
                                                        <li>
                                                            <a href=""
                                                               class="hover:underline transition duration-150 hover:text-blue-500">{ director.person.name }</a>
                                                        </li>
                                                    {/each}
                                                </ul>
                                            </div>
                                        {/if}
                                    {/if}
                                    {#if content.genres !== null}
                                        <div class="space-y-2">
                                            <p class="font-bold">Жанры</p>
                                            <div class="flex flex-wrap gap-2">
                                                {#each content.genres as genre}
                                                    <span class="inline-flex items-center gap-x-1.5 py-1.5 px-3 rounded-full text-xs font-medium bg-blue-800/30 text-blue-500 hover:bg-blue-500 hover:text-white transition duration-150">{ genre.title }</span>
                                                {/each}
                                            </div>
                                        </div>
                                    {/if}
                                    {#if content.countries !== null}
                                        <div class="space-y-2">
                                            <p class="font-bold">Страны</p>
                                            <div class="flex flex-wrap gap-2">
                                                {#each content.countries as country}
                                                    <span class="inline-flex items-center gap-x-1.5 py-1.5 px-3 rounded-full text-xs font-medium bg-blue-800/30 text-blue-500 hover:bg-blue-500 hover:text-white transition duration-150">{ country.title }</span>
                                                {/each}
                                            </div>
                                        </div>
                                    {/if}
                                </div>
                            </div>
                            <div class="grid grid-cols-6 gap-x-12 gap-y-4 py-8">
                                <div class="col-span-5">
                                    <div class="flex items-end justify-between">
                                        <h4 class="text-2xl">Комментарии (48)</h4>
                                    </div>
                                </div>
                                <div class="col-span-6 mb-8">
                                    <div class="border border-slate-800 rounded-lg p-4 w-full">
                                        <Form/>
                                    </div>
                                    <div class="pt-8">
                                        <ReviewsSection/>
                                    </div>
                                </div>
                            </div>
                        </div>
                    </div>
                </section>
            </div>
            {#if content.similar !== undefined}
                <div class="col-span-1">
                    <div class="col-span-2 space-y-4">
                        <h4 class="text-2xl text-white">Похожие фильмы</h4>
                        <!--<CatalogGrid grid="{ 2 }"/>-->
                    </div>
                </div>
            {/if}
        </div>
    </section>
{/await}