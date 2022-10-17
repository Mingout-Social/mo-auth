package middleware

import (
	"net/http"

	"github.com/Mingout-Social/mo-auth/lib"

	"github.com/Mingout-Social/mo-auth/controllers"

	"github.com/gin-gonic/gin"
)

func AuthenticateUserToken(ctx *gin.Context) {
	token := ctx.GetHeader("x-user-token")
	os := ctx.GetHeader("x-os")

	userId, err := lib.VerifyToken(token, os)

	if err != nil {
		ctx.JSON(http.StatusForbidden, controllers.ErrorResponse{
			Error:        true,
			ErrorMessage: err.Error(),
		})
		ctx.Abort()
	}

	ctx.Set("userId", userId)
	ctx.Next()
}
