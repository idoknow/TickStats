import { fetchWrapper } from "./utils"

class ChartForm {
    constructor() {
        this.resetForm()
    }

    async create() {
        let appid = this.data.appid
        if (appid === '') {
            throw new Error('appid is required')
        }
        return await fetchWrapper(`/api/account/app/${appid}/chart/new`, {
            method: 'POST',
            body: JSON.stringify(this.data)
        })
    }

    async update() {
        if (!this.data.chart_id) {
            throw new Error('chart_id is required')
        }
        if (!this.data.appid === '') {
            throw new Error('appid is required')
        }
        return await fetchWrapper(`/api/account/app/${this.data.appid}/chart/${this.data.chart_id}`, {
            method: 'PUT',
            body: JSON.stringify(this.data)
        })
    }

    async delete() {
        if (!this.data.chart_id) {
            throw new Error('chart_id is required')
        }
        return await fetchWrapper(`/api/account/app/${this.data.appid}/chart/${this.data.chart_id}`, {
            method: 'DELETE',
            body: JSON.stringify(this.data)
        })
    }

    resetForm() {
        this.data = {
            chart_id: '',
            chart_name: '',
            key_name: '',
            chart_type: '',
            public: false,
            description: '',
            appid: '',
            extra_config: {}
        }
    }
}

export { ChartForm }