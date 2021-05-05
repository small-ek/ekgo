import Layout from "../layout/index.vue";

export default [
    {
        path: '/login',
        name: '登录',
        meta: {
            title: '登录',
        },
        component: () => import('../views/login/index.vue'),
    },
    {
        path: "/",
        name: "首页",
        component: Layout,
        children: [
            {
                path: '',
                component: () => import('../views/index/index.vue'),
                meta: {
                    title: 'Home',
                }
            }
        ],
    },
    {
        path: "/home",
        name: "Home",
        component: () => import("../views/home.vue"),
    },
]
