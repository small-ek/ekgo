import {createStore} from 'vuex'

const file = import.meta.globEager("./modules/*.js")
var modules = {}
Object.keys(file).forEach((key) => {
    var fileName = key.split("/")
    var firstFileName = fileName[2].split(".")
    var defaults = file[key]["default"]
    defaults["namespaced"] = true
    modules[firstFileName[0]] = defaults
})
export default createStore({
    modules: modules,
})
