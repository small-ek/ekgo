package response

import "log"

const (
	ERROR   = 403
	SUCCESS = 200
)

//返回
type Write struct {
	Code  int         `json:"code"`
	Msg   string      `json:"msg"`
	Data  interface{} `json:"data"`
	Error string      `json:"error"`
}

//分页参数
type PageParam struct {
	CurrentPage int         `form:"current_page"`
	PageSize    int         `form:"page_size"`
	Total       int64         `form:"total"`
	Filter      interface{} `form:"filter"`
	Order       string      `form:"order"`
}

//分页返回
type Page struct {
	Total int64         `json:"total"`
	List  interface{} `json:"list"`
}

//错误输出
func ErrorResponse(err error) *Write {
	return &Write{
		Code:  ERROR,
		Msg:   "参数错误",
		Error: err.Error(),
	}
}

//成功返回
func Success(msg string, data ...interface{}) *Write {
	var lenData = len(data)

	if lenData == 1 {
		return &Write{Code: SUCCESS, Msg: msg, Data: data[0]}
	} else if lenData > 1 {
		return &Write{Code: SUCCESS, Msg: msg, Data: data}
	}

	return &Write{Code: SUCCESS, Msg: msg}
}

//错误返回,第二个参数传参返回给前端并会打印
func Fail(msg string, err ...interface{}) *Write {

	if len(err) > 0 {
		log.Println(err[0])
		return &Write{Code: ERROR, Msg: msg, Error: err[0].(string), Data: ""}
	}

	return &Write{Code: ERROR, Msg: msg, Data: ""}
}
