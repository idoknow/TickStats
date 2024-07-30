<template>
    <v-dialog max-width="500">
        <template v-slot:activator="{ props: activatorProps }">
            <v-list-item v-bind:="activatorProps" style="cursor: pointer; margin-bottom: 24px" border="opacity-50 md"
                lines="two" max-width="600" min-width="600" rounded="lg" variant="flat">
                <div style="display: flex; align-items: center; justify-content: center; width: 100%; height: 100%;">
                    <v-icon size="48" color="#777">mdi-plus</v-icon>
                </div>
            </v-list-item>
        </template>

        <template v-slot:default="{ isActive }">
            <v-card title="Create a application">

                <v-card-text>
                    <v-text-field v-model="appName" label="App name" outlined></v-text-field>
                </v-card-text>

                <v-card-actions>
                    <v-spacer></v-spacer>
                    <v-btn @click="createApplication(isActive)">
                        Create
                    </v-btn>
                    <v-btn text="Close" @click="; isActive.value = false;"></v-btn>
                </v-card-actions>
            </v-card>
        </template>
    </v-dialog>


    <v-snackbar v-model="toast.show" :color="toast.color" :timeout="toast.timeout">
        {{ toast.text }}
    </v-snackbar>


</template>

<script setup>
import { ref } from 'vue';


var appName = ref('');

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

const createApplication = (isActive) => {
    fetch('https://ts.lwl.lol/api/account/app/new', {
        method: 'POST',
        credentials: 'include',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify({
            name: appName.value,
        })
    }).then((response) => {
        if (response.ok) {
            isActive.value = false;
            appName.value = '';
            makeToast('Application created successfully', 'success');
            this.$emit('create');
        } else {
            makeToast('Failed to create application', 'error');
        }
    });
}
</script>
