<script>
    const rating = ["☹️", "😕", "😐", "🙂", "🤯"]

    let userRate = 0;
</script>

<form action="">
    <div class="space-y-4">
        <div class="w-full">
            <textarea placeholder="Напишите Ваш отзыв здесь..." class="rounded-md transition duration-150 text-gray-500 border border-slate-800/80 w-full h-24 block bg-slate-900/50 focus:outline-none focus:border-slate-700 py-3 px-4"></textarea>
        </div>
        <div class="flex gap-x-4 max-w-2xl">
            <div class="w-full">
                <input type="text" placeholder="Имя" class="rounded-md transition duration-150 text-gray-500 w-full bg-slate-900/50 border border-slate-800/80 py-2 px-4 focus:outline-none focus:border-slate-700">
            </div>
            <div class="w-3/4">
                <button type="button" class="py-2 px-4 inline-flex items-center gap-x-2 text-sm font-semibold rounded-lg border border-transparent bg-blue-600 text-white hover:bg-blue-700 disabled:opacity-50 disabled:pointer-events-none">
                    Добавить отзыв
                </button>
            </div>
        </div>
    </div>
</form>