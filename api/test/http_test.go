package test

import (
	"ekgo/api/ek/http"
	"log"
	"testing"
)

func TestHttp(t *testing.T) {
	log.SetFlags(log.Llongfile | log.LstdFlags)
	var http = http.NewClient("http://127.0.0.1:95/")
	if result, err := http.Get(); err == nil {
		log.Println(string(result))
	}
}
