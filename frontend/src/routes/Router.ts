import { createWebHistory, createRouter } from "vue-router";
import LandingPage from "../pages/LandingPage.vue";

const routes = [
    {
        path: "/",
        component: LandingPage
    }
];

export const Router = createRouter({
    history: createWebHistory(),
    routes
})