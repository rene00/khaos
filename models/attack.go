package models

type Attack struct {
	Model
	Name         string `gorm:"unique;not null"`
	AttackTypeID uint   `sql:"type:integer REFERENCES attack_type(id)" gorm:"not null"`
	AttackType   AttackType
}

type AttackType struct {
	Model
	Name string `gorm:"unique;not null"`
}

func GetAttack(attackID uint) (Attack, error) {
	var attack Attack
	err := db.Where("id = ?", attackID).First(&attack).Error
	if err != nil {
		return attack, err
	}
	err = db.Where("id = ?", attack.AttackTypeID).First(&attack.AttackType).Error
	if err != nil {
		return attack, err
	}
	return attack, nil
}
