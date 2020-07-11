<template>
    <div>
        <el-card class="page-card pagination" shadow="hover" v-show="isShow">
            <el-pagination background @size-change="handleSizeChange" @current-change="handleCurrentChange" :current-page.sync="currentPage"
                           :page-sizes="[10,15,50,100,200,300,400,500,600,700,800,900,1000]" :page-size="pageZize" layout="total, sizes, prev, pager, next, jumper"
                           :total="total">
            </el-pagination>
        </el-card>

    </div>
</template>
<script>
    export default {
        name: 'Pagination',
        props: {
            url: String,
            params: Object,
            refresh: Boolean,

        },
        data() {
            return {
                pageZize: 10,
                currentPage: 1,
                total: 0,
                isShow: false
            }
        },
        created: function () {
            this.paginate();
        },
        methods: {
            paginate: function () {
                if (this.url) {
                    const self = this;
                    this.params['page_size'] = this.pageZize;
                    this.params['current_page'] = this.currentPage;
                    this.$http.get(this.url, {
                        params: this.params
                    }).then(function (response) {
                        const result = response.data;
                        self.total = result.data.total
                        self.total > 0 ? self.isShow = true : self.isShow = false;
                        self.$emit('return', result);
                    })
                }
            },
            handleCurrentChange: function (val) { //当前页
                this.currentPage = val;
                this.paginate();
            },
            handleSizeChange: function (val) { //多少条
                this.pageZize = val
                this.paginate();
            },
            search: function () {
                this.currentPage = 1;
                this.paginate();
            }
        }
    }
</script>
<style scoped>
    .pagination {
        text-align: center;
    }

    .page-card {
        margin-top: 10px;
        overflow: auto;
    }
</style>
