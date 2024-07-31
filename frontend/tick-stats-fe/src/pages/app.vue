<template>
    <AppBar />

    <div class="content">
        <h1 class="gradient index-title" v-if="accountName !== ''">{{ accountName }}/{{ appName }} Stats
        </h1>
        <!-- <v-btn @click="mock" variant="plain">DEBUG: MOCK RANDOM DATA</v-btn> -->
        <v-dialog max-width="800">
            <template v-slot:activator="{ props: activatorProps }">
                <v-btn style="margin: 32px;" v-bind:="activatorProps" @click="onCreatingChart = true;">Create
                    Chart</v-btn>
            </template>

            <template v-slot:default="{ isActive }">
                <v-card title="Create Chart">

                    <div class="create-chart-container">
                        <div style="flex: 1; width: 100%">
                            <v-card-text>
                                <v-text-field readonly v-model="appId" label="App Id" variant="outlined"></v-text-field>
                                <v-text-field v-model="newChart.chart_name" label="Chart Name"
                                    variant="outlined"></v-text-field>
                                <v-text-field v-model="newChart.description" label="Description(optional)"
                                    variant="outlined"></v-text-field>
                                <v-radio-group v-model="newChart.chart_type" row @change="createChangeChart">
                                    <v-radio label="Simple Line Chart" value="simple_line"></v-radio>
                                    <v-radio label="Pie Chart" value="simple_pie"></v-radio>
                                </v-radio-group>
                                <v-checkbox v-model="newChart.public" label="Public" color="primary"></v-checkbox>
                                <v-text-field v-model="newChart.key_name" label="Key Name"
                                    variant="outlined"></v-text-field>
                            </v-card-text>
                        </div>

                        <div style="flex:2; margin-top:16px" v-if="newChart.chart_type!=''">
                            <h3>Chart Demo</h3>
                            <div id="chart-demo" style="width: 100%; height: 300px"></div>
                            <h3>Data Example</h3>
                            <pre>{{ pushMetricExample[newChart.chart_type] }}</pre>
                        </div>

                    </div>

                    <v-card-actions>
                        <v-spacer></v-spacer>
                        <v-btn text @click="createChart(isActive)">
                            Create
                        </v-btn>
                        <v-btn text="Close" @click="isActive.value=false; onCloseCreateChartDialog()">Close</v-btn>
                    </v-card-actions>
                </v-card>
            </template>
        </v-dialog>
    </div>

    <v-snackbar v-model="toast.show" :color="toast.color" :timeout="toast.timeout">
        {{ toast.text }}
    </v-snackbar>

</template>

<script>
import AppBar from '@/components/AppBar.vue';
import moment from 'moment';
import * as echarts from 'echarts';

export default {
    components: {
        AppBar
    },
    data() {
        return {
            toast: {
                show: false,
                text: '',
                color: 'primary',
                timeout: 3000,
            },
            chartInstance: {},
            chartData: null,
            appId: '',
            accountName: '',
            appName: '',
            newChart: {
                appid: '',
                chart_name: '',
                key_name: '',
                description: '',
                chart_type: '',
                public: false
            },
            simpleLineChartOptionModel: {
                tooltip: {
                    trigger: 'axis'
                },
                calculable: true,
                xAxis: {
                    type: 'category',
                    axisLabel: {
                        formatter: function (value) {
                            return moment(value).format('MM-DD HH:mm');
                        }
                    }
                },
                yAxis: {
                    type: 'value'
                },
                color: ['#21CBF3'],
                dataZoom: [
                    {
                        show: true,
                        realtime: true,
                    }
                ],
            },
            simplePieChartOptionModel: {
                tooltip: {
                    trigger: 'item',
                    formatter: '{a} <br/>{b} : {c} ({d}%)'
                },
                padAngle: 5,
                itemStyle: {
                    borderRadius: 5
                },
                legend: {
                    orient: 'vertical',
                    left: 'left',
                    data: []
                },
            },
            chartOptions: [],
            onCreatingChart: false,
            mock_os: ['Windows', 'Linux', 'MacOS'],
            pushMetricExample: {
                '': '',
                'simple_line': `
{
    "metrics_data": {
        "used_count": 2,
        ...
    }
}
                `,
                'simple_pie': `
{
    "metrics_data": {
        "os_name": "windows",
        ...
    }
}
                `
            }
        }
    },
    mounted() {
        this.appId = this.$route.query.id;
        this.getCharts();

        window.addEventListener('resize', () => {
            for (let key in this.chartInstance) {
                this.chartInstance[key].resize();
            }
        });
    },
    methods: {
        makeToast(text, color = 'primary', timeout = 3000) {
            this.toast.text = text;
            this.toast.color = color;
            this.toast.timeout = timeout;
            this.toast.show = true;
        },
        createChart(isActive) {
            this.newChart.appid = this.appId;
            fetch(`https://ts.lwl.lol/api/account/app/${this.appId}/chart/new`, {
                method: 'POST',
                credentials: 'include',
                headers: {
                    'Content-Type': 'application/json'
                },
                body: JSON.stringify(this.newChart)
            }).then(response => {
                if (response.ok) {
                    isActive.value = false;
                    this.makeToast('Chart created successfully');
                    this.getCharts();
                    this.newChart = {
                        chart_name: '',
                        key_name: '',
                        chart_type: '',
                        public: false,
                        description: '',
                        appid: ''
                    };
                    this.onCreatingChart = false
                }
            }).catch(error => {
                console.error(error);
                this.makeToast('Failed to create chart', 'error');
            });
        },
        createChangeChart() {
            console.log(this.newChart.chart_type);
            this.removeDemoChart()
            // demo
            if (this.newChart.chart_type === 'simple_line') {
                this.updateChart('chart-demo', {
                    ...this.simpleLineChartOptionModel,
                    title: {
                        text: this.newChart.chart_name
                    },
                    series: [
                        {
                            type: 'line',
                            data: Array.from({ length: 10 }, () => Math.floor(Math.random() * 100))
                        }
                    ],
                    _chart_type: 'simple_line'
                });
            } else if (this.newChart.chart_type === 'simple_pie') {
                this.updateChart('chart-demo', {
                    ...this.simplePieChartOptionModel,
                    title: {
                        text: this.newChart.chart_name
                    },
                    series: [
                        {
                            name: this.newChart.chart_name,
                            type: 'pie',
                            data: Array.from({ length: 3 }, (_, i) => {
                                return {
                                    name: this.mock_os[i],
                                    value: Math.floor(Math.random() * 100)
                                }
                            })
                        }
                    ],
                    _chart_type: 'simple_pie'
                });
            }
        },
        onCloseCreateChartDialog() {
            console.log('close');
            this.onCreatingChart = false;
            this.newChart = {
                chart_name: '',
                key_name: '',
                chart_type: '',
                public: false,
                description: '',
                appid: ''
            };
            // remove demo chart
            this.removeDemoChart();
        },
        removeDemoChart() {
            this.chartInstance['chart-demo'] && this.chartInstance['chart-demo'].dispose();
            this.chartInstance['chart-demo'] && delete this.chartInstance['chart-demo'];
        },

        getCharts() {
            fetch(`https://ts.lwl.lol/api/account/app/${this.appId}/chart`, {
                credentials: 'include'
            }).then(response => response.json())
                .then(data => {
                    this.chartData = data;

                    if (this.chartData.length > 0) {
                        this.accountName = this.chartData[0].account_name;
                        this.appName = this.chartData[0].app_name;
                        this.getMetrics();
                    }
                }).catch(error => {
                    console.error(error);
                    this.makeToast('Failed to get charts', 'error');
                });
        },
        getMetrics() {
            for (let i = 0; i < this.chartData.length; i++) {
                fetch(`https://ts.lwl.lol/api/metric/${this.appId}?key_name=${this.chartData[i].key_name}&chart_type=${this.chartData[i].chart_type}`)
                    .then((response) => response.json())
                    .then((data) => {
                        if (this.chartData[i].chart_type === 'simple_line') {
                            this.chartOptions.push({
                                ...this.simpleLineChartOptionModel,
                                title: {
                                    text: this.chartData[i].chart_name
                                },
                                series: [
                                    {
                                        type: 'line',
                                        data: data.map((item) => {
                                            return [item.k, item.v];
                                        })
                                    }
                                ],
                                _chart_type: 'simple_line'
                            });
                        } else if (this.chartData[i].chart_type === 'simple_pie') {
                            this.chartOptions.push({
                                ...this.simplePieChartOptionModel,
                                title: {
                                    text: this.chartData[i].chart_name
                                },
                                series: [
                                    {
                                        name: this.chartData[i].chart_name,
                                        type: 'pie',
                                        data: data.map((item) => {
                                            return {
                                                name: item.k,
                                                value: item.v
                                            }
                                        })
                                    }
                                ],
                                _chart_type: 'simple_pie'
                            });
                        }
                        this.updateChart(this.chartData[i].chart_name, this.chartOptions[this.chartOptions.length - 1]);
                    });
            }

        },
        mock() {
            fetch(`https://ts.lwl.lol/api/metric/${this.appId}`, {
                method: 'POST',
                credentials: 'include',
                headers: {
                    'Content-Type': 'application/json'
                },
                body: JSON.stringify({
                    metrics_data: {
                        "cnt_usage": Math.floor(Math.random() * 100),
                        "os_name": this.mock_os[Math.floor(Math.random() * this.mock_os.length)]
                    }
                })
            }).then(response => {
                if (response.ok) {
                    this.getMetrics();
                }
            }).catch(error => {
                console.error(error);
            });
        },
        updateChart(divId, option) {
            // create element if not exist
            if (!document.getElementById(divId)) {
                const chartDiv = document.createElement('div');
                chartDiv.id = divId;
                chartDiv.style.width = '100%';
                chartDiv.style.height = '350px';
                chartDiv.style.marginTop = '16px';
                document.querySelector('.content').appendChild(chartDiv);
            }
            let chart = null;
            if (this.chartInstance[divId]) {
                chart = this.chartInstance[divId];
            } else {
                chart = echarts.init(document.getElementById(divId));
            }
            chart.setOption(option);
            this.chartInstance[divId] = chart;
        }
    }
}
</script>

<style>
.gradient {
    background: linear-gradient(45deg, #2196F3, #21CBF3);
    -webkit-background-clip: text;
    -webkit-text-fill-color: transparent;
}

.content {
    padding: 32px 128px;
    display: flex;
    flex-direction: column;
    align-items: center;
}

.index-title {
    font-size: 64px;
}

.create-chart-container {
    display: flex;
    flex-direction: row;
    align-items: center;
    overflow-y: scroll;
    gap:16px
}

@media (max-width: 600px) {
    .content {
        padding: 32px 16px;
    }

    .index-title {
        font-size: 48px;
    }

    .create-chart-container {
        flex-direction: column;
        
    }
}
</style>