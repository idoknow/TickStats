<template>
  <AppBar />
  <div class='index-main'>
    <h1 class="gradient index-title">Tick Stats</h1>
    <p style="font-size: 22px; color: #777">A simple & easy Metric Collector</p>
    <v-btn color="primary" size="x-large" style="margin-top: 24px;" @click="$router.push('/dashboard')">
      Get Started
    </v-btn>
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
}


</style>

