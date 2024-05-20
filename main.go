package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"miiwo/skyfarer/handler"
	"miiwo/skyfarer/middleware"
	"github.com/joho/godotenv"
	"miiwo/skyfarer/backend/models"
)

func main() {
	// INIT
	router := gin.Default()
	godotenv.Load(".env")
	models.ConnectToMariaDB()

	// ROUTES
	v1 := router.Group("v1", middleware.CorsConfig, middleware.ValidateAPIKey)
	v1.GET("/weapons", handler.FetchWeaponsFromDB)

	router.GET("/", handler.FetchWeapons)
	router.GET("/auth/key", func (gctx *gin.Context) {
		gctx.IndentedJSON(http.StatusOK, gin.H{"apiKey": middleware.GenerateAPIKey()})
	})

	// RUN
	router.Run("192.168.68.88:8084")
}

