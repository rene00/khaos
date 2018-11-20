package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rene00/khaos/pkg/app"
	"github.com/rene00/khaos/pkg/e"
)

func GetAttacks(c *gin.Context) {
	appG := app.Gin{C: c}
	data := make(map[string]interface{})
	data["attacks"] = "test"
	appG.Response(http.StatusOK, e.SUCCESS, data)
}
