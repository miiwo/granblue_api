package models

type WeaponSkill struct {
	ID                  int `json:"-" gorm:"primaryKey"`
	Name                string
	Description         string
	BoostType           string
	StatAffected        string
	SkillPercLvlOne     string `json:"SkillLvlOne"`
	SkillPercLvlTen     string `json:"SkillLvlTen"`
	SkillPercLvlFifteen string `json:"SkillLvlFifteen"`
	SkillPercLvlTwenty  string `json:"SkillLvlTwenty"`
}
