package lib

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

//post请求
func Post(url string, data map[string]interface{}, header map[string]string) []byte {
	configdata, _ := json.Marshal(data)
	post_data := bytes.NewBuffer(configdata)
	req, _ := http.NewRequest("POST", url, post_data)
	client := &http.Client{}

	req.Header.Set("Content-Type", "application/json")
	if len(header) > 0 {
		for k, v := range header {
			req.Header.Set(k, v)
		}
	}

	resp, _ := client.Do(req)
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println(body)
	if err != nil {
		log.Print(url + "请求失败" + err.Error())
	}
	return body
}

//post请求
func Delete(url string, data map[string]interface{}, header map[string]string) []byte {
	configdata, _ := json.Marshal(data)
	post_data := bytes.NewBuffer(configdata)
	req, _ := http.NewRequest("DELETE", url, post_data)
	client := &http.Client{}

	req.Header.Set("Content-Type", "application/json")
	if len(header) > 0 {
		for k, v := range header {
			req.Header.Set(k, v)
		}
	}

	resp, _ := client.Do(req)
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println(body)
	if err != nil {
		log.Print(url + "请求失败" + err.Error())
	}
	return body
}

//put请求
func Put(url string, data map[string]interface{}, header map[string]string) []byte {
	configdata, _ := json.Marshal(data)
	post_data := bytes.NewBuffer(configdata)
	req, _ := http.NewRequest("PUT", url, post_data)
	client := &http.Client{}

	req.Header.Set("Content-Type", "application/json")
	if len(header) > 0 {
		for k, v := range header {
			req.Header.Set(k, v)
		}
	}

	resp, _ := client.Do(req)
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println(body)
	if err != nil {
		log.Print(url + "请求失败" + err.Error())
	}
	return body
}

//get请求
func Get(url string, header map[string]string) []byte {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	if len(header) > 0 {
		for k, v := range header {
			req.Header.Set(k, v)
		}
	}
	resp, _ := client.Do(req)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Print(url + "请求失败" + err.Error())
	}
	return body
}
