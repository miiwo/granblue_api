package middleware

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func CorsConfig(gctx *gin.Context) {
	gctx.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	gctx.Writer.Header().Set("Access-Control-Allow-Methods", "GET")
	gctx.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

	if gctx.Request.Method == "OPTIONS" {
		gctx.AbortWithStatus(http.StatusNoContent)
		return
	}

	gctx.Next()
}