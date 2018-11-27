package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/rene00/khaos/pkg/app"
	"github.com/rene00/khaos/pkg/e"
	"github.com/rene00/khaos/pkg/util"
	"github.com/rene00/khaos/service/ping_service"
	"net/http"
	"regexp"
)

// PingError is the Ping Error struct.
type PingError struct {
	Message string
}

func Ping(c *gin.Context) {
	appG := app.Gin{C: c}

	data := make(map[string]interface{})
	data["response"] = "pong"

	authHeader := c.Request.Header.Get("Authorization")
	authRegex := regexp.MustCompile("^Bearer\\s([0-9a-zA-Z\\.\\_\\-]+)$")
	match := authRegex.FindStringSubmatch(authHeader)

	if len(match) != 2 {
		appG.Response(http.StatusBadRequest, e.ERROR_PING, &PingError{Message: "Failed to parse Authorization header for token"})
		return
	}

	authToken := match[1]
	authTokenData, _ := util.ParseToken(authToken)

	pingService := ping_service.Ping{
		Username: util.DecryptString(authTokenData.Username),
	}

	err := pingService.Add()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_PING, &PingError{Message: "Failed to add ping to datastore"})
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, data)
}
