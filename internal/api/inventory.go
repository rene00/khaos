package api

import (
	"encoding/json"
	// "fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/rene00/khaos/internal/khaos"
	"github.com/rene00/khaos/models"
)

func Inventory(router *gin.RouterGroup, conf *khaos.Config) {
	router.POST("/inventory", func(c *gin.Context) {

		// A slice of Inventory structs. The client will be submitting
		// an array of inventory items.
		var i []models.Inventory

		// Get raw body as it will contain an array of objects which
		// gins BindJSON method cant handle.
		// https://github.com/gin-gonic/gin/issues/715
		data, _ := c.GetRawData()

		if err := json.Unmarshal(data, &i); err != nil {
			c.JSON(400, gin.H{"error": "Failed to bind"})
			return
		}

		// Get AuthID from gin context. If it is not set,
		// something went wrong within the SetAuthID middleware
		// so we best handle this gracefully.
		authID, ok := c.Get("AuthID")
		if !ok {
			c.JSON(400, gin.H{"error": "Failed to find AuthID"})
			return
		}

		// Iterate over the inventory items within the slice and
		// add them to the database.
		for _, v := range i {
			v.Add(authID.(int))
		}

		c.JSON(http.StatusOK, i)
	})
}
