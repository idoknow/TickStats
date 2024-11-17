<template>
    <h3>{{ chartData.chart_name }}</h3>
    <v-data-table :items="items" density="compact" :headers="headers" :loading="loading" style="margin-top: 8px;"></v-data-table>
    

</template>

<script>
import { fetchWrapper } from '@/assets/utils';

export default {
    props: {
        chartData: Object
    },
    data() {
        return {
            initLoaded: false,
            loading: false,
            items: [],
            headers: []
        };
    },
    mounted() {
        this._getMetric();

        // Set headers
        this.headers.push({
            title: 'Time',
            align: 'start',
            sortable: true,
            key: 'k',
        })
        let key_names = this.chartData.key_name.split(',');
        for (let i = 0; i < key_names.length; i++) {
            this.headers.push({
                title: key_names[i],
                align: 'start',
                sortable: false,
                key: 'v' + i,
            })
        }
    },
    methods: {
        async _getMetric() {
            this.loading = true;
            await this.getMetric(this.chartData.appid, this.chartData.chart_id, this.chartData).then(() => {
                this.initLoaded = true;
                this.loading = false;
            });
        },
        async getMetric(appid, chartid, chartData) {
            await fetchWrapper(`/api/metric/${appid}/${chartid}`, {
                method: 'GET',
            }).then((data) => {
                for (let i = 0; i < data.length; i++) {
                    data[i].k = new Date(data[i].k*1000).toLocaleString();
                    for (let j = 0; j < data[i].v.length; j++) {
                        data[i]['v' + j] = data[i].v[j];
                        delete data[i].v[j];
                    }
                }
                this.items = data;

            }).catch((error) => {
                console.error(error);
            });
        },
    },
};
</script>