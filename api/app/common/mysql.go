package common

import (
	"ekgo/api/boot/db"
	"ekgo/api/lib/response"
	"github.com/gin-gonic/gin"
)

type TableAll struct {
	TableName      string `json:"table_name"`
	TableRows      int    `json:"table_rows"`
	TableCollation string `json:"table_collation"`
	TableComment   string `json:"table_comment"`
}

//查询所有表名称
func GetAllTable(this *gin.Context) {
	_, _, _, _, _, database, _ := db.GetDbConfig()
	var result = []TableAll{}

	db.Master.Raw("SELECT TABLE_NAME AS table_name,TABLE_ROWS AS table_rows,TABLE_COLLATION AS table_collation,TABLE_COMMENT AS table_comment FROM INFORMATION_SCHEMA.Tables WHERE table_schema = ?", database).Find(&result)
	this.SecureJSON(200, response.Success("成功", result))
}

type Table struct {
	ColumnName             string `json:"column_name"`
	CharacterMaximumLength string `json:"character_maximum_length"`
	DataType               string `json:"data_type"`
	ColumnKey              string `json:"column_key"`
	ColumnComment          string `json:"column_comment"`
}

//查询单表信息
func GetTable(this *gin.Context) {
	_, _, _, _, _, database, _ := db.GetDbConfig()
	var result = []Table{}

	db.Master.Raw("SELECT COLUMN_NAME AS column_name,CHARACTER_MAXIMUM_LENGTH AS character_maximum_length,DATA_TYPE AS data_type,COLUMN_KEY AS column_key,COLUMN_COMMENT AS column_comment FROM `information_schema`.`columns` WHERE TABLE_SCHEMA = ? AND table_name = ? ORDER BY ORDINAL_POSITION", database, this.Param("name")).Find(&result)
	this.SecureJSON(200, response.Success("成功", result))
}
