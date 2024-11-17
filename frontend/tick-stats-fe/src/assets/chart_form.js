import { fetchWrapper } from "./utils"

class ChartForm {
    constructor() {
        this.resetForm()
    }

    async validate() {
        if (this.data.chart_name === '') {
            throw new Error('chart_name is required')
        }
        if (this.data.chart_type === '') {
            throw new Error('chart_type is required')
        }
        if (this.data.appid === '') {
            throw new Error('appid is required')
        }
        if (this.key_name_input.length === 0) {
            throw new Error('key_name is required')
        }
        for (let key of this.key_name_input) {
            // 不能含有特殊字符
            if (!/^[a-zA-Z0-9_]+$/.test(key)) {
                throw new Error('key_name can only contain letters, numbers and underscores')
            }
        }
        this.data.key_name = this.key_name_input.join(',')
        return true
    }

    async create() {
        let appid = this.data.appid
        await this.validate()
        return await fetchWrapper(`/api/account/app/${appid}/chart/new`, {
            method: 'POST',
            body: JSON.stringify(this.data)
        })
    }

    async update() {
        if (!this.data.chart_id) {
            throw new Error('chart_id is required')
        }
        this.validate()
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
        this.key_name_input = []
    }
}

export { ChartForm }