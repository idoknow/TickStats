/**
 * main.js
 *
 * Bootstraps Vuetify and other plugins then mounts the App`
 */

// Plugins
import { registerPlugins } from '@/plugins'

// Components
import App from './App.vue'

// Composables
import { createApp } from 'vue'

const app = createApp(App)

const baseUrl = "https://ts.lwl.lol"
app.config.globalProperties.baseUrl = baseUrl;

registerPlugins(app)

app.mount('#app')
