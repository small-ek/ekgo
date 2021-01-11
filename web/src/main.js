import {createApp} from 'vue'
import App from './App.vue'
import router from './router'
import store from './store'
import Antd from 'ant-design-vue'
import {initI18n} from '@/utils/i18n'
import './theme/index.less'
const i18n = initI18n('CN', 'US')
const app = createApp(App).use(router).use(store)

app.use(Antd).use(i18n).mount('#app')