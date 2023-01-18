package boot

import (
	"flag"
	"github.com/gin-gonic/gin"
	"net/http"
	g "oj/app/global"
	"oj/app/route"
)

func ServerSetup() {
	config := g.Config.Server
	flag.Parse()

	gin.SetMode(config.Mode)
	routers := route.InitRoute()
	server := &http.Server{
		Addr:    config.Addr(),
		Handler: routers,

		ReadTimeout:       config.GetReadTimeout(),
		ReadHeaderTimeout: 0,
		WriteTimeout:      config.GetWriteTimeout(),
		IdleTimeout:       0,
		MaxHeaderBytes:    1 << 20, // 16mb

	}

	g.Logger.Infof("server running on %s ...", config.Addr())
	g.Logger.Errorf(server.ListenAndServe().Error())
}
