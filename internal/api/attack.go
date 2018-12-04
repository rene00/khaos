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
			c.JSON(404, gin.H{"message": "No attack found"})
			return
		}

		c.JSON(http.StatusOK, attacks)
	})
}
