//路由拦截
import router from '../router/index.js'
import NProgress from 'nprogress' // Progress 进度条
import {local} from './storage.js'

//拦截器
router.beforeEach((to, from, next) => {
    //滚动条加载
    NProgress.start()
    //判断是否登录
    const user = local.get('user')
    if ((user == null || user == false) && to.meta.auth == undefined) {
        localStorage.clear()
        next({
            path: '/login',
            query: {
                redirect: to.fullPath
            }
        })
    }

    if (to.content) {
        let meta = document.createElement('meta');
        meta.content = to.content;
        document.getElementsByTagName('head')[0].appendChild(meta);
    }
    //路由发生变化修改页面title
    if (to.meta.title) {
        document.title = to.meta.title;
    }

    next()
});

router.afterEach(() => {
    NProgress.done() // 结束进度条
})
