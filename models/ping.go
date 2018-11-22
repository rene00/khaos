package models

type Ping struct {
	Model
	Username string `json:"username"`
}

func AddPing(data map[string]interface{}) error {
	ping := Ping{
		Username: data["username"].(string),
	}
	if err := db.Create(&ping).Error; err != nil {
		return err
	}
	return nil
}
