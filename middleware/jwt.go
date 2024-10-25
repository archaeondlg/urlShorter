package middleware

import (
	"errors"
	"project/model/errorCode"
	"project/model/response"
	"project/utils"

	"github.com/gin-gonic/gin"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := utils.GetToken(c)
		if token == "" {
			response.Error(errorCode.NOAUTH, c)
			c.Abort()
			return
		}
		// TODO 黑名单验证

		claims, err := utils.ParseToken(token)
		if err != nil {
			if errors.Is(err, utils.ErrTokenExpired) {
				response.Error(errorCode.Expired, c)
				utils.ClearToken(c)
				c.Abort()
				return
			}
			response.Error(errorCode.Invalid, c)
			utils.ClearToken(c)
			c.Abort()
			return
		}

		c.Set("auth", claims.BaseClaims)
		// 自动刷新token
		utils.AutoRefreshToken(claims, c)
		c.Next()
	}
}
