import Vue from 'vue'
import App from './App.vue'
import router from './router'
import store from './store'
import VueI18n from 'vue-i18n'; //语言
import {messages} from './lang/i18n'; //语言包
import './common/auth' //路由登录授权拦截和加载进度条
import ElementUI from 'element-ui'; //导入ui
import './assets/theme/index.css'; // 默认主题
import 'element-ui/lib/theme-chalk/display.css'; //响应式隐藏
import './assets/css/style.css' //本地样式
import http from './common/http' //请求封装
import {config} from './common/config' //配置
import func from './common/func.js' //公共函数
import submit from './common/submit.js' //公共提交操作

const i18n = new VueI18n({ //默认中文
    locale: sessionStorage.getItem("language") ? sessionStorage.getItem("language") : config.language,
    messages
});

Vue.use(ElementUI, {
    i18n: (key, value) => i18n.t(key, value),
    size: sessionStorage.getItem("layout_size") ? sessionStorage.getItem("layout_size") : config.layoutSize
});
Vue.config.productionTip = false

Vue.prototype.$Func = func
Vue.prototype.$Submit = submit
Vue.prototype.$http = http
Vue.prototype.$Config = config
new Vue({
    router,
    i18n,
    store,
    render: h => h(App),
    beforeCreate() { //用于非父子组件通信
        Vue.prototype.$bus = this
    }
}).$mount('#app')
