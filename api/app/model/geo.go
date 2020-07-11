package model

//地理表
type Geo struct {
	ID           int    `json:"id" uri:"id"`   //编码
	Pid          int    `json:"pid"`           //上级编码
	Deep         int    `json:"deep"`          //深度
	Name         string `json:"name"`          //名称
	PinyinPrefix string `json:"pinyin_prefix"` //拼音首字母
	Pinyin       string `json:"pinyin"`        //拼音
	ExtId        int64  `json:"ext_id"`        //编码
	ExtName      string `json:"ext_name"`      //全称
}

func (Geo) TableName() string {
	return "s_geo"
}
