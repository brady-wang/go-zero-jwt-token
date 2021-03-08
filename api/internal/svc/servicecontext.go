package svc

import (
	"github.com/tal-tech/go-zero/rest"
	"user/api/internal/config"
	"user/api/internal/middleware"
)

type ServiceContext struct {
	Config config.Config
	Example rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		Example: middleware.NewExampleMiddleware().Handle,
	}
}
