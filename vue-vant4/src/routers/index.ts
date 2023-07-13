import {createRouter, createWebHashHistory, Router} from 'vue-router'

const routes = [
    {
        path: '/',
        redirect: '/home/new',
        component: () => import('../views/Main.vue'),
        children: [
            {
                path: 'home',
                component: () => import('../views/main/Home.vue'),
                children: [
                    {
                        path: 'new',
                        component: () => import('../components/home/New.vue')
                    },
                    {
                        path: 'sell',
                        component: () => import('../components/home/Sell.vue')
                    },
                    {
                        path: 'want',
                        component: () => import('../components/home/Want.vue')
                    }
                ]
            },
            {
                path: 'atcl',
                component: () => import('../views/main/Atcl.vue')
            },
            {
                path: 'cmdty',
                component: () => import('../views/main/Cmdty.vue')
            },
            {
                path: 'msg',
                component: () => import('../views/main/Msg.vue')
            },
            {
                path: 'user',
                component: () => import('../views/main/User.vue')
            }
        ]
    },
    {
        path: '/login',
        component: () => import('../views/Login.vue')
    }
]

const router: Router = createRouter({history: createWebHashHistory(), routes})

export default router