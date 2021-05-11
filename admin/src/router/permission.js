import NProgress from "nprogress";
import store from "../store/index"
import router from "../router/index";

const Layout = () => import("../layout/index.vue")
const setTitle = title => {
    // 如有i18n在这里修改
    document.title = `ek-admin-${title}`
}
/**
 * 为所有路由添加component视图
 * @param routes
 */
export const setRouteComponent = (routes, basePath = '/') => {
    let res = []
    for (const item of routes) {
        const data = {
            path: item.path,
            component: () => import('@/views' + item.path + ".vue"),
            meta: {
                title: item.title,
                status: item.status
            }
        }

        if (item.title && item.path && item.path != '/') {
            res.push(data)
        }

        if (item.children) {
            const tempRoutes = setRouteComponent(item.children, data.path)
            if (tempRoutes.length >= 1) {
                res = [...res, ...tempRoutes]
            }
        }
    }

    return res
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
            console.log('========== 开始加载用户权限菜单 ==========')
            const menu = setRouteComponent(store.state.routes.menu)
            console.log(menu)

            router.addRoute({ path: '/user/index', component: () => import('../views/user/index.vue') })
            console.log(router.getRoutes())
        }
        next()
    }
}