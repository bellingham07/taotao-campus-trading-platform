package svc

import (
	"go-go-zero/service/cmdty/cmd/api/internal/config"
)

type ServiceContext struct {
	Config config.Config
	//auth   rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		//auth:   middleware.NewAuthMiddleware().Handle,
	}
}
