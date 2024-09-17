'use strict';

import axios from "axios";
import {env} from "$env/dynamic/public";

export class Content {
    params = "";

    constructor() {
        this.axios = axios.create({
            baseURL: `${env.PUBLIC_API_BASE_URL}/content`,
            timeout: 1000,
            headers: {
                'Content-Type': 'application/json',
            }
        });
    }

    async get() {
        return await this.axios.get(`/${this.params};`).then(r => r.data.data)
    }

    async count() {
        return await this.axios.get(`/count${this.params};`).then(r => r.data.data.count)
    }

    async findBySlug(slug) {
        return await this.axios.get(`/${slug}`).then(r => r.data.data.item)
    }

    applyParams(f) {
        this.params = "?"

        for (const key in f) {
            if (f.hasOwnProperty(key)) {
                this.params += `${key}==${f[key]}${Object.keys(f).indexOf(key) === Object.keys(f).length - 1 ? "" : ";"}`
            }
        }

        return this
    }
}