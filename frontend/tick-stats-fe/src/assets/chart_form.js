import { fetchWrapper } from "./utils"
import { chartsPresetConfigs } from "./utils"

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
        let keys = this.key_name_input.join(',')
        if (keys.length > 255) {
            throw new Error('key_name is too long')
        }

        if (this.key_name_input.length > 1) {
            for (let preset of chartsPresetConfigs) {
                console.log(preset)
                if (this.data.chart_type === preset.chart_type) {
                    if (!preset.multiple_keys) {
                        throw new Error('multiple keys are not allowed for this chart type')
                    }
                }
            }
        }

        this.data.key_name = keys

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