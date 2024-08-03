<template>
    <AppBar nav="1" />
    <div class='index-main'>
        <h1 class="gradient index-title">Worlds' Stats</h1>
        <p style="font-size: 22px; color: #777; margin-bottom:32px">See the public stats.</p>
        <v-alert v-model="showErrAlert" type="info" variant="tonal" closable style="margin: 16px 0px">
            {{ errAlert }}
        </v-alert>
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
    max-width: 600px;
    margin: 0 auto;
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
import { fetchWrapper } from '@/assets/utils';

export default {
    name: 'World',
    components: {
        AppBar,
        ApplicationItem
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
        this.getPublicApps();
    },
    methods: {
        getPublicApps() {
            this.loading = true;
            fetchWrapper('/api/stats/apps', {
                method: 'GET',
            }).then((data) => {
                this.apps = data;
            }).catch((err) => {
                this.showErrAlert = true;
                this.errAlert = 'Something went wrong, please try again later: ' + err;
            }).finally(() => {
                this.loading = false;
            });
        }
    }
}

</script>