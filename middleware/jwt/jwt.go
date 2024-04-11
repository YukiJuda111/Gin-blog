package jwt

import (
	"net/http"

	"github.com/YukiJuda111/go-gin-blog/pkg/e"
	"github.com/YukiJuda111/go-gin-blog/pkg/logging"
	util "github.com/YukiJuda111/go-gin-blog/pkg/utile"
	"github.com/gin-gonic/gin"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		var data interface{}

		code = e.SUCCESS
		token := c.Query("token")
		if token == "" {
			code = e.INVALID_PARAMS
		} else {
			_, err := util.ParseToken(token)
			if err != nil {
				code = e.ERROR_AUTH_CHECK_TOKEN_FAIL
			}
		}

		if code != e.SUCCESS {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": code,
				"msg":  e.GetMsg(code),
				"data": data,
			})
			logging.Warn("Token错误")
			c.Abort()
			return
		}

		c.Next()
	}
}
