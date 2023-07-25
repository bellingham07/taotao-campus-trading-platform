import {createRouter, createWebHashHistory, Router} from 'vue-router'

const routes = [
    {
        path: '/',
        redirect: '/home/sell',
        component: () => import('../views/Main.vue'),
        children: [
            {
                path: 'home',
                component: () => import('../views/main/Home.vue'),
                children: [
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
                path: 'cmdty/:id',
                component: () => import('../views/cmdty/Info.vue'),
                // children: [
                //     {
                //         path: 'info',
                //         name: 'info',
                //         component: () => import('../views/cmdty/Info.vue')
                //     }
                // ]
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