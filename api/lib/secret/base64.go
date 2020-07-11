package secret

import (
	"encoding/base64"
	"log"
)

//base64加密
func Encode(str [] byte) string {
	encodeString := base64.StdEncoding.EncodeToString(str)
	return encodeString
}

//base64解密
func Decode(str string) string {
	decodeBytes, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		log.Print("base64解密失败" + err.Error())
	}
	return string(decodeBytes)
}
