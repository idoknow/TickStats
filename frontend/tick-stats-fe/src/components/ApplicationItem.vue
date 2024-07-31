<template>
    <v-list-item  border="opacity-50 md" lines="two" max-width="600" class="dash-item"
         rounded="lg" variant="flat">
        
        <div style="display: flex; align-items: center; gap: 16px">
            <span style="font-size:32px">{{ app.emoji }}</span>
            <div>
                <v-list-item-title>
                    <span class="text-h6">{{ app.name }}</span>
                </v-list-item-title>
        
                <v-list-item-subtitle :opacity="isDelete ? .8 : undefined">
                    <v-scroll-y-reverse-transition mode="out-in">
                        <div v-if="isDelete" key="subscribed" class="text-error text-caption">
                            Are you sure to delete?
                        </div>
        
                        <div v-else key="not-subscribed" class="text-caption">
                            {{ app.app_id }}
                        </div>
                    </v-scroll-y-reverse-transition>
                </v-list-item-subtitle>
            </div>    
        </div>


        <template v-slot:append>
            <v-fade-transition mode="out-in">
                <v-btn :key="`subscribe-${isDelete}`" :border="`thin ${isDelete ? 'error' : 'success'}`"
                    :color="isDelete ? 'success' : 'error'" :prepend-icon="isDelete ? 'mdi-close' : 'mdi-delete'"
                    :slim="isDelete" :text="isDelete ? 'Cancel' : 'Delete'" :variant="isDelete ? 'plain' : 'tonal'"
                    class="me-2 text-none" size="small" flat @click="isDelete = !isDelete"></v-btn>
            </v-fade-transition>

            <v-fade-transition mode="out-in">
                <v-btn :key="`info-${isDelete}`" :color="isDelete ? 'error' : 'primary'"
                    :prepend-icon="isDelete ? 'mdi-check' : 'mdi-open-in-new'"
                    :text="isDelete ? 'Delete' : 'Visit'" class="text-none" size="small" variant="flat"
                    @click="isDelete ? deleteApplication() : visitApplication()"
                    flat></v-btn>
            </v-fade-transition>
        </template>
    </v-list-item>
</template>

<script setup>
import { shallowRef } from 'vue'

const isDelete = shallowRef(false)
</script>

<script>
export default {
    name: 'ApplicationItem',
    props: {
        app: {
            type: Object,
            required: true
        }
    },
    methods: {
        deleteApplication() {
            fetch(`https://ts.lwl.lol/api/account/app/${this.app.app_id}`, {
                method: 'DELETE',
                credentials: 'include'
            }).then(response => {
                if (response.ok) {
                    this.$emit('delete')
                }
            }).catch(error => {
                console.error(error)
            })
        },
        visitApplication() {
            this.$router.push(`/app?id=${this.app.app_id}`)
        }
    }
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