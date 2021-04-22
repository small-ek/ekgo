import { createApp } from "vue";
import Antd from 'ant-design-vue';
import App from "./App.vue";
import Router from "./router";
import Store from "./store";

import "ant-design-vue/dist/antd.less";
import "./assets/css/index.less";

const app = createApp(App);
app.config.productionTip = false;
app.use(Antd);
app.use(Store);
app.use(Router);
app.mount("#app");

export default app
