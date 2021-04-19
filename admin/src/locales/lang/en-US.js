import { getModules } from '@/locales/tool'
import antLocal from 'ant-design-vue/es/locale/en_US'
const requireENContext = require.context('./en-US/', false, /[\w+].(js)$/)

const enUS = getModules(requireENContext)

export default {
  ...enUS,
  antLocal
}
