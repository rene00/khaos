package models

import (
	"errors"
	"time"
)

type Campaign struct {
	Model

	AuthID uint `sql:"type:integer REFERENCES auth(id)" gorm:"not null"`

	AttackID uint `sql:"type:integer REFERENCES attack(id)" gorm:"not null"`
	Attack   Attack

	ResourceID uint `sql:"type:integer REFERENCES resource(id)" gorm:"not null"`
	Resource   Resource

	CampaignStatusID uint `sql:"type:integer REFERENCES campaign_status(id)" gorm:"not null"`
	CampaignStatus   CampaignStatus

	CampaignResultID uint `sql:"type:integer REFERENCES campaign_result(id)"`
	CampaignResult   CampaignResult

	// BeginDate
	BeginDate *time.Time

	// EndDate
	EndDate *time.Time

	// Enabled
	Enabled bool `gorm:"default:false"`
}

type CampaignStatus struct {
	Model
	Name string `gorm:"unique;not null"`
}

type CampaignResult struct {
	Model
	Name string `gorm:"unique;not null"`
}

// Save the campaign to the datastore.
func (c *Campaign) Save() (*Campaign, error) {
	if err := db.Save(c).Error; err != nil {
		return c, err
	}
	return c, nil
}

func GetCampaign(campaignID uint) (Campaign, error) {
	var campaign Campaign
	if err := db.Where("id = ?", campaignID).First(&campaign).Error; err != nil {
		return campaign, errors.New("Campaign does not exist")
	}

	// Get Attack
	attack, err := GetAttack(campaign.AttackID)
	if err != nil {
		return campaign, err
	}
	campaign.Attack = attack

	// Get Resource
	resource, err := GetResource(campaign.ResourceID)
	if err != nil {
		return campaign, err
	}
	campaign.Resource = resource

	// Get Status
	campaignStatus, err := GetCampaignStatus(campaign.CampaignStatusID)
	if err != nil {
		return campaign, err
	}
	campaign.CampaignStatus = campaignStatus

	// Get Result
	campaignResult, err := GetCampaignResult(campaign.CampaignResultID)
	if err != nil {
		return campaign, err
	}
	campaign.CampaignResult = campaignResult

	return campaign, nil
}

func GetCampaigns(authID uint) ([]Campaign, error) {
	var campaigns []Campaign
	db.Where(&Campaign{AuthID: authID}).Find(&campaigns)
	c := campaigns[:0]
	for _, campaign := range campaigns {
		campaign, err := GetCampaign(campaign.ID)
		if err != nil {
			return campaigns, err
		}
		c = append(c, campaign)
	}
	return c, nil
}

func GetCampaignStatusByName(campaignStatusName string) (CampaignStatus, error) {
	var campaignStatus CampaignStatus
	err := db.Where("name = ?", campaignStatusName).First(&campaignStatus).Error
	if err != nil {
		return campaignStatus, errors.New("Failed getting campaign status")
	}
	return campaignStatus, nil
}

func GetCampaignStatus(campaignStatusID uint) (CampaignStatus, error) {
	var campaignStatus CampaignStatus
	err := db.Where("id = ?", campaignStatusID).First(&campaignStatus).Error
	if err != nil {
		return campaignStatus, errors.New("Failed getting campaign status")
	}
	return campaignStatus, nil
}

func GetCampaignResult(campaignResultID uint) (CampaignResult, error) {
	var campaignResult CampaignResult
	err := db.Where("id = ?", campaignResultID).First(&campaignResult).Error
	if err != nil {
		return campaignResult, errors.New("Failed getting campaign result")
	}
	return campaignResult, nil
}

func (c *Campaign) EnableDisable(authID uint, enabled bool) error {
	c.Enabled = enabled
	if err := db.Update(c).Error; err != nil {
		return err
	}
	return nil
}
