package backend

/*import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)*/
import (
	"strings"
)

type WeaponMemory struct {
	ID			string
	Name		string
	Element		string
	WeaponType 	string
}

var weapons = []WeaponMemory {
	{ID: "1", Name: "Bahamut Blade", Element: "Dark", WeaponType: "katana"},
	{ID: "2", Name: "Knight of Ice", Element: "Water", WeaponType: "dagger"},
	{ID: "3", Name: "Phoenix's Torch", Element: "Fire", WeaponType: "staff"},
}

// Retrieve at max 10 from the database
func RetrieveWeapons() []WeaponMemory {
	return weapons[:min(10, len(weapons))]
}

// Retrieve at max 10 from the database
func RetrieveWeaponsMax(max int) []WeaponMemory {
	if max > len(weapons) {
		return weapons
	} else {
		return weapons[:max]
	}
}

func RetrieveWeaponsByPartialName(name string) []WeaponMemory {
	results := make([]WeaponMemory, 0)

	// Search backend
	for i := range weapons {
		if strings.Contains(strings.ToLower(weapons[i].Name), name) {
			results = append(results, weapons[i])
		}
	}


	return results
}

/*type Weapon struct {
	ID			string	`gorm:"primaryKey"`
	Name		string	`gorm:"unique"`
	Element		string	
	WeaponType 	string	
}

var weapons = []weapon {
	{ID: "1", Name: "Bahamut Blade", Element: "Dark", WeaponType: "katana"},
	{ID: "2", Name: "Knight of Ice", Element: "Water", WeaponType: "dagger"},
	{ID: "3", Name: "Phoenix's Torch", Element: "Fire", WeaponType: "staff"},
}

func connectToMariaDB() (*gorm.DB, error) {
	dburl := "user:pass@tcp(localhost:8084)/dbname?charset=utf8mb4*parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dburl), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	return db, nil
}*/

/*func readWeapon(db *gorm.DB, name string) (*Weapon, error) {
	var weapon Weapon
	result := db.First(&weapon, name)

	if result.Error != nil {
		return nil, result.Error
	}

	return &weapon, nil
}*/
