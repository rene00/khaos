package api

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/rene00/khaos/internal/khaos"
	"github.com/rene00/khaos/models"
)

func Resource(router *gin.RouterGroup, conf *khaos.Config) {
	router.POST("/resource", func(c *gin.Context) {
		var resources []models.Resource

		data, _ := c.GetRawData()
		if err := json.Unmarshal(data, &resources); err != nil {
			c.JSON(400, gin.H{"error": "Failed to bind"})
			return
		}

		authID, ok := c.Get("AuthID")
		if !ok {
			c.JSON(400, gin.H{"error": "Failed to find AuthID"})
			return
		}

		for _, resource := range resources {
			resource.Add(authID.(uint))
		}

		c.JSON(http.StatusOK, resources)
	})
}

func ResourceType(router *gin.RouterGroup, conf *khaos.Config) {
	router.GET("/resourcetype", func(c *gin.Context) {
		resourceTypes, _ := models.GetResourceTypes()
		c.JSON(http.StatusOK, resourceTypes)
	})
	router.GET("/resourcetype/:resourceTypeName", func(c *gin.Context) {
		resourceTypeName := c.Param("resourceTypeName")
		resourceType, err := models.GetResourceType(resourceTypeName)
		if err != nil {
			c.AbortWithStatusJSON(404, gin.H{"error": "resource type not found"})
			return
		}
		c.JSON(http.StatusOK, resourceType)
	})
}
