import Vue from 'vue'
import Vuex from 'vuex'

Vue.use(Vuex)

export default new Vuex.Store({
	size: 'default', //字体大小
	menu: [], //菜单
	Breadcrumb: [], //面包屑
	user:{},//用户信息
	state: {},
	mutations: {},
	actions: {},
	modules: {}
})
