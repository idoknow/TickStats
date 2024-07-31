<template>
  <AppBar nav="-1"/>
  <div class='index-main'>
    <svg style="margin-bottom: 16px" t="1722403535188" class="icon" viewBox="0 0 1024 1024" version="1.1" xmlns="http://www.w3.org/2000/svg" p-id="5323" width="150" height="150"><path d="M42.666667 853.333333h938.666666v85.333334H42.666667zM341.333333 87.466667h85.333334V768h-85.333334zM597.333333 512h85.333334v256h-85.333334zM85.333333 384h85.333334v384H85.333333zM853.333333 256h85.333334v512h-85.333334z" p-id="5324" fill="#2196F3"></path></svg>
    <h1 class="gradient index-title">Tick Stats</h1>
    <p style="font-size: 22px; color: #777">A simple & easy Metric Collector</p>
    <div style="display: flex; align-items:center; margin-top: 24px; gap: 16px">
      <v-btn variant="outlined" color="grey" size="x-large" @click="$router.push('/help')">Tutorial</v-btn>
      <v-btn variant="flat" color="primary" size="x-large"  @click="$router.push('/dashboard')">
        My Apps
      </v-btn>
    </div>

    
    <div class="code-wrap">
      <pre class="index-code"><span style="font-weight:bold">POST</span> https://ts.lwl.lol/api/metric/:appid</pre>
    </div>

    <v-row style="margin-top: 64px; gap: 16px" justify="center">
      <v-col cols="12" md="6" lg="4">
        <v-card class="text-center">
          <v-card-title style="margin-top: 16px;">
            <v-icon size="64" color="#777">mdi-rocket</v-icon>
          </v-card-title>
          <v-card-text> 
            <h2>Easy to use</h2>
            <p style="margin-top:16px">Just post your metric data via HTTP API</p>
          </v-card-text>
        </v-card>
      </v-col>
      <v-col cols="12" md="6" lg="4">
        <v-card class="text-center">
          <v-card-title style="margin-top: 16px;">
            <v-icon size="64" color="#777">mdi-shield-check</v-icon>
          </v-card-title>
          <v-card-text>
            <h2>High reliability</h2>
            <p style="margin-top:16px">Highly available and reliable service</p>
          </v-card-text>
        </v-card>
      </v-col>
    </v-row>

    
  </div>

</template>

<script setup>
import AppBar from '@/components/AppBar.vue';

</script>

<script>

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

.index-main {
  height: 100%;
  display: flex; 
  align-items: center; 
  padding: 64px; 
  flex-direction: column;
}

.index-title {
  font-size: 64px;
}

.index-code {
  font-size: 16px;
  color: #777;
}

.code-wrap {
  margin-top: 24px;
  background-color: #f1f1f1;
  padding: 4px;
  border-radius: 4px;
}

.text-center {
  text-align: center;
  min-height: 200px;
  min-width: 200px;
}

@media (max-width: 600px) {
  .index-main {
    padding: 32px;
    padding-top: 64px;
  }

  .index-title {
    font-size: 48px;
  }

  .index-code {
    display: none;
  }
}


</style>

