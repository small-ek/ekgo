package test

import (
	"bytes"
	"ekgo/api/lib/conv"
	"encoding/gob"
	"fmt"
	"log"
	"testing"
)

func TestSum(t *testing.T) {

	var test int = 123
	var total = conv.String(test)
	fmt.Println(total)
}

func TestBytes(t *testing.T) {
	var data int = 12112
	var test = conv.Bytes(data)
	log.Println(test)
	log.Println(conv.Int(test))
}

type User struct {
	Test1 string `json:"test1"`
	Test2 string `json:"test2"`
}

func TestStruct(t *testing.T) {
	var user = User{Test1: "99999999999"}
	var params2 = map[string]string{"test1": "123", "test2": "456"}
	conv.Struct(&user, params2)
	log.Println(user)
	log.Println(user.Test1)
}

type P struct {
	X, Y, Z int
	Name    string
}

type Q struct {
	X, Y int32
	Name string
}

func TestGob(t *testing.T) {
	var network bytes.Buffer
	enc := gob.NewEncoder(&network)
	dec := gob.NewDecoder(&network)
	// Encode (send) the value.
	err := enc.Encode(P{3, 4, 5, "Pythagoras"})
	/*var test = map[string]interface{}{"1221": "1221"}
	err := enc.Encode(test)*/
	if err != nil {
		log.Fatal("encode error:", err)
	}
	// Decode (receive) the value.
	var q Q
	err = dec.Decode(&q)
	if err != nil {
		log.Fatal("decode error:", err)
	}
	log.Println(q)
	log.Println(q.Name)
	log.Println(q.X)
	log.Println(q.Y)
	/*fmt.Printf("%q: {%d,%d}\n", q.Name, *q.X, *q.Y)*/
}
