import {local} from "../../utils/storage";
import config from "../../config/index.config.js"

const state = () => ({
    token: local.get("token") ? local.get("token") : "",
    userInfo: local.get('user_info') ? local.get("user_info") : "",
})
const getters = {}
const mutations = {
    setUserInfo(state, userInfo) {
        var row = JSON.parse(JSON.stringify(userInfo))
        if (row[config.tokenName]) {
            delete row[config.tokenName]
        }
        state.userInfo = row
        local.set("user_info", row, config.loginTimeout)
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
                commit('setToken', data[config.tokenName])
                commit('setUserInfo', data)
            } else {
            }
        } catch (e) {
            console.log(e)
        }
    },
}
export default {state, getters, mutations, actions}
