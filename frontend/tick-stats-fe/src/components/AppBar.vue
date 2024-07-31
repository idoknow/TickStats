<template>

    <v-app-bar color="primary">
        <v-app-bar-title style="font-weight: 1000; cursor: pointer"
            @click="$router.push('/')">TickStats</v-app-bar-title>

        <template v-slot:prepend>
            <v-app-bar-nav-icon @click="menuRail = !menuRail; mobileDrawer = !mobileDrawer"></v-app-bar-nav-icon>
        </template>

        <template v-slot:append>
            <v-dialog v-model="loginDialog" max-width="500" v-if="account.name == ''">
                <template v-slot:activator="{ props: activatorProps }">
                    <v-btn v-bind:="activatorProps" variant="plain">Login</v-btn>
                </template>

                <template v-slot:default="{ isActive }">
                    <v-card :title="AuthTitle[isNewUser]">

                        <v-card-text>
                            <v-text-field v-model="credentials.username" label="Username" variant="outlined"
                                </v-text-field>
                                <v-text-field v-model="credentials.email" label="Email" variant="outlined"
                                    v-if="isNewUser"></v-text-field>
                                <v-text-field v-model="credentials.password" label="Password" variant="outlined"
                                    type="password"></v-text-field>

                                <a style="text-decoration: underline; cursor: pointer" @click="isNewUser = !isNewUser">
                                    {{
                                        AuthHint[isNewUser] }} </a>
                        </v-card-text>

                        <v-card-actions>
                            <v-spacer></v-spacer>
                            <v-btn text @click="login(isActive)">
                                {{ AuthTitle[isNewUser] }}
                            </v-btn>
                            <v-btn text="Close"
                                @click="credentials.username = ''; credentials.password = ''; isActive.value = false;"></v-btn>
                        </v-card-actions>
                    </v-card>
                </template>
            </v-dialog>
            <span v-else style="margin-right: 16px;"> {{ account.name }} </span>
        </template>
    </v-app-bar>

    <v-navigation-drawer v-if="!isMobile" permanent :rail="menuRail" app>
        <v-divider></v-divider>
        <v-list nav>
            <v-list-item v-for="item in items" :key="item.title" :prepend-icon="item.icon" :title="item.title"
                :value="item.value" @click="router(item.value)" :active="nav == item.value" </v-list-item>

        </v-list>
    </v-navigation-drawer>
    <v-navigation-drawer v-else temporary v-model="mobileDrawer" app>
        <v-list nav>
            <v-list-item v-for="item in items" :key="item.title" :prepend-icon="item.icon" :title="item.title"
                :value="item.value" @click="router(item.value)" :active="nav == item.value" </v-list-item>
            <v-divider></v-divider>
        </v-list>
    </v-navigation-drawer>

    <v-snackbar v-model="toast.show" :color="toast.color" :timeout="toast.timeout">
        {{ toast.text }}
    </v-snackbar>
</template>

<script setup>
import { ref } from 'vue';

const AuthTitle = ref({
    true: 'Sign up',
    false: 'Login',
})
const AuthHint = ref({
    true: 'Already have an account?',
    false: 'Don\'t have an account?',
})
const isNewUser = ref(false);
const credentials = ref({
    username: '',
    email: '',
    password: '',
});

const toast = ref({
    show: false,
    text: '',
    color: 'primary',
    timeout: 3000,
});
const makeToast = (text, color = 'primary', timeout = 3000) => {
    toast.value.text = text;
    toast.value.color = color;
    toast.value.timeout = timeout;
    toast.value.show = true;
};

const login = (isActive) => {
    if (isNewUser.value) {
        fetch('https://ts.lwl.lol/api/account/register', {
            method: 'POST',
            credentials: 'include',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({
                name: credentials.value.username,
                email: credentials.value.email,
                password: credentials.value.password,
            }),
        }).then((response) => {
            if (response.status === 200) {
                makeToast('Register successful', 'success');
                isActive.value = false;
                isNewUser.value = false;
                credentials.value.email = '';
                credentials.value.username = '';
                credentials.value.password = '';
            } else {
                makeToast('Register failed', 'error');
            }
        }).catch((err) => {
            makeToast(err, 'error');
        });
    } else {
        fetch('https://ts.lwl.lol/api/account/login', {
            method: 'POST',
            credentials: 'include',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({
                username: credentials.value.username,
                password: credentials.value.password,
            }),
        }).then((response) => {
            if (response.status === 200) {
                makeToast('Login successful', 'success');
                isActive.value = false;
                auth();
            } else {
                makeToast('Login failed', 'error');
            }

        }).catch((err) => {
            makeToast(err, 'error');
        });
    }
};

</script>

<script>
export default {
    data() {
        return {
            loginDialog: false,
            isMobile: false,
            mobileDrawer: false,
            menuRail: false,
            items: [
                { title: 'Home', icon: 'mdi-home', value: -1 },
                { title: 'Your Applications', icon: 'mdi-view-dashboard', value: 0 },
                { title: 'World Stats', icon: 'mdi-chart-bar', value: 1 },
                { title: 'Settings', icon: 'mdi-cog', value: 2 },
                { title: 'Sign in/up', icon: 'mdi-account', value: 3 },
            ],  
            account: {
                name: '',
                email: '',
            }
        };
    },
    props: {
        nav: {
            type: Number,
            default: -1,
        },
    },
    mounted() {
        this.checkScreenWidth();
        window.addEventListener('resize', this.checkScreenWidth);
        this.auth();
    },
    methods: {
        auth() {
            fetch('https://ts.lwl.lol/api/account/auth', {
                method: 'GET',
                credentials: 'include',
            }).then((response) => {
                if (response.status === 200) {
                    response.json().then((data) => {
                        account.name = data.name;
                        account.email = data.email;
                    });
                } else {
                    // makeToast('Not logged in', 'error');
                    this.$emit('error', 'You need to login to use this service :)');
                }
            }).catch((err) => {
                // makeToast(err, 'error');
                this.$emit('error', 'Something went wrong, please try again later: ' + err);
            });
        },
        checkScreenWidth() {
            this.isMobile = window.innerWidth <= 768;
        },
        router(value) {
            switch (value) {
                case -1:
                    this.$router.push('/');
                    break;
                case 0:
                    this.$router.push('/dashboard');
                    break;
                case 1:
                    this.$router.push('/world');
                    break;
                case 2:
                    this.$router.push('/settings');
                    break;
                case 3:
                    this.loginDialog = true;
                    break;
            }
        },
    },
};

</script>