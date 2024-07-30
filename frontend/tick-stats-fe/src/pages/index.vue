<template>
  <AppBar />
  <div class='index-main'>
    <h1 class="gradient index-title">Tick Stats</h1>
    <p style="font-size: 22px; color: #777">A simple & easy Metric Collector</p>
    <v-btn color="primary" size="x-large" style="margin-top: 24px;" @click="$router.push('/dashboard')">
      Get Started
    </v-btn>
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
    padding-top: 64px;
  }

  .index-title {
    font-size: 48px;
  }
}

</style>

