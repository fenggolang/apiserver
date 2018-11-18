package middleware

import (
	"apiserver/handler"
	"apiserver/pkg/errno"
	"apiserver/pkg/token"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		// Parse the json web token.
		if _, err := token.ParseRequest(context); err != nil {
			handler.SendResponse(context, errno.ErrTokenInvalid, nil)
			context.Abort()
			return
		}

		context.Next()
	}
}
