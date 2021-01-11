import {createRouter, createWebHistory} from 'vue-router'
import TabsView from '@/layouts/tabs/TabsView'

const routes = [
    {
        path: '/404',
        name: '404',
        component: () => import('@/views/exception/404'),
    },
    {
        path: '/403',
        name: '403',
        component: () => import('@/views/exception/403'),
    },
    {
        path: '/',
        name: 'Home',
        component: TabsView,
        redirect: '/login',
        children: [
            {
                path: 'demo',
                name: '演示页',
                meta: {
                    icon: 'file-ppt'
                },
                component: () => import('@/views/demo/index')
            }]
    },
    {
        path: '/about',
        name: 'About',
        // route level code-splitting
        // this generates a separate chunk (about.[hash].js) for this route
        // which is lazy-loaded when the route is visited.
        component: () => import(/* webpackChunkName: "about" */ '../views/About.vue')
    },
    {
        path: '/login',
        name: 'login',
        component: () => import(/* webpackChunkName: "about" */ '../views/login/index.vue')
    }
]

const router = createRouter({
    history: createWebHistory(process.env.BASE_URL),
    routes
})

export default router
