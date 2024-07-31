import { defineStore } from 'pinia'

export const useGlobalStore = defineStore('global', {
  state: () => ({
    baseUrl: 'http://localhost:8080',
    account: {
        name: '',
        email: '',
    }
  }),
  actions: {
    updateState(payload) {
        this.$patch(payload)
    }
  }
})