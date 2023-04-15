package config

import (
	"zero-template/common/confz"

	"github.com/zeromicro/go-zero/rest"
)

type JwtAuth struct {
	AccessSecret string
	AccessExpire int64
}

type Config struct {
	rest.RestConf
	HealthPath string
	JwtAuth    JwtAuth
	Security   confz.SecurityConf
	Doc        struct {
		Enable bool   `json:",default=false"`
		Url    string `json:",default=/docs"`
		File   string `json:",default=docs/swagger.json"`
	}
	Mysql struct {
		Source string
	}
}
