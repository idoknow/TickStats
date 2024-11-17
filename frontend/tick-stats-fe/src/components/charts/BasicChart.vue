<template>
    <div v-show="!initLoaded" style="display: flex; justify-content: center; align-items: center; height: 350px; flex-direction: column">
        <small>Hang On...</small>
        <v-progress-linear
        style="width: 50%"
        color="deep-purple-accent-4"
        height="6"
        indeterminate
        rounded
        ></v-progress-linear>        
    </div>
    <div v-show="initLoaded" :id="chartData.chart_id" style="margin-top: 16px; height: 350px; min-width: 100%; ">
    </div>
</template>

<script>
import { Chart } from '@/assets/chart';
import { fetchWrapper, fillingTimeData, chartsPresetConfigs } from '@/assets/utils';

export default {
    props: {
        chartData: Object
    },
    data() {
        return {
            initLoaded: false,
            chart: null,
        };
    },
    mounted() {
        // 当div的宽度发生变化时，重新渲染图表
        new ResizeObserver(() => {
            if (this.chart) {
                this.chart.resize();
            }
        }).observe(document.getElementById(this.chartData.chart_id));

        this.getMetric(this.chartData.appid, this.chartData.chart_id, this.chartData).then(() => {
            if (this.chart) this.chart.dispose();
            this.chart = new Chart(this.chartData.chart_id)
            this.chart.initChart(this.chartData.option);
            this.initLoaded = true;
        });
    },
    methods: {
        getChartConfig(chartType) {
            for (let i = 0; i < chartsPresetConfigs.length; i++) {
                if (chartsPresetConfigs[i].chart_type === chartType) {
                    return chartsPresetConfigs[i];
                }
            }
            return null;
        },
        async getMetric(appid, chartid, chartData) {
            await fetchWrapper(`/api/metric/${appid}/${chartid}`, {
                method: 'GET',
            }).then((data) => {
                let _display_type = 'default';
                if (chartData.extra_config.only_represent_number) {
                    _display_type = 'plain_number';
                }
                if (chartData.chart_type === 'simple_line') {
                    chartData.option = {
                        ...this.getChartConfig('simple_line').option_model,
                        title: {text: chartData.chart_name},
                        series: [{
                                showSymbol: false,
                                type: 'line',
                                data: fillingTimeData(data)
                            }],
                        _display_type: _display_type,
                    }
                } else if (chartData.chart_type === 'simple_pie') {
                    chartData.option = {
                        ...this.getChartConfig('simple_pie').option_model,
                        title: {text: chartData.chart_name},
                        series: [{
                            name: chartData.chart_name,
                            type: 'pie',
                            data: data.map((item) => {
                                return {
                                    name: item.k,
                                    value: item.v
                                }
                            })
                        }],
                    }
                }
            }).catch((error) => {
                console.error(error);
                this.makeToast('Failed to get metrics', 'error');
            });
        },
    }
};
</script>