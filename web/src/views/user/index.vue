<template>
    <div class="main">
        <!-- 搜索 -->
        <el-card class="box-card" shadow="hover">
            <el-form label-position="top" :model="where">
                <el-row :gutter="10" v-show="search.status" class="transition-box">
                    <el-col :xs="24" :sm="12" :md="6" :lg="6" :xl="4">
                        <el-form-item label="用户名">
                            <el-input placeholder="请输入" v-model="where.filter[0][2]" clearable></el-input>
                        </el-form-item>
                    </el-col>
                    <el-col :xs="24" :sm="12" :md="6" :lg="6" :xl="4">
                        <el-form-item label="邮箱">
                            <el-input placeholder="请输入" v-model="where.filter[1][2]" clearable></el-input>
                        </el-form-item>
                    </el-col>
                </el-row>
            </el-form>
            <el-row :gutter="10" class="button-row">
                <!--<el-button @click="onEdit()" icon="el-icon-plus" size="small" type="primary">{{$t('add')}}</el-button>-->
                <el-button @click="onSearch()" icon="el-icon-search" v-loading.fullscreen.lock="fullscreen_loading" type="success" size="small">{{$t('search')}}</el-button>
                <el-button @click="$Func.clearSearch(where)" icon="el-icon-close" size="small">{{$t('clear_search')}}</el-button>
                <el-button @click="$Func.searchSwitch(search)" icon="el-icon-sort" size="small">{{search.status==true?$t('collapse'):$t('expand')}}</el-button>
            </el-row>

            <!-- 表格 -->
            <div class="table-box">
                <el-table :data="list" stripe class="table">
                    <el-table-column prop="id" label="标识" width="50">
                    </el-table-column>
                    <el-table-column prop="username" label="用户名" show-overflow-tooltip width="150">
                    </el-table-column>
                    <el-table-column prop="email" label="邮箱" show-overflow-tooltip width="150">
                    </el-table-column>
                    <el-table-column prop="full_name" label="姓名" show-overflow-tooltip>
                    </el-table-column>
                    <el-table-column prop="status" label="状态">
                        <template slot-scope="scope">
                            <el-tag size="mini" v-if="scope.row.status=='true'">启用</el-tag>
                            <el-tag size="mini" type="danger" v-if="scope.row.status=='false'">禁用</el-tag>
                        </template>
                    </el-table-column>
                    <el-table-column prop="created_at" show-overflow-tooltip label="创建时间">
                    </el-table-column>
                    <el-table-column prop="Setting" label="设置" width="250">
                        <template slot-scope="item">
                            <el-button type="text" icon="el-icon-edit" @click="onEdit(item.row)">
                                {{$t('edit')}}
                            </el-button>
                            <el-button type="text" class="red" icon="el-icon-delete" @click="onDelete(item.row,item.$index)">
                                {{$t('del')}}
                            </el-button>
                        </template>
                    </el-table-column>
                </el-table>
            </div>

            <!-- 分页 -->
            <Page url="/admin/user" :params="where" ref="page" @return="getPage">
            </Page>
        </el-card>
        <!-- 表单 -->
        <el-dialog :title="method=='POST'?'添加用户':'编辑用户'" :fullscreen="this.$Func.isFullscreen()" lock-scroll :visible.sync="dialog_status">
            <el-form :model="form" :rules="rules" ref="form">
                <el-form-item label="状态" prop="status">
                    <el-radio-group v-model="form.status">
                        <el-radio label="true">启用</el-radio>
                        <el-radio label="false">禁用</el-radio>
                    </el-radio-group>
                </el-form-item>
            </el-form>
            <div slot="footer">
                <el-button @click="dialog_status = false">{{$t('cancel')}}</el-button>
                <el-button type="primary" @click="onSubmit('form')" :loading="loading">{{ loading ? $t('wait') : $t('submit') }}</el-button>
            </div>
        </el-dialog>
    </div>
</template>

<script>
    import Page from '../../components/Page/Index.vue'

    export default {
        components: {Page},
        data() {
            return {
                url: 'admin/user',
                search: {status: true},//搜索栏状态
                fullscreen_loading: false,//搜索全屏加载
                loading: false,//表单加载
                dialog_status: false,//弹出状态
                method: "POST",
                where: {
                    filter: [
                        ["username", "like", ""],
                        ["email", "like", ""],
                    ]
                },
                form: {},
                list: [],
                //校验规则
                rules: {}
            }
        },
        props: {},
        created() {
        },
        mounted() {

        },
        watch() {
        },
        methods: {
            //搜索开关
            onSearchSwitch() {
                this.$Func.searchSwitch(this)
            },
            //搜索
            onSearch() {
                this.$Func.search(this)
            },
            /**
             * 获取分页
             * @param result
             */
            getPage(result) {
                this.list = result.data.list;
            },
            /**
             * 编辑
             * @param row
             */
            onEdit(row) {
                if (row) {
                    this.form = JSON.parse(JSON.stringify(row));
                    this.method = "PUT";
                } else {
                    this.method = "POST";
                    this.form = {};
                }
                this.dialog_status = true

                /*
                //跳转
                var data;
                if (row) {
                    data = {id: row.id}
                }
                this.$Func.jump(this, '/role/edit', data)
                */
            },
            /**
             * 提交表单
             * @param formName
             */
            onSubmit(formName) {
                let form = this.form
                const url = this.method == "POST" ? this.url : this.url + "/" + form.id
                this.$Submit.submitClose(this, formName, url, form, this.method)
            },
            /**
             * 删除
             * @param row 参数
             * @param index 索引
             */
            onDelete(row, index) {
                this.$Submit.deleteSubmit(this, 'admin/user/' + row.id, index)
            }
        }
    }
</script>
