import { getModules } from '@/locales/tool'
import antLocal from "ant-design-vue/es/locale/zh_CN";

const requireZHContext = require.context('./zh-CN/', false, /[\w+].(js)$/)

const zhCN = getModules(requireZHContext)

export default {
  ...zhCN,
  antLocal
}
