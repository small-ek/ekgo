import Menu from '../menu/index.json';
import {config} from './config';

const Layout = () => import("../layout/Main.vue")

//自动加载路由
export default {
	autoRoute() {
		if (config.isLocalMenu == true) {
			var result = this.generateRoutes(Menu)
			return {
				path: '/admin',
				component: Layout,
				children: result
			};
		}
	},
	generateRoutes(routes, basePath = '/') {
		let res = []

		for (const router of routes) {
			const data = {
				path: router.path,
				component: () => import('../views' + router.path),
				meta: {
					title: router.title,
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
}
