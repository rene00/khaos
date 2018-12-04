package models

import (
	log "github.com/sirupsen/logrus"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"

	"github.com/rene00/khaos/internal/khaos"
	"github.com/rene00/khaos/pkg/util"
)

var db *gorm.DB

type Model struct {
	ID uint `gorm:"primary_key" json:"id"`
}

func Setup(conf *khaos.Config) {
	var err error
	db, err = gorm.Open(conf.DatabaseType, conf.DatabaseURI)
	if err != nil {
		log.Println(err)
	}

	if conf.Debug {
		log.Debug("Setting detailed database logging")
		db.LogMode(true)
	}
	db.SingularTable(true)

	db.AutoMigrate(&Auth{})

	password, _ := util.HashPassword(conf.AdminPassword)
	adminUser := &Auth{
		Username: conf.AdminUsername,
		Password: password,
	}
	if err := db.Where("username = ?", conf.AdminUsername).First(&adminUser).Error; err != nil {
		db.Create(adminUser)
	}

	db.AutoMigrate(&Ping{})

	db.AutoMigrate(&Inventory{})

	db.AutoMigrate(&Resource{})
	db.AutoMigrate(&ResourceType{})
	for _, v := range []string{"AWSInstance", "K8SNode", "K8SPod"} {
		db.Create(&ResourceType{Name: v})
	}

	db.AutoMigrate(&AttackType{})
	platformAttackType := &AttackType{Name: "Platform"}
	db.Create(&platformAttackType)

	db.AutoMigrate(&Attack{})
	terminateInstanceAttack := &Attack{
		Name:       "TerminateInstance",
		AttackType: *platformAttackType,
	}
	rebootInstanceAttack := &Attack{
		Name:       "RebootInstance",
		AttackType: *platformAttackType,
	}
	db.Create(&terminateInstanceAttack)
	db.Create(&rebootInstanceAttack)

	db.AutoMigrate(&Campaign{})

	db.AutoMigrate(&CampaignStatus{})
	for _, v := range []string{"NotStarted", "Started", "Completed"} {
		db.Create(&CampaignStatus{Name: v})
	}

	db.AutoMigrate(&CampaignResult{})
	for _, v := range []string{"Failed", "Success"} {
		db.Create(&CampaignResult{Name: v})
	}

}
