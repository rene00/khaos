package api

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/rene00/khaos/internal/khaos"
	"github.com/rene00/khaos/models"
	// "log"
)

func Campaign(router *gin.RouterGroup, conf *khaos.Config) {
	router.GET("/campaign", func(c *gin.Context) {
		authID, ok := c.Get("AuthID")
		if !ok {
			c.AbortWithStatusJSON(400, gin.H{"error": "Failed to get AuthID"})
			return
		}

		campaigns, err := models.GetCampaigns(authID.(uint))
		if err != nil {
			c.AbortWithStatusJSON(400, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, campaigns)
	})
}
