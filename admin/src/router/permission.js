import NProgress from "nprogress";
import store from "../store/index"
import router from "../router/index";

const Layout = () => import("../layout/index.vue")
const setTitle = title => {
    // 如有i18n在这里修改
    document.title = `ek-admin-${title}`
}
let res = []
/**
 * 为所有路由添加component视图
 * @param routes
 */
export const setRouteComponent = (routes) => {
    console.log(routes)
    routes.forEach(row => {

        if (row.children && row.children.length > 0) {
            row.component = Layout
            setRouteComponent(row.children)
        } else {
            row.component = () => import('../views/' + row.path + '.vue')
        }
    })
}
/**
 * 路由拦截器
 * @param to
 * @param from
 * @param next
 * @returns {Promise<void>}
 */
export const permission = async (to, from, next) => {
    //配置路由加载动画效果
    NProgress.start();
    const {meta: {title = ''}} = to
    setTitle(title);
    if (!to.fullPath.includes('/login') && !store.state.user.token) {
        next({path: '/login'})
    } else {

        if (!router.getRoutes().map(it => it.path).includes(to.fullPath)) {
            var menus = store.state.routes.menu
            setRouteComponent(menus)
            for (let i = 0; i < menus.length; i++) {
                var value = menus[i]
                router.addRoute(value)
            }
            next(to.fullPath)
        }
        next()
    }
}