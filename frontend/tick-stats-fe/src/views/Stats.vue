<template>
    <AppBar />

    <div class="content">

        <h1 class="gradient index-title" v-if="accountName !== ''">{{ accountName }}/{{ appName }} Stats
        </h1>

        <div>
            <v-progress-circular v-if="loadingCharts" color="primary" indeterminate></v-progress-circular>
        </div>

        <v-dialog max-width="600">
            <template v-slot:activator="{ props: activatorProps }">
                <v-btn v-bind:="activatorProps" v-if="isOwner" variant="plain" style="margin-left: 16px;">Manage Charts</v-btn>
            </template>

            <template v-slot:default="{ isActive }">
                <v-card title="Manage Charts">
                    <v-card-text>
                        <v-list>
                            <v-list-item v-for="(chart, index) in chartData" :key="index">
                                <div style="display: flex; justify-content: space-between">
                                    <v-list-item-title><span style="font-weight: bold;">{{ chart.chart_name }}</span>@{{ chart.chart_id }}</v-list-item-title>
                                    <div>
                                        <v-btn :key="index" variant="plain"
                                        @click="editChat(chart)">Edit</v-btn>
                                        <v-btn :key="index" :loading="loading" variant="plain"
                                        @click="deleteChart(chart)">Delete</v-btn>
                                    </div>
                                </div>
                                <v-divider></v-divider>
                            </v-list-item>
                        </v-list>
                        <span v-if="chartData.length === 0">No charts yet</span>
                    </v-card-text>
                    <v-card-actions>
                        <v-spacer></v-spacer>
                        <v-btn text="Close" @click="isActive.value = false;">Close</v-btn>
                    </v-card-actions>
                </v-card>
            </template>
        </v-dialog>


        <v-alert v-model="showErrAlert" type="info" variant="tonal" closable style="margin: 16px 0px; width: 100%;">
            {{ errAlert }}
        </v-alert>

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
                    <div
                        style="display: flex; justify-content: center; margin: 8px; flex-direction: column; align-items: center">
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

        <v-dialog max-width="1000" v-model="creatChartDialog">
            <template v-slot:activator="{ props: activatorProps }">
                <v-fab icon="mdi-plus" color="primary" size="52" style="position: fixed; right: 80px; bottom: 52px; z-index: 1000;"
                    v-bind:="activatorProps" @click="onFabClicked"></v-fab>
            </template>

            <template v-slot:default="{ isActive }">
                <v-card title="Chart">
                    <div class="create-chart-container">
                        <div style="flex: 1; width: 100%;">
                            <v-card-text style="margin-top: 16px;">
                                <v-text-field v-model="newChart.data.chart_name" label="Chart Name"
                                    variant="outlined"></v-text-field>
                                <v-text-field v-model="newChart.data.description" label="Description(optional)"
                                    variant="outlined"></v-text-field>
                                <v-radio-group v-model="newChart.data.chart_type" row
                                    @change="createChangeChart('chart-demo2', newChart.data.chart_type)">
                                    <v-radio v-for="(chart, index) in chartsPresetConfigs" :key="index"
                                        :label="chart.title" :value="chart.chart_type"></v-radio>
                                </v-radio-group>

                                <!-- extra options -->
                                <div v-for="(ops, index) in selectedChartConfig.extra_config" :key="index">
                                    <div v-if="ops.type === 'bool'">
                                        <small>{{ ops.description }}</small>
                                        <v-checkbox v-model="newChart.data.extra_config[ops.name]" :label="ops.name"
                                            color="primary">
                                        </v-checkbox>
                                    </div>
                                    <div v-else-if="ops.type === 'selectable'">
                                        <small>{{ ops.description }}</small>
                                        <v-select v-model="newChart.data.extra_config[ops.name]" :items="ops.options"
                                            :label="ops.name" variant="outlined">
                                        </v-select>
                                    </div>
                                </div>
                                <small>Public charts can be viewed by anyone</small>
                                <v-checkbox v-model="newChart.data.public" label="Public" color="primary"></v-checkbox>
                                <small>Used to query metrics, corresponds to the key name in the `metrics_data`</small>
                                <v-text-field v-model="newChart.data.key_name" label="Key Name"
                                    variant="outlined"></v-text-field>

                                <p>App Id: {{ appId }}</p>
                                <p v-if="newChart.data.chart_id" style="margin-bottom: 16px;">Chart Id: {{ newChart.data.chart_id }}</p>
                            </v-card-text>
                        </div>

                        <div style="flex:2; margin-top:16px" v-if="newChart.data.chart_type != ''">
                            <h3>Chart Demo</h3>
                            <div id="chart-demo2" style="width: 100%; height: 300px"></div>
                            <h3>Data Example</h3>
                            <pre>{{ metricPushExample }}</pre>
                        </div>

                    </div>
                    <v-card-actions>
                        <v-spacer></v-spacer>
                        <v-btn v-if="newChart.data.chart_id" text @click="updateChart(isActive)" :loading="loading">Update</v-btn>
                        <v-btn v-else text @click="createChart(isActive)" :loading="loading">
                            Create
                        </v-btn>
                        <v-btn text="Close" @click="isActive.value = false; onCloseCreateChartDialog()">Close</v-btn>
                    </v-card-actions>
                </v-card>
            </template>
        </v-dialog>

        <div v-for="(chart, index) in chartData" :key="chart.chart_id" style="width: 100%;">
            <SimplePie v-if="chart.chart_type == 'simple_pie'" :chart-data="chart"></SimplePie>
            <BasicChart v-else :chart-data="chart"></BasicChart>
        </div>
    </div>

    <v-snackbar v-model="toast.show" :color="toast.color" :timeout="toast.timeout">
        {{ toast.text }}
    </v-snackbar>

</template>

<script>
import AppBar from '@/components/AppBar.vue';

import { fetchWrapper, fillingTimeData, chartsPresetConfigs } from '@/assets/utils';
import { useGlobalStore } from '@/stores/global';
import { ChartForm } from '@/assets/chart_form';
import { Chart } from '@/assets/chart';
import SimplePie from '@/components/charts/SimplePie.vue';
import BasicChart from '@/components/charts/BasicChart.vue';

export default {
    components: {
        AppBar,
        SimplePie,
        BasicChart,
    },
    computed: {
        metricPushExample() {
            // for creating a new chart
            if (this.newChart.data.chart_type === '') return '';
            for (let i = 0; i < this.chartsPresetConfigs.length; i++) {
                if (this.chartsPresetConfigs[i].chart_type === this.newChart.data.chart_type) {
                    return this.chartsPresetConfigs[i].metric_example;
                }
            }
            return '';
        },
        selectedChartConfig() {
            // for creating a new chart
            for (let i = 0; i < this.chartsPresetConfigs.length; i++) {
                if (this.chartsPresetConfigs[i].chart_type === this.newChart.data.chart_type) {
                    return this.chartsPresetConfigs[i];
                }
            }
            return {
                demo_data: [],
                extra_config: []
            };
        }
    },
    data() {
        return {
            loading: false,
            loadingCharts: false,
            toast: {
                show: false,
                text: '',
                color: 'primary',
                timeout: 3000,
            },
            demoChartInstance: null,
            chartData: [],
            appId: '',
            accountName: '',
            appName: '',
            newChart: new ChartForm(),
            createChartTab: 0,
            chartOptions: [],
            showErrAlert: false,
            errAlert: '',
            showCreateChartTutorial: false,
            chartTutorialRendering: false,
            creatChartDialog: false,
            isOwner: false,
            global: useGlobalStore(),
            chartsPresetConfigs: chartsPresetConfigs
        }
    },
    mounted() {
        this.appId = this.$route.params.appid;
        this.getCharts();
    },
    methods: {
        onFabClicked() {
            this.creatChartDialog = true;
            this.newChart.resetForm();
        },
        makeToast(text, color = 'primary', timeout = 3000) {
            this.toast.text = text;
            this.toast.color = color;
            this.toast.timeout = timeout;
            this.toast.show = true;
        },
        async createChart(isActive) {
            this.newChart.data.appid = this.appId;
            // fill extra configs if not set
            for (let i = 0; i < this.selectedChartConfig.extra_config.length; i++) {
                if (!this.newChart.data.extra_config[this.selectedChartConfig.extra_config[i].name]) {
                    this.newChart.data.extra_config[this.selectedChartConfig.extra_config[i].name] = this.selectedChartConfig.extra_config[i].default;
                }
            }
            try {
                this.loading = true;
                await this.newChart.create();
                isActive.value = false;
                this.makeToast('Chart created successfully');
                this.getCharts();
            } catch (error) {
                this.makeToast('Failed to create chart: ' + error, 'error');
            } finally {
                this.loading = false;
            }
        },
        async updateChart(isActive) {
            try {
                this.loading = true;
                await this.newChart.update();
                isActive.value = false;
                this.makeToast('Chart updated successfully');
                this.getCharts();
            } catch (error) {
                this.makeToast('Failed to update chart: ' + error, 'error');
            } finally {
                this.loading = false;
            }
        },
        createChangeChart(elementId, chartType) {
            console.log(elementId, chartType);
            this.removeDemoChart()
            this.newChart.data.chart_type = chartType;
            if (chartType === 'simple_line') {
                this.demoChartInstance = new Chart(elementId)
                this.demoChartInstance.initChart({
                    ...this.selectedChartConfig.option_model,
                    series: [{
                            showSymbol: false,
                            type: 'line',
                            data: this.selectedChartConfig.demo_data
                        }],
                    grid: {
                        top: '16px',
                        containLabel: true
                    },
                });
            } else if (chartType === 'simple_pie') {
                this.demoChartInstance = new Chart(elementId)
                this.demoChartInstance.initChart({
                    ...this.selectedChartConfig.option_model,
                    series: [{
                            name: this.newChart.data.chart_name,
                            type: 'pie',
                            data: this.selectedChartConfig.demo_data
                        }],
                })
            }
        },
        onCloseCreateChartDialog() {
            console.log('close');
            this.newChart.resetForm();
            // remove demo chart
            this.removeDemoChart();
        },
        removeDemoChart() {
            if (this.demoChartInstance) {
                this.demoChartInstance.dispose();
                this.demoChartInstance = null;
            }
        },
        clearCharts() {
        },
        async getCharts() {
            this.clearCharts();
            this.loadingCharts = true;
            let url = `/api/stats/${this.appId}/charts` // only public charts
            await this.global.getAccount().then(() => {
                for (let i = 0; i < this.global.account.account_apps.length; i++) {
                    if (this.global.account.account_apps[i].app_id === this.appId) {
                        url = `/api/account/app/${this.appId}/chart` // all charts
                        this.isOwner = true;
                        break;
                    }
                }
            }).catch((err) => {
                console.error(err);
            });

            let _chartData = [];

            await fetchWrapper(url, {
                method: 'GET',
            }).then((data) => {
                _chartData = data.chart;
                this.accountName = data.account_name;
                this.appName = data.app_name;
                this.showErrAlert = false;
            }).catch((error) => {
                console.error(error);
                this.makeToast('Failed to get charts: ' + error, 'error');
            }).finally(() => {
                this.loadingCharts = false;
            });

            if (_chartData.length == 0) {
                // user don't have any charts
                this.handleEmptyApp();
            } else {
                this.chartData = _chartData;
            }
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
            this.newChart.data.chart_type = t;
            setTimeout(() => {
                this.createChangeChart('chart-demo2', t);
            }, 500);
        },
        async deleteChart(chart) {
            this.loading = true;
            this.newChart.data = {
                chart_id: chart.chart_id,
                chart_name: chart.chart_name,
                key_name: chart.key_name,
                chart_type: chart.chart_type,
                public: chart.public,
                description: chart.description,
                appid: this.appId,
                extra_config: chart.extra_config
            }
            try {
                await this.newChart.delete();
                this.makeToast('Chart deleted successfully');
                this.getCharts();
            } catch (error) {
                this.makeToast('Failed to delete chart', 'error');
            } finally {
                this.loading = false;
            }
        },
        editChat(chart) {
            this.creatChartDialog = true;
            this.newChart.data = {
                chart_id: chart.chart_id,
                chart_name: chart.chart_name,
                key_name: chart.key_name,
                chart_type: chart.chart_type,
                public: chart.public,
                description: chart.description,
                appid: this.appId,
                extra_config: chart.extra_config
            };
            setTimeout(() => {
                this.createChangeChart('chart-demo2', chart.chart_type);
            }, 500);
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
    justify-content: center;
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