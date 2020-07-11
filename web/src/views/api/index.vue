<template>
    <div class="main">
        <el-card class="box-card" shadow="hover">
            <!-- 搜索表单 -->
            <el-form label-position="top" :model="where" v-show="search.status">
                <el-row :gutter="10" class="transition-box">
                    <el-col :xs="24" :sm="12" :md="8" :lg="6" :xl="4">
                        <el-form-item label="分组">
                            <el-input placeholder="请输入" v-model="where.filter[0][2]" clearable></el-input>
                        </el-form-item>
                    </el-col>
                    <el-col :xs="24" :sm="12" :md="8" :lg="6" :xl="4">
                        <el-form-item label="名称">
                            <el-input placeholder="请输入" v-model="where.filter[1][2]" clearable></el-input>
                        </el-form-item>
                    </el-col>
                    <el-col :xs="24" :sm="12" :md="8" :lg="6" :xl="4">
                        <el-form-item label="路径">
                            <el-input placeholder="请输入" v-model="where.filter[2][2]" clearable></el-input>
                        </el-form-item>
                    </el-col>
                    <el-col :xs="24" :sm="12" :md="8" :lg="6" :xl="4">
                        <el-form-item label="请求类型">
                            <el-select class="widen" v-model="where.filter[3][2]" placeholder="请选择" clearable>
                                <el-option v-for="item in method_list" :key="item.value" :label="item.label" :value="item.value">
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
                    <el-table-column prop="id" label="标识" width="50">
                    </el-table-column>
                    <el-table-column prop="group" label="分组" show-overflow-tooltip width="150">
                    </el-table-column>
                    <el-table-column prop="title" label="名称" show-overflow-tooltip width="150">
                    </el-table-column>
                    <el-table-column prop="path" label="路径" show-overflow-tooltip width="150">
                    </el-table-column>
                    <el-table-column prop="method" label="请求类型" show-overflow-tooltip width="150">
                    </el-table-column>
                    <el-table-column prop="created_at" show-overflow-tooltip label="创建时间">
                        <template slot-scope="item">
                            {{$Func.formatDate(item.row.created_at,"YYYY-mm-dd HH:MM:SS")}}
                        </template>
                    </el-table-column>
                    <el-table-column prop="Setting" label="操作" width="150">
                        <template slot-scope="item">
                            <el-button type="text" @click="onEdit(item.row)">
                                {{$t('edit')}}
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
            <Page url="/admin/api" :params="where" ref="page" @return="getPage">
            </Page>
        </el-card>
        <!--弹出-->
        <template>
            <!-- 表单 -->
            <el-dialog :title="method=='POST'?'添加':'编辑'" :fullscreen="this.$Func.isFullscreen()" width="60%" lock-scroll :visible.sync="dialog_status">
                <el-form :model="form" :rules="rules" ref="form" :label-position="this.$Func.formPosition()" label-width="110px">
                    <el-form-item label="分组" prop="group">
                        <el-input v-model="form.group"></el-input>
                    </el-form-item>
                    <el-form-item label="名称" prop="title">
                        <el-input v-model="form.title"></el-input>
                    </el-form-item>
                    <el-form-item label="路径" prop="path">
                        <el-input v-model="form.path"></el-input>
                    </el-form-item>
                    <el-form-item label="请求类型" prop="method">
                        <el-radio-group v-model="form.method">
                            <el-radio v-for="item in method_list" :key="item.value" :label="item.value">{{item.label}}</el-radio>
                        </el-radio-group>
                    </el-form-item>
                </el-form>
                <div slot="footer">
                    <el-button @click="dialog_status = false">{{$t('cancel')}}</el-button>
                    <el-button type="primary" @click="onSubmit('form')" :loading="loading">{{ loading ? $t('wait') : $t('submit') }}</el-button>
                </div>
            </el-dialog>
        </template>
    </div>
</template>

<script>
    import Page from '../../components/Page/Index.vue'

    export default {
        components: {
            Page
        },
        data() {
            return {
                url: "admin/api",
                method: "POST",
                loading: false,
                fullscreen_loading: false,
                dialog_status: false,
                search: {
                    status: true
                },
                form: {
                    method: ''
                },
                list: [],
                method_list: [
                    {"label": "GET", "value": "GET"},
                    {"label": "POST", "value": "POST"},
                    {"label": "PUT", "value": "PUT"},
                    {"label": "DELETE", "value": "DELETE"},
                ],
                where: {
                    filter: [
                        ["group", "like", ""],
                        ["title", "like", ""],
                        ["path", "like", ""],
                        ["method", "=", ""],
                    ]
                },
                //校验规则
                rules: {
                    title: [{
                        required: true,
                        message: '不能为空',
                        trigger: 'blur'
                    }],
                    path: [{
                        required: true,
                        message: '不能为空',
                        trigger: 'blur'
                    }],
                    group: [{
                        required: true,
                        message: '不能为空',
                        trigger: 'blur'
                    }],
                    method: [{
                        required: true,
                        message: '不能为空',
                        trigger: 'blur'
                    }],
                }
            }
        },
        created() {
        },
        methods: {
            getPage(result) {
                this.list = result.data.list;
            },
            onSearch() {
                this.$Func.search(this)
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
            onSubmit(formName) {
                let form = this.form
                const url = this.method == "POST" ? this.url : this.url + "/" + form.id
                this.$Submit.submitClose(this, formName, url, form, this.method)
            },
            onDelete(row, index) {
                this.$Submit.deleteSubmit(this, this.url + "/" + row.id, index)
            }
        }
    }
</script>
