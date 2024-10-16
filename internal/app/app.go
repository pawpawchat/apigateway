package app

import (
	"github.com/gin-gonic/gin"
	"github.com/monami/apigateway/internal/rest"
)

func Run() {
	r := gin.Default()
	rest.SetupRoutes(r)
	r.Run()
}
