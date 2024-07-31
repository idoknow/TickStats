<template>
    <v-dialog max-width="500">
        <template v-slot:activator="{ props: activatorProps }">
            <v-list-item  v-bind:="activatorProps" class="dash-item" border="opacity-50 md"
                lines="two" max-width="600" rounded="lg" variant="flat">
                <div style="display: flex; align-items: center; justify-content: center; width: 100%; height: 100%; flex-direction:column">
                    <v-icon size="48" color="#777">mdi-plus</v-icon>
                    <p class="text-caption" style="color: #a3a3a3">Create a new one</p>
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

<script>
import EmojiSelector from './EmojiSelector.vue';

export default {
    name: 'EmptyApplication',
    data() {
        return {
            rules: {
                required: value => !!value || 'Required.',
            },
            toast: {
                show: false,
                text: '',
                color: 'primary',
                timeout: 3000,
            },
            newApp: {
                name: '',
                public: false,
                emoji: '🚀',
            }
        }
    },
    components: {
        EmojiSelector
    },
    methods: {
        makeToast(text, color = 'primary', timeout = 3000) {
            this.toast.text = text;
            this.toast.color = color;
            this.toast.timeout = timeout;
            this.toast.show = true;
        },
        createApplication(isActive) {
            fetch('https://ts.lwl.lol/api/account/app/new', {
                method: 'POST',
                credentials: 'include',
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify(this.newApp)
            }).then((response) => {
                if (response.ok) {
                    isActive.value = false;
                    this.newApp = {
                        name: '',
                        public: false,
                        emoji: '🚀',
                    };
                    this.makeToast('Application created successfully 🎉', 'success');
                    this.$emit('create');
                } else {
                    this.makeToast('Failed to create application', 'error');
                }
            });
        }
    }
}
</script>


<style>

.dash-item {
    cursor: pointer; 
    margin-bottom: 8px; 
    min-width:100%;
}

@media (max-width: 600px) {
}

</style>