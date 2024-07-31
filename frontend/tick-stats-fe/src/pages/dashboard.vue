<template>
    <AppBar nav="0" />
    <div class="content">
        <h1 style="text-align: left; margin-bottom: 32px;">Your Apps</h1>
        <EmptyApplication @create="fetchApps" />
        <ApplicationItem style="margin-top: 16px;" v-for="app in apps" :app="app" :key="app.name" @delete="fetchApps" />
        <v-progress-circular v-if="loading" color="primary" indeterminate></v-progress-circular>
        <p v-if="apps.length == 0 && !loading"> Hmm... You don't have any apps yet. </p>
    </div>

</template>

<script>
import AppBar from '@/components/AppBar.vue';
import ApplicationItem from '@/components/ApplicationItem.vue';
import EmptyApplication from '@/components/EmptyApplication.vue';


export default {
    name: 'Dashboard',
    components: {
        AppBar,
        ApplicationItem,
        EmptyApplication
    },
    data() {
        return {
            apps: [],
            loading: false
        };
    },
    mounted() {
        this.fetchApps();
    },
    methods: {
        fetchApps() {
            this.loading = true;
            fetch('https://ts.lwl.lol/api/account/app', { credentials: 'include' })
                .then(response => response.json())
                .then(data => {
                    let apps = []
                    for (let i = 0; i < data.length; i++) {
                        apps.push({
                            name: data[i].name,
                            description: "",
                            emoji: data[i].emoji,
                            app_id: data[i].app_id,
                            created_time: data[i].created_time,
                            updated_time: data[i].updated_time,
                        })
                    }
                    this.loading = false;
                    this.apps = apps;
                }).catch(error => {
                    console.error(error);
                    this.loading = false;
                });
        }
    }
}

</script>

<style>

.content {
    padding: 128px; 
    display: flex; 
    flex-direction: column; 
    align-items: center;
}

@media (max-width: 600px) {
    .content {
        padding: 32px;
        padding-top: 128px;
    }
}

</style>