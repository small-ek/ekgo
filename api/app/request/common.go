package request

//分页请求结构体
//Paging request structure
type Page struct {
	CurrentPage int         `form:"current_page"`
	PageSize    int         `form:"page_size"`
	Total       int         `form:"total"`
	Filter      interface{} `form:"filter"`
	Order       string      `form:"order"`
}
