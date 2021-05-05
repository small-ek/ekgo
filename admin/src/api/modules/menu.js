import http from '../http'

const api = {
    url: "admin/menus",
    left_menu: "admin/left_menu"
}

export function getUserMenu(token) {
    var url = api.left_menu
    return http.axios.request({
        method: 'GET',
        url,
        headers: {"authorization": token},
    })
}

export default {
    getUserMenu
}
