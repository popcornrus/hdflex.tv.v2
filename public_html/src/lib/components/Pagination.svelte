<script>
    export let total = 0,
        limit = 0;

    export let onChange = (offset) => {};

    let page = 1;

    const goForward = () => {
        page++;
        onChange((page - 1) * limit);
    }

    const goBack = () => {
        page--;
        onChange((page - 1) * limit);
    }

    const changePage = (e) => {
        const result = parseInt(e.target.value.replaceAll(/\D/g, ''))

        if (isNaN(result)) {
            e.target.value = 1;
            return;
        }

        if (result < 1) {
            e.target.value = 1;
            return;
        }

        if (result > total / limit) {
            e.target.value = (total / limit + 1).toFixed(0);
            return;
        }

        e.target.value = result;
        page = result;

        onChange((result - 1) * limit);
    }
</script>

<nav class="flex items-center gap-x-2" class:hidden={ (total / limit) < 2 }>
    <button
            on:click={ goBack }
            disabled="{ page === 1 }"
            type="button" class="min-h-[38px] min-w-[38px] py-2 px-2.5 inline-flex justify-center items-center gap-x-2 text-sm rounded-lg text-gray-800 hover:bg-gray-100 focus:outline-none focus:bg-gray-100 disabled:opacity-50 disabled:pointer-events-none dark:text-white dark:hover:bg-white/10 dark:focus:bg-white/10">
        <svg class="flex-shrink-0 size-3.5" xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
            <path d="m15 18-6-6 6-6"></path>
        </svg>
        <span aria-hidden="true" class="sr-only">Назад</span>
    </button>
    <div class="flex items-center gap-x-1">
        <input
                bind:value={ page }
                on:change={ changePage }
                class="bg-transparent w-[3rem] text-center min-h-[2.375rem] min-w-[2.375rem] border border-gray-200 text-gray-800 py-2 px-3 text-sm rounded-lg focus:outline-none focus:bg-gray-50 disabled:opacity-50 disabled:pointer-events-none dark:border-neutral-700 dark:text-white dark:focus:bg-white/10" />
        <span class="min-h-[2.375rem] flex justify-center items-center text-gray-500 py-2 pl-1.5 text-sm dark:text-neutral-500">из</span>
        <span class="min-h-[2.375rem] flex justify-center items-center text-gray-500 py-2 text-sm dark:text-neutral-500">{ (total / limit).toFixed(0) }</span>
    </div>
    <button
            on:click={ goForward }
            disabled="{ page === parseInt((total / limit).toFixed(0))}"
            type="button" class="min-h-[38px] min-w-[38px] py-2 px-2.5 inline-flex justify-center items-center gap-x-2 text-sm rounded-lg text-gray-800 hover:bg-gray-100 focus:outline-none focus:bg-gray-100 disabled:opacity-50 disabled:pointer-events-none dark:text-white dark:hover:bg-white/10 dark:focus:bg-white/10">
        <span aria-hidden="true" class="sr-only">Вперёд</span>
        <svg class="flex-shrink-0 size-3.5" xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
            <path d="m9 18 6-6-6-6"></path>
        </svg>
    </button>
</nav>