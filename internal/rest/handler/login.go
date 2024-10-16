package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/monami/apigateway/internal/rest/domain/dto"
)

func LoginHandler() func(*gin.Context) {
	var input dto.LoginInput
	return func(ctx *gin.Context) {
		if err := ctx.ShouldBindJSON(&input); err != nil {
			//failed to parse the request body
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		if input.Username != "1" {
			//failed to authorize
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"error": "login works only for username \"1\"",
			})
			return
		}

		ctx.JSON(200, dto.LoginOutput{
			DummyToken: "dummy_token",
		})
	}
}
