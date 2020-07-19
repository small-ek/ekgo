package test

import (
	"log"
	"testing"
)

type Btn struct {
	Name      string            `json:"name"`
	Type      string            `json:"type"`
	Url       string            `json:"url,omitempty"`
	SubButton []Btn             `json:"sub_button,omitempty"` //值为空时 直接忽略
	UnShow    string            `json"-"`                     //忽略字段
	Test3     map[string]string `json:"test_3"`
}

type Menu struct {
	Button []Btn `json:"button"`
}

func TestStructs(t *testing.T) {
	var test3 = map[string]string{"name": "name"}
	var SubButton = []Btn{
		{Name: "a1", Type: "view", Url: "https://www.qq.com/a1"},
		{Name: "a2", Type: "view", Url: "https://www.qq.com/a2"},
		{Name: "a3", Type: "view", Url: "https://www.qq.com/a3"},
	}

	var Btn = []Btn{
		{Name: "home", Type: "view", Url: "https://www.qq.com/auth", Test3: test3},
		{Name: "tool", SubButton: SubButton},
		{Name: "other", SubButton: SubButton},
	}
	jsonData := &Menu{
		Button: Btn,
	}

	log.Println(jsonData)

}

type A struct {
	Name string `json:"name"`
	B    []B    `json:"name"`
}
type B struct {
	Name string `json:"name"`
	C    []C    `json:"test_3"`
}
type C struct {
	Name map[string]string `json:"name"`
}
type Page struct {
	Total int         `json:"total"`
	Data  interface{} `json:"data"`
}

func TestStructs2(t *testing.T) {
	var C = []C{
		{Name: map[string]string{"111": "test"}},
	}
	var B = []B{
		{Name: "123", C: C},
		{Name: "456"},
		{Name: "789"},
	}

	var test = []A{
		{Name: "22222222222", B: B},
	}

	var page = &Page{
		Total: 0,
		Data:  test,
	}
	/*rows, _ := db.Query("select * from test where id = ?", 2)
	for rows.Next() {
		line := data{}
		err = rows.Scan(&test)
		log.Println(line)
	}*/
	log.Println(page)
}
