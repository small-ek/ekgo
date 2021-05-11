<template>
  <div id="login">
    <div class="login-form">
      <a-form :label-col="labelCol" :wrapper-col="wrapperCol" :model="form" :rules="rules" ref="formRef">
        <a-form-item>
          <img class="logo" src="../../assets/image/logo.png"/>
          <div class="head">Ekgo Admin</div>
          <div class="desc"></div>
        </a-form-item>
        <a-form-item name="username">
          <a-input placeholder="用户名" v-model:value="form.username" autocomplete="off"/>
        </a-form-item>
        <a-form-item name="password">
          <a-input placeholder="密 码" v-model:value="form.password" type="password" @keyup.enter="onSubmit"
                   autocomplete="off"/>
        </a-form-item>
        <!--                <a-form-item>-->
        <!--                  <a-checkbox :checked="true">-->
        <!--                    记住我-->
        <!--                  </a-checkbox>-->
        <!--                  <a class="forgot" href=""> 忘记密码 </a>-->
        <!--                </a-form-item>-->
        <a-form-item :wrapper-col="{ span: 24 }">
          <a-button :loading="loading" type="primary" @click="onSubmit">
            登录
          </a-button>
        </a-form-item>
      </a-form>
    </div>
  </div>
</template>
<script>
import {defineComponent, reactive, ref, toRaw} from 'vue';
import {message} from 'ant-design-vue';
import user from "../../api/modules/user.js"
import menu from "../../api/modules/menu.js"
import {useStore} from "vuex";
import {useRouter} from "vue-router";

export default defineComponent({
  setup() {
    const formRef = ref();
    const loading = ref(false);
    const {dispatch, commit} = useStore();
    const router = useRouter();
    const form = reactive({
      username: "",
      password: ""
    });

    const rules = {
      username: [
        {
          required: true,
          message: '请输入用户名',
          trigger: 'blur',
        },
      ],
      password: [
        {
          required: true,
          message: '请输入密码',
          trigger: 'blur',
        },
      ]
    };

    //提交表单
    const onSubmit = () => {
      loading.value = true
      formRef.value.validate().then(() => {
        //登录
        user.login(toRaw(form)).then(async result => {
          if (result.code == 200) {
            await dispatch('user/login', result)
            //设置菜单
            await menu.getUserMenu(result.data["authorization"]).then(async res => {
              await dispatch('routes/setRouter', res.data.data)
              setTimeout(() => {
                router.push("/")
              }, 200)
            })
          } else {
            loading.value = false
            message.error(result.msg)
          }
        })
      }).catch(error => {
        loading.value = false
        console.log('error', error);
      });
    };

    return {
      formRef,
      loading: loading,
      labelCol: {span: 6},
      wrapperCol: {span: 24},
      form,
      rules,
      onSubmit
    };
  },
});
</script>

<style lang="less" scoped>
#login {
  width: 100%;
  height: 100%;
  background: url(../../assets/image/background.svg);
  background-size: cover;

  .login-form {
    margin: 0 auto;
    width: 340px;
    min-height: 20px;
    padding-top: 150px;

    .ant-input {
      border-radius: 4px;
      line-height: 42px;
      height: 42px;
    }

    .ant-btn {
      width: 100%;
      height: 42px;
      line-height: 42px;
    }
  }

  .logo {
    width: 60px !important;
    margin-top: 10px !important;
    margin-bottom: 10px !important;
    margin-left: 20px !important;
    border: none;
    background-color: transparent;
  }

  .head {
    width: 300px;
    font-size: 30px !important;
    font-weight: 700 !important;
    margin-left: 20px !important;
    line-height: 60px !important;
    margin-top: 10px !important;
    position: absolute !important;
    display: inline-block !important;
    height: 60px !important;
    color: #36b368;
  }

  .desc {
    width: 100% !important;
    text-align: center !important;
    color: gray !important;
    height: 60px !important;
    line-height: 60px !important;
  }

  .forgot {
    float: right;
  }
}
</style>
