package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (server *Server) login(ctx *gin.Context) {
	var req LoginRequest
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	if req.Username != "admin" || req.Password != "secret" {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "invalid username or password"})
		return
	}
	ctx.SetCookie("token", "secrettoken", 3600, "", "", false, true)
	ctx.JSON(http.StatusOK, gin.H{"message": "login success"})

}

func (server *Server) logout(ctx *gin.Context) {
	ctx.SetCookie("token", "", -1, "", "", false, true)
	ctx.JSON(http.StatusOK, gin.H{"message": "logout success"})
}

func (server *Server) checkAuthentication(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "authenticated"})
}
