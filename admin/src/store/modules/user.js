import {local} from "../../utils/storage";
import config from "../../config/index.config.js"

const state = () => ({
    token: local.get("token") ? local.get("token") : "",
    userInfo: local.get('user_info') ? local.get("user_info") : "",
})
const getters = {}
const mutations = {
    setUserInfo(state, userInfo) {
        state.userInfo = userInfo
        local.set("user_info", userInfo, config.loginTimeout)
    },
    setToken(state, token) {
        state.token = token
        local.set("token", token, config.loginTimeout)
    }
}
const actions = {
    async login({commit}, row) {
        try {
            const {code, data} = row
            if (code === 200) {
                commit('setUserInfo', data)
                commit('setToken', data["authorization"])
            } else {
            }
        } catch (e) {
            console.log(e)
        }
    },
}
export default {state, getters, mutations, actions}
