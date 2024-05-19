package handler

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"miiwo/skyfarer/backend"
	"miiwo/skyfarer/backend/models"
)

type weaponJSONMap struct {
	ID			string	`json:"id"`
	Name		string	`json:"name"`
	Element		string	`json:"element"`
	WeaponType 	string	`json:"type"`
}

func FetchWeaponsFromDB(gctx *gin.Context) {
	queryparams := gctx.Request.URL.Query()
	tempfilters := make(map[string]interface{})
	for key, value := range queryparams {
		tempfilters[key] = value[0] // Grabs the first value out of the query parameters if there is duplicates, make it so that there can only be one kind in later iteration
	}
	results, err := models.GetWeaponsByQuery(tempfilters)

	if (err != nil) {
		gctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Server encountered an error"})
	} else {
		gctx.IndentedJSON(http.StatusOK, results)
	}
}

func FetchWeapons(gctx *gin.Context) {
	// Query Parameters
	queryparams := gctx.Request.URL.Query()

	// Grab from backend
	var backendresults []backend.WeaponMemory
	if queryparams.Get("name") != "" {
		// If name is specified, only search by it
		backendresults = backend.RetrieveWeaponsByPartialOrFullName(queryparams.Get("name"))

	} else if len(queryparams) >= 1 {
		// Parse through the rest of the parameters
		// Make an else statement for if an unused query parameter is used
		tempfilters := make(map[string]string)
		for key, value := range queryparams {
			tempfilters[key] = value[0] // Grabs the first value out of the query parameters if there is duplicates
		}
		backendresults = backend.RetrieveWeaponsWithConditions(tempfilters)

	} else {
		// If no other conditions, return all weapons
		backendresults = backend.RetrieveWeapons()
	}

	// Parse to JSON mapping to send back to requester
	results := make([]weaponJSONMap, 0)

	for i := range backendresults {
		results = append(results, remapModel(backendresults[i]))
	}


	// Return result
	gctx.IndentedJSON(http.StatusOK, results)
}

func remapModel(wep backend.WeaponMemory) weaponJSONMap {
	return weaponJSONMap {
		ID: wep.ID, 
		Name: wep.Name, 
		Element: wep.Element, 
		WeaponType: wep.WeaponType,
	}
}