import NProgress from "nprogress";
import store from "../store/index"

const setTitle = title => {
    // 如有i18n在这里修改
    document.title = `ek-admin-${title}`
}
/**
 * 权限控制
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
        next()
    }
}