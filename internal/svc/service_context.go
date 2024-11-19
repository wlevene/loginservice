package svc

import (
	"github.com/wlevene/loginservice/internal/config"
	"github.com/wlevene/loginservice/internal/dao"
	"github.com/wlevene/loginservice/internal/middleware"
	"github.com/zeromicro/go-zero/rest"
)

type ServiceContext struct {
	Config         config.Config
	NopeMiddleware rest.Middleware
	AuthMiddleware rest.Middleware
	PeriodLimit    rest.Middleware
	RABCMiddleware rest.Middleware

	Dao *dao.Dao
}

func NewServiceContext(c config.Config) *ServiceContext {
	dao := dao.NewDao(c.Mongo.DataSource, c.Mongo.DataBase)

	return &ServiceContext{
		Config:         c,
		NopeMiddleware: middleware.NewNopeMiddleware().Handle,
		AuthMiddleware: middleware.NewAuthMiddleware().Handle,
		PeriodLimit:    middleware.NewPeriodLimitMiddleware(c).Handle,
		RABCMiddleware: middleware.NewRABCMiddleware().Handle,
		Dao:            dao,
	}
}
