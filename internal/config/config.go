package config

import (
	"github.com/zeromicro/go-zero/rest"
)

type Config struct {
	rest.RestConf

	Mongo struct {
		DataSource string
		DataBase   string
	}

	Redis struct {
		Addr string
	}

	JwtAuth struct {
		AccessSecret string
	}
}
