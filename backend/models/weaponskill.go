package models

type WeaponSkill struct {
	ID                  int `json:"-" gorm:"primaryKey"`
	Name                string
	Description         string
	BoostType           string
	StatAffected        string
	SkillPercLvlOne     string
	SkillPercLvlTen     string
	SkillPercLvlFifteen string
	SkillPercLvlTwenty  string
}
