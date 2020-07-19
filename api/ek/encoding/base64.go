package encoding

import (
	"encoding/base64"
	"log"
)

//Encode bytes using BASE64 algorithm
func Encode(str [] byte) string {
	encodeString := base64.StdEncoding.EncodeToString(str)
	return encodeString
}

//Use BASE64 algorithm to decode string
func Decode(str string) string {
	decodeBytes, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		log.Print("base64解密失败" + err.Error())
	}
	return string(decodeBytes)
}
