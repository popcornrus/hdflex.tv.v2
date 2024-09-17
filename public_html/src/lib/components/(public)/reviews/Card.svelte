<div class="space-y-4">
    <div class="flex w-full items-center justify-between">
        <div class="flex items-center gap-x-4">
            <div class="w-max">
                <span class="inline-flex items-center justify-center size-[46px] text-sm font-semibold leading-none rounded-full bg-gray-50 text-gray-500 dark:bg-white/10 dark:text-white">
                  AC
                </span>
            </div>
            <div class="">
                <p class="text-gray-400 text-lg font-semibold">user#123321</p>
                <p class="text-gray-400 text-sm">28.04.2024 13:56</p>
            </div>
        </div>
        <div class="flex text-gray-400">
            {#each Array(5) as k, i}
                <span class:text-yellow-500={ i < 3 }>
                    <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" class="w-6 h-6">
                        <path stroke-linecap="round" stroke-linejoin="round" d="M11.48 3.499a.562.562 0 0 1 1.04 0l2.125 5.111a.563.563 0 0 0 .475.345l5.518.442c.499.04.701.663.321.988l-4.204 3.602a.563.563 0 0 0-.182.557l1.285 5.385a.562.562 0 0 1-.84.61l-4.725-2.885a.562.562 0 0 0-.586 0L6.982 20.54a.562.562 0 0 1-.84-.61l1.285-5.386a.562.562 0 0 0-.182-.557l-4.204-3.602a.562.562 0 0 1 .321-.988l5.518-.442a.563.563 0 0 0 .475-.345L11.48 3.5Z" />
                    </svg>
                </span>
            {/each}
        </div>
    </div>
    <div class="bg-slate-900 py-3 px-4 rounded-md text-gray-300">
        Lorem ipsum dolor sit amet, consectetur adipisicing elit. Cum illum inventore iure quis quo rerum? A commodi cum deserunt excepturi, ipsa itaque iusto laboriosam molestiae quidem repudiandae sequi soluta unde?
    </div>
</div>