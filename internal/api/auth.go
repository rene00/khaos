package api

import (
	"encoding/base64"
	"net/http"
	"strings"

	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"

	"fmt"
	"github.com/rene00/khaos/internal/khaos"
	"github.com/rene00/khaos/models"
	"github.com/rene00/khaos/pkg/app"
	"github.com/rene00/khaos/pkg/e"
	"github.com/rene00/khaos/pkg/util"
)

// Credentials is a struct which holds the clients submitted username and
// password. The validation schema is included as a struct tag. The credentials
// are base64 decoded from the Authorization header (basic auth).
type Credentials struct {
	Username string `valid:"Required; MaxSize(50)"`
	Password string `valid:"Required; MaxSize(50)"`
}

// AuthError is the Auth Error struct.
type AuthError struct {
	Message string
}

func (e *AuthError) Error() string {
	return fmt.Sprintf("%s", e.Message)
}

// DecodeAuthorization takes a standard Authorization basic auth header string
// (without the realm). The function will attempt to base64 decode the string
// and return a Credentials struct if successful.
func DecodeAuthorization(header string) (*Credentials, error) {

	if len(header) <= 10 {
		return nil, &AuthError{Message: "Invalid Authorization Header"}
	}

	value, err := base64.StdEncoding.DecodeString(header[6:])
	if err != nil {
		return nil, &AuthError{Message: "Failed to decode Authorization Header"}
	}

	credentials := string(value)
	i := strings.Index(credentials, ":")
	return &Credentials{Username: credentials[:i], Password: credentials[i+1:]}, nil
}

func Auth(router *gin.Engine, conf *khaos.Config) {
	router.GET("/auth", func(c *gin.Context) {
		appG := app.Gin{C: c}
		valid := validation.Validation{}

		h := c.Request.Header.Get("Authorization")

		credentials, err := DecodeAuthorization(h)
		if err != nil {
			c.AbortWithStatusJSON(400, gin.H{"error": "Failed1"})
			return
		}

		ok, err := valid.Valid(credentials)
		if err != nil || !ok {
			c.AbortWithStatusJSON(400, gin.H{"error": "Failed2"})
			return
		}

		if success, err := models.Authenticate(credentials.Username, credentials.Password); err != nil || !success {
			c.AbortWithStatusJSON(401, gin.H{"error": "Failed3"})
			return
		}

		token, err := util.GenerateToken(credentials.Username, credentials.Password)
		if err != nil {
			appG.Response(http.StatusOK, e.ERROR_AUTH_TOKEN, nil)
			return
		}

		appG.Response(http.StatusOK, e.SUCCESS, map[string]string{
			"token": token,
		})
	})
}
