package api

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/rene00/khaos/internal/khaos"
	"github.com/rene00/khaos/models"
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

	router.PUT("/campaign/:campaignID", func(c *gin.Context) {
		authID, ok := c.Get("AuthID")
		if !ok {
			c.AbortWithStatusJSON(400, gin.H{"error": "Failed to get AuthID"})
			return
		}

		campaignID, err := strconv.ParseUint(c.Param("campaignID"), 10, 32)
		if err != nil {
			c.AbortWithStatusJSON(400, gin.H{"error": "Bad campaign id"})
			return
		}

		campaign, err := models.GetCampaign(uint(campaignID))
		if err != nil {
			c.AbortWithStatusJSON(400, gin.H{"error": err.Error()})
			return
		}

		if campaign.AuthID != authID {
			c.AbortWithStatusJSON(401, gin.H{"error": "Not owner of campaign"})
			return
		}

		if len(c.Query("enabled")) > 0 {
			enabled, err := strconv.ParseBool(c.Query("enabled"))
			if err != nil {
				c.AbortWithStatusJSON(400, gin.H{"error": "Enabled must be true or false"})
				return
			}
			campaign.Enabled = enabled
		}

		if len(c.Query("status")) > 0 {
			status := c.Query("status")
			campaignStatus, err := models.GetCampaignStatusByName(status)
			if err != nil {
				c.AbortWithStatusJSON(400, gin.H{"error": "Status not supported"})
				return
			}
			campaign.CampaignStatusID = campaignStatus.ID
			campaign.CampaignStatus = campaignStatus
		}

		if _, err = campaign.Save(); err != nil {
			c.AbortWithStatusJSON(400, gin.H{"error": "Failed to update campaign"})
			return
		}

		// Re-retrieve the campaign from the datastore so it can be
		// returned in its most up-to-date state.
		campaign, _ = models.GetCampaign(uint(campaignID))

		c.JSON(http.StatusOK, campaign)
	})
}
