import Layout from '@/layout/index.vue'

export default [
  {
    path: '/',
    redirect: "/dashboard/console",
    hidden: true,
  },
  {
    path: '/login',
    component: () => import('@/view/account/login.vue'),
    hidden: true,
  },
  {
    path: '/account',
    name: 'account',
    component: Layout,
    hidden: true,
    children: [{
      path: 'center',
      name: 'account-center',
      meta: { key: 'center', title: '个人中心', icon: 'DashboardOutlined'},
      component: () => import('@/view/account/center.vue')
    }]
  },
  //首页
  {
    path: '/dashboard',
    name: 'dashboard',
    meta: { key: '1', title: '工作空间', icon: 'HomeOutlined' },
    component: Layout,
    children: [
      {
        path: 'console',
        name: 'dashboard-console',
        meta: { key: 'console', title: '分析页', icon: 'DashboardOutlined', fixed: true },
        component: () => import('@/view/dashboard/console.vue'),
      }
    ]
  }, {
    path: '/form',
    name: 'form',
    component: Layout,
    meta: { key: '18', title: '表单页面', icon: 'FormOutlined' },
    children: [
      {
        path: 'advancedForm',
        name: 'advanced-form',
        meta: { key: '21', title: '高级表单', icon: 'DatabaseOutlined' },
        component: () => import('@/view/form/advancedForm.vue')
      }, {
        path: 'baseForm',
        name: 'base-form',
        meta: { key: '19', title: '基础表单', icon: 'DatabaseOutlined' },
        component: () => import('@/view/form/baseForm.vue')
      }, {
        path: 'stepForm',
        name: 'step-form',
        meta: { key: '20', title: '分步表单', icon: 'DatabaseOutlined' },
        component: () => import('@/view/form/stepForm.vue')
      }
    ]
  },
  //列表
  {
    path: '/list',
    name: 'list',
    meta: { key: '6', title: '列表页面', icon: 'UnorderedListOutlined' },
    component: Layout,
    redirect: "/list/baseList",
    children: [
      {
        path: 'baseList',
        name: 'base-list',
        meta: { key: '7', title: '基础列表', icon: 'DatabaseOutlined' },
        component: () => import('@/view/list/baseList.vue'),
      }, {
        path: 'cardList',
        name: 'card-list',
        meta: { key: '12', title: '卡片列表', icon: 'DatabaseOutlined' },
        component: () => import('@/view/list/cardList.vue'),
      }, {
        path: 'newsList',
        name: 'news-list',
        meta: { key: '13', title: '图文列表', icon: 'DatabaseOutlined' },
        component: () => import('@/view/list/newsList.vue'),
      }, {
        path: 'tableList',
        name: 'table-list',
        meta: { key: '14', title: '查询表格', icon: 'DatabaseOutlined' },
        component: () => import('@/view/list/tableList.vue'),
      }
    ]
  }, {
    path: '/result',
    name: 'result-menu',
    meta: { key: '15', title: '结果页面', icon: 'TagOutlined' },
    redirect: "/result/success",
    component: Layout,
    children: [
      {
        path: 'success',
        name: 'result-success',
        meta: { key: '17', title: '成功', icon: 'DatabaseOutlined' },
        component: () => import('@/view/result/success.vue'),
      }, {
        path: 'failure',
        name: 'result-failure',
        meta: { key: '16', title: '失败', icon: 'DatabaseOutlined' },
        component: () => import('@/view/result/failure.vue'),
      }
    ]
  }, {
    path: '/error',
    name: 'error',
    meta: { key: '8', title: '错误页面', icon: 'StopOutlined' },
    component: Layout,
    redirect: "/error/403",
    children: [
      {
        path: '403',
        name: 'error-403',
        meta: { key: '9', title: '403', icon: 'DatabaseOutlined' },
        component: () => import('@/view/error/403.vue'),
      }, {
        path: '404',
        name: 'error-404',
        meta: { key: '10', title: '404', icon: 'DatabaseOutlined' },
        component: () => import('@/view/error/404.vue'),
      }, {
        path: '500',
        name: 'error-500',
        meta: { key: '11', title: '500', icon: 'DatabaseOutlined' },
        component: () => import('@/view/error/500.vue'),
      }
    ]
  }
]
