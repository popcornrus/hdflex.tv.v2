<script>
import Card from "$lib/components/(public)/reviews/Card.svelte";
</script>

<section class="space-y-8">
    {#each Array(5) as _, i}
        <Card />
    {/each}
</section>