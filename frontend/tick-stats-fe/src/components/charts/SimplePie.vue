<template>
    <div v-show="!initLoaded" style="display: flex; justify-content: center; align-items: center; height: 350px; flex-direction: column">
        <small>Hang On...</small>   
        <v-progress-linear
        color="deep-purple-accent-4"
        height="6"
        style="width: 50%"
        indeterminate
        rounded
        ></v-progress-linear>    
         
    </div>
    <div v-show="initLoaded" :id="chartData.chart_id" style="margin-top: 16px; height: 300px; min-width: 100%; display: flex; justify-content: center; align-items: center;">
    </div>
    <div v-show="initLoaded" style="display: flex; flex-direction: column; align-items: center">
        <div>
            <small style="margin-left: 8px">From <span style="font-weight: bold;">{{
                humanReadableTime(valueTimeSeries[value[0]]) }}</span> to <span style="font-weight: bold;">{{
                        humanReadableTime(valueTimeSeries[value[1]]) }}</span></small>
            <v-btn :loading="loading" variant="outlined" size="x-small" style="margin-left: 16px" @click="_getMetric(valueTimeSeries[value[0]],
                valueTimeSeries[value[1]])">Apply</v-btn>
        </div>
        <v-range-slider :min="min" :max="max" style="width: 90%; max-width: 700px;" step="1" color="primary"
            v-model="value" strict></v-range-slider>
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
            loading: false,
            min: 0,
            max: 2 * 24 * 7,
            value: [0, 100],
            StepSec: 30 * 60,
            valueTimeSeries: [],
            chart: null
        };
    },
    mounted() {
        // 当div的宽度发生变化时，重新渲染图表
        new ResizeObserver(() => {
            if (this.chart) {
                this.chart.resize();
            }
        }).observe(document.getElementById(this.chartData.chart_id));

        this.generateTimeStampSeries(this.StepSec);
        this.value = [this.valueTimeSeries.length - 48, this.valueTimeSeries.length - 1];
        this._getMetric(this.valueTimeSeries[this.value[0]], this.valueTimeSeries[this.value[1]]);
    },
    methods: {
        async _getMetric(from, to) {
            this.loading = true;
            await this.getMetric(this.chartData.appid, this.chartData.chart_id, this.chartData, from, to).then(() => {
                if (this.chart) this.chart.dispose();
                this.chart = new Chart(this.chartData.chart_id)
                this.chart.initChart(this.chartData.option);
                this.initLoaded = true;
                this.loading = false;
            });
        },
        humanReadableTime(timestamp) {
            let date = new Date(timestamp * 1000);
            return date.toLocaleString();
        },
        generateTimeStampSeries(stepSec) {
            let currentTimeStampSeconds = Math.floor(Date.now() / 1000);
            let series = [];
            for (let i = this.min; i < this.max; i++) {
                series.push(currentTimeStampSeconds - i * stepSec);
            }
            series.reverse();
            this.valueTimeSeries = series;
        },
        getChartConfig(chartType) {
            for (let i = 0; i < chartsPresetConfigs.length; i++) {
                if (chartsPresetConfigs[i].chart_type === chartType) {
                    return chartsPresetConfigs[i];
                }
            }
            return null;
        },
        async getMetric(appid, chartid, chartData, from=0, to=0) {
            await fetchWrapper(`/api/metric/${appid}/${chartid}?from=${from}&to=${to}`, {
                method: 'GET',
            }).then((data) => {
                let _display_type = 'default';
                if (chartData.extra_config.only_represent_number) {
                    _display_type = 'plain_number';
                }
                if (chartData.chart_type === 'simple_line') {
                    chartData.option = {
                        ...this.getChartConfig('simple_line').option_model,
                        title: { text: chartData.chart_name },
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
                        title: { text: chartData.chart_name },
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
    },
};
</script>