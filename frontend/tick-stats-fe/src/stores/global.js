import { defineStore } from 'pinia'

export const useGlobalStore = defineStore('global', {
  state: () => ({
    // baseUrl: window.location.href.replace(new RegExp(window.location.pathname + '$'), ''),
    baseUrl: 'http://localhost:8080',
    account: {
        id: '',
        name: '',
        email: '',
    },
    account_apps: [],
  }),
  actions: {
    updateState(payload) {
        this.$patch(payload)
    }
  }
})