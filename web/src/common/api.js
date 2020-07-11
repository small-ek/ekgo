import http from './http.js'

/**
 * 获取菜单
 * @param data
 * @returns {Promise<AxiosResponse<any>>}
 */
export const getMenu = data => {
    return http.request({
        method: "GET",
        url: 'admin/menu',
        data: data
    })
};
/**
 * 获取角色
 * @param data
 * @returns {Promise<AxiosResponse<any>>}
 */
export const getApi = data => {
    return http.request({
        method: "GET",
        url: 'admin/api',
        params: {page_size: 1000}
    })
};
/**
 * 获取地理省市区
 * @param pid
 * @returns {Promise<AxiosResponse<any>>}
 */
export const getGeo = pid => {
    return http.request({
        method: "GET",
        url: 'geo/' + pid,
    })
};
/**
 * 获取所有角色
 * @param pid
 * @returns {Promise<AxiosResponse<any>>}
 */
export const getRole = data => {
    return http.request({
        method: "GET",
        url: 'admin/role',
        params: {page_size: 1000}
    })
};
/**
 * 获取当前系统所有表
 * @param pid
 * @returns {Promise<AxiosResponse<any>>}
 */
export const getAllTable = data => {
    return http.request({
        method: "GET",
        url: 'admin/get_all_table',
        params: data
    })
};
/**
 * 获取当前表字段信息
 * @param name
 * @returns {Promise<AxiosResponse<T>>}
 */
export const getTable = name => {
    return http.request({
        method: "GET",
        url: 'admin/get_table/' + name,
    })
};
/**
 * 检测登录
 * @param data
 * @returns {Promise<AxiosResponse<T>>}
 */
export const checkLogin = data => {
    return http.request({
        method: "GET",
        url: 'admin/check_login',
    })
};