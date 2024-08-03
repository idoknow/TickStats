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
        <ApplicationItem style="margin-top: 8px;" v-for="app in apps" :app="app" :key="app.name" from="dashboard" @delete="fetchApps" />
        <v-progress-circular v-if="loading" color="primary" indeterminate></v-progress-circular>
        <p v-if="apps.length == 0 && !loading"> Hmm... You don't have any apps yet. </p>

    </div>

</template>

<script>
import AppBar from '@/components/AppBar.vue';
import ApplicationItem from '@/components/ApplicationItem.vue';
import EmptyApplication from '@/components/EmptyApplication.vue';
import { fetchWrapper } from '@/assets/utils';
import { useGlobalStore } from '@/stores/global';


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
            errAlert: '',
            global: useGlobalStore()
        };
    },
    mounted() {
        this.fetchApps();
    },
    methods: {
        fetchApps() {
            this.loading = true;
            fetchWrapper('/api/account/app', {
                method: 'GET',
            }).then((data) => {
                this.apps = data;
                this.global.updateState({
                    account_apps: data
                });
            }).catch((err) => {
                this.showErrAlert = true;
                if (err.status === 401) {
                    this.errAlert = 'You need to sign in to view your apps.';
                } else {
                    this.errAlert = 'Something went wrong, please try again later: ' + err;
                }
            }).finally(() => {
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