export default {
    //导出
    exportList(header, data, name) {
        require.ensure([], function () {
            const {
                export_json_to_excel
            } = require("./excel/Export2Excel"); //引入文件

            export_json_to_excel(header, data, name);
        });
    },
    formatJson(filterVal, jsonData) {
        return jsonData.map(v => filterVal.map(j => v[j]));
    }
}
