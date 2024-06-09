package models

import (
	"strconv"

	"gorm.io/gorm"
)

type Weapon struct {
	// gorm.Model (only used for populating a struct w/ id, createdat, updatedat, deleted at)
	ID               int    `json:"-" gorm:"primaryKey;"`
	Name             string `gorm:"unique"`
	Element          string
	WepType          string `json:"WeaponType"`
	WepDesc          string `json:"Description"`
	CaName           string `json:"OugiName"`
	CaDesc           string `json:"OugiDesc"`
	LvlOneAtk        int
	LvlHundredAtk    int
	LvlOnefiftyAtk   int
	LvlTwohundredAtk int
	Skills           []*WeaponSkill `json:"Skills,omitempty" gorm:"many2many:weapon_skills_relationship;"`
}

func GetWeaponsByQuery(filters map[string]interface{}) ([]Weapon, error) {
	var weapons []Weapon
	var pagination int = 0
	var err error = nil
	var dbctx *gorm.DB = DB.Preload("Skills").Table("Weapons")

	// Remove pagination from filters if it exists
	if filters["pagination"] != nil {
		pagination, err = strconv.Atoi(filters["pagination"].(string))
		delete(filters, "pagination")
	}

	if err != nil {
		return weapons, err
	}

	// Do partial match on name if present
	if filters["name"] != nil {
		dbctx = dbctx.Where("name LIKE ?", "%"+filters["name"].(string)+"%")
		delete(filters, "name")
	}

	err = dbctx.Where(filters).Limit(20).Offset(20 * pagination).Find(&weapons).Error

	return weapons, err
}

func GetWeaponByID(id string) (Weapon, error) {
	var weapon Weapon
	var err error = nil

	err = DB.Preload("Skills").Table("Weapons").Where("name = ?", id).Find(&weapon).Error

	return weapon, err
}
