<script>
    import {Genre} from "$lib/model/Genre.js";
    import {Country} from "$lib/model/Country.js";

    import {onMount} from "svelte";

    const genresModel = new Genre()
    const countriesModel = new Country()

    let genres = {items: []}
    let countries = {items: []}

    onMount(async () => {
        genres = await genresModel.get()
        countries = await countriesModel.get()
    })
</script>

<header class="py-4 bg-slate-900/50 px-8 relative z-50 backdrop-blur-2xl rounded-b-xl">
    <div class="flex justify-between items-center gap-x-12 relative">
        <a href="/" class="flex items-end gap-x-1 font-bold">
            <div class="text-xl text-white">HDFLEX</div>
            <div class="text-sm text-blue-500">TV</div>
        </a>
        <div class="flex justify-between text-white w-full">
            <ul class="flex justify-between w-96 mx-auto">
                <li class="group hover:after:block after:hidden after:left-0 after:w-full after:h-12 after:absolute">
                    <div class="flex items-center gap-x-2 group-hover:text-blue-500 cursor-pointer">
                        <span>Фильмы</span>
                        <span class="transition group-hover:rotate-180">
                            <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" class="w-3 h-3">
                                <path stroke-linecap="round" stroke-linejoin="round" d="m19.5 8.25-7.5 7.5-7.5-7.5" />
                            </svg>
                        </span>
                    </div>
                    <div class="absolute top-[calc(100%+2rem)] backdrop-blur-xl dark:bg-slate-950/95 inset-x-0 rounded-xl py-4 px-8 shadow-lg shadow-slate-900 group-hover:opacity-100 opacity-0 group-hover:visible invisible transition duration-300 transform z-10 delay-300">
                        <div class="flex">
                            <div class="grid gap-x-4 w-full gap-y-4">
                                <div class="space-y-4 border-b border-gray-500 pb-4">
                                    <h3 class="font-bold text-xl">Жанры</h3>
                                    <ul class="text-sm space-y-2 columns-6">
                                        {#each genres.items as genre}
                                           <li class="text-gray-400">
                                               <a href="" class="hover:text-blue-500 cursor-pointer transition" style="text-transform:capitalize">
                                                   { genre.title }
                                               </a>
                                           </li>
                                        {/each}
                                    </ul>
                                </div>
                                <div class="grid grid-cols-4">
                                    <div class="space-y-4 col-span-3">
                                        <h3 class="font-bold text-xl">Страны</h3>
                                        <div class="max-h-72 overflow-y-scroll">
                                            <ul class="text-sm space-y-2 columns-3">
                                                {#each countries.items as country}
                                                    <li class="text-gray-400">
                                                        <a href="" class="hover:text-blue-500 cursor-pointer transition">
                                                            { country.title }
                                                        </a>
                                                    </li>
                                                {/each}
                                            </ul>
                                        </div>
                                    </div>
                                    <div class="space-y-4">
                                        <h3 class="font-bold text-xl">Годы</h3>
                                        <ul class="text-sm space-y-2">
                                            {#each [2024, 2023, 2022, 2021, 2020, 2019] as year}
                                                <li class="text-gray-400">
                                                    <a href="" class="hover:text-blue-500 cursor-pointer transition">
                                                        { year } год
                                                    </a>
                                                </li>
                                            {/each}
                                        </ul>
                                    </div>
                                </div>
                            </div>
                            <div class="w-1/4 border-l border-gray-500 ps-8">
                                <div class="w-full border-e-2 border-gray-200 dark:border-gray-700">
                                    <nav class="-me-0.5 flex flex-col space-y-3">
                                        <a class="py-1 pe-4 inline-flex items-center gap-2 border-e-2 border-transparent hover:border-blue-500 text-sm whitespace-nowrap text-gray-300 hover:text-blue-600 focus:outline-none focus:text-blue-800" href="#" aria-current="page">
                                            Каталог
                                        </a>
                                        <a class="py-1 pe-4 inline-flex items-center gap-2 border-e-2 border-transparent hover:border-blue-500 text-sm whitespace-nowrap text-gray-300 hover:text-blue-600 focus:outline-none focus:text-blue-600" href="#">
                                            Подборки
                                        </a>
                                        <a class="py-1 pe-4 inline-flex items-center gap-2 border-e-2 border-transparent hover:border-blue-500 text-sm whitespace-nowrap text-gray-300 hover:text-blue-600 focus:outline-none focus:text-blue-600" href="#">
                                            Сейчас смотрят
                                        </a>
                                        <a class="py-1 pe-4 inline-flex items-center gap-2 border-e-2 border-transparent hover:border-blue-500 text-sm whitespace-nowrap text-gray-300 hover:text-blue-600 focus:outline-none focus:text-blue-600" href="#">
                                            Популярные
                                        </a>
                                        <a class="py-1 pe-4 inline-flex items-center gap-2 border-e-2 border-transparent hover:border-blue-500 text-sm whitespace-nowrap text-gray-300 hover:text-blue-600 focus:outline-none focus:text-blue-600" href="#">
                                            В ожидании
                                        </a>
                                    </nav>
                                </div>
                            </div>
                        </div>
                    </div>
                </li>
            </ul>
        </div>
        <div class="flex gap-x-4 items-center w-96">
            <div class="flex rounded-lg shadow-sm relative">
                <input type="text" class="py-3 px-4 block w-full border-gray-200 shadow-sm rounded-lg text-sm focus:z-10 focus:border-blue-500 focus:ring-blue-500 disabled:opacity-50 disabled:pointer-events-none dark:bg-slate-800 dark:border-gray-700 dark:text-gray-400 dark:focus:ring-gray-600">
                <button type="button" class="absolute w-[2.875rem] h-[2.875rem] flex-shrink-0 inline-flex justify-center items-center gap-x-2 text-sm font-semibold rounded-s-md border border-transparent text-white disabled:opacity-50 disabled:pointer-events-none">
                    <svg class="flex-shrink-0 size-4" xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                        <circle cx="11" cy="11" r="8"></circle>
                        <path d="m21 21-4.3-4.3"></path>
                    </svg>
                </button>
            </div>

            <div class="text-white">
                <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" class="w-6 h-6">
                    <path stroke-linecap="round" stroke-linejoin="round" d="M16.5 10.5V6.75a4.5 4.5 0 1 0-9 0v3.75m-.75 11.25h10.5a2.25 2.25 0 0 0 2.25-2.25v-6.75a2.25 2.25 0 0 0-2.25-2.25H6.75a2.25 2.25 0 0 0-2.25 2.25v6.75a2.25 2.25 0 0 0 2.25 2.25Z" />
                </svg>
            </div>
        </div>
    </div>
</header>