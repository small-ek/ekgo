/**
 * 静态路由
 */
import Layout from "@/layout";

//alwaysShow 需要在meta下也配置一个
export default [
  {
    path: "/dashboard",
    name: "dashboard",
    meta: {
      key: "1",
      title: "工作空间",
      icon: "HomeOutlined",
      alwaysShow: true
    },
    alwaysShow: true,
    component: Layout,
    children: [
      {
        path: "console",
        name: "dashboard-console",
        meta: {
          key: "console",
          title: "分析页",
          icon: "DashboardOutlined",
          fixed: true
        },
        component: () => import("@/view/dashboard/console.vue")
      }
    ]
  },
  {
    path: "/form",
    name: "form",
    component: Layout,
    meta: { key: "18", title: "表单页面", icon: "FormOutlined" },
    children: [
      {
        path: "baseForm",
        name: "base-form",
        meta: { key: "19", title: "基础表单", icon: "DatabaseOutlined" },
        component: () => import("@/view/form/baseForm.vue")
      },
      {
        path: "stepForm",
        name: "step-form",
        meta: { key: "20", title: "分步表单", icon: "DatabaseOutlined" },
        component: () => import("@/view/form/stepForm.vue")
      }
    ]
  },
  // 列表
  {
    path: "/list",
    name: "list",
    meta: { key: "6", title: "列表页面", icon: "UnorderedListOutlined" },
    component: Layout,
    redirect: "/list/baseList",
    children: [
      {
        path: "baseList",
        name: "base-list",
        meta: { key: "7", title: "基础列表", icon: "DatabaseOutlined" },
        component: () => import("@/view/list/baseList.vue")
      },
      {
        path: "cardList",
        name: "card-list",
        meta: { key: "12", title: "卡片列表", icon: "DatabaseOutlined" },
        component: () => import("@/view/list/cardList.vue")
      },
      {
        path: "newsList",
        name: "news-list",
        meta: { key: "13", title: "图文列表", icon: "DatabaseOutlined" },
        component: () => import("@/view/list/newsList.vue")
      },
      {
        path: "tableList",
        name: "table-list",
        meta: { key: "14", title: "查询表格", icon: "DatabaseOutlined" },
        component: () => import("@/view/list/tableList.vue")
      }
    ]
  },
  {
    path: "/result",
    name: "result-menu",
    meta: { key: "15", title: "结果页面", icon: "TagOutlined" },
    redirect: "/result/success",
    component: Layout,
    children: [
      {
        path: "success",
        name: "result-success",
        meta: { key: "17", title: "成功", icon: "DatabaseOutlined" },
        component: () => import("@/view/result/success.vue")
      },
      {
        path: "failure",
        name: "result-failure",
        meta: { key: "16", title: "失败", icon: "DatabaseOutlined" },
        component: () => import("@/view/result/failure.vue")
      }
    ]
  }
];
