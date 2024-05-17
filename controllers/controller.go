package controllers

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

type WeaponMemory struct {
	ID			string	`json:"id"`
	Name		string	`json:"name"`
	Element		string	`json:"element"`
	WeaponType 	string	`json:"type"`
}

var weapons = []WeaponMemory {
	{ID: "1", Name: "Bahamut Blade", Element: "Dark", WeaponType: "katana"},
	{ID: "2", Name: "Knight of Ice", Element: "Water", WeaponType: "dagger"},
	{ID: "3", Name: "Phoenix's Torch", Element: "Fire", WeaponType: "staff"},
}

func FetchWeapons(gctx *gin.Context) {
	// Query Parameter
	// wepname := gctx.DefaultQuery("name", "None")
	// ele := gctx.Query("element")

	// Grab from backend

	// Return result
	gctx.IndentedJSON(http.StatusOK, weapons)
}