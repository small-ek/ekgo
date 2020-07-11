import {serviceTemplate} from './template/service'
import {saveAs} from 'file-saver';
import {modelTemplate} from "./template/model";
import {controllerTemplate} from "./template/controller";
import {routeTemplate} from "./template/route";
import {vueTemplate} from "./template/vue";

/**
 * 导出模板文件
 * @param data
 * @param table_name
 * @param type
 */
export const exportTemplate = (data, table_name, table_comment, type) => {
    let fileName = "";

    let template = '';

    var humpTable = convTable(table_name)
    //模板导出类型
    switch (type) {
        case "vue":
            fileName = "index.vue"
            template = vueTemplate(data, table_name, table_comment)
            break;
        case "model":
            fileName = getRidTable(table_name) + ".go"
            template = modelTemplate(data, humpTable, table_name, table_comment)
            break;
        case "service":
            fileName = humpTable + "Service.go"

            template = serviceTemplate(humpTable, data)
            break;
        case "controller":
            fileName = getRidTable(table_name) + ".go"
            template = controllerTemplate(table_name, humpTable)
            break;
        case "route":
            fileName = "route.go"
            template = routeTemplate(table_name, humpTable, table_comment)
            break;
    }
    var blob = new Blob([template], {type: "text/plain;charset=utf-8"});
    saveAs(blob, fileName);
};
//表大小转换
export const convTable = (table_name) => {
    if (table_name) {
        let newTable = getRidTable(table_name)
        //下划线转驼峰
        return convHump(newTable)
    }
    return table_name
}
//去掉s b r开头关键字
export const getRidTable = (table_name) => {
    if (table_name) {
        let str = table_name.split("_")
        let newTable = table_name;
        //s代表系统表前缀,b代表业务前缀,r代表中间表前缀
        if (str[0] == "s" || str[0] == "b" || str[0] == "r") {
            newTable = str.slice(1, str.length)
            newTable = newTable.join('_');
        }
        //下划线转驼峰
        return newTable
    }
    return table_name
}
//小写转换驼峰
export const convHump = (column_name) => {
    let hump = "_" + column_name
    hump = hump.replace(/\_(\w)/g, function (all, letter) {
        return letter.toUpperCase();
    });
    return hump
}