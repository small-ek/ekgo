import {createRouter, createWebHistory} from "vue-router";
import Layout from '../layout/index.vue'

const routes = [
  {
    path: '/login',
    name: '登录',
    meta: {},
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
];

export default createRouter({
  history: createWebHistory(),
  routes,
});