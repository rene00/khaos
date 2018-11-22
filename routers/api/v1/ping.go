package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rene00/khaos/pkg/app"
	"github.com/rene00/khaos/pkg/e"
	"github.com/rene00/khaos/service/ping_service"
)

type PingBody struct {
	Username string `json:"username"`
}

func Ping(c *gin.Context) {
	var pingBody PingBody
	appG := app.Gin{C: c}

	data := make(map[string]interface{})
	data["response"] = "pong"

	c.BindJSON(&pingBody)

	pingService := ping_service.Ping{
		Username: pingBody.Username,
	}

	err := pingService.Add()
	if err != nil {
		appG.Response(http.StatusOK, e.ERROR_PING, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, data)
}
