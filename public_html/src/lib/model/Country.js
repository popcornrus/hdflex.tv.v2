'use strict';

import axios from "axios";
import {env} from "$env/dynamic/public";

export class Country {
    constructor() {
        this.axios = axios.create({
            baseURL: `${env.PUBLIC_API_BASE_URL}/countries`,
            timeout: 1000,
            headers: {
                'Content-Type': 'application/json',
            }
        });
    }

    async get() {
        return await this.axios.get(`/`).then(r => r.data.data)
    }
}