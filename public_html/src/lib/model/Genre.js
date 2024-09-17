'use strict';

import axios from "axios";
import {env} from "$env/dynamic/public";

export class Genre {
    constructor() {
        this.axios = axios.create({
            baseURL: `${env.PUBLIC_API_BASE_URL}/genres`,
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