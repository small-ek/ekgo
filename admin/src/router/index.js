import {createRouter, createWebHistory} from "vue-router";
import routes from './basic_routes.js'
import {permission} from "./permission.js";
import "nprogress/nprogress.css";
import NProgress from "nprogress";

const router = createRouter({
    history: createWebHistory(),
    routes,
})

router.beforeEach(permission)

router.afterEach((to, from) => {
    NProgress.done();
})
export default router;