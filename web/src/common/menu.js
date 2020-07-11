import Menu from '../menu/index.json';
import {config} from './config';
import http from './http.js'
import func from './func.js'
import {local} from './storage.js'

const Layout = () => import("../layout/Main.vue")

//自动加载路由
export default {
    /**
     * 同步获取API接口
     */
    async getApiMenu() {
        var menu = []
        let self = this;
        await http.get('admin/left_menu')
            .then(function (response) {
                var data = response.data.data
                //树形菜单递归
                var Menu = func.toTree(data)
                self.setMenu(Menu)
                menu = Menu
            })
        return menu;
    },
    /**
     * 存储菜单
     * @param {Object} data
     */
    setMenu(data) {
        local.set("menu", data, 60)
    },
    /**
     * 存储菜单
     * @param {Object} data
     */
    getMenu(name) {
        var menu = local.get(name)
        return menu
    },
    //判断是否本地菜单
    menuList() {
        //获取本地缓存菜单
        var menu = this.getMenu('menu');
        if (menu != null && menu != false) {//如果有缓存取缓存
            return menu
        } else if (config.isLocalMenu == true) { //加载本地菜单
            this.setMenu(Menu)
            return Menu;
        } else if (config.isLocalMenu == false) { //加载服务端菜单
            return this.getApiMenu()
        }
    },
    /**
     * 自动加载路由
     */
    async autoRoute() {
        var menu = await this.menuList()
        if (menu) {
            var result = await this.generateRoutes(menu)
            return {
                path: '/admin',
                component: Layout,
                children: result
            };
        }
    },
    /**
     * 解析菜单路由
     */
    generateRoutes(routes, basePath = '/') {
        let res = []

        for (const router of routes) {

            const data = {
                path: router.path,
                component: () => import('../views' + router.path),
                meta: {
                    title: router.title,
                    status: router.status
                }
            }

            if (router.title && router.path && router.path != '/') {

                res.push(data)
            }

            if (router.children) {
                const tempRoutes = this.generateRoutes(router.children, data.path)
                if (tempRoutes.length >= 1) {
                    res = [...res, ...tempRoutes]
                }
            }
        }
        return res
    },
    /**
     * 设置菜单样式
     * @param self
     */
    setMenuTheme(self) {
        let theme = self.$Config.theme
        self.menu_background = theme.menu_background
        self.menu_text_color = theme.menu_text_color
        self.menu_active_text_color = theme.menu_active_text_color
        self.menu_active_background = theme.menu_active_background
    }
}
