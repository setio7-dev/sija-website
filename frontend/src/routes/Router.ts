import { createWebHistory, createRouter } from "vue-router";
import LandingPage from "../pages/LandingPage.vue";
import Login from "../pages/Login.vue";

const routes = [
    {
        path: "/",
        component: LandingPage
    },
    {
        path: "/masuk",
        component: Login
    },
];

export const Router = createRouter({
    history: createWebHistory(),
    routes
})