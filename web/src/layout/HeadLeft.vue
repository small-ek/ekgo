<template>
    <div>
        <el-breadcrumb separator=">" class="header-btn">
            <template v-for="(item,index) in Breadcrumb">
                <el-breadcrumb-item :key="index">{{item}}</el-breadcrumb-item>
            </template>
        </el-breadcrumb>
    </div>
</template>

<script>
    import getMenu from '../common/menu.js';

    export default {
        data() {
            return {
                search: '',
                menu: [],
                menuOptions: [],
                Breadcrumb: [],
                list: [],
            };
        },
        components: {},
        created() {
            const self = this;
            //设置下拉框菜单
            this.getMenu().then(function (result) {
                self.menuOptions = self.generateRoutes(result);
                //获取面包屑
                self.getBreadcrumb()
            })

        },
        computed: {},
        watch: {
            $route(route) {
                // if you go to the redirect page, do not update the breadcrumbs
                if (route.path.startsWith('/redirect/')) {
                    return
                }
                this.getBreadcrumb()
            }
        },
        methods: {
            //获取面包屑
            getBreadcrumb() {
                var Breadcrumb = this.menuOptions.find(item => item.path == this.$route.fullPath)
                if (Breadcrumb) {
                    this.Breadcrumb = Breadcrumb.title
                    // this.$store.Breadcrumb = Breadcrumb.title
                } else if (this.$store.Breadcrumb) {
                    this.Breadcrumb = this.$store.Breadcrumb
                } else if (this.$route.meta.title) {
                    this.Breadcrumb = [this.$route.meta.title]
                } else {
                    this.Breadcrumb = ['首页']
                }
            },
            async getMenu() {
                return await getMenu.menuList()
            },
            //过滤菜单合并国际化
            generateRoutes(routes, basePath = '/', prefixTitle = []) {
                let res = []

                for (const router of routes) {

                    const data = {
                        path: router.path,
                        title: [...prefixTitle]
                    }
                    if (router.title) {
                        data.title = [...data.title, router.title]
                        if (router.path && router.status == 'true') {
                            res.push(data)
                        }
                    }

                    if (router.children) {
                        const tempRoutes = this.generateRoutes(router.children, data.path, data.title)
                        if (tempRoutes.length >= 1) {
                            res = [...res, ...tempRoutes]
                        }
                    }
                }
                return res
            },

        }
    }
</script>

