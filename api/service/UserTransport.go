package service

import (
	"context"
	"ekgo/api/ek/conv"
	"encoding/json"
	"errors"
	"net/http"
)

func DecodeUser(this context.Context, h *http.Request) (interface{}, error) {
	var uid = h.URL.Query().Get("uid")
	if uid != "" {
		return UserRequest{Uid: conv.Int(uid)}, nil
	}
	return nil, errors.New("参数不正确")
}

func EncodeUser(this context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
