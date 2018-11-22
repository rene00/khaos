package jwt

import (
	"net/http"
	"regexp"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"

	"github.com/rene00/khaos/pkg/e"
	"github.com/rene00/khaos/pkg/util"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		var data interface{}

		code = e.SUCCESS

		authHeader := c.Request.Header.Get("Authorization")
		r, _ := regexp.Compile("^Bearer (.+)$")

		match := r.FindStringSubmatch(authHeader)

		if len(match) == 0 {
			code = e.UNAUTHORIZED
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": code,
				"msg":  e.GetMsg(code),
				"data": data,
			})

			c.Abort()
			return
		}

		token := match[1]

		if len(token) == 0 {
			code = e.ERROR_AUTH_TOKEN
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": code,
				"msg":  e.GetMsg(code),
				"data": data,
			})

			c.Abort()
			return
		}

		_, err := util.ParseToken(token)

		if err != nil {
			switch err.(*jwt.ValidationError).Errors {
			case jwt.ValidationErrorExpired:
				code = e.ERROR_AUTH_CHECK_TOKEN_TIMEOUT
			default:
				code = e.ERROR_AUTH_CHECK_TOKEN_FAIL
			}
		}

		if code != e.SUCCESS {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": code,
				"msg":  e.GetMsg(code),
				"data": data,
			})

			c.Abort()
			return
		}

		c.Next()
	}
}
