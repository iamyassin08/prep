import Keycloak from "keycloak-js";
import {keycloakJsConfig, serviceConfig} from "@/keycloak/config";
import type {KeycloakService} from "@/keycloak/types";
import type {UserStoreReturnType} from "@/stores/userStore";
import {Service} from "@/keycloak/service";

let s : KeycloakService | null = null;

export function createKeycloakInstance(): Keycloak {
    return new Keycloak(keycloakJsConfig)
}
export function serviceFactory( userStore: UserStoreReturnType): KeycloakService {
    if(s === null){
            s = new Service(createKeycloakInstance(), userStore, serviceConfig)
    }
    return s
}
