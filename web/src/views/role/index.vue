<template>
    <div class="main">
        <!-- 按钮 -->
        <el-card class="box-card" shadow="hover">
            <!-- 搜索表单 -->
            <el-form label-position="top" :model="where" v-show="search.status">
                <el-row :gutter="10" class="transition-box">
                    <el-col :xs="24" :sm="12" :md="8" :lg="6" :xl="4">
                        <el-form-item label="角色名称">
                            <el-input placeholder="请输入" v-model="where.filter[0][2]" clearable></el-input>
                        </el-form-item>
                    </el-col>
                </el-row>
            </el-form>
            <!--按钮-->
            <el-row :gutter="10" class="button-row">
                <el-button @click="onEdit()" icon="el-icon-plus" size="small" type="primary">{{$t('add')}}</el-button>
                <el-button @click="onSearch()" icon="el-icon-search" v-loading.fullscreen.lock="loading" size="small" type="success">{{$t('search')}}</el-button>
                <el-button @click="$Func.clearSearch(where)" icon="el-icon-close" size="small">{{$t('clear_search')}}</el-button>
                <el-button @click="$Func.searchSwitch(search)" icon="el-icon-sort" size="small">{{search.status==true?$t('collapse'):$t('expand')}}</el-button>
            </el-row>
            <!-- 表格 -->
            <div class="table-box">
                <el-table :data="list" stripe>
                    <el-table-column prop="id" label="标识" width="50">
                    </el-table-column>
                    <el-table-column prop="name" label="角色名称" show-overflow-tooltip width="150">
                    </el-table-column>
                    <el-table-column prop="created_at" show-overflow-tooltip label="创建时间">
                        <template slot-scope="item">
                            {{$Func.formatDate(item.row.created_at,"YYYY-mm-dd HH:MM:SS")}}
                        </template>
                    </el-table-column>
                    <el-table-column prop="Setting" label="设置" fixed="right" width="180">
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
            <Page url="/admin/role" :params="where" ref="page" @return="getPage">
            </Page>
        </el-card>
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
                url: "admin/role",
                search: {
                    status: true
                },
                loading: false,
                where: {
                    filter: [
                        ["name", "like", ""],
                    ]
                },
                form: {
                    method: ''
                },
                list: [],
            }
        },
        created() {

        },
        methods: {
            getPage(result) {
                this.list = result.data.list;
            },
            onSearchSwitch() {
                this.$Func.searchSwitch(this)
            },
            onSearch() {
                this.$Func.search(this)
            },
            onEdit(row) {
                var data;
                row ? data = {id: row.id} : ''
                this.$Func.jump(this, '/role/edit', data)
            },
            onDelete(row, index) {
                this.$Submit.deleteSubmit(this, this.url + "/" + row.id, index)
            }
        }
    }
</script>
