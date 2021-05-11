import {local} from "../../utils/storage";
import config from "../../config/index.config";
import Menu from '../../config/menu.json';
import func from "../../utils/func";

const state = () => ({
    router: local.get("router") ? local.get("router") : [],
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
        state.router = result
        local.set("router", result, config.loginTimeout)
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
            commit('setMenu', row)
            commit('setRouters', row)
        } catch (e) {
            console.log(e)
        }
    },
}
export default {state, getters, mutations, actions}
