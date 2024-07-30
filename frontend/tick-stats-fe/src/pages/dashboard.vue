<template>
    <AppBar/>
    <div style="padding: 128px; display: flex; flex-direction: column; align-items: center;">
        <!-- h1 靠左 -->
        <h1 style="text-align: left; margin-bottom: 32px;">Your Apps</h1>
        <EmptyApplication @create="fetchApps"/>
        <ApplicationItem v-for="app in apps" :app="app" :key="app.name" @delete="fetchApps"/>
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
            apps: []
        };
    },
    mounted() {
        this.fetchApps();
    },
    methods: {
        fetchApps() {
            fetch('https://ts.lwl.lol/api/account/app', {credentials: 'include'})
            .then(response => response.json())
            .then(data => {
                let apps = []
                for (let i = 0; i < data.length; i++) {
                    apps.push({
                        name: data[i].name,
                        description: "",
                        app_id: data[i].app_id,
                        created_time: data[i].created_time,
                        updated_time: data[i].updated_time,
                    })
                }
                this.apps = apps;
            }).catch(error => {
                console.error(error);
            });
        }
    }
}

</script>