<template>
    <div class="main">
        <el-card class="box-card" shadow="hover">
            <!-- 搜索表单 -->
            <el-form label-position="top" :model="where" v-show="search.status">
                <el-row :gutter="10" class="transition-box">
                    <el-col :xs="24" :sm="12" :md="8" :lg="6" :xl="4">
                        <el-form-item label="用户名">
                            <el-input placeholder="请输入" v-model="where.filter[0][2]" clearable></el-input>
                        </el-form-item>
                    </el-col>
                    <el-col :xs="24" :sm="12" :md="8" :lg="6" :xl="4">
                        <el-form-item label="姓名">
                            <el-input placeholder="请输入" v-model="where.filter[1][2]" clearable></el-input>
                        </el-form-item>
                    </el-col>
                    <el-col :xs="24" :sm="12" :md="8" :lg="6" :xl="4">
                        <el-form-item label="状态">
                            <el-select class="widen" v-model="where.filter[2][2]" placeholder="请选择" clearable>
                                <el-option v-for="item in status" :key="item.value" :label="item.label" :value="item.value">
                                </el-option>
                            </el-select>
                        </el-form-item>
                    </el-col>
                    <el-col :xs="24" :sm="12" :md="8" :lg="6" :xl="4">
                        <el-form-item label="类型">
                            <el-select class="widen" v-model="where.filter[3][2]" placeholder="请选择" clearable>
                                <el-option v-for="item in super_list" :key="item.value" :label="item.label" :value="item.value">
                                </el-option>
                            </el-select>
                        </el-form-item>
                    </el-col>
                </el-row>
            </el-form>
            <!--按钮-->
            <el-row :gutter="10" class="button-row">
                <el-button @click="onEdit()" icon="el-icon-plus" size="small" type="primary">{{$t('add')}}</el-button>
                <el-button @click="onSearch()" icon="el-icon-search" v-loading.fullscreen.lock="fullscreen_loading" size="small" type="success">{{$t('search')}}</el-button>
                <el-button @click="$Func.clearSearch(where)" icon="el-icon-close" size="small">{{$t('clear_search')}}</el-button>
                <el-button @click="$Func.searchSwitch(search)" icon="el-icon-sort" size="small">{{search.status==true?$t('collapse'):$t('expand')}}</el-button>
            </el-row>
            <!-- 表格 -->
            <div class="table-box">
                <el-table :data="list" stripe>
                    <el-table-column prop="id" label="标识" width="85" show-overflow-tooltip>
                    </el-table-column>
                    <el-table-column prop="username" label="用户名" width="150" show-overflow-tooltip>
                    </el-table-column>
                    <el-table-column prop="name" label="姓名" width="150" show-overflow-tooltip>
                    </el-table-column>
                    <el-table-column prop="status" label="类型" show-overflow-tooltip>
                        <template slot-scope="scope">
                            <el-tag size="mini" v-if="scope.row.super=='true'">超级管理员</el-tag>
                            <el-tag size="mini" v-if="scope.row.super=='false'">普通用户</el-tag>
                        </template>
                    </el-table-column>
                    <el-table-column prop="status" label="状态" show-overflow-tooltip>
                        <template slot-scope="scope">
                            <el-tag size="mini" v-if="scope.row.status=='true'">启用</el-tag>
                            <el-tag size="mini" type="danger" v-if="scope.row.status=='false'">禁用</el-tag>
                        </template>
                    </el-table-column>
                    <el-table-column prop="created_at" label="创建时间" show-overflow-tooltip>
                    </el-table-column>
                    <el-table-column prop="Setting" label="操作" width="250">
                        <template slot-scope="item">
                            <el-button type="text" @click="onEdit(item.row)">
                                {{$t('edit')}}
                            </el-button>
                            <el-divider direction="vertical"></el-divider>
                            <el-button type="text" @click="onUpdatePassword(item.row)">
                                {{$t('change_password')}}
                            </el-button>
                            <el-divider direction="vertical"></el-divider>
                            <el-button type="text" class="red" @click="onDelete(item.row,item.$index)">
                                {{$t('del')}}
                            </el-button>
                        </template>
                    </el-table-column>
                </el-table>
            </div>
            <!-- 分页 -->
            <Page url="/admin/admin" :params="where" ref="page" @return="getPage">
            </Page>
        </el-card>

        <!--弹出-->
        <template>
            <!-- 表单 -->
            <el-dialog :title="method=='POST'?'添加管理员':'编辑管理员'" :fullscreen="this.$Func.isFullscreen()" width="50%" lock-scroll :visible.sync="dialog_status">
                <el-form :model="form" :rules="rules" ref="form" :label-position="this.$Func.formPosition()" label-width="110px">
                    <el-form-item label="用户名" prop="username">
                        <el-input v-model="form.username"></el-input>
                    </el-form-item>
                    <el-form-item label="密码" v-show="method=='POST'?true:false" prop="password">
                        <el-input type="password" v-model="form.password" autocomplete="off" show-password></el-input>
                    </el-form-item>
                    <el-form-item label="姓名" prop="name">
                        <el-input v-model="form.name"></el-input>
                    </el-form-item>
                    <el-form-item label="类型" prop="super">
                        <el-radio-group v-model="form.super">
                            <el-radio v-for="item in super_list" :key="item.value" :label="item.value">{{item.label}}</el-radio>
                        </el-radio-group>
                    </el-form-item>
                    <el-form-item label="所属角色" prop="admin_name" v-if="form.super=='false'">
                        <el-select v-model="form.role_id" placeholder="请选择">
                            <el-option v-for="item in role" :key="item.id" :label="item.name" :value="item.id">
                            </el-option>
                        </el-select>
                    </el-form-item>
                    <el-form-item label="状态" prop="status">
                        <el-radio-group v-model="form.status">
                            <el-radio v-for="item in status" :key="item.value" :label="item.value">{{item.label}}</el-radio>
                        </el-radio-group>
                    </el-form-item>
                </el-form>
                <div slot="footer">
                    <el-button @click="dialog_status = false">{{$t('cancel')}}</el-button>
                    <el-button type="primary" @click="onSubmit('form')" :loading="loading">{{ loading ? $t('wait') : $t('submit') }}</el-button>
                </div>
            </el-dialog>
            <!-- 重置密码 -->
            <el-dialog title="修改密码" :fullscreen="this.$Func.isFullscreen()" width="55%" lock-scroll :visible.sync="dialog_status2">
                <el-form :model="form" :rules="rules" ref="form" :label-position="this.$Func.formPosition()" label-width="120px">
                    <el-form-item label="用户名" prop="username">
                        <el-input v-model="form.username" :disabled="true"></el-input>
                    </el-form-item>
                    <el-form-item label="密码" prop="password">
                        <el-input type="password" v-model="form.password" autocomplete="off" show-password></el-input>
                    </el-form-item>
                </el-form>
                <div slot="footer">
                    <el-button @click="dialog_status2 = false">{{$t('cancel')}}</el-button>
                    <el-button type="primary" @click="onSubmit('form','admin/admin')" :loading="loading">{{ loading ? $t('wait') : $t('submit') }}</el-button>
                </div>
            </el-dialog>
        </template>
    </div>
</template>

<script>
    import Page from '../../components/Page/Index.vue'
    import {getRole} from "../../common/api";

    export default {
        components: {Page},
        data() {
            return {
                url: "admin/admin",
                method: "POST",
                fullscreen_loading: false,
                loading: false,
                search_status: true,
                dialog_status: false,
                dialog_status2: false,
                search: {
                    status: true
                },
                where: {
                    filter: [
                        ["username", "like", ""],
                        ["name", "like", ""],
                        ["status", "=", ""],
                        ["super", "=", ""],
                    ]
                },
                form: {
                    role_id: '',
                    super: ''
                },
                status: [
                    {"label": "启用", "value": "true"},
                    {"label": "禁用", "value": "false"},
                ],
                super_list: [
                    {"label": "超级管理员", "value": "true"},
                    {"label": "角色管理员", "value": "false"},
                ],
                list: [],
                //校验
                rules: {
                    username: [{
                        required: true,
                        message: '用户名不能为空',
                        trigger: 'blur'
                    }],
                    name: [{
                        required: true,
                        message: '姓名不能为空',
                        trigger: 'blur'
                    }],
                    status: [{
                        required: true,
                        message: '状态不能为空',
                        trigger: 'blur'
                    }],
                    super: [{
                        required: true,
                        message: '类型不能为空',
                        trigger: 'blur'
                    }]
                }
            }
        },
        created() {
            const self = this;
            getRole().then(function (response) {
                const result = response.data.data
                self.role = result.list
            })
        },
        mounted() {
        },
        methods: {
            onSearch() {
                this.$Func.search(this)
            },
            getPage(result) {
                this.list = result.data.list;
            },
            onEdit(row) {
                if (row) {
                    this.form = JSON.parse(JSON.stringify(row));
                    this.method = "PUT";
                } else {
                    this.method = "POST";
                    this.form = {};
                }
                this.dialog_status = true
            },
            onUpdatePassword(row) {
                if (row) {
                    this.form = JSON.parse(JSON.stringify(row));
                    this.method = "PUT";
                } else {
                    this.method = "POST";
                    this.form = {};
                }
                this.dialog_status2 = true
            },
            onSubmit(formName) {
                let form = this.form
                const url = this.method == "POST" ? this.url : this.url + "/" + form.id
                this.$Submit.submitClose(this, formName, url, form, this.method)
                this.dialog_status2 = false
            },
            onDelete(row, index) {
                this.$Submit.deleteSubmit(this, this.url + "/" + row.id, index)
            }
        }
    }
</script>
