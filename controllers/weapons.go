package controllers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"miiwo/skyfarer/backend"
)

type weaponJSONMap struct {
	ID			string	`json:"id"`
	Name		string	`json:"name"`
	Element		string	`json:"element"`
	WeaponType 	string	`json:"type"`
}

func FetchWeapons(gctx *gin.Context) {
	// Query Parameters
	
	wepname := gctx.Query("name")
	//queryparams := gctx.Request.URL.Query()

	// Grab from backend
	var backendresults []backend.WeaponMemory
	if wepname != "" {
		// If name is specified, only search by it
		backendresults = backend.RetrieveWeaponsByPartialName(wepname)
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