package middleware

import (
	"gin/app/common/response"
	"gin/app/services"
	"gin/global"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func JwtAuth(GuardName string) gin.HandlerFunc {
	return func(context *gin.Context) {
		tokenStr := context.Request.Header.Get("Authorization")
		if tokenStr == "" {
			response.TokenFail(context)
			context.Abort()
			return
		}
		tokenStr = tokenStr[len(services.TokenType)+1:]

		// Token 解析校验
		token, err := jwt.ParseWithClaims(tokenStr, &services.CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(global.App.Config.Jwt.Secret), nil
		})
		if err != nil || services.JwtService.IsInBlacklist(tokenStr) {
			response.TokenFail(context)
			context.Abort()
			return
		}

		claims := token.Claims.(*services.CustomClaims)

		// Token 发布者校验
		if claims.Issuer != GuardName {
			response.TokenFail(context)
			context.Abort()
			return
		}

		context.Set("token", token)
		context.Set("id", claims.Id)
	}
}
