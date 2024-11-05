export interface KeycloakJsConfigType {
  url: string
  realm: string
  clientId: string
}

export interface ServiceConfigType {
  refreshTokenMilliseconds: number
}

export const keycloakJsConfig : KeycloakJsConfigType = {
  url: import.meta.env.VITE_KEYCLOAK_URL ?? "",
  realm: import.meta.env.VITE_KEYCLOAK_REALM ?? "",
  clientId: import.meta.env.VITE_KEYCLOAK_CLIENT_ID ?? "",
}

export const serviceConfig: ServiceConfigType = {
  refreshTokenMilliseconds: import.meta.env.VITE_APP_REFRESH_TOKEN_MS ?? 50000,
}