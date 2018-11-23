package models

import (
	"github.com/jinzhu/gorm"
	"github.com/rene00/khaos/pkg/util"
)

type Auth struct {
	Model
	Username string `gorm:"unique;not null" json:"username"`
	Password string `json:"password"`

	Pings []Ping
}

func CheckAuth(username, password string) (bool, error) {
	var auth Auth
	err := db.Where(&Auth{Username: username}).First(&auth).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}

	if true != util.CheckPasswordHash(password, auth.Password) {
		return false, err
	}

	if auth.ID > 0 {
		return true, nil
	}

	return false, nil
}
