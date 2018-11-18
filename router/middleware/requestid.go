package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/satori/go.uuid"
)

// 在请求和返回的 Header 中插入 X-Request-Id（X-Request-Id 值为 32 位的 UUID，用于唯一标识一次 HTTP 请求）
// 请求中插入X-Request-Id中间件
func RequestId() gin.HandlerFunc {
	return func(context *gin.Context) {
		// Check for incoming header, use it if exists
		requestId := context.Request.Header.Get("X-Request-Id")

		// Create request id with UUID4
		if requestId == "" {
			u4, _ := uuid.NewV4()
			requestId = u4.String()
		}

		// Expose it for use in the application
		context.Set("X-Request-Id", requestId)

		// Set X-Request-Id header
		context.Writer.Header().Set("X-Request-Id", requestId)
		context.Next()
	}
}
