import {local} from "../../utils/storage";
import config from "../../config/index.config";
import Menu from '../../config/menu.json';
import func from "../../utils/func";

const state = () => ({
    routesList: local.get("router") ? local.get("router") : [],
    routers: local.get("router") ? local.get("router") : []
})
const getters = {}
const mutations = {
    setRoutesList(state, result) {
        if (config.isLocalMenu == true) {
            state.routesList = Menu
        } else {
            state.routesList = func.toTree(result.data)
        }
        state.router = Menu
        local.set("routers", state.routers, Menu)
        local.set("router", state.routesList, config.loginTimeout)
    },
    setRoutes(state){

    }
}
const actions = {}
export default {state, getters, mutations, actions}
