package server

import (
	"github.com/gin-gonic/gin"

	"github.com/rene00/khaos/internal/api"
	"github.com/rene00/khaos/internal/khaos"
	"github.com/rene00/khaos/middleware/jwt"
)

func registerRoutes(app *gin.Engine, conf *khaos.Config) {

	// Authentication endpoint
	// app.GET("/auth", getAuth)
	// app.GET("/auth", api.Auth)
	api.Auth(app, conf)

	// JSON-REST API Version 1
	v1 := app.Group("/api/v1")
	v1.Use(jwt.JWT())
	{
		api.Ping(v1, conf)
	}
}
