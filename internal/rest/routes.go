package rest

import (
	"github.com/gin-gonic/gin"
	"github.com/monami/apigateway/internal/rest/handler"
)

func SetupRoutes(r gin.IRouter) {
	api := r.Group("/api")
	{
		api.POST("/login", handler.LoginHandler())
	}
}
