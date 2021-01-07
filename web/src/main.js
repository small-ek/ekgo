import {createApp} from 'vue'
import App from './App.vue'
import router from './router'
import store from './store'
import Antd from 'ant-design-vue'
import './theme/index.less'

const app = createApp(App).use(router).use(store)

app.use(Antd).mount('#app')