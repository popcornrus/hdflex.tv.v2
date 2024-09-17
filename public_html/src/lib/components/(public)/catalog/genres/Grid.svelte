<div class="max-w-7xl mx-auto text-white space-y-8">
    <h2 class="text-3xl font-semibold text-white/80">Жанры</h2>
    <div class="">
        <div class="grid grid-cols-5 gap-4">
            {#each Array(9) as genre}
                <div class="rounded-xl overflow-hidden relative group shadow-xl">
                    <img src="https://flixtv.volkovdesign.com/main/img/category/1.jpg" class="brightness-75 group-hover:scale-125 group-hover:brightness-50 transition delay-150" alt="">
                    <div class="absolute bottom-2 px-3 left-0 w-full">
                        <div class="flex items-center justify-between">
                            <p class="text-xl">Животные</p>
                            <div class="py-1 px-3 bg-blue-950/80 backdrop-blur rounded-lg">
                                240
                            </div>
                        </div>
                    </div>
                </div>
            {/each}
            <div class="rounded-xl overflow-hidden relative group shadow-xl bg-gray-950 group">
                <div class="flex items-center justify-center w-full h-full">
                    <span class="text-slate-800/80 group-hover:text-slate-600 transition duration-300">
                        <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" class="w-24 h-24">
                          <path stroke-linecap="round" stroke-linejoin="round" d="M6.75 12a.75.75 0 1 1-1.5 0 .75.75 0 0 1 1.5 0ZM12.75 12a.75.75 0 1 1-1.5 0 .75.75 0 0 1 1.5 0ZM18.75 12a.75.75 0 1 1-1.5 0 .75.75 0 0 1 1.5 0Z" />
                        </svg>
                    </span>
                </div>
            </div>
        </div>
    </div>
</div>