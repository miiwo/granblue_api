package middleware

import (
	"fmt"
	"strings"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	//"github.com/golang-jwt/jwt/v5"
	"miiwo/skyfarer/backend"
)

func GenerateAPIKey() string {
	key := uuid.New().String()
	fmt.Println(key)
	return key
}

func ValidateAPIKey(gctx *gin.Context) {
	key := parseBearerToken(gctx.GetHeader("Authorization"))

	if key == "" {
		gctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Needs an API Key"})
	} else if backend.ValidateAPIKey(key) {
		gctx.Next()
	} else {
		gctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid API Key"})
	}
}

func parseBearerToken(authreqheader string) string {
	if strings.HasPrefix(authreqheader, "Bearer ") {
		splits := strings.Split(authreqheader, " ")

		if len(splits) == 2 {
			return splits[1]
		}
	}

	return ""
}

func AuthRequired (gctx *gin.Context) {
	fmt.Println("Went through auth function")
	gctx.Next()
}