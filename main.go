package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"miiwo/skyfarer/handler"
	"miiwo/skyfarer/auth"
	"github.com/joho/godotenv"
	"miiwo/skyfarer/backend/models"
)

func main() {
	// INIT
	router := gin.Default()
	godotenv.Load(".env")
	models.ConnectToMariaDB()

	// ROUTES
	v1 := router.Group("v1", auth.ValidateAPIKey)
	v1.GET("/weapons", handler.FetchWeaponsFromDB)

	router.GET("/", handler.FetchWeapons)
	router.GET("/auth/key", func (gctx *gin.Context) {
		gctx.IndentedJSON(http.StatusOK, gin.H{"apiKey": auth.GenerateAPIKey()})
	})

	// RUN
	router.Run("localhost:8084")
}

