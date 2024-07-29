<template>
    <AppBar />

    <div class="content" style="padding: 32px 128px; display: flex; flex-direction: column; align-items: center;">
        <h1 class="gradient" v-if="accountName !== ''" style="font-size: 56px;">{{ accountName }}/{{ appName }} Stats
        </h1>
        <v-btn @click="mock" variant="plain">DEBUG: MOCK RANDOM DATA</v-btn>
        <v-dialog max-width="500">
            <template v-slot:activator="{ props: activatorProps }">
                <v-btn style="margin: 32px;" v-bind:="activatorProps">Create Chart</v-btn>
            </template>

            <template v-slot:default="{ isActive }">
                <v-card title="Create Chart">

                    <v-card-text>
                        <v-text-field readonly v-model="appId" label="App Id" outlined></v-text-field>
                        <v-text-field v-model="newChart.chartName" label="Chart Name" outlined></v-text-field>
                        <v-radio-group v-model="newChart.chartType" row>
                            <v-radio label="Simple Line Chart" value="simple_line"></v-radio>
                            <v-radio label="Pie Chart" value="simple_pie"></v-radio>
                        </v-radio-group>
                        <v-text-field v-model="newChart.keyName" label="Key Name" outlined></v-text-field>
                    </v-card-text>

                    <v-card-actions>
                        <v-spacer></v-spacer>
                        <v-btn text @click="createChart(isActive)">
                            Create
                        </v-btn>
                        <v-btn text="Close" @click="isActive.value = false;"></v-btn>
                    </v-card-actions>
                </v-card>
            </template>
        </v-dialog>
    </div>



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
            chartInstance: {},
            mock_os: ['Windows', 'Linux', 'MacOS'],
            chartData: null,
            appId: '',
            accountName: '',
            appName: '',
            newChart: {
                chartName: '',
                keyName: '',
                chartType: '',
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
                    borderRadius: 10
                },
                legend: {
                    orient: 'vertical',
                    left: 'left',
                    data: []
                },
            },
            chartOptions: []
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
        createChart(isActive) {
            fetch(`https://ts.lwl.lol/api/account/app/${this.appId}/chart/new`, {
                method: 'POST',
                credentials: 'include',
                headers: {
                    'Content-Type': 'application/json'
                },
                body: JSON.stringify({
                    appid: this.appId,
                    chart_name: this.newChart.chartName,
                    key_name: this.newChart.keyName,
                    chart_type: this.newChart.chartType
                })
            }).then(response => {
                if (response.ok) {
                    isActive.value = false;
                }
            }).catch(error => {
                console.error(error);
            });
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
                                _chartType: 'simple_line'
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
                                _chartType: 'simple_pie'
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
</style>