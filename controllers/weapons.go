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
	// Query Parameter
	// wepname := gctx.DefaultQuery("name", "None")
	// ele := gctx.Query("element")

	// Grab from backend
	var backendresults []backend.WeaponMemory = backend.RetrieveWeapons()
	results := make([]weaponJSONMap, 0)

	// Parse to JSON mapping to send back to requester
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