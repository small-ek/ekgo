export const vueTemplate = (data, newTable, table_comment) => {
    var serachHtml = ''//搜索模板
    var filter = ''//过滤数据
    var tableHtml = ''//表格html
    var formHtml = ''//表单模板
    var filterIndex = 0//过滤函数索引
    console.log(data)
    data.forEach(function (value, index) {
        //是否表格展示
        if (value.isShow == true) {
            tableHtml = tableHtml + `<el-table-column prop="` + value.column_name + `" label="` + value.column_comment + `" show-overflow-tooltip></el-table-column>\n`
        }
        //是否搜索
        if (value.isSearch == true) {

            //搜索条件脚本
            let condition = value.condition != '' ? value.condition : 'like';
            filter = filter + `["` + value.column_name + `", "` + condition + `", ""],\n`
            //搜索页面
            serachHtml = serachHtml + `<el-col :xs="24" :sm="12" :md="6" :lg="6" :xl="4">
                        <el-form-item label="` + value.column_comment + `">
                            <el-input placeholder="请输入" v-model="where.filter[` + filterIndex + `][2]" clearable></el-input>
                        </el-form-item>
                    </el-col>`
            filterIndex++
        }
        formHtml = formHtml + `<el-form-item label="` + value.column_comment + `" prop="` + value.column_name + `">
                    <el-input v-model="form.` + value.column_name + `" clearable></el-input>
                </el-form-item>`
    })
    return `<template>
    <div class="main">
        <el-card class="box-card" shadow="hover">
            <!-- 搜索 -->
            <el-form label-position="top" :model="where">
                <el-row :gutter="10" v-show="search.status" class="transition-box">
                    ` + serachHtml + `
                </el-row>
            </el-form>
            <!-- 按钮 -->
            <el-row :gutter="10" class="button-row">
                <el-button @click="onEdit()" icon="el-icon-plus" size="small" type="primary">{{$t('add')}}</el-button>
                <el-button @click="onSearch()" icon="el-icon-search" v-loading.fullscreen.lock="fullscreen_loading" type="success" size="small">{{$t('search')}}</el-button>
                <el-button @click="$Func.clearSearch(where)" icon="el-icon-close" size="small">{{$t('clear_search')}}</el-button>
                <el-button @click="$Func.searchSwitch(search)" icon="el-icon-sort" size="small">{{search.status==true?$t('collapse'):$t('expand')}}</el-button>
            </el-row>
            <!-- 表格 -->
            <div class="table-box">
                <el-table :data="list" stripe class="table">
                ` + tableHtml + `
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
            <Page :url="url" :params="where" ref="page" @return="getPage">
            </Page>
        </el-card>
        <!-- 表单 TODO -->
        <el-dialog :title="method=='POST'?'添加` + table_comment + `':'编辑` + table_comment + `'" :fullscreen="this.$Func.isFullscreen()" width="55%" lock-scroll :visible.sync="dialog_status">
            <el-form :model="form" :rules="rules" ref="form">
                ` + formHtml + `
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
                url: "admin/` + newTable + `",
                method: "POST",
                fullscreen_loading: false,//搜索全屏加载
                loading: false,//表单加载
                dialog_status: false,//弹出状态
                search: {status: true},//搜索状态
                where: {
                    filter: [
                        ` + filter + `
                    ]
                },
                form: {},
                list: [],
                //校验规则 TODO
                rules: {}
            }
        },
        created() {
        },
        methods: {
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
`
}
export const options = (formType, options) => {
    //静态数据
    if (options.type == "true") {
        list.forEach(function (value) {

        })
    } else {

    }
}
//表单列表
export const formList = (formType, column_comment, column_name, options) => {
    var html = '';
    switch (formType) {
        //文本框
        case "input":
            html = `<el-col :xs="24" :sm="24" :md="24" :lg="24" :xl="12">
                       <el-form-item label="` + column_comment + `" prop="` + column_name + `">
                            <el-input placeholder="请输入" v-model="form.` + column_name + `" clearable></el-input>
                        </el-form-item>
                    </el-col>`
            break;
        //文本域
        case "textarea":
            html = `<el-col :xs="24" :sm="24" :md="24" :lg="24" :xl="12">
                       <el-form-item label="` + column_comment + `" prop="` + column_name + `">
                            <el-input type="textarea" :rows="3" placeholder="请输入" v-model="form.` + column_name + `" clearable></el-input>
                        </el-form-item>
                    </el-col>`
            break;
        //计数器
        case "input_number":
            html = `<el-col :xs="24" :sm="24" :md="24" :lg="24" :xl="12">
                       <el-form-item label="` + column_comment + `" prop="` + column_name + `">
                             <el-input-number v-model="form.` + column_name + `" controls-position="right"  :min="1" :max="500"></el-input-number>
                        </el-form-item>
                    </el-col>`
            break;
        //开关
        case "switch":
            html = `<el-col :xs="24" :sm="24" :md="24" :lg="24" :xl="12">
                       <el-form-item label="` + column_comment + `" prop="` + column_name + `">
                             <el-switch v-model="form.` + column_name + `" active-color="#13ce66"></el-switch>
                        </el-form-item>
                    </el-col>`
            break;
        //复选框
        case "checkbox":
            html = `<el-col :xs="24" :sm="24" :md="24" :lg="24" :xl="12">
                       <el-form-item label="` + column_comment + `" prop="` + column_name + `">
                            <el-checkbox v-for="row in ` + column_comment + `_list" v-model="form.` + column_name + `" :label="row.value" :key="row.value">{{row.key}}</el-checkbox>
                        </el-form-item>
                    </el-col>`
            break;
        //单选框
        case "radio":
            html = `<el-col :xs="24" :sm="24" :md="24" :lg="24" :xl="12">
                       <el-form-item label="` + column_comment + `" prop="` + column_name + `">
                            <el-radio v-for="row in ` + column_comment + `_list" v-model="form.` + column_name + `" :label="row.value" :key="row.value">{{row.key}}</el-radio>
                        </el-form-item>
                    </el-col>`
            break;
        //单选框
        case "radio":
            html = `<el-col :xs="24" :sm="24" :md="24" :lg="24" :xl="12">
                       <el-form-item label="` + column_comment + `" prop="` + column_name + `">
                            <el-radio v-for="row in ` + column_comment + `_list" v-model="form.` + column_name + `" :label="row.value" :key="row.value">{{row.key}}</el-radio>
                        </el-form-item>
                    </el-col>`
            break;
        //下拉框
        case "select":
            html = `<el-col :xs="24" :sm="24" :md="24" :lg="24" :xl="12">
                       <el-form-item label="` + column_comment + `" prop="` + column_name + `">
                       <el-select v-model="value" placeholder="请选择">
                        <el-option v-for="item in options" :key="item.value" :label="item.label":value="item.value"></el-option>
                        </el-select>
                            <el-radio v-for="row in ` + column_comment + `_list" v-model="form.` + column_name + `" :label="row.value" :key="row.value">{{row.key}}</el-radio>
                        </el-form-item>
                    </el-col>`
            break;
    }
}
