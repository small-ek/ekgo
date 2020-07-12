<template>
    <div class="main">
        <el-alert title="当前操作,不影响任何表结构和表的变化,下载到对应代码目录结构里面,如果较为复杂并且需要关联查询子表可以建立虚拟表视图方式" type="error" effect="dark"></el-alert>
        <el-card class="box-card" shadow="hover">
            <!-- 搜索表单 -->
            <el-form label-position="top" :model="where" v-show="search.status">
                <el-row :gutter="10" class="transition-box">
                    <el-col :xs="24" :sm="12" :md="8" :lg="6" :xl="4">
                        <el-form-item label="数据表">
                            <el-select class="widen" v-model="where.table_name" placeholder="请选择" clearable filterable @change="onchangeTable">
                                <el-option v-for="item in table_list" :key="item.table_name" :label="item.table_name" :value="item.table_name">
                                    <span style="float: left">{{ item.table_comment }}</span>
                                    <span style="float: right; color: #8492a6; font-size: 13px">{{ item.table_name }}</span>
                                </el-option>
                            </el-select>
                        </el-form-item>
                    </el-col>
                </el-row>
            </el-form>
            <!--按钮-->
            <el-row :gutter="10" class="button-row">
                <el-button @click="onEdit(dialog)" icon="el-icon-plus" size="small" type="primary">{{$t('add')}}</el-button>
                <el-button @click="$Func.searchSwitch(search)" icon="el-icon-sort" size="small">{{search.status==true?$t('collapse'):$t('expand')}}</el-button>
                <el-dropdown style="margin-left: 10px" @command="onCreate">
                    <el-button size="small">
                        生成代码<i class="el-icon-arrow-down el-icon--right"></i>
                    </el-button>
                    <el-dropdown-menu slot="dropdown">
                        <el-dropdown-item command="vue">vue模板分页</el-dropdown-item>
                        <el-dropdown-item command="model">model模板</el-dropdown-item>
                        <el-dropdown-item command="service">service模板</el-dropdown-item>
                        <el-dropdown-item command="controller">controller模板</el-dropdown-item>
                        <el-dropdown-item command="route">router路由模板</el-dropdown-item>
                    </el-dropdown-menu>
                </el-dropdown>
            </el-row>
            <!--表单-->
            <el-table :data="list" ref="multipleTable" row-key="id" default-expand-all :tree-props="{children: 'children', hasChildren: 'hasChildren'}" stripe>
                <el-table-column prop="column_comment" width="150" label="字段注释" show-overflow-tooltip></el-table-column>
                <el-table-column prop="column_name" width="150" label="字段名称" show-overflow-tooltip></el-table-column>
                <el-table-column prop="data_type" label="字段类型" show-overflow-tooltip></el-table-column>
                <!--<el-table-column prop="character_maximum_length" label="字段长度" show-overflow-tooltip></el-table-column>-->
                <el-table-column label="头部搜索" show-overflow-tooltip>
                    <tempalte slot-scope="scope">
                        <el-switch v-model="scope.row.isSearch" active-text="是" inactive-text="否"></el-switch>
                    </tempalte>
                </el-table-column>
                <el-table-column label="搜索条件" show-overflow-tooltip>
                    <tempalte slot-scope="scope">
                        <el-select v-model="scope.row.condition" clearable placeholder="请选择">
                            <el-option v-for="item in condition" :key="item.value" :label="item.value" :value="item.value">
                            </el-option>
                        </el-select>
                    </tempalte>
                </el-table-column>
                <el-table-column label="列表展示" show-overflow-tooltip>
                    <tempalte slot-scope="scope">
                        <el-switch v-model="scope.row.isShow" active-text="是" inactive-text="否"></el-switch>
                    </tempalte>
                </el-table-column>
                <el-table-column label="表单是否显示" show-overflow-tooltip>
                    <tempalte slot-scope="scope">
                        <el-switch v-model="scope.row.isFormShow" active-text="是" inactive-text="否"></el-switch>
                    </tempalte>
                </el-table-column>
                <el-table-column label="表单是否必填" show-overflow-tooltip>
                    <tempalte slot-scope="scope">
                        <el-switch v-model="scope.row.required" active-text="是" inactive-text="否"></el-switch>
                    </tempalte>
                </el-table-column>
                <el-table-column label="表单文本类型" show-overflow-tooltip>
                    <tempalte slot-scope="scope">
                        <el-select v-model="scope.row.formType" clearable placeholder="请选择">
                            <el-option v-for="item in formTypeList" :key="item.label" :label="item.value" :value="item.label">
                            </el-option>
                        </el-select>
                    </tempalte>
                </el-table-column>
                <el-table-column label="操作" width="190">
                    <template slot-scope="scope">
                        <tempalte v-if="scope.row.formType=='checkbox'||scope.row.formType=='select'||scope.row.formType=='radio'||scope.row.formType=='transfer'">
                            <el-button @click="onEdit(option,scope.row,scope.$index)" size="small" type="text">编辑选项数据</el-button>
                            <el-divider direction="vertical"></el-divider>
                        </tempalte>
                        <el-button @click="onEdit(dialog,scope.row,scope.$index)" size="small" type="text">{{$t('edit')}}</el-button>
                        <el-divider direction="vertical"></el-divider>
                        <el-button type="text" class="red" @click="onDelete(scope.row,scope.$index)">删除</el-button>
                    </template>
                </el-table-column>
            </el-table>
        </el-card>
        <!--添加管理或者字段-->
        <template>
            <el-dialog :title="method=='POST'?'添加字段':'编辑字段'" :fullscreen="this.$Func.isFullscreen()" width="60%" lock-scroll :visible.sync="dialog.status">
                <el-alert title="关联表需要先生成,采用预加载的方式如果要搜索关联表内容,需要手动判断查询关联表内容得到关联标识在Where In的方式在主表查询即可满足" type="error" effect="dark"></el-alert>
                <el-form :model="form" :rules="rules" ref="form" :label-position="this.$Func.formPosition()" label-width="150px">
                    <el-form-item label="是否关联表" prop="isSearch">
                        <el-switch v-model="form.association_table" active-color="#7892fd">
                        </el-switch>
                    </el-form-item>
                    <el-form-item label="字段注释" prop="column_comment">
                        <el-input v-model="form.column_comment"></el-input>
                    </el-form-item>

                    <!--已关联-->
                    <template v-if="form.association_table==true">
                        <el-form-item label="当前关联的字段" prop="column_name">
                            <el-select class="widen" v-model="form.column_name" placeholder="请选择" clearable filterable>
                                <el-option v-for="item in list" :key="item.column_name" :label="item.column_name" :value="item.column_name">
                                    <span style="float: left">{{ item.column_name }}</span>
                                    <span style="float: right; color: #8492a6; font-size: 13px">{{ item.column_comment }}</span>
                                </el-option>
                            </el-select>
                        </el-form-item>
                        <el-form-item label="被引用的表" prop="table_name">
                            <el-select class="widen" v-model="form.table_name" placeholder="请选择" clearable filterable @change="onAssociationTable">
                                <el-option v-for="item in table_list" :key="item.table_name" :label="item.table_name" :value="item.table_name">
                                    <span style="float: left">{{ item.table_name }}</span>
                                    <span style="float: right; color: #8492a6; font-size: 13px">{{ item.table_comment }}</span>
                                </el-option>
                            </el-select>
                        </el-form-item>
                        <el-form-item label="被引用的字段" prop="field">
                            <el-select class="widen" v-model="form.field" placeholder="请选择" clearable filterable>
                                <el-option v-for="item in association_table_list" :key="item.column_name" :label="item.column_name" :value="item.column_name">
                                    <span style="float: left">{{ item.column_name }}</span>
                                    <span style="float: right; color: #8492a6; font-size: 13px">{{ item.column_comment }}</span>
                                </el-option>
                            </el-select>
                        </el-form-item>
                        <el-form-item label="关联关系" prop="relation">
                            <el-select class="widen" v-model="form.relation" placeholder="请选择" clearable filterable>
                                <el-option v-for="item in relation_list" :key="item.value" :label="item.value" :value="item.value">
                                    <span style="float: left">{{ item.value }}</span>
                                    <span style="float: right; color: #8492a6; font-size: 13px">{{ item.label }}</span>
                                </el-option>
                            </el-select>
                        </el-form-item>
                    </template>
                    <!--未关联-->
                    <template v-if="form.association_table==false">
                        <el-form-item label="字段名称" prop="column_name">
                            <el-input v-model="form.column_name"></el-input>
                        </el-form-item>
                        <el-form-item label="字段类型" prop="isSearch">
                            <el-select v-model="form.data_type" placeholder="请选择" clearable>
                                <el-option v-for="item in data_type" :key="item.value" :label="item.value" :value="item.value">
                                </el-option>
                            </el-select>
                        </el-form-item>
                        <el-form-item label="头部搜索" prop="isSearch">
                            <el-switch v-model="form.isSearch" active-color="#7892fd">
                            </el-switch>
                        </el-form-item>
                        <el-form-item label="搜索条件" prop="condition">
                            <el-select v-model="form.condition" placeholder="请选择" clearable>
                                <el-option v-for="item in condition" :key="item.value" :label="item.value" :value="item.value">
                                </el-option>
                            </el-select>
                        </el-form-item>
                    </template>
                </el-form>
                <div slot="footer">
                    <el-button @click="dialog.status = false">{{$t('cancel')}}</el-button>
                    <el-button type="primary" @click="onSubmit('form')" :loading="loading">{{ loading ? $t('wait') : $t('submit') }}</el-button>
                </div>
            </el-dialog>
        </template>
        <!--添加选项数据-->
        <template>
            <el-dialog title="编辑选项数据" :fullscreen="this.$Func.isFullscreen()" width="70%" lock-scroll :visible.sync="option.status">
                <el-form :model="form" :rules="rules" ref="form" :label-position="this.$Func.formPosition()" label-width="110px">
                    <el-form-item label="数据类型" prop="isSearch">
                        <el-radio v-model="form.options.type" label="true">静态数据</el-radio>
                        <el-radio v-model="form.options.type" label="false">动态数据</el-radio>
                    </el-form-item>
                    <!--动态数据-->
                    <template v-if="form.options.type=='false'">
                        <el-form-item label="请求地址" prop="url">
                            <el-input v-model="form.options.url"></el-input>
                        </el-form-item>
                        <el-form-item label="键名" prop="key">
                            <el-input v-model="form.options.key"></el-input>
                        </el-form-item>
                        <el-form-item label="值名" prop="value">
                            <el-input v-model="form.options.value"></el-input>
                        </el-form-item>
                    </template>
                    <!--静态数据-->
                    <template v-if="form.options.type=='true'">
                        <el-form-item label="选项数据" prop="isSearch">
                            <el-button style="width: 100%" type="primary" @click="onPushType(form.options)" plain>添加选项</el-button>
                            <el-table :data="form.options.list" border>
                                <el-table-column label="选项名" show-overflow-tooltip>
                                    <tempalte slot-scope="scope">
                                        <el-input v-model="scope.row.key"></el-input>
                                    </tempalte>
                                </el-table-column>
                                <el-table-column label="选项值" show-overflow-tooltip>
                                    <tempalte slot-scope="scope">
                                        <el-input v-model="scope.row.value"></el-input>
                                    </tempalte>
                                </el-table-column>
                                <el-table-column prop="value" label="操作">
                                    <tempalte slot-scope="scope">
                                        <el-button type="text" class="red" @click="onDeleteOption(scope.$index)">删除</el-button>
                                    </tempalte>
                                </el-table-column>
                            </el-table>
                        </el-form-item>
                    </template>
                </el-form>
                <div slot="footer">
                    <el-button @click="option.status = false">{{$t('cancel')}}</el-button>
                    <el-button type="primary" @click="onSubmit('form')" :loading="loading">{{ loading ? $t('wait') : $t('submit') }}</el-button>
                </div>
            </el-dialog>
        </template>
    </div>
</template>
<script>
    import {exportTemplate} from "../../common/template";
    import {getAllTable, getTable} from "../../common/api";

    export default {
        name: 'menu',
        components: {},
        data() {
            return {
                index: 0,
                loading: false,
                search: {
                    status: true
                },
                option: {
                    status: false
                },
                dialog: {
                    status: false
                },
                list: [],
                where: {
                    table_name: '',
                    table_comment: '',
                },
                form: {
                    association_table: true,
                    field: '',
                    options: {
                        type: '',
                        list: []
                    }
                },
                association_table_list: [],
                table_list: [],
                condition: [
                    {"value": 'like'},
                    {"value": 'notlike'},
                    {"value": 'rlike'},
                    {"value": '='},
                    {"value": '<>'},
                    {"value": '>'},
                    {"value": '<'},
                    {"value": 'in'},
                ],
                data_type: [
                    {"value": "varchar"},
                    {"value": "char"},
                    {"value": "int"},
                    {"value": "bigint"},
                    {"value": "smallint"},
                    {"value": "tinyint"},
                    {"value": "decimal"},
                    {"value": "double"},
                    {"value": "text"},
                    {"value": "longtext"},
                    {"value": "timestamp"},
                    {"value": "time"},
                    {"value": "enum"},
                    {"value": "json"},
                ],
                formTypeList: [
                    {"label": "input", "value": "文本"},
                    {"label": "textarea", "value": "文本域"},
                    {"label": "input_number", "value": "计数器"},
                    {"label": "switch", "value": "开关"},
                    {"label": "checkbox", "value": "多选框"},
                    {"label": "radio", "value": "单选框"},
                    {"label": "select", "value": "下拉框"},
                    {"label": "edit", "value": "富文本"},
                    {"label": "upload", "value": "上传图片组件"},
                    {"label": "upload_video", "value": "上传视频组件"},
                    {"label": "TimePicker", "value": "时间选择器"},
                    {"label": "TimeScopePicker", "value": "时间范围选择器"},
                    {"label": "DateTimePicker", "value": "日期时间选择器"},
                    {"label": "DateTimeScopePicker", "value": "日期时间范围选择器"},
                ],
                relation_list: [
                    {"label": "一对一", "value": "has_one"},
                    {"label": "一对多", "value": "has_many"},
                ],
                //校验规则
                rules: {
                    column_name: [{
                        required: true,
                        message: '字段不能为空',
                        trigger: 'blur'
                    }],
                }
            }
        },
        created() {
            const self = this;
            getAllTable().then(function (result) {
                self.table_list = result.data.data;
            })
        },
        methods: {
            //提交数据
            onSubmit(formName) {
                let self = this;

                this.$Submit.submitVerify(this, formName, function () {
                    if (self.index > -1) {
                        var list = JSON.parse(JSON.stringify(self.list))
                        list[self.index] = self.form
                        self.list = list
                    } else {
                        self.list.push(self.form)
                    }
                    self.option.status = false;
                    self.dialog.status = false;
                    self.loading = false;
                })
            },
            //编辑
            onEdit(status, row, index) {
                if (row) {
                    row.options.list == undefined ? row.options.list = [] : ""
                    row.options.type == undefined ? row.options.type = "" : ""
                    row.association_table == undefined ? this.form.association_table = false : ''
                    this.form = JSON.parse(JSON.stringify(row));
                    this.index = index;
                    this.method = "PUT";
                } else {
                    this.index = -1;
                    this.method = "POST";
                    this.form = {association_table: true, field: '', options: {type: "", list: []}}
                }
                status.status = true
            },
            onImport() {
                const self = this;
                const table_name = self.where.table_name;

                if (table_name) {
                    getTable(table_name).then(function (result) {
                        result.data.data.forEach(function (value) {
                            value["isSearch"] = false;
                            value["isShow"] = true;
                            value["isFormShow"] = true;
                            value["required"] = false;
                            value["options"] = {};
                            value["formType"] = "input";
                            value["condition"] = "";
                        })
                        self.list = result.data.data;
                    })
                }
            },
            //选择表
            onchangeTable(row) {
                const self = this;

                if (row == "") {
                    self.list = [];
                }

                this.table_list.forEach(function (value) {
                    if (row == value.table_name) {
                        self.where["table_comment"] = value.table_comment
                    }
                })
                this.onImport()
            },
            //获取关联字段
            onAssociationTable(table_name) {
                if (table_name) {
                    const self = this;
                    getTable(table_name).then(function (result) {
                        self.association_table_list = result.data.data
                        self.form.field = ""
                    })
                }
            },
            //创建
            onCreate(type) {
                let list = this.list;
                let table_name = this.where.table_name
                let table_comment = this.where.table_comment

                if (table_name == "" || list.length == 0) {
                    this.$message.error('请选择数据表并导入当前页面');
                    return;
                }
                exportTemplate(list, table_name, table_comment, type);
            },
            //删除表单
            onDelete(row, index) {
                this.list.splice(index, 1)
            },
            //添加选项类型
            onPushType(options) {
                options.list.push({key: "选项" + options.list.length, value: "选项" + options.list.length})
            },
            //删除选项
            onDeleteOption(index) {
                this.form.options.list.splice(index, 1)
            }
        }
    }
</script>
