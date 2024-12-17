package rpc

import (
	"sync"

	"gomall/app/frontend/conf"
	"gomall/rpc_gen/kitex_gen/user/userservice"

	frontendUtils "gomall/app/frontend/utils"

	"github.com/cloudwego/kitex/client"
	consul "github.com/kitex-contrib/registry-consul"
)

var (
	UserClient userservice.Client
	once       sync.Once
)

func InitClient() {
	once.Do(func() {
		initUserClient()
	})
}

func initUserClient() {
	// r, err := consul.NewConsulResolver(conf.GetConf().Hertz.RegistryAddr)
	// frontendUtils.MustHandleError(err)
	// UserClient, err = userservice.NewClient("user", client.WithResolver(r))
	// frontendUtils.MustHandleError(err)
	r, err := consul.NewConsulResolver(conf.GetConf().Hertz.RegistryAddr)
	frontendUtils.MustHandleError(err)
	UserClient, err = userservice.NewClient("user", client.WithResolver(r))
	frontendUtils.MustHandleError(err)
}
