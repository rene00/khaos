package server

import (
	"github.com/gin-gonic/gin"

	"github.com/rene00/khaos/internal/api"
	"github.com/rene00/khaos/internal/khaos"
	"github.com/rene00/khaos/middleware/jwt"
	"github.com/rene00/khaos/middleware/setauthid"
)

func registerRoutes(app *gin.Engine, conf *khaos.Config) {

	// Authentication endpoint
	api.Auth(app, conf)

	// JSON-REST API Version 1
	v1 := app.Group("/api/v1")
	v1.Use(jwt.JWT())
	v1.Use(setauthid.SetAuthID())
	{
		api.Ping(v1, conf)
		api.Inventory(v1, conf)
		api.Campaign(v1, conf)
		api.Attack(v1, conf)
		api.Resource(v1, conf)
		api.ResourceType(v1, conf)
	}
}
