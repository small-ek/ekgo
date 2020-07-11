import Vue from 'vue'
import VueRouter from 'vue-router'
import route from '../common/menu';
import {config} from '../common/config'

Vue.use(VueRouter)

const Layout = () => import('../layout/Main.vue')


const routes = [


    {
        path: '/login',
        component: () => import('../views/login/index'),
        meta: {
            title: "Login",
            auth: false
        }
    },
    {
        path: '/redirect',
        component: Layout,
        children: [{
            path: '/redirect/:path*',
            auth: false,
            component: () => import('../views/redirect/index')
        }]
    },
    {
        path: '/',
        component: Layout,
        children: [{
            path: '',
            component: () => import('../views/home/index'),
            meta: {
                title: '首页',
            }
        }, {
            path: 'index',
            component: () => import('../views/home/index'),
            meta: {
                title: '首页',
            }
        }]
    },
    /* {
         path: '/401',
         component: () => import('../views/error/401'),
         meta: {
             title: "Login",
             auth: false
         }
     },*/
    /* {
        path: '*',
        component: () => import('../views/error/404'),
        meta: {
            title: "Login",
            auth: false
        }
    }, */
]


const router = new VueRouter({
    mode: 'history',
    base: process.env.BASE_URL,
    routes
})

//根据菜单自动加载路由
var path = config.routesWhiteList;
var fullPath = window.location.pathname;

if (path.includes(fullPath) == false) {
    route.autoRoute().then(function (result) {

        router.addRoutes([result])
    })
}

export default router
