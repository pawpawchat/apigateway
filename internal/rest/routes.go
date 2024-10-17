package rest

import (
	"github.com/gin-gonic/gin"
	"github.com/monami/apigateway/internal/rest/handler"
	"github.com/monami/apigateway/internal/rest/middleware"
)

func SetupRoutes(r gin.IRouter) {
	r.Use(middleware.CORSMiddleware())

	api := r.Group("/api")
	api.POST("/login", handler.LoginHandler())
	api.POST("/register", handler.RegisterHandler())

	r.GET("ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"message": "pong"})
	})
}
