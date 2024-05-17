package main

import (
	"github.com/gin-gonic/gin"
	"miiwo/skyfarer/controllers"
)

func main() {
	router := gin.Default()

	// AUTH
	/*authorized := router.Group("/weapons", gin.BasicAuth(gin.Accounts{
		"foo": "bar",
		"user2": "orchid",
	}))*/

	// ROUTES
	router.GET("/weapons", controllers.FetchWeapons)
	//authorized.Use(AuthRequired) {
	//	authorized.GET("/auth/weapons", getWeapons)
	//}
	

	router.Run("localhost:8084")
}

/*func AuthRequired (gctx *gin.Context) {
	authorization := gctx.GetHeader("Authorization")
	gctx.Next()
	/*if strings.HasPrefix(authorization, "Bearer ") {
		splits := strings.Split(authorization, " ")
		if len(splits != 2) {
			// say it is an invalid header for not having a certain number of auth fields. This one should only check for Bearer token
		}
		claims, err := nil
		gctx.Set(ClaimsContextVar, claims)
	}
}*/