package models

type Ping struct {
	Model
	AuthID int
}

func AddPing(data map[string]interface{}) error {
	var auth Auth
	err := db.Where(&Auth{Username: data["username"].(string)}).First(&auth).Error
	if err != nil {
		return err
	}
	ping := Ping{
		AuthID: auth.ID,
	}
	if err := db.Create(&ping).Error; err != nil {
		return err
	}
	return nil
}
