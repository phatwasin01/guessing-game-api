package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func authMiddleWare(ctx *gin.Context) {
	cookie, err := ctx.Cookie("token")
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "authentication required"})
		return
	}
	if cookie != "secrettoken" {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
		return
	}
	ctx.Next()
}
