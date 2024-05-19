package models

type WeaponSkill struct {
	ID			int			`json:"-" gorm:"primaryKey"`
	Name 		string
	Description string
}