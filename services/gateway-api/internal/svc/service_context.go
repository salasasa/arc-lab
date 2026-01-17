// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package svc

import (
	"gateway-api/internal/config"

	"auth-rpc/auth_client"

	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config  config.Config
	AuthRpc auth_client.Auth
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:  c,
		AuthRpc: auth_client.NewAuth(zrpc.MustNewClient(c.AuthRpc)),
	}
}
