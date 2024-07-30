<template>
    <AppBar />
    <div class="index-main">
        <h1 class="gradient index-title">How to use?</h1>
        <h2>Post Data</h2>
        <p class="text">Post your metric data via <span class="quote">`POST https://mc.lwl.lol/api/metric/:appid`</span> </p>
        <p class="text">Body should be a JSON object with <span class="quote">`metrics_data`</span> field
        </p>
        <p class="text">metrics_data LIMIT (Current version):</p>
        <ul class="text">
            <li>1. should be a JSON object</li>
            <li>2. key should be a string</li>
            <li>3. value should be a number or JSON object(max depth 2)</li>
        </ul>

        <p class="text">Example</p>
        <pre style="font-size: 18px; color: #333; margin-top: 20px">
{
    "metrics_data": {
        "cpu": 0.5,
        "memory": {
            "used": 1024,
            "total": 2048
        },
        "os_name": "windows"
    }
}
        </pre>

        <h2>Data Format of Chart</h2>

        <p class="text">Currently we support 2 types of chart.</p>

        <ul class="text">
            <li>1. Simple Line Chart</li>
            <li>2. Simple Pie Chart</li>
            <li>3. ... We are developing more types of chart.</li>
        </ul>

        <p class="text">Every chart should bind a key in the metric data you posted. </p>
        <h3 style="margin-top: 16px;">Simple Line Chart</h3>
        <p class="text">The value should be a number</p>
        <p>Currently, we will aggregate and retrieve numeric data using a 30-minute time bucket. </p>
        </br>
        <h3 style="margin-top: 16px;">Simple Pie Chart</h3>
        <p class="text">The value should be a String</p>
        <p>Currently we'll count the number of each value in the last 1 hour in the metric data you posted.</p>

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
    font-size: 17px;
    color: #333;
    background-color: #eeeeee;
    padding: 4px;
    border-radius: 4px;
    font-weight: bold;
}


.index-main {
    display: flex; 
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
      padding: 32px;
    }
  
    .index-title {
      font-size: 48px;
    }
  }
</style>