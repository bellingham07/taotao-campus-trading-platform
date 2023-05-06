import {createRouter,createWebHashHistory} from 'vue-router';
import LoginForm from "@/components/login/LoginForm";
import MainLayout from "@/layout/MainLayout";

const routes = [
    {
        path:'/login',
        component:LoginForm
    },
    {
        path: '/tao',
        component: MainLayout
    }
]


export default createRouter({history:createWebHashHistory(),routes})

