<template>
    <AppBar nav="0" @error="(err) => { showErrAlert = true; errAlert = err; }" />
    <div class="dash-content">
        <div style="margin-bottom: 16px;">
            <h1>Your Apps</h1>
            <p class="text-caption" style="color: #a3a3a3;">Manage your apps here</p>
        </div>

        <v-alert v-model="showErrAlert" type="info" variant="tonal" closable style="margin: 16px 0px">
            {{ errAlert }}
        </v-alert>

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
            loading: false,
            showErrAlert: false,
            errAlert: ''
        };
    },
    mounted() {
        this.fetchApps();
        this.apps.push({
            name: "test",
            description: "test",
            emoji: "🚀",
            app_id: "test",
            created_time: "test",
            updated_time: "test",
        })
        this.apps.push({
            name: "test",
            description: "test",
            emoji: "😄",
            app_id: "test",
            created_time: "test",
            updated_time: "test",
        })
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
.dash-content {
    padding: 32px;
    display: flex;
    flex-direction: column;
    align-items: center;
    text-align: center;
    justify-content: center;
    max-width: 600px;
    margin: 0 auto;
}

@media (max-width: 600px) {
    .dash-content {
        padding: 16px;
        padding-top: 32px;
    }
}
</style>