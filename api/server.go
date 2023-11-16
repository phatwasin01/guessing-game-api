package api

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Server struct {
	router *gin.Engine
}

func NewServer() (*Server, error) {
	router := gin.Default()
	corsConfig := cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		AllowCredentials: true,
	}
	router.Use(cors.New(corsConfig))
	server := &Server{
		router: router,
	}
	authRoutes := router.Group("/").Use(authMiddleWare)
	router.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(200, "ok")
	})
	router.POST("/login", server.login)
	router.POST("/logout", server.logout)
	authRoutes.POST("/guess", server.guessNumber)
	// check authentication
	authRoutes.GET("/check", server.checkAuthentication)

	return server, nil
}

func (server *Server) Start() error {
	return server.router.Run()
}
func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
