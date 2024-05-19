package models

type Weapon struct {
	// gorm.Model (only used for populating a struct w/ id, createdat, updatedat, deleted at)
	ID			int				`json:"-" gorm:"primaryKey;"`
	Name		string			`gorm:"unique"`
	Element		string	
	WepType 	string			`json:"WeaponType"`
	Skills		[]*WeaponSkill	`json:"Skills,omitempty" gorm:"many2many:weapon_skills_relationship;"`
}

func GetWeaponsByQuery(filters map[string]interface{}) ([]Weapon, error) {
	var weapons []Weapon
	// Sanitize the filters in calling function. Throw error if something happens in the DB
	err := DB.Preload("Skills").Table("Weapons").Where(filters).Find(&weapons).Error

	return weapons, err
}