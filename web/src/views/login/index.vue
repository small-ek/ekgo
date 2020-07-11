<template>
    <div class="main">
        <div class="login">
            <el-card class="box-card">
                <div class="login-head">{{$t("manage")}}</div>
                <el-row :gutter="10" style="padding: 40px 30px 80px">
                    <el-form :model="form" :rules="rules" ref="form" label-width="0px" size="default">
                        <el-form-item prop="username">
                            <el-input v-model="form.username" :placeholder="$t('username')" @keyup.enter.native="submitForm('form')" suffix-icon="el-icon-edit" clearable>
                            </el-input>
                        </el-form-item>
                        <el-form-item prop="password">
                            <el-input v-model="form.password" :placeholder="$t('password')" @keyup.enter.native="submitForm('form')" type="password" show-password>
                            </el-input>
                        </el-form-item>
                        <el-col :xs="24" :sm="24" :md="24" :lg="24" :xl="24" style="margin-top: 35px">
                            <el-button @click="submitForm('form')" :loading="loading" type="primary" style="width: 100%;height: 45px;font-size: 14px;color:white">
                                {{$t("login")}}
                            </el-button>
                        </el-col>
                    </el-form>
                </el-row>
            </el-card>
        </div>
        <Canvas></Canvas>
    </div>
</template>
<script>
    import {local} from '../../common/storage'
    import {Base64Encode} from "../../common/secret";
    import Canvas from '../../components/Canvas/vue-particle-line'

    export default {
        name: 'login',
        components: {
            Canvas
        },
        data() {
            return {
                form: { //表单中的参数
                    username: '',
                    password: '',
                },
                loading: false,
                //校验规则
                rules: {
                    username: [{
                        required: true,
                        message: '请输入用户名',
                        trigger: 'blur'
                    },],
                    password: [{
                        required: true,
                        message: '请输入密码',
                        trigger: 'blur'
                    },],
                },
            }
        },
        mounted() {

        },
        methods: {
            submitForm(formName) {
                const self = this;
                this.$refs[formName].validate((valid) => {
                    if (valid) {
                        self.loading = true
                        var form = JSON.parse(JSON.stringify(this.form))
                        form.password = Base64Encode(form.password)
                        this.$http.post("admin/login", form).then(function (response) {
                            const result = response.data;
                            self.loading = false

                            if (result.code == 200) {
                                self.$message({
                                    message: self.$t('login_success'),
                                    type: 'success'
                                })
                                local.set('user', result.data, 86400);
                                self.$store.user = result.data

                                if (self.$route.query.redirect) {
                                    window.location.href = self.$route.query.redirect
                                } else {
                                    window.location.href = '/'
                                }

                            }

                            self.$message({
                                message: result.msg,
                                type: 'error'
                            })
                        }).catch(function () {
                            self.loading = false
                        })
                    } else {
                        return false;
                    }
                });
            },
            resetForm(formName) {
                this.$refs[formName].resetFields();
            },

        }
    }
</script>

<style scoped>
    .main {
        background: url("../../assets/images/logo-background.svg"), linear-gradient(-180deg, #BCC5CE 0%, #929EAD 98%), radial-gradient(at top left, rgba(255, 255, 255, 0.30) 0%, rgba(0, 0, 0, 0.30) 100%);
        background-blend-mode: screen;
        background-repeat: no-repeat;
        background-size: 100% 100%;
        width: 100%;
        height: 100%;
        position: absolute;
        z-index: 20;
    }

    .login {
        z-index: 100;
        margin: 130px auto 0 auto;
        height: 440px;
        max-width: 400px;
        opacity: 0.84;
        background-image: linear-gradient(to right, rgba(0, 0, 0, 0) 0%, rgba(0, 0, 0, 0.4) 5%, rgba(0, 0, 0, 0.6) 10%, rgba(0, 0, 0, 0.6) 90%, rgba(0, 0, 0, 0.4) 95%, rgba(0, 0, 0, 0) 100%);
    }

    .login-head {
        font-size: 25px;
        text-align: center;
        font-weight: bold;
        color: #343434;
        text-shadow: -1px -1px #ccc;
        padding: 15px 0;
        width: 100%;
        margin: 0 auto;
    }
</style>
