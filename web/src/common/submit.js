export default {
    /**
     * 通用删除提示
     * @param parameter
     * @param self
     */
    getList: function (self, callback) {
        self.$confirm(self.$i18n.t('delete_tip'), self.$i18n.t('warning'), {
            confirmButtonText: self.$i18n.t('ok'),
            cancelButtonText: self.$i18n.t('cancel'),
            type: 'warning'
        }).then(() => {
            callback()
        }).catch(() => {

        });
    },
    /**
     * 通用删除提示
     * @param parameter
     * @param self
     */
    deletePrompt: function (self, callback) {
        self.$confirm(self.$i18n.t('delete_tip'), self.$i18n.t('warning'), {
            confirmButtonText: self.$i18n.t('ok'),
            cancelButtonText: self.$i18n.t('cancel'),
            type: 'warning'
        }).then(() => {
            callback()
        }).catch(() => {

        });
    },
    /**
     * 通用删除提交
     * @param self this
     * @param url  删除地址
     * @param index list所用
     * @param callback 传入回调方法只执行回掉
     */
    deleteSubmit(self, url, index, callback) {
        const that = this;
        this.deletePrompt(self, function () {
            self.$http.request({
                method: "DELETE",
                url: url,
            }).then(function (response) {
                if (callback == undefined) {
                    that.delResponse(self, response.data, index)
                } else {
                    callback(response)
                }
            })
        })
    },
    /**
     * 通用删除公用返回
     * @param self
     * @param result
     */
    delResponse(self, result, index) {
        if (result.code == 200) {
            self.list.splice(index, 1)
            self.$message({
                type: 'success',
                message: self.$i18n.t('del_success')
            });
        }
    },
    /**
     * 通用提交表单验证
     * @param {Object} self 当前对象
     * @param {Object} formName 表单名称
     * @param {Object} callback 回调方法
     */
    submitVerify(self, formName, callback) {
        self.loading = true
        self.$refs[formName].validate((valid) => {
            if (valid) {
                callback()
            } else {
                self.loading = false
                return false
            }
        })
    },
    /**
     * 通用提交表单验证并关闭 弹出状态dialog_status
     * @param self this
     * @param formName 表单名称
     * @param url 请求地址
     * @param form 请求表单
     * @param method 请求类型
     * @param callback 请求回调,不存在执行默认关闭弹出
     */
    submitClose(self, formName, url, form, method = 'POST', callback) {
        let that = this;
        this.submitVerify(self, formName, function () {
            self.loading = true;
            self.$http.request({
                method: method,
                url: url,
                data: form
            }).then(function (response) {
                const result = response.data
                if (callback == undefined) {
                    that.responseClose(self, result)
                } else {
                    callback(response)
                }
            }).catch(function () {
                self.loading = false;
            })
        })
    },
    /**
     * 通用关闭弹出
     * @param self
     * @param result
     */
    responseClose(self, result) {
        self.loading = false
        self.$message({
            message: result.msg,
            type: result.code == 200 ? 'success' : 'error'
        })

        if (result.code == 200) {
            self.dialog_status = false;
            self.$refs.page.search()
        }
    },
    /**
     * 通用提交弹出表单验证并返回上一页
     * @param self this
     * @param formName 表单名称
     * @param url 请求地址
     * @param form 表单数据
     * @param method 请求类型
     * @param callback 请求回调,不存在执行默认返回
     */
    submitBack(self, formName, url, form, method, callback) {
        let that = this;
        this.submitVerify(self, formName, function () {
            self.loading = true;
            self.$http.request({
                method: method,
                url: url,
                data: form
            }).then(function (response) {
                const result = response.data
                if (callback == undefined) {
                    that.responseBack(self, result)
                } else {
                    callback(response)
                }
            }).catch(function () {
                self.loading = false;
            })
        })
    },
    /**
     * 通用返回上一页
     * @param self
     * @param result
     */
    responseBack(self, result) {
        self.loading = false
        self.$message({
            message: result.msg,
            type: result.code == 200 ? 'success' : 'error'
        })

        if (result.code == 200) {
            self.$router.go(-1);
        }
    },

}