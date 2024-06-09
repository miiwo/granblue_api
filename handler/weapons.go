package handler

import (
	"miiwo/skyfarer/backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func FetchWeaponsFromDB(gctx *gin.Context) {
	queryparams := gctx.Request.URL.Query()
	tempfilters := make(map[string]interface{})
	for key, value := range queryparams {
		tempfilters[key] = value[0] // Grabs the first value out of the query parameters if there is duplicates, make it so that there can only be one kind in later iteration
	}
	results, err := models.GetWeaponsByQuery(tempfilters)

	if err != nil {
		gctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Server encountered an error fetch weapons by query."})
	} else {
		gctx.IndentedJSON(http.StatusOK, results)
	}
}

func FetchWeaponByKeyFromDB(gctx *gin.Context) {
	name := gctx.Param("name")

	result, err := models.GetWeaponByID(name)

	if err != nil {
		gctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Server encountered an error fetching one weapon."})
	} else {
		gctx.IndentedJSON(http.StatusOK, result)
	}
}
