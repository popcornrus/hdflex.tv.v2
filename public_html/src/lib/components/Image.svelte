<script>
    import {onMount} from "svelte";
    import {env} from "$env/dynamic/public";

    let element = null;

    export let src = "";

    let loaded = false;
    let failed = false;

    const roundToNearestFive = (number) => {
        return Math.ceil(number / 5) * 5;
    }

    let img;

    const GetImageSrc = () => {
        if ($$props.width && $$props.height) {
            return `${env.PUBLIC_IMAGE_BASE_URL}/fr/${roundToNearestFive($$props.width)}/${roundToNearestFive($$props.height)}/${src}`;
        }

        const rect = element.getBoundingClientRect()

        if ($$props.width && !$$props.height) {
            return `${env.PUBLIC_IMAGE_BASE_URL}/fr/${roundToNearestFive($$props.width)}/${roundToNearestFive(element.getBoundingClientRect().height)}/${src}`;
        }

        if (!$$props.width && $$props.height) {
            return `${env.PUBLIC_IMAGE_BASE_URL}/fr/${roundToNearestFive(element.getBoundingClientRect().width)}/${roundToNearestFive($$props.height)}/${src}`;
        }

        return `${env.PUBLIC_IMAGE_BASE_URL}/fr/${roundToNearestFive(rect.width)}/${roundToNearestFive(rect.height)}/${src}`;
    }

    const loadImage = () => {
        img = new Image();
        img.src = GetImageSrc();

        img.onload = () => {
            loaded = true;
        }
        img.onerror = () => {
            failed = true;
        }
    }

    onMount(() => {
        loadImage()
    })

    $: if (img?.src !== undefined && src !== img.src) {
        if (src.length > 0) {
            loadImage()
        } else {
            loaded = false;
            failed = true
        }
    }
</script>

<div bind:this={element} class="w-full h-full border border-slate-700 rounded-2xl overflow-hidden">
    {#if loaded}
        <img src="{img.src}" alt="{src}" loading="lazy" class="w-full h-full object-cover" />
    {:else if failed}
        <div class="h-full flex items-center justify-center text-gray-600 hover:text-slate-700 duration-150 transition" title="{src}">
            <svg fill="currentColor" viewBox="0 0 32 32" xmlns="http://www.w3.org/2000/svg" class="w-24 h-24">
                <path d="M30,3.4141,28.5859,2,2,28.5859,3.4141,30l2-2H26a2.0027,2.0027,0,0,0,2-2V5.4141ZM26,26H7.4141l7.7929-7.793,2.3788,2.3787a2,2,0,0,0,2.8284,0L22,19l4,3.9973Zm0-5.8318-2.5858-2.5859a2,2,0,0,0-2.8284,0L19,19.1682l-2.377-2.3771L26,7.4141Z"/>
                <path d="M6,22V19l5-4.9966,1.3733,1.3733,1.4159-1.416-1.375-1.375a2,2,0,0,0-2.8284,0L6,16.1716V6H22V4H6A2.002,2.002,0,0,0,4,6V22Z"/>
                <rect fill="none" width="32" height="32"/>
            </svg>
        </div>
    {/if}
</div>