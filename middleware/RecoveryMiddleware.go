package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"zbangbang/gin-vue-app/response"
)

// 返回错误信息
func RecoveryMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				response.Fail(ctx, nil, fmt.Sprint(err))
			}
		}()

		ctx.Next()
	}
}
