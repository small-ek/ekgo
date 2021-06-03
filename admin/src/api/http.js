import axios from "axios";
import config from "../config/index.config"
import {message} from 'ant-design-vue';
import store from "../store";
//post请求头
var headers = {}

if (store.state.user.token != "") {
    headers[config.tokenName] = store.state.user.token
}

var http = axios.create({
    baseURL: config.url,
    timeout: config.timeout,
    headers: headers
});

//添加请求拦截器
http.interceptors.request.use(
    config => {
        return config;
    }, error => {
        console.log(error);
        console.log("请求错误")
        return Promise.reject(error);
    }
);

//添加响应拦截器
http.interceptors.response.use(
    response => {
        if (response.status && response.data.code === config.invalidCode) {
            message.error('网络请求失败，请刷新重试');
            //window.location.href = '/login'
        }
        return response;
    }, error => {
        console.log(error)
        console.log("拦截错误")
        return Promise.reject(error);
    }
);
export default {
    setHeaders(key, value) {
        headers[key] = value
        http = axios.create({
            baseURL: config.url,
            timeout: config.timeout,
            headers: headers
        })
    },
    get(url, data) {
        return new Promise((resolve, reject) => {
            http.request({
                method: 'GET',
                url,
                params: data,
            }).then(res => {
                resolve(res.data)
            }).catch(err => {
                reject(err)
            })
        })
    },
    post(url, data) {
        return new Promise((resolve, reject) => {
            http.request({
                method: 'POST',
                url,
                data: data,
            }).then(res => {
                resolve(res.data)
            }).catch(err => {

                reject(err)
            });
        })
    },
    put(url, data) {
        return new Promise((resolve, reject) => {
            http.request({
                method: 'PUT',
                url,
                data: data,
            }).then(res => {
                resolve(res.data)
            }).catch(err => {
                reject(err)
            });
        })
    },
    patch(url, data) {
        return new Promise((resolve, reject) => {
            http.request({
                method: 'PATCH',
                url,
                data: data,
            }).then(res => {
                resolve(res.data)
            }).catch(err => {
                reject(err)
            });
        })
    },
    deletes(url, data) {
        return new Promise((resolve, reject) => {
            http.request({
                method: 'DELETE',
                url,
                data: data,
            }).then(res => {
                resolve(res.data)
            }).catch(err => {
                reject(err)
            });
        })
    },
    axios: http
};