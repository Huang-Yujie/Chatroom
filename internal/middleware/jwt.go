package middleware

import (
	"github.com/Huang-Yujie/Chatroom/pkg/auth"
	"github.com/Huang-Yujie/Chatroom/pkg/errcode"
	"github.com/Huang-Yujie/Chatroom/pkg/response"
	"github.com/gin-gonic/gin"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		token, exists := c.GetQuery("token")
		code := errcode.Success
		if !exists || token == "" {
			code = errcode.InvalidParams
		} else {
			claims, err := auth.ParseToken(token)
			if err != nil {
				code = errcode.UnauthorizedTokenError
			} else {
				c.Set("UserID", claims.ID)
			}
		}
		if code != errcode.Success {
			r := response.NewResponse(c)
			r.ToErrorResponse(code)
			c.Abort()
			return
		}
		c.Next()
	}
}
