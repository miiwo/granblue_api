package main

import (
	"miiwo/skyfarer/backend/models"
	"miiwo/skyfarer/handler"
	"miiwo/skyfarer/middleware"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// INIT
	godotenv.Load(".env")
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.ForwardedByClientIP = true
	router.SetTrustedProxies([]string{"127.0.0.1", "10.0.0.0/8"})
	models.ConnectToMariaDB()

	// ROUTES
	v1 := router.Group("v1", middleware.CorsConfig, middleware.ValidateAPIKey)
	v1.GET("/weapons/:name", handler.FetchWeaponByKeyFromDB)
	v1.GET("/weapons", handler.FetchWeaponsFromDB)

	// Just to show routes are working without authentication
	router.GET("/ping", func(gctx *gin.Context) {
		gctx.IndentedJSON(200, gin.H{"message": "pong"})
	})

	// RUN
	router.Run(os.Getenv("BASE_URL"))
	// HTTPS
	//log.Fatal(autotls.Run(router, os.Getenv("CUSTOM_DOMAIN")))
}
