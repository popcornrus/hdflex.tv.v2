<script>
    import "$lib/app.css";
    import "$lib/index.js";

    import Header from '$lib/components/Header.svelte'
</script>

<div class="max-w-7xl mx-auto">
    <Header />
</div>

<main>
    <slot />
</main>

<footer class="h-96 bg-slate-950">

</footer>