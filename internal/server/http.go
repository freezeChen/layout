package server

import (
	"gitee.com/bethink1501/24on-library/xgin"
	"github.com/freezeChen/layout/internal/conf"
	"github.com/freezeChen/layout/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/transport/http"
)

// NewHTTPServer new a HTTP server.
func NewHTTPServer(c *conf.Server, greeter *service.GreeterService, logger log.Logger) *http.Server {

	var opts = []http.ServerOption{
		http.Middleware(middleware.Chain(
			tracing.Server(),
			logging.Server(logger),
		)),
	}
	router := gin.Default()
	router.Use(func(ctx *gin.Context) {

	}, xgin.Middlewares(
		recovery.Recovery(),
		logging.Server(logger),
	))

	if c.Http.Network != "" {
		opts = append(opts, http.Network(c.Http.Network))
	}
	if c.Http.Addr != "" {
		opts = append(opts, http.Address(c.Http.Addr))
	}
	if c.Http.Timeout != 0 {
		opts = append(opts, http.Timeout(c.Http.Timeout.Duration()))
	}
	srv := http.NewServer(opts...)

	srv.HandlePrefix("/", router)

	//v1.RegisterGreeterHTTPServer(srv, greeter)
	return srv
}
