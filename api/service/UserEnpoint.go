package service

import (
	"context"
	"github.com/go-kit/kit/endpoint"
)

type UserRequest struct {
	Uid int `json:"uid"`
}
type UserResponse struct {
	Result string `json:"result"`
}

func GenUserEnpoint(userService UserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		var r = request.(UserRequest)
		var result = userService.GetName(r.Uid)
		return UserResponse{Result: result}, nil
	}
}
