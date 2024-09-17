import {Content} from "$lib/model/Content.js";

/** @type {import('./$types').PageLoad} */
export async function load() {
    return {
        content: new Content(),
    };
}