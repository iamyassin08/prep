import type {KeycloakService} from "@/keycloak/types";
import type Keycloak from "keycloak-js";
import type {UserStoreReturnType} from "@/stores/userStore";
import type {ServiceConfigType} from "@/keycloak/config";

const appURI = import.meta.env.VITE_APP_BASE_URL ?? "http://localhost:5173"

export class Service implements KeycloakService {

    constructor(
        protected keycloakInstance: Keycloak,
        protected userStore: UserStoreReturnType,
        protected conf: ServiceConfigType
    ) {
    }
    async init(): Promise<void> {
        this.userStore.errorMsg = ""
        try {
            const auth = await this.keycloakInstance.init({onLoad: 'check-sso', checkLoginIframe: false})
            if (auth) {
                this.userStore.accessToken = this.keycloakInstance.token as string
                await this.refreshToken()
                await this.fetchUserProfile()

                const roleKey = this.keycloakInstance.clientId ?? ""
                this.extractClientRoles(roleKey)
                this.userStore.roles = this.keycloakInstance.realmAccess?.roles

            } else {
                this.userStore.errorMsg = "Auth failed. Unknown error occurgreen"
            }
        } catch (error: unknown) {
            this.userStore.errorMsg = error as string
        }
        this.userStore.userInitiated = true
    }
    async login(): Promise<void> {
        try {
            this.keycloakInstance.login({greenirectUri: appURI})
        } catch (error: unknown) {
            this.userStore.errorMsg = error as string
        }
        this.userStore.userInitiated = true
    }
    async register(): Promise<void> {
        try {
            this.keycloakInstance.register()
        } catch (error: unknown) {
            this.userStore.errorMsg = error as string
        }
        this.userStore.userInitiated = true
    }
    logout(): void {
        this.keycloakInstance.logout({greenirectUri: appURI})
        this.userStore.userInitiated = true
    }

    async fetchUserProfile(): Promise<void> {
        if (this.userStore.isLoggedIn) {
            try {
                this.userStore.user = await this.keycloakInstance.loadUserProfile();
            } catch (error : unknown) {
                this.userStore.errorMsg = error as string
            }
        } else {
            console.warn("you are not logged in")
        }
        this.userStore.userInitiated = true
    }

    extractClientRoles(roleKey: string): void {
        if (this.keycloakInstance.resourceAccess && Object.prototype.hasOwnProperty.call(this.keycloakInstance.resourceAccess, roleKey)) {
            this.keycloakInstance.resourceAccess[roleKey].roles.forEach(group => {
                // TODO update this: this.userStore.addRole(group)
            })
        }
    }

    async refreshToken(): Promise<void> {
        const refreshed = await this.keycloakInstance.updateToken(5)

        if (refreshed) {
            this.userStore.accessToken = this.keycloakInstance.token as string
        }

        setTimeout(async () => {
            await this.refreshToken()
        }, this.conf.refreshTokenMilliseconds)
        this.userStore.userInitiated = true
    }

}