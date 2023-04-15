package main

import (
	"flag"

	"zero-template/api/internal/config"
	"zero-template/api/internal/handler"
	"zero-template/api/internal/svc"
	"zero-template/common/confz"
	"zero-template/common/restz"
	"zero-template/common/swag"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/api.yaml", "the config file")
var showInfo = flag.Bool("i", false, "show info")

//	@title			API文档
//	@version		1.0
//	@description	API仅当 HTTP Status 为 200 OK 时访问成功，访问成功时的响应格式为 {\"code\":200, \"msg\":\"\", \"data\":{}}, code 为 200 表示业务请求成功，此时可以访问 data 字段（如果 API 本身没有响应，那么 data 字段不存在）；code 不为 200 表示业务请求失败，此时没有 data 字段。<br/><br/>
//	@description	API 应当包含以下请求头
//	@description	<ul>
//	@description	<li> client: 客户端类型 android/ios/web </li>
//	@description	</ul>

//	@securityDefinitions.apiKey	ApiKeyAuth
//	@in							header
//	@name						Authorization
//	@description				用户端API Key
//	@authorizationurl			/api/v1/auth/firebase

func main() {
	flag.Parse()

	var c config.Config
	confz.MustLoad(*configFile, &c)

	if *showInfo {
		logx.Infow("config loaded\n", logx.Field("config", c))
		return
	}

	var opts []rest.RunOption
	if c.Doc.Enable {
		opts = append(opts, swag.WithSwaggerHandler(c.Doc.File, c.Doc.Url))
	}
	server := rest.MustNewServer(c.RestConf, opts...)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)
	restz.AddHealthCheck(server, c.HealthPath)

	logx.Infof("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
