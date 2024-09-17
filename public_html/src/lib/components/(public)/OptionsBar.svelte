<script>
    export let data;
    let sort = 'created_at';

    export let category = 0;
    export let onChange = (o, s) => {};

    const setSortBy = (s) => {
        sort = s;
        onChange(0, s);
    }

    const setCategory = (c) => {
        category = c;
        onChange(0, sort);
    }
</script>

<div class="py-4 mb-8 border-b border-slate-800 relative z-20">
    <div class="flex items-center justify-between">
        <div class="p-2 rounded-xl text-white/80">
            <ul class="flex flex-col sm:flex-row gap-x-4">
                <li
                    class:border-b={sort === 'created_at'}
                    class:text-white={sort === 'created_at'}>
                    <button on:click={ () => setSortBy('created_at') }>
                        <span class="block w-full text-sm">
                            Последние поступления
                        </span>
                    </button>
                </li>
                <li
                    class:border-b={sort === 'popularity'}
                    class:text-white={sort === 'popularity'}>
                    <button on:click={ () => setSortBy('popularity') }>
                        <span class="block w-full text-sm">
                            Популярные
                        </span>
                    </button>
                </li>
                <li
                    class:border-b={sort === 'last_watch_at'}
                    class:text-white={sort === 'last_watch_at'}>
                    <button on:click={ () => setSortBy('last_watch_at') }>
                        <span class="block w-full text-sm">
                            Сейчас смотрят
                        </span>
                    </button>
                </li>
            </ul>
        </div>
        <div class="text-white/80">
            <ul class="flex gap-x-4">
                <li
                    class:border-b={category === 0}
                    class:text-white={category === 0}>
                    <button on:click={ () => setCategory(0) }>
                        <span class="block w-full text-sm">
                            Все
                        </span>
                    </button>
                </li>
                <li
                    class:border-b={category === 1}
                    class:text-white={category === 1}>
                    <button on:click={ () => setCategory(1) }>
                        <span class="block w-full text-sm">
                            Фильмы
                        </span>
                    </button>
                </li>
                <li
                    class:border-b={category === 2}
                    class:text-white={category === 2}>
                    <button on:click={ () => setCategory(2) }>
                        <span class="block w-full text-sm">
                            Сериалы
                        </span>
                    </button>
                </li>
                <li class="relative group"
                    class:border-b={ [3,4].includes(category) }
                    class:text-white={ [3,4].includes(category) }>
                    <span class="w-full text-sm cursor-pointer group-focus:border-b">
                        Аниме
                    </span>
                    <div class="absolute top-full inset-x-0 mx-auto -translate-x-1/3 w-max pt-2 opacity-0 h-0 overflow-y-hidden group-hover:opacity-100 group-hover:h-auto shadow transition duration-300">
                        <div class="bg-slate-900 py-3 px-4 rounded shadow">
                            <div class="flex gap-x-4">
                                <button
                                        class="pb-0.5"
                                        class:border-b={ category === 3 }
                                        on:click={ () => setCategory(3) }>
                                    <span class="block w-full text-sm">
                                        Фильмы
                                    </span>
                                </button>
                                <button
                                        class="pb-0.5"
                                        class:border-b={ category === 4 }
                                        on:click={ () => setCategory(4) }>
                                    <span class="block w-full text-sm">
                                        Сериалы
                                    </span>
                                </button>
                            </div>
                        </div>
                    </div>
                </li>
                <li
                        class:border-b={category === 5}
                        class:text-white={category === 5}>
                    <button on:click={ () => setCategory(5) }>
                        <span class="block w-full text-sm">
                            ТВ шоу
                        </span>
                    </button>
                </li>
            </ul>
        </div>
    </div>
</div>