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
    for (let i = 0; i < routes.length; i++) {
        var value = routes[i]
        const data = {
            path: value.path,
            component: () => import('../views/user/index.vue'),
            meta: {
                title: value.title,
                status: value.status
            }
        }

        if (value.title && value.path && value.parent_id == 0) {
            res.push(data)
        }else{

        }

        if (value.children) {
            const tempRoutes = setRouteComponent(value.children, data.path)
            console.log(tempRoutes)
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

        console.log(to.fullPath)
        console.log(router.getRoutes().map(it => it.path).includes(to.fullPath))

        if (!router.getRoutes().map(it => it.path).includes(to.fullPath)) {
            console.log('========== 开始加载用户权限菜单 ==========')
            var menus = store.state.routes.menu
            console.log(menus)
            const menu = setRouteComponent(menus)
            console.log(menu)
            router.addRoute({
                path: '/',
                name: "用户",
                meta: {title: '登录',},
                component: Layout,
                children: menu
            })
            console.log(router.getRoutes())
        }
        next()
    }
}