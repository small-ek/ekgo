import { createStore } from 'vuex'

const files = import.meta.globEager("./modules/*.js")
Object.keys(files).forEach((key) => {
  files[key]["default"]['namespaced'] = true
})
export default createStore({
  files,
})
