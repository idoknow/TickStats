<template>
    <v-dialog max-width="500">
        <template v-slot:activator="{ props: activatorProps }">
            <v-list-item  v-bind:="activatorProps" class="dash-item" border="opacity-50 md"
                lines="two" max-width="600" rounded="lg" variant="flat">
                <div style="display: flex; align-items: center; justify-content: center; width: 100%; height: 100%;">
                    <v-icon size="48" color="#777">mdi-plus</v-icon>
                </div>
            </v-list-item>
        </template>

        <template v-slot:default="{ isActive }">
            <v-card title="Create a application">
                
                <v-card-text>
                    <emoji-selector style="margin-top: 16px; margin-bottom: 24px;" @select="emoji => newApp.emoji = emoji"></emoji-selector>
                    <p style="font-size: 52px; text-align:center; margin-bottom: 16px;">{{ newApp.emoji }}</p>
                    <v-text-field v-model="newApp.name" label="App name" variant="outlined" :rules="[rules.required]"></v-text-field>
                    <v-checkbox v-model="newApp.public" label="Public" color="primary"></v-checkbox>

                    <small>* Public applications are visible to everyone. You can modify these settings later.</small>

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
import EmojiSelector from './EmojiSelector.vue';

var newApp = ref({
    name: '',
    public: false,
    emoji: '🚀',
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

const createApplication = (isActive) => {
    fetch('https://ts.lwl.lol/api/account/app/new', {
        method: 'POST',
        credentials: 'include',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify(newApp.value)
    }).then((response) => {
        if (response.ok) {
            isActive.value = false;
            newApp.value = {
                name: '',
                public: false,
                emoji: '🚀',
            };
            makeToast('Application created successfully 🎉', 'success');
            this.$emit('create');
        } else {
            makeToast('Failed to create application', 'error');
        }
    });
}
</script>

<script>
export default {
    name: 'EmptyApplication',
    data() {
        return {
            rules: {
                required: value => !!value || 'Required.',
            }
        }
    },
    components: {
        EmojiSelector
    },
}
</script>


<style>

.dash-item {
    cursor: pointer; 
    margin-bottom: 24px; 
    min-width:600px
}

@media (max-width: 600px) {
    .dash-item {
        min-width: 360px;
    }
}

</style>