import {createApp} from 'vue'
import App from './App.vue'
import router from './router'
import Antd from 'ant-design-vue';
import 'ant-design-vue/dist/antd.css';
import '@/vab'
import store from './store'

const app = createApp(App).use(store);
app.config.productionTip = false;
app.use(router).use(Antd).mount('#app')
