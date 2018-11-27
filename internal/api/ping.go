package api

import (
	"github.com/gin-gonic/gin"
	"github.com/rene00/khaos/internal/khaos"
	"github.com/rene00/khaos/models"
	"github.com/rene00/khaos/pkg/util"
	"net/http"
	"regexp"
)

func Ping(router *gin.RouterGroup, conf *khaos.Config) {
	router.POST("/ping", func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		authRegex := regexp.MustCompile("^Bearer\\s([0-9a-zA-Z\\.\\_\\-]+)$")
		match := authRegex.FindStringSubmatch(authHeader)

		if len(match) != 2 {
			c.AbortWithStatusJSON(400, gin.H{"error": "Failed to parse token"})
			return
		}

		authToken := match[1]
		authTokenData, _ := util.ParseToken(authToken)

		if err := models.AddPing(util.DecryptString(authTokenData.Username)); err != nil {
			c.AbortWithStatusJSON(400, gin.H{"error": "Failed to add ping event"})
			return
		}

		c.JSON(http.StatusOK, http.Response{})
	})
}
