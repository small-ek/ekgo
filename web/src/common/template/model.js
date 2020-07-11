import {convHump, convTable} from "../template";

export const modelTemplate = (data, humpTable, table_name, table_comment) => {
    var field = ''
    var onImport = '';//导入

    data.forEach(function (value) {
        //小写转驼峰
        let hump = convHump(value.column_name)
        var table_name = convTable(value.table_name)

        if (table_name != undefined) {
            var lowercaseTableName = table_name.toLowerCase();
        }

        //数据类型转换
        let data_type = switchDataType(value.data_type)
        //条件判断追加
        if (value.column_name == "id" && value.association_table == undefined) {
            field = field + '\t' + hump + '              ' + data_type + '        `gorm:"PRIMARY_KEY" json:"id" uri:"id"` //标识\n'
        } else if (value.column_name == "deleted_at") {//删除时间
            onImport = '\t"time"\n'
            data_type = "*time.Time"
            field = field + '\t' + hump + '              ' + data_type + '        `json:"deleted_at"` //' + value.column_comment + '\n'
        } else if (value.association_table == true && value.relation == "has_one") {//一对一关联表
            field = field + '\t' + table_name + '         ' + table_name + '      `json:"' + lowercaseTableName + '" gorm:"FOREIGNKEY:' + value.column_name + ';ASSOCIATION_FOREIGNKEY:' + value.field + '"` //' + value.column_comment + '\n'
        } else if (value.association_table == true && value.relation == "has_many") {//一对多关联表
            field = field + '\t' + table_name + '         []' + table_name + '      `json:"' + lowercaseTableName + '_list" gorm:"FOREIGNKEY:' + value.field + ';ASSOCIATION_FOREIGNKEY:' + value.column_name + '"` //' + value.column_comment + '\n'
        } else {
            field = field + '\t' + hump + '              ' + data_type + '        `json:"' + value.column_name + '"` //' + value.column_comment + '\n'
        }
    })

    return 'package model\n' +
        '\n' +
        'import (\n' +
        '\t"github.com/jinzhu/gorm"\n' +
        onImport +
        ')\n' +
        '\n' +
        '//' + table_comment + '\n' +
        'type ' + humpTable + ' struct {\n' +
        field +
        '}\n' +
        '\n' +
        'func (' + humpTable + ') TableName() string {\n' +
        '\treturn "' + table_name + '"\n' +
        '}\n' +
        '//查询之后\n' +
        'func (this *' + humpTable + ') AfterFind(scope *gorm.Scope) error {\n' +
        '\treturn nil\n' +
        '}\n' +
        '\n' +
        '//创建之前\n' +
        'func (this *' + humpTable + ') BeforeCreate(scope *gorm.Scope) error {\n' +
        '\treturn nil\n' +
        '}\n' +
        '\n' +
        '//创建之后\n' +
        'func (this *' + humpTable + ') AfterCreate(scope *gorm.Scope) error {\n' +
        '\treturn nil\n' +
        '}\n' +
        '\n' +
        '//更新之前\n' +
        'func (this *' + humpTable + ') BeforeUpdate(scope *gorm.Scope) error {\n' +
        '\treturn nil\n' +
        '}\n' +
        '\n' +
        '//更新之后\n' +
        'func (this *' + humpTable + ') AfterUpdate(scope *gorm.Scope) error {\n' +
        '\treturn nil\n' +
        '}\n' +
        '\n' +
        '//删除之前\n' +
        'func (this *' + humpTable + ') BeforeDelete(scope *gorm.Scope) error {\n' +
        '\treturn nil\n' +
        '}\n' +
        '\n' +
        '//删除之后\n' +
        'func (this *' + humpTable + ') AfterDelete(scope *gorm.Scope) error {\n' +
        '\treturn nil\n' +
        '}'
}
/**
 * 数据库转go数据类型
 * @param type
 * @returns {string}
 */
export const switchDataType = (type) => {
    var result = "string";
    switch (type) {
        case "int":
            result = "int"
            break;
        case "smallint":
            result = "int"
            break;
        case "tinyint":
            result = "int"
            break;
        case "bigint":
            result = "int64"
            break;
        case "float":
            result = "float64"
            break;
        case "decimal":
            result = "float64"
            break;
        case "double":
            result = "float64"
            break;
        case "varchar":
            result = "string"
            break;
        case "char":
            result = "string"
            break;
        case "longtext":
            result = "string"
            break;
        case "enum":
            result = "string"
            break;
        case "json":
            result = "Json"
            break;
        case "timestamp":
            result = "*Time"
            break;
        case "datetime":
            result = "*Time"
            break;
        case "date":
            result = "*Time"
            break;
    }
    return result
}