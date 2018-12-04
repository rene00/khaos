package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rene00/khaos/internal/khaos"
	"github.com/rene00/khaos/models"
)

func Ping(router *gin.RouterGroup, conf *khaos.Config) {
	router.POST("/ping", func(c *gin.Context) {
		authID, ok := c.Get("AuthID")
		if !ok {
			c.AbortWithStatusJSON(400, gin.H{"error": "Failed to get AuthID"})
			return
		}

		ping := models.Ping{AuthID: authID.(uint)}
		if err := ping.Add(); err != nil {
			c.AbortWithStatusJSON(400, gin.H{"error": "Failed to add Ping event"})
			return
		}

		c.JSON(http.StatusOK, map[string]string{"message": "pong"})
	})
}
