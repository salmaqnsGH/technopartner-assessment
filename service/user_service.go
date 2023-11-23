package service

import (
	"context"
	"technopartner/test/model/web"
)

type UserService interface {
	Create(ctx context.Context, req web.UserCreateRequest) web.UserResponse
	Login(ctx context.Context, req web.UserLoginRequest) web.UserResponse
}
