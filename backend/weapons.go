package backend

/*import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)*/
import (
	"strings"
	"os"
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
	{ID: "4", Name: "Phoenix's Claw", Element: "Fire", WeaponType: "melee"},
	{ID: "5", Name: "Unsigned Kaneshige", Element: "Light", WeaponType: "katana"},
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

func RetrieveWeaponsByPartialOrFullName(name string) []WeaponMemory {
	results := make([]WeaponMemory, 0)

	// Search database
	for i := range weapons {
		if strings.Contains(strings.ToLower(weapons[i].Name), strings.ToLower(name)) {
			results = append(results, weapons[i])
		}
	}


	return results
}

func RetrieveWeaponsWithConditions(filters map[string]string) []WeaponMemory {
	results := make([]WeaponMemory, 0)

	for i := range weapons {
		shouldAdd := true
		for key, value := range filters {
			fieldVal := grabWeaponMemoryKeyValue(weapons[i], key) 
			if strings.ToLower(fieldVal) != value {
				shouldAdd = false
				break
			}
		}

		if shouldAdd {
			results = append(results, weapons[i])
		}
		
	}

	
	return results
}

func ValidateAPIKey(key string) bool {
	return checkKeys(key)
}

func checkKeys(rawKey string) bool {
	isValid := true
	apiKey := os.Getenv("API_KEY")

	if len(apiKey) != len(rawKey) {
		return false
	}

	// Goes through all the characters to try and avoid timing attack
	for i := range apiKey {
		if rawKey[i] != apiKey[i] {
			isValid = false
		}
	}

	return isValid
}

func grabWeaponMemoryKeyValue(item WeaponMemory, key string) string {
	if key == "name" {
		return item.Name
	} else if key == "element" {
		return item.Element
	} else if key == "weptype" {
		return item.WeaponType
	} else {
		return ""
	}
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
