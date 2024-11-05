export interface KeycloakService {
  init() : Promise<void>
  login() : Promise<void>
  register() : Promise<void>
  fetchUserProfile() : Promise<void>
  logout(): void
}