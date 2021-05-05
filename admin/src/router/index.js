import {createRouter, createWebHistory} from "vue-router";
import routes from './basic_routes.js'
import {permission} from "./permission.js";

const router = createRouter({
    history: createWebHistory(),
    routes,
})
router.beforeEach(permission)
export default router;