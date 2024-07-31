<template>
    <AppBar />
    <div class="index-main">
        <h1 class="gradient index-title">How to use?</h1>
        <v-timeline side="end" style="width: 100%;">
            <v-timeline-item dot-color="primary" size="small" >
                <div>
                    <strong class="me-4">0. Sign in</strong>
                    <div class="text-caption">
                        Click the <span class="quote">`Login`</span> button at the top right corner.
                    </div>
                </div>
            </v-timeline-item>
            <v-timeline-item dot-color="primary" size="small" >
                <div>
                    <strong class="me-4">1. Create Your Application</strong>
                    <div class="text-caption">
                        Create your application at the dashboard page. <a href="/dashboard">Go to Dashboard</a>
                    </div>
                </div>
            </v-timeline-item>

            <v-timeline-item dot-color="primary" size="small" >
                <div>
                    <strong class="me-4">2. Create a new chart</strong>
                    <div class="text-caption">
                        Click <span class="quote">`Visit`</span> button at the dashboard page, you will be navigated to
                        the chart page.
                    </div>
                    <div class="text-caption">
                        Currently, we support 2 types of chart.
                    </div>

                </div>
            </v-timeline-item>

            <v-timeline-item dot-color="primary" size="small" >
                <div>
                    <strong class="me-4">3. Post your metric data</strong>
                    <div class="text-caption">
                        Post your metric data via <span class="quote">`POST https://ts.lwl.lol/api/metric/:appid`</span>
                    </div>
                    <div class="text-caption">
                        Body should be a JSON object with <span class="quote">`metrics_data`</span> field. 
                    </div>
                    <div class="text-caption">
                        <span class="quote">`metrics_data`</span> should be a JSON object, and the value should be a number, string or JSON object(max depth 2).
                    </div>
                    <v-card style="margin-top: 8px; max-width: 600px; min-height: 200px">
                        <v-tabs v-model="chartTab" align-tabs="start" color="deep-purple-accent-4">
                            <v-tab v-for="(chart, index) in charts" :key="index" :value="index">
                                {{ chart.name }}
                            </v-tab>
                        </v-tabs>
                        <v-tabs-window v-model="chartTab">
                            <v-tabs-window-item v-for="n in 3" :key="n" :value="n">
                                <div style="padding: 16px" class="text-caption">
                                    <p>{{ charts[chartTab].description }}</p>
                                    <p v-if="charts[chartTab].dataExample != ''">Data Example</p>
                                    <pre style="color: #333; margin-top: 20px">{{ charts[chartTab].dataExample }}</pre>
                                </div>
                            </v-tabs-window-item>
                        </v-tabs-window>
                    </v-card>
                </div>
            </v-timeline-item>
        </v-timeline>
    </div>


</template>

<script>
import AppBar from '@/components/AppBar.vue';
export default {
    name: 'Index',
    components: {
        AppBar
    },
    data() {
        return {
            items: [
                {
                    id: 1,
                    color: 'info',
                    icon: 'mdi-plus',
                },
                {
                    id: 2,
                    color: 'error',
                    icon: 'mdi-chart-line',
                },
            ],
            chartTab: 0,
            charts: [
                {
                    name: "Simple Line Chart",
                    description: "The value should be a number. Currently, we will aggregate and retrieve numeric data using a 30-minute time bucket.",
                    dataExample: `{
    "metrics_data": {
        "usage_cnt": 0.5,
    }
}`
                },
                {
                    name: "Simple Pie Chart",
                    description: "The value should be a String. Currently we'll count the number of each value in the last 1 hour in the metric data you posted.",
                    dataExample: `{
    "metrics_data": {
        "os_name": "windows",
    }
}`
                },
                {
                    name: "...",
                    description: "We will support more types of charts soon.",
                    dataExample: ""
                }
                    
            ]
        };
    },
    mounted() {
    },
    methods: {
    }
}
</script>

<style>
.gradient {
    background: linear-gradient(45deg, #2196F3, #21CBF3);
    -webkit-background-clip: text;
    -webkit-text-fill-color: transparent;
}

.text {
    font-size: 17px;
    color: #333;
    margin-top: 16px;
}

.quote {
    color: #2196F3;
    background-color: #eeeeee;
    padding: 2px;
    border-radius: 4px;
    font-weight: bold;
}


.index-main {
    display:flex;
    align-items: center;
    justify-content: center;
    padding: 64px;
    flex-direction: column;
}

.index-title {
    font-size: 64px;
}
@media (max-width: 600px) {
    .index-main {
        padding: 8px;
    }

    .index-title {
        margin: 16px;
        font-size: 36px;
    }
}
</style>