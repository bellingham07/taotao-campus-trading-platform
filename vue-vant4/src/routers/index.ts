import {createRouter, createWebHashHistory, Router} from 'vue-router'
import Index from "../views/Home.vue";
import Login from "../views/Login.vue";

const routes = [
    {
        path: '/',
        component: Index
    },
    {
        path: '/login',
        component: Login
    }
]

const router: Router = createRouter({history: createWebHashHistory(), routes})

export default router