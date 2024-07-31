<template>
    <AppBar />
    <div class='index-main'>
        <h1 class="gradient index-title">Worlds' Stats</h1>
        <p style="font-size: 22px; color: #777; margin-bottom:32px">See the public stats.</p>

        <ApplicationItem v-for="app in apps" :key="app.app_id" :app="app" />
    </div>


</template>

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

<script>
import AppBar from '@/components/AppBar.vue';
import ApplicationItem from '@/components/ApplicationItem.vue';

export default {
    name: 'World',
    components: {
        AppBar,
        ApplicationItem
    },
    data() {
        return {
            apps: [],
            loading: false
        };
    },
    mounted() {
        this.getPublicApps();
    },
    methods: {
        getPublicApps() {
            this.loading = true;
            fetch('https://ts.lwl.lol/api/stats/apps', { credentials: 'include' })
                .then(response => response.json())
                .then(data => {
                    this.apps = data;
                    this.loading = false;
                }).catch(error => {
                    console.error(error);
                    this.loading = false;
                });
        }
    }
}

</script>