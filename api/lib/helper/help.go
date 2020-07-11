package helper

import (
	"math/rand"
	"net/url"
	"reflect"
	"time"
)

//map[string][]string转map对象
func PostToMap(str url.Values) map[string]string {
	maps := make(map[string]string)
	for key, val := range str {
		for _, subVal := range val {
			maps[key] = subVal
		}
	}
	return maps
}

//RandStringRunes 返回随机字符串
func RandomString(n int) string {
	var letterRunes = []rune("1234567890abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	rand.Seed(time.Now().UnixNano())
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

//截取字符串
func Substring(source string, start, end int) string {
	rs := []rune(source)
	length := len(rs)
	if start < 0 {
		start = 0
	}
	if end > length {
		end = length
	}
	return string(rs[start:end])
}

//搜索切片索引
func SearchSliceInt(array []int, value int) int {
	for k, v := range array {
		if v == value {
			return k
		}
	}
	return -1
}

//搜索切片索引
func SearchSliceString(array []string, value string) int {
	for k, v := range array {
		if v == value {
			return k
		}
	}
	return -1
}

//更具索引删除切片
func DeleteSliceInt(array []int, index int) []int {
	array = append(array[:index], array[index+1:]...)
	return array
}

//更具索引删除切片
func DeleteSliceString(array []string, index int) []string {
	array = append(array[:index], array[index+1:]...)
	return array
}

//过滤切片最大数字
func FilterMaxNumberUint64(str []uint64, max uint64) []uint64 {
	var result []uint64
	for _, v := range str {
		if v < max {
			result = append(result, v)
		}
	}
	return result
}

//struct数组转Map数组
func StructToListStrAny(array interface{}) []map[string]interface{} {
	var result []map[string]interface{}
	model := reflect.TypeOf(array)
	arrayValue := reflect.ValueOf(array)
	fileNum := model.Elem().NumField()
	count := arrayValue.Len()

	for i := 0; i < count; i++ {
		value_m := map[string]interface{}{
		}
		for j := 0; j < fileNum; j++ {
			value := arrayValue.Index(i).Field(j).Interface()
			value_m[model.Elem().Field(j).Name] = value
		}
		result = append(result, value_m)
	}
	return result
}

//struct输入数组中某个单一列的值
func ArrayColumn(array interface{}, column string) []interface{} {
	var result []interface{}
	model := reflect.TypeOf(array)
	if model.Kind() == reflect.Ptr {
		model = model.Elem()
	}
	arrayValue := reflect.ValueOf(array)
	//来获取其包含的字段的总数
	fileNum := model.Elem().NumField()
	count := arrayValue.Len()

	for i := 0; i < count; i++ {
		for j := 0; j < fileNum; j++ {
			//从0开始获取Student所包含的key
			if model.Elem().Field(j).Name == column {
				value := arrayValue.Index(i).Field(j).Interface()
				result = append(result, value)
			}
		}
	}
	return result
}
