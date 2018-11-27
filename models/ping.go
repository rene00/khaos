package models

type Ping struct {
	Model
	AuthID int
}

func AddPing(username string) error {
	var auth Auth
	err := db.Where(&Auth{Username: username}).First(&auth).Error
	if err != nil {
		return err
	}
	if err := db.Create(&Ping{AuthID: auth.ID}).Error; err != nil {
		return err
	}
	return nil
}
