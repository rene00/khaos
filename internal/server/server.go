package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/rene00/khaos/internal/khaos"
	"github.com/rene00/khaos/models"
)

func Start(conf *khaos.Config) {

	models.Setup(conf)

	if conf.ServerMode != "" {
		gin.SetMode(conf.ServerMode)
	} else if conf.Debug == false {
		gin.SetMode(gin.ReleaseMode)
	}
	app := gin.Default()
	registerRoutes(app, conf)
	app.Run(fmt.Sprintf("%s:%d", conf.ServerIP, conf.ServerPort))
}
