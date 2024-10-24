import { defineStore } from 'pinia'
import { fetchWrapper } from '@/assets/utils'

export const useGlobalStore = defineStore('global', {
  state: () => ({
    baseUrl: window.location.href.replace(new RegExp(window.location.pathname + '$'), ''),
    // baseUrl: 'http://localhost:8080',
    account: {
        loaded: false,
        id: '',
        name: '',
        email: '',
        account_apps: [],
    },
  }),
  actions: {
    updateState(payload) {
        this.$patch(payload)
    },
    async getAccount() {
      if (this.account.loaded) {
        return Promise.resolve(this.account)
      } else {
        const data = await fetchWrapper('/api/account/auth')
        this.updateState({
          account: {
            loaded: true,
            id: data.id,
            name: data.name,
            email: data.email,
            account_apps: data.apps,
          }
        })
        return this.account
      }
    },
    clearAccount() {
      this.updateState({
        account: {
          loaded: false,
          id: '',
          name: '',
          email: '',
          account_apps: [],
        }
      })
    }
  }
})
