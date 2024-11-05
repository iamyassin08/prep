import { defineStore } from 'pinia'
import type { KeycloakProfile } from 'keycloak-js'

export type UserStoreReturnType = ReturnType<typeof useUserStore>

export type UserStore = {
  user: KeycloakProfile | null
  userInitiated: boolean
  accessToken: string | null
  roles: string[] | undefined
  darkMode: boolean
  errorMsg: string
};

export const useUserStore = defineStore({
  id: 'user',
  state: (): UserStore => ({
    user: null,
    accessToken: '',
    userInitiated: false,
    roles: [],
    darkMode: false, 
    errorMsg: '',
  }),
  persist: false,
  getters: {
    isLoggedIn(): boolean {
      return !!this.accessToken;
    },
    isAdmin(): boolean {
      if (this.roles?.includes('admin')) {
        return true
      } else {
        return false
      }
    },
    isSeller(): boolean {
      if (this.roles?.includes('seller')) {
        return true
      } else {
        return false
      }
    },
    isBuyer(): boolean {
      if (this.roles?.includes('buyer')) {
        return true
      } else {
        return false
      }
    },
  },
// TODO update this:
  // actions: {
  //     addRole(role: string) : void {
  //         this.roles.push(role)
  //     },

  // }
})
