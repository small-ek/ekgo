package secret

import (
	"github.com/satori/go.uuid"
)

//生成uuid
func GetUUID() string {
	uuid := uuid.NewV4().String()
	return uuid
}
