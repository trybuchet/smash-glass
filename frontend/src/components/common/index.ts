import AppLoader from "./AppLoader.vue";

export default {
    AppLoader,
};

declare module 'vue' {
    export interface GlobalComponents {
        AppLoader: typeof AppLoader;
    }
}