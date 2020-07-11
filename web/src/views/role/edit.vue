<template>
    <div class="main">
        <el-card class="box-card form" shadow="hover">
            <el-form :model="form" :rules="rules" ref="form" :label-position="this.$Func.formPosition()" label-width="110px">
                <el-form-item label="角色名称" prop="name">
                    <el-input v-model="form.name"></el-input>
                </el-form-item>

                <el-form-item label="菜单权限">
                    <el-tree :data="menu" show-checkbox node-key="id" ref="tree" :default-expand-all="true" :default-checked-keys="form.menu_id"
                             :props="defaultProps">
                    </el-tree>
                </el-form-item>

                <el-form-item label="api权限">
                    <el-tree :data="api" show-checkbox node-key="id" ref="tree2" :default-expand-all="true" :default-checked-keys="form.api_id"
                             :props="defaultProps">
                    </el-tree>
                </el-form-item>

                <el-form-item>
                    <el-button type="primary" size="medium" icon="el-icon-edit" @click="onSubmit('form')" :loading="loading">{{ loading ? $t('wait') : $t('submit') }}
                    </el-button>
                    <el-button type="primary" plain size="medium" icon="el-icon-back" @click="$router.back(-1)">返回
                    </el-button>
                </el-form-item>
            </el-form>
        </el-card>
    </div>
</template>

<script>
    import {getApi, getMenu} from "../../common/api";

    export default {
        data() {
            return {
                url: "admin/role",
                form: {
                    menu_id: [],
                    api_id: []
                },
                menu_all: [],
                menu: [],
                api: [],
                method: "POST",
                loading: false,
                defaultProps: {
                    children: 'children',
                    label: 'title'
                },
                //校验规则
                rules: {
                    name: [{
                        required: true,
                        message: '不能为空',
                        trigger: 'blur'
                    }]
                }
            }
        },
        created: function () {
            this.get()
        },
        methods: {
            onSubmit(formName) {
                const self = this;
                const form = self.form
                const url = this.method == "POST" ? this.url : this.url + "/" + form.id
                var menu = self.$refs.tree.getCheckedKeys()

                var menu_all = self.menu_all
                //获取父级
                var parent_menu_id = []

                menu.forEach(function (item) {
                    menu_all.forEach(function (item2) {
                        if (item2['id'] == item && parent_menu_id.includes(item2.parent_id) == false) {
                            parent_menu_id.push(item2['parent_id'])
                        }
                    })
                })
                parent_menu_id = self.$Func.deleteArray(parent_menu_id, 0)

                //获取api
                var api = self.$refs.tree2.getCheckedKeys()
                api = api.filter(function (item) {
                    return item != null;
                });

                form.menu_id = menu
                form.api_id = api
                form.parent_menu_id = parent_menu_id
                this.$Submit.submitBack(this, formName, url, form, this.method)
            },
            get() {
                const self = this
                //获取菜单
                getMenu({id: 1}).then(function (response) {
                    self.menu_all = response.data.data
                    self.menu = self.$Func.toTree(response.data.data)
                })
                //获取所有api
                getApi().then(function (response) {
                    self.api = self.$Func.groupBy(response.data.data.list, 'group')
                })

                var id = this.$route.query.id

                if (id) {
                    this.method = 'PUT'
                    this.$http.get('admin/role/' + id).then(function (response) {
                        const result = response.data.data
                        self.form = result
                    })
                }
            }
        }
    }
</script>
