package models

import (
	"encoding/base64"
	"log"
	"os"
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
	Image64          string         `gorm:"-"` // Image field to send to client, does not exist in DB. Base64
	PictureSmall     string         `json:"-"` // Don't send image link to client, only for internal use
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

	// Fail fast
	if err != nil {
		return weapons, err
	}

	// Do partial match on primary match: name,ougi_desc,weapon_skill if present
	if filters["name"] != nil {
		dbctx = dbctx.Where("name LIKE ?", "%"+filters["name"].(string)+"%")
		delete(filters, "name")
	} else if filters["ca_desc"] != nil {
		dbctx = dbctx.Where("ca_desc LIKE ?", "%"+filters["ca_desc"].(string)+"%")
		delete(filters, "ca_desc")
	} else if filters["skill"] != nil {
		dbctx = dbctx.Where("id in (select weapon_id as id from weapon_skills_relationship where weapon_skill_id in (select id as weapon_skill_id from weapon_skills where name like ? or description like ?))", "%"+filters["skill"].(string)+"%", "%"+filters["skill"].(string)+"%")
		delete(filters, "skill")
	}

	// DB Query
	err = dbctx.Where(filters).Limit(20).Offset(20 * pagination).Find(&weapons).Error
	if err != nil {
		return weapons, err
	}

	for i := range weapons {
		err = fetchImageInternal(&weapons[i])
		if err != nil {
			return weapons, err
		}
	}

	return weapons, err
}

func GetWeaponByID(id string) (Weapon, error) {
	var weapon Weapon
	var err error = nil

	err = DB.Preload("Skills").Table("Weapons").Where("name = ?", id).Find(&weapon).Error
	if err != nil {
		return weapon, err
	}

	err = fetchImageInternal(&weapon)

	return weapon, err
}

func fetchImageInternal(weapon *Weapon) error {
	var err error = nil

	// Get images based on path
	if weapon.PictureSmall == "" {
		// I don't have all the images so return without doing anything
		return nil
	}

	imageFile, err := os.ReadFile(os.Getenv("IMAGE_DIR") + weapon.PictureSmall)
	if err != nil {
		log.Print(err.Error())
		return err
	}

	// Convert to base64
	var base64Image string = base64.StdEncoding.EncodeToString(imageFile)

	// Attach to returnable result
	weapon.Image64 = base64Image

	return nil
}
