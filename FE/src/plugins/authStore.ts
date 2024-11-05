// import { serviceFactory } from "@/keycloak/factory";
// import { useUserStore } from "@/stores/userStore";


// const userAuthStorePlugin = {
//     install(app :any, option: any) {
//         const store = useUserStore(option.pinia)
//         app.config.globalProperties.$store = store
//         const keycloakService = serviceFactory(store);
//         keycloakService.init()   
//     }
// }

// export default userAuthStorePlugin;