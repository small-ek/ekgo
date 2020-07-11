<template>
    <div class="main">
        <el-row :gutter="10">
            <el-col :xs="24" :sm="24" :md="12" :lg="10" :xl="8">
                <el-card class="box-card" shadow="hover" style="min-height: calc(88vh)">
                    <div class="form" style="border: 1px solid #e6ebf5;padding: 10px">
                        <el-form :model="form" :rules="rules" ref="form" label-width="25%">
                            <el-form-item label="上级菜单">
                                <el-cascader v-model="form.parent_id" :options="parent_menu" :props="{ checkStrictly: true }" clearable></el-cascader>
                            </el-form-item>
                            <el-form-item label="菜单名称" prop="title">
                                <el-input v-model="form.title"></el-input>
                            </el-form-item>
                            <el-form-item label="菜单路径" prop="route">
                                <el-input v-model="form.path"></el-input>
                            </el-form-item>
                            <el-form-item label="排序" prop="sort">
                                <el-input v-model="form.sort"></el-input>
                            </el-form-item>
                            <el-form-item label="图标" prop="icon">
                                <el-input v-model="form.icon"></el-input>
                                <el-button type="text" @click="outerVisible = true">选择图标</el-button>
                                <el-dialog title="图标" :visible.sync="outerVisible">
                                    <Icon @return="getIcon"></Icon>
                                </el-dialog>
                            </el-form-item>
                            <el-form-item label="状态" prop="status">
                                <template v-for="(value) in status">
                                    <el-radio v-model="form.status" :label="value.id" size="small">
                                        {{value.name}}
                                    </el-radio>
                                </template>
                            </el-form-item>
                            <el-form-item>
                                <el-button type="primary" size="small" icon="el-icon-edit" @click="onSubmit('form','admin/menu')">
                                    提交
                                </el-button>
                                <el-button type="primary" size="small" icon="el-icon-loading" @click="onCancel()">
                                    取消
                                </el-button>
                            </el-form-item>
                        </el-form>
                    </div>
                </el-card>
            </el-col>

            <el-col :xs="24" :sm="24" :md="12" :lg="14" :xl="16">
                <el-card class="box-card" shadow="hover" style="min-height: calc(88vh)">
                    <el-table :data="list" ref="multipleTable" row-key="id" default-expand-all :tree-props="{children: 'children', hasChildren: 'hasChildren'}" stripe>
                        <el-table-column prop="icon" width="70" label="图标">
                            <template slot-scope="scope">
                                <i :class="scope.row.icon"></i>
                            </template>
                        </el-table-column>
                        <el-table-column prop="title" width="150" label="菜单名称" show-overflow-tooltip></el-table-column>
                        <el-table-column prop="path" label="菜单路径" show-overflow-tooltip></el-table-column>
                        <el-table-column prop="sort" width="60" label="排序" show-overflow-tooltip></el-table-column>
                        <el-table-column prop="status" width="70" label="状态">
                            <template slot-scope="scope">
                                <el-tag v-if="scope.row.status=='true'">显示</el-tag>
                                <el-tag type="danger" v-if="scope.row.status=='false'">隐藏</el-tag>
                            </template>
                        </el-table-column>
                        <el-table-column label="操作" width="150">
                            <template slot-scope="scope">
                                <el-button type="text" @click="onEdit(scope.row)">{{$t('edit')}}</el-button>
                                <el-divider direction="vertical"></el-divider>
                                <el-button type="text" class="red" @click="onDelete(scope.row,scope.$index)">
                                    {{$t('del')}}
                                </el-button>
                            </template>
                        </el-table-column>
                    </el-table>
                </el-card>
            </el-col>
        </el-row>
    </div>
</template>
<script>
    import Icon from '../../components/Icon/Index.vue'

    export default {
        name: 'menu',
        components: {Icon},
        data() {
            return {
                list: [],
                parent_menu: [],
                method: "POST",
                outerVisible: false,
                form: {},
                status: [{
                    id: 'true',
                    name: '显示'
                }, {
                    id: 'false',
                    name: '隐藏'
                }],
                //校验规则
                rules: {
                    title: [{
                        required: true,
                        message: '名称不能为空',
                        trigger: 'blur'
                    }],
                    status: [{
                        required: true,
                        message: '状态不能为空',
                        trigger: 'blur'
                    }]
                }
            }
        },
        created() {
            //获取数据
            this.getList();
        },
        methods: {
            onCancel() {
                this.form = {}
                this.method = "POST"
            },
            onEdit(data) {
                this.method = "PUT"
                this.form = data
            },
            getIcon(data) {
                this.form.icon = data
                this.outerVisible = false
            },
            onDelete(row, index) {
                const self = this;
                if (row.children) {
                    self.$message({
                        type: 'error',
                        message: "请删除子项后再删除"
                    });
                    return false
                }
                this.$Submit.deleteSubmit(this, 'admin/menu/' + row.id, index, function (response) {
                    const result = response.data
                    if (result.code == 200) {
                        self.getList();
                        self.$message({
                            type: 'success',
                            message: self.$i18n.t('del_success')
                        });
                    }
                })
            },
            onSubmit(formName, url) {
                let self = this
                let form = JSON.parse(JSON.stringify(this.form))
                form.sort = Number(this.form.sort)
                if (typeof (this.form.parent_id) == "object") {
                    form.parent_id = form.parent_id.pop() || 0
                }
                url = this.method == "POST" ? url : url + "/" + form.id

                this.$Submit.submitClose(this, formName, url, form, this.method, function (response) {
                    const result = response.data
                    self.$message({
                        message: result.msg,
                        type: result.code == 200 ? 'success' : 'error'
                    })
                    if (result.code == 200) {
                        self.form = {}
                        self.method = "POST"
                        self.getList()
                    }
                })
            },
            getList: function () {
                const self = this
                this.$http.get('admin/menu', {
                    params: this.where
                }).then(function (response) {
                    var result = response.data.data;
                    var parent_data = []
                    //父级菜单
                    result.forEach(function (item) {
                        parent_data.push({
                            value: item.id,
                            label: item.title,
                            parent_id: item.parent_id
                        })
                    })

                    self.parent_menu = self.$Func.toTree(parent_data, 'value')
                    self.list = self.$Func.toTree(result)
                })
            },
        }
    }
</script>
