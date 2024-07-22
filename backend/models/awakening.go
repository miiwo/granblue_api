package models

type Awakening struct {
	ID           int `json:"-" gorm:"primaryKey"`
	Name         string
	StatAffected string
	Strength     string
}
