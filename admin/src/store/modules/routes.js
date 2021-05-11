import {local} from "../../utils/storage";
import config from "../../config/index.config";
import Menu from '../../config/menu.json';
import func from "../../utils/func";

const state = () => ({
    routers: local.get("routers") ? local.get("routers") : [],
    menu: local.get("menu") ? local.get("menu") : []
})
const getters = {}
const mutations = {
    /**
     * 设置菜单
     * @param state
     * @param result
     */
    setMenu(state, result) {
        var list = []
        if (config.isLocalMenu == true) {
            list = Menu
        } else {
            list = func.toTree(result)
        }
        state.menu = list
        local.set("menu", list, config.loginTimeout)
    },
    /**
     * 设置路由列表
     * @param state
     */
    setRouters(state, result) {
        state.routers = result
        local.set("routers", result, config.loginTimeout)
    }
}
const actions = {
    /**
     * 设置路由
     * @param commit
     * @param row
     * @returns {Promise<void>}
     */
    async setRouter({commit}, row) {
        try {
            //这里可能浏览器兼容性问题，不支持localStorage
            commit('setMenu', row)
            commit('setRouters', row)
        } catch (e) {
            console.error(e)
        }
    },
}
export default {state, getters, mutations, actions}
