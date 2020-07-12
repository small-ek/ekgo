import {convHump, getRidTable} from "../template";

export const vueTemplate = (data, newTable, table_comment) => {
    var serachHtml = ''; //搜索模板
    var filter = '';     //过滤数据
    var tableHtml = '';  //表格html
    var formHtml = '';   //表单模板
    var filterIndex = 0; //过滤函数索引
    var script = '';     //调用脚本
    var func = '';       //方法
    var createFunc = ''; //调用方法
    var returnData = ''; //data数据
    var components = ''; //组件引入
    var formData = '';   //表单数据
    var required = '';     //必填
    var url = 'admin/' + getRidTable(newTable)

    data.forEach(function (value, index) {
        //表单是否显示
        if (value.isFormShow == true) {
            var getformList = formList(value.formType, value.column_comment, value.column_name, value.options)
            formHtml = formHtml + getformList["html"];
            script = script + getformList["script"];
            func = func + getformList["func"];
            createFunc = createFunc + getformList["createFunc"];
            returnData = returnData + getformList["data"];
            components = components + getformList["components"];
            formData = formData + getformList['formData']
        }

        //表格是否展示
        if (value.isShow == true) {
            tableHtml = tableHtml + `<el-table-column prop="` + value.column_name + `" label="` + value.column_comment + `" show-overflow-tooltip></el-table-column>\n`
        }
        //是否必填
        if (value.required == true) {
            required = required + `` + value.column_name + `: [{required: true,message: '不能为空',trigger: 'blur'}],\n`
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
        <template>
            <el-dialog :title="method=='POST'?'添加` + table_comment + `':'编辑` + table_comment + `'" :fullscreen="this.$Func.isFullscreen()" width="55%" lock-scroll :visible.sync="dialog_status">
            <el-row :gutter="15">
                <el-form :model="form" :rules="rules" ref="form" label-width="125px">
                    ` + formHtml + `
                </el-form>
            </el-row>
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
    ` + script + `
    export default {
        components: {Page,` + components + `},
        data() {
            return {
                url: "` + url + `",
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
                form: {
                    ` + formData + `
                },
                list: [],
                ` + returnData + `
                //校验规则 TODO
                rules: {
                    ` + required + `
                }
            }
        },
        created() {
            ` + createFunc + `
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
                    this.$refs['form'] != undefined ? this.$refs['form'].resetFields() : ''
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
            },
            ` + func + `
        }
    }
</script>
`
}

//表单列表
export const formList = (formType, column_comment, column_name, options) => {
    var html = '';//页面
    var script = ''//引入脚本

    var getFormScript = formScript(options, column_name + "_list")//脚本创建
    var func = getFormScript["func"];//方法
    var createFunc = getFormScript["createFunc"];//调用方法名
    var data = getFormScript["data"];//data数据
    var components = ''//组件引入
    var formData = ''//表单json
    switch (formType) {
        //文本框
        case "input":
            html = `<el-col :span="24">
                       <el-form-item label="` + column_comment + `" prop="` + column_name + `">
                            <el-input placeholder="请输入" v-model="form.` + column_name + `" clearable></el-input>
                        </el-form-item>
                    </el-col>`
            break;
        //文本域
        case "textarea":
            html = `<el-col :span="24">
                       <el-form-item label="` + column_comment + `" prop="` + column_name + `">
                            <el-input type="textarea" :rows="3" placeholder="请输入" v-model="form.` + column_name + `" clearable></el-input>
                        </el-form-item>
                    </el-col>`
            break;
        //计数器
        case "input_number":
            html = `<el-col :span="24">
                       <el-form-item label="` + column_comment + `" prop="` + column_name + `">
                             <el-input-number v-model="form.` + column_name + `" controls-position="right"  :min="1" :max="500"></el-input-number>
                        </el-form-item>
                    </el-col>`
            break;
        //开关
        case "switch":
            html = `<el-col :span="24">
                       <el-form-item label="` + column_comment + `" prop="` + column_name + `">
                             <el-switch v-model="form.` + column_name + `" active-color="#13ce66"></el-switch>
                        </el-form-item>
                    </el-col>`
            break;
        //复选框
        case "checkbox":
            var item = ''
            //静态数据
            if (options.type == "true") {
                item = `<el-checkbox v-for="row in ` + column_name + `_list" :label="row.value" :key="row.key">{{row.key}}</el-checkbox>`
            } else {
                item = `<el-checkbox v-for="row in ` + column_name + `_list" :label="` + options.key + `" :key="` + options.value + `">{{row.` + options.key + `}}</el-checkbox>`
            }
            html = html = `<el-col :span="24">
                       <el-form-item label="` + column_comment + `" prop="` + column_name + `">
                            <el-checkbox-group v-model="form.` + column_name + `">
                                ` + item + `
                            </el-checkbox-group>
                        </el-form-item>
                    </el-col>`
            formData = `` + column_name + `:[],`
            break;
        //单选框
        case "radio":
            var item = ''
            //静态数据
            if (options.type == "true") {
                item = `<el-radio v-for="row in ` + column_name + `_list" v-model="form.` + column_name + `" :label="row.value" :key="row.value">{{row.key}}</el-radio>`
            } else {
                item = `<el-radio v-for="row in ` + column_name + `_list" v-model="form.` + column_name + `" :label="` + options.key + `" :key="` + options.value + `">{{row.` + options.key + `}}</el-radio>`
            }

            html = html = `<el-col :span="24">
                       <el-form-item label="` + column_comment + `" prop="` + column_name + `">
                            ` + item + `
                        </el-form-item>
                    </el-col>`
            break;
        //下拉框
        case "select":
            var item = ''
            //静态数据
            if (options.type == "true") {
                item = `<el-option v-for="row in ` + column_name + `_list"  :label="row.key" :key="row.value" :value="row.value"></el-option>`
            } else {
                item = `<el-option v-for="row in ` + column_name + `_list" v-model="form.` + column_name + `" :label="` + options.key + `" :key="` + options.value + `" :value="row.` + options.key + `"></el-option>`
            }

            html = html = `<el-col :span="24">
                       <el-form-item label="` + column_comment + `" prop="` + column_name + `">
                            <el-select v-model="form.` + column_name + `" placeholder="请选择">
                                 ` + item + `
                            </el-select>
                        </el-form-item>
                    </el-col>`
            break;
        //编辑器
        case "edit":
            html = `<el-col :span="24">
                       <el-form-item label="` + column_comment + `" prop="` + column_name + `">
                            <Editor :model.sync="form.` + column_name + `" :max="2000"></Editor>
                       </el-form-item>
                    </el-col>`
            script = `import Editor from '../../components/Editor/Index.vue';`
            components = `Editor,`
            break;
        //上传组件
        case "upload":
            html = `<el-col :span="24">
                        <el-form-item label="` + column_comment + `" prop="` + column_name + `">
                            <UploadImage :fileList="form.` + column_name + `" rules="image" :limit="5"></UploadImage>
                        </el-form-item>
                    </el-col>`
            script = `import UploadImage from '../../components/UploadImage/Index.vue';`
            components = `UploadImage,`
            break;
        //上传视频组件
        case "upload_video":
            html = `<el-col :span="24">
                          <el-form-item label="` + column_comment + `" prop="` + column_name + `">
                              <UploadImage :fileList="form.` + column_name + `" rules="video" :limit="1" :max="100000"></UploadImage>
                          </el-form-item>
                    </el-col>`
            script = `import UploadImage from '../../components/UploadImage/Index.vue';`
            components = `UploadImage,`
            break;
        //时间选择器
        case "TimePicker":
            html = `<el-col :span="24">
                        <el-form-item label="` + column_comment + `" prop="` + column_name + `">
                            <el-time-picker v-model="form.` + column_name + `" :picker-options="{'selectableRange':'00:00:00-23:59:59'}" placeholder="请选择时间"></el-time-picker>
                       </el-form-item>
                    </el-col>`
            break;
        //时间选择器
        case "TimeScopePicker":
            html = `<el-col :span="24">
                        <el-form-item label="` + column_comment + `" prop="` + column_name + `">
                            <el-time-picker is-range v-model="form.` + column_name + `" range-separator="至" start-placeholder="开始时间" end-placeholder="结束时间" placeholder="选择时间范围">
                            </el-time-picker>
                        </el-form-item>
                    </el-col>`
            break;
        //日期选择器
        case "DateTimePicker":
            html = `<el-col :span="24">
                        <el-form-item label="` + column_comment + `" prop="` + column_name + `">
                            <el-date-picker v-model="form.` + column_name + `" type="date"placeholder="选择日期"></el-date-picker>
                        </el-form-item>
                    </el-col>`
            break;
        //日期选择器
        case "DateTimeScopePicker":
            html = `<el-col :span="24">
                        <el-form-item label="` + column_comment + `" prop="` + column_name + `">
                            <el-date-picker v-model="form.` + column_name + `" type="daterange" range-separator="至" start-placeholder="开始日期"end-placeholder="结束日期"></el-date-picker>
                        </el-form-item>
                    </el-col>`
            break;
    }
    return {"html": html, "script": script, "func": func, "createFunc": createFunc, "data": data, 'components': components, 'formData': formData};
}
//表单动态数据和静态数据处理
export const formScript = (options, column_name) => {
    //数据
    var data = '';
    //方法
    var func = '';
    //创建调用方法
    var createFunc = ''
    //静态数据
    if (options.type == "true") {
        options.list.forEach(function (value) {
            data = data + `{"key": "` + value.key + `", "value": "` + value.value + `"},\n`
        })
        data = `` + column_name + `:[\n` + data + `],`
    } else if (options.type == "false") {
        data = `` + column_name + `:[],`
        let humpColumeName = convHump(column_name)
        createFunc = `this.get` + humpColumeName + `()`
        func = `get` + humpColumeName + `() {
                const self = this;
                this.$http.get('` + options.url + `').then(function (response) {
                    self.` + column_name + `= response.data.data.list
                })
            },\n`
    }
    return {
        data: data,
        func: func,
        createFunc: createFunc,
    }
}