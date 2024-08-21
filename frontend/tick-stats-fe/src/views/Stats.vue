<template>
    <AppBar />

    <div class="content">
        <h1 class="gradient index-title" v-if="accountName !== ''">{{ accountName }}/{{ appName }} Stats
        </h1>

        <v-alert v-model="showErrAlert" type="info" variant="tonal" closable style="margin: 16px 0px; width: 100%;">
            {{ errAlert }}
        </v-alert>

        <v-btn @click="mockData" variant="plain" v-if="devMode">DEBUG: MOCK RANDOM DATA</v-btn>

        <!-- Tutorial of creating a chart -->
        <v-card style="width: 100%;" v-if="showCreateChartTutorial">
            <v-tabs v-model="createChartTab" align-tabs="center" color="deep-purple-accent-4">
                <v-tab v-for="(chart, index) in chartsPresetConfigs" :key="index" :value="index"
                    @click="createChangeChart('chart-demo1', chart.chart_type)">
                    {{ chart.title }}
                </v-tab>
            </v-tabs>

            <v-tabs-window v-model="createChartTab">
                <v-tabs-window-item v-for="n in 3" :key="n" :value="n">
                    <div style="display: flex; justify-content: center; margin: 8px; flex-direction: column; align-items: center">
                        <p class="text-caption">{{ chartsPresetConfigs[createChartTab].description }}</p>
                        <v-progress-circular v-if="chartTutorialRendering" color="primary"
                            indeterminate></v-progress-circular>
                    </div>

                    <div id="chart-demo1" style="width: 100%; height: 250px; margin-bottom: 16px"></div>

                    <div style="display: flex; justify-content: flex-end; margin: 16px;">
                        <v-btn variant="plain" @click="onTutorialCreateBtnClicked">Create ></v-btn>
                    </div>
                </v-tabs-window-item>
            </v-tabs-window>
        </v-card>

        <!-- <v-btn @click="mock" variant="plain">DEBUG: MOCK RANDOM DATA</v-btn> -->
        <v-dialog max-width="800" v-model="creatChartDialog">
            <template v-slot:activator="{ props: activatorProps }">
                <v-fab icon="mdi-plus" color="primary" size="52" style="position: fixed; right: 80px; bottom: 52px;"
                    v-bind:="activatorProps"></v-fab>
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
                                <v-radio-group v-model="newChart.chart_type" row
                                    @change="createChangeChart('chart-demo2', newChart.chart_type)">
                                    <v-radio v-for="(chart, index) in chartsPresetConfigs" :key="index" :label="chart.title"
                                        :value="chart.chart_type"></v-radio>
                                </v-radio-group>

                                <!-- extra options -->
                                 <div v-for="(ops, index) in selectedChartConfig.extra_configs" :key="index">
                                    <div v-if="ops.type === 'bool'">
                                        <small>{{ops.description}}</small>
                                        <v-checkbox v-model="newChart.extra_configs[ops.name]" :label="ops.name"
                                            color="primary">
                                        </v-checkbox>
                                    </div>
                                    <div v-else-if="ops.type === 'selectable'">
                                        <small>{{ops.description}}</small>
                                        <v-select v-model="newChart.extra_configs[ops.name]" :items="ops.options"
                                         :label="ops.name" variant="outlined">
                                        </v-select>
                                    </div>
                                </div>

                                <v-checkbox v-model="newChart.public" label="Public" color="primary"></v-checkbox>
                                <v-text-field v-model="newChart.key_name" label="Key Name"
                                    variant="outlined"></v-text-field>
                            </v-card-text>
                        </div>

                        <div style="flex:2; margin-top:16px" v-if="newChart.chart_type != ''">
                            <h3>Chart Demo</h3>
                            <div id="chart-demo2" style="width: 100%; height: 300px"></div>
                            <h3>Data Example</h3>
                            <pre>{{ metricPushExample }}</pre>
                        </div>

                    </div>

                    <v-card-actions>
                        <v-spacer></v-spacer>
                        <v-btn text @click="createChart(isActive)">
                            Create
                        </v-btn>
                        <v-btn text="Close" @click="isActive.value = false; onCloseCreateChartDialog()">Close</v-btn>
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
import * as echarts from 'echarts';
import { fetchWrapper, fillingTimeData, chartsPresetConfigs } from '@/assets/utils';
import { useGlobalStore } from '@/stores/global';

export default {
    components: {
        AppBar,
    },
    computed: {
        metricPushExample() {
            // for creating a new chart
            if (this.newChart.chart_type === '') return '';
            for (let i = 0; i < this.chartsPresetConfigs.length; i++) {
                if (this.chartsPresetConfigs[i].chart_type === this.newChart.chart_type) {
                    return this.chartsPresetConfigs[i].metric_example;
                }
            }
            return '';
        },
        selectedChartConfig() {
            // for creating a new chart
            for (let i = 0; i < this.chartsPresetConfigs.length; i++) {
                if (this.chartsPresetConfigs[i].chart_type === this.newChart.chart_type) {
                    return this.chartsPresetConfigs[i];
                }
            }
            return {
                demo_data: [],
                extra_configs: []
            };
        }
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
                public: false,
                extra_configs: {}
            },
            createChartTab: 0,
            chartOptions: [],
            mock_os: ['Windows', 'Linux', 'MacOS'],
            showErrAlert: false,
            errAlert: '',
            showCreateChartTutorial: false,
            chartTutorialRendering: false,
            creatChartDialog: false,
            global: useGlobalStore(),
            devMode: false,
            chartsPresetConfigs: chartsPresetConfigs
        }
    },
    mounted() {
        this.appId = this.$route.params.appid;
        this.getCharts();

        window.addEventListener('resize', () => {
            for (let key in this.chartInstance) {
                this.chartInstance[key].resize();
            }
        });

        if (this.$route.query.dev === 'true') this.devMode = true;
    },
    methods: {
        mockData() {
            fetchWrapper(`/api/metric/${this.appId}`, {
                method: 'POST',
                body: JSON.stringify({
                    metrics_data: {
                        used_count: Math.floor(Math.random() * 100),
                        os_name: this.mock_os[Math.floor(Math.random() * 3)]
                    }
                })
            }).then(() => {
                this.makeToast('Mock data posted successfully');
                this.getCharts();
            }).catch((error) => {
                console.error(error);
                this.makeToast('Failed to post mock data', 'error');
            });
        },

        getChartConfig(chartType) {
            for (let i = 0; i < this.chartsPresetConfigs.length; i++) {
                if (this.chartsPresetConfigs[i].chart_type === chartType) {
                    return this.chartsPresetConfigs[i];
                }
            }
            return null;
        },

        makeToast(text, color = 'primary', timeout = 3000) {
            this.toast.text = text;
            this.toast.color = color;
            this.toast.timeout = timeout;
            this.toast.show = true;
        },
        createChart(isActive) {
            this.newChart.appid = this.appId;
            // fill extra configs if not set
            for (let i = 0; i < this.selectedChartConfig.extra_configs.length; i++) {
                if (!this.newChart.extra_configs[this.selectedChartConfig.extra_configs[i].name]) {
                    this.newChart.extra_configs[this.selectedChartConfig.extra_configs[i].name] = this.selectedChartConfig.extra_configs[i].default;
                }
            }

            fetchWrapper(`/api/account/app/${this.appId}/chart/new`, {
                method: 'POST',
                body: JSON.stringify(this.newChart),
            }).then(() => {
                isActive.value = false;
                this.makeToast('Chart created successfully');
                this.showCreateChartTutorial = false;
                this.getCharts();
                this.newChart = {
                    chart_name: '',
                    key_name: '',
                    chart_type: '',
                    public: false,
                    description: '',
                    appid: '',
                    extra_configs: {}
                };
            }).catch((error) => {
                console.error(error);
                this.makeToast('Failed to create chart', 'error');
            });
        },
        createChangeChart(elementId, chartType) {
            console.log(elementId, chartType);
            this.removeDemoChart(elementId)
            this.newChart.chart_type = chartType;
            // demo
            if (chartType === 'simple_line') {
                this.updateChart(elementId, {
                    ...this.selectedChartConfig.option_model,
                    series: [
                        {
                            showSymbol: false,
                            type: 'line',
                            data: this.selectedChartConfig.demo_data
                        }
                    ],
                    grid: {
                        top: '16px',
                        containLabel: true
                    },
                });
            } else if (chartType === 'simple_pie') {
                this.updateChart(elementId, {
                    ...this.selectedChartConfig.option_model,
                    series: [
                        {
                            name: this.newChart.chart_name,
                            type: 'pie',
                            data: this.selectedChartConfig.demo_data
                        }
                    ],
                });
            }
        },
        onCloseCreateChartDialog() {
            console.log('close');
            this.newChart = {
                chart_name: '',
                key_name: '',
                chart_type: '',
                public: false,
                description: '',
                appid: '',
                extra_configs: {}
            };
            // remove demo chart
            this.removeDemoChart();
        },
        removeDemoChart(elementId) {
            this.chartInstance[elementId] && this.chartInstance[elementId].dispose();
            this.chartInstance[elementId] && delete this.chartInstance[elementId];
        },
        getCharts() {
            console.log(this.global.account_apps)
            let url = `/api/stats/${this.appId}/charts` // only public charts
            for (let i = 0; i < this.global.account_apps.length; i++) {
                if (this.global.account_apps[i].app_id === this.appId) {
                    url = `/api/account/app/${this.appId}/chart` // all charts
                    break;
                }
            }

            fetchWrapper(url, {
                method: 'GET',
            }).then((data) => {
                this.chartData = data.chart;
                this.accountName = data.account_name;
                this.appName = data.app_name;
                if (this.chartData.length > 0) {
                    this.getMetrics();
                } else {
                    // user don't have any charts
                    this.handleEmptyApp();
                }
            }).catch((error) => {
                console.error(error);
                this.makeToast('Failed to get charts: ' + error, 'error');
            });
        },
        handleEmptyApp() {
            this.showErrAlert = true;
            this.errAlert = 'You don\'t have any charts yet, create one now!';
            this.showCreateChartTutorial = true;
            this.chartTutorialRendering = true;
            setTimeout(() => {
                // wait for the element to be created
                this.createChangeChart('chart-demo1', this.chartsPresetConfigs[0].chart_type);
                this.chartTutorialRendering = false;
            }, 500);
        },
        onTutorialCreateBtnClicked() {
            let t = this.chartsPresetConfigs[this.createChartTab].chart_type;
            this.creatChartDialog = true;
            this.newChart.chart_type = t;
            setTimeout(() => {
                this.createChangeChart('chart-demo2', t);
            }, 500);
        },
        getMetrics() {
            for (let i = 0; i < this.chartData.length; i++) {
                fetchWrapper(`/api/metric/${this.appId}?key_name=${this.chartData[i].key_name}&chart_type=${this.chartData[i].chart_type}`, {
                    method: 'GET',
                }).then((data) => {
                    if (this.chartData[i].chart_type === 'simple_line') {
                        this.updateChart(this.chartData[i].chart_name, {
                            ...this.getChartConfig('simple_line').option_model,
                            title: {
                                text: this.chartData[i].chart_name
                            },
                            series: [
                                {
                                    showSymbol: false,
                                    type: 'line',
                                    data: fillingTimeData(data)
                                }
                            ],
                        });
                    } else if (this.chartData[i].chart_type === 'simple_pie') {
                        this.updateChart(this.chartData[i].chart_name, {
                            ...this.getChartConfig('simple_pie').option_model,
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
                        });
                    }
                }).catch((error) => {
                    console.error(error);
                    this.makeToast('Failed to get metrics', 'error');
                });
            }

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
    padding: 32px 0px;
    display: flex;
    flex-direction: column;
    align-items: center;
    width: 100%;
    max-width: 900px;
    margin: 0 auto;
}

.index-title {
    font-size: 40px;
    text-align: center;
}

.create-chart-container {
    display: flex;
    flex-direction: row;
    align-items: center;
    overflow-y: scroll;
    gap: 16px
}

@media (max-width: 600px) {
    .content {
        padding: 32px 16px;
    }

    .index-title {
        font-size: 32px;
    }

    .create-chart-container {
        flex-direction: column;
    }
}
</style>