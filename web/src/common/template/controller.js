import {config} from "../config";

export const controllerTemplate = (name, humpTable) => {
    return `package ` + name + `

import (
\t"` + config.projectName + `/api/app/model"
\t"` + config.projectName + `/api/app/service"
\t"` + config.projectName + `/api/boot/db"
\t"` + config.projectName + `/api/lib"
\t"github.com/gin-gonic/gin"
)

//接口服务
var Interface service.` + humpTable + `Interface

//分页
func Index(this *gin.Context) {
\tvar param = response.PageParam{CurrentPage: 1, PageSize: 10}
\tthis.ShouldBindQuery(&param)
\tparam.Filter = this.QueryArray("filter[]")
\tInterface = &service.` + humpTable + `{PageParam: param, Db: db.Master}

\tthis.SecureJSON(200, Interface.Index())

}

//查询
func Read(this *gin.Context) {
\tvar data = model.` + humpTable + `{}
\tthis.ShouldBindUri(&data)
\tInterface = &service.` + humpTable + `{Model: data, Db: db.Master}

\tthis.SecureJSON(200, Interface.Show())
}

//创建
func Store(this *gin.Context) {
\tvar data = model.` + humpTable + `{}
\tthis.ShouldBind(&data)
\tInterface = &service.` + humpTable + `{Model: data, Db: db.Master}

\tthis.SecureJSON(200, Interface.Store())
}

//修改
func Update(this *gin.Context) {
\tvar data = model.` + humpTable + `{}
\tthis.ShouldBind(&data)
\tInterface = &service.` + humpTable + `{Model: data, Db: db.Master}

\tthis.SecureJSON(200, Interface.Update())
}

//删除
func Delete(this *gin.Context) {
\tvar data = model.` + humpTable + `{}
\tthis.ShouldBindUri(&data)
\tInterface = &service.` + humpTable + `{Model: data, Db: db.Master}

\tthis.SecureJSON(200, Interface.Delete())
}
`
}
