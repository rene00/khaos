package api

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/rene00/khaos/internal/khaos"
	"github.com/rene00/khaos/models"
)

func Attack(router *gin.RouterGroup, conf *khaos.Config) {
	router.GET("/attack/:attackID", func(c *gin.Context) {
		attackID, _ := strconv.ParseUint(c.Param("attackID"), 10, 32)
		attacks, err := models.GetAttack(uint(attackID))
		if err != nil {
			c.AbortWithStatusJSON(404, gin.H{"error": "attack not found"})
			return
		}

		c.JSON(http.StatusOK, attacks)
	})
}
