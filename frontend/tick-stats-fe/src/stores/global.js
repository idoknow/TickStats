import { defineStore } from 'pinia'

export const useGlobalStore = defineStore('global', {
  state: () => ({
    baseUrl: window.location.href.replace(/\/$/, ''),
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