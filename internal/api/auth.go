package api

import (
	"encoding/base64"
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/rene00/khaos/internal/khaos"
	"github.com/rene00/khaos/models"
	"github.com/rene00/khaos/pkg/util"
)

// Credentials is a struct which holds the clients submitted username and
// password.
type Credentials struct {
	Username string
	Password string
}

// DecodeAuthorization takes a standard Authorization basic auth header string
// (without the realm). The function will attempt to base64 decode the string
// and return a Credentials struct if successful.
func DecodeAuthorization(header string) (*Credentials, error) {

	if len(header) <= 10 {
		return nil, errors.New("Invalid Authorization Header")
	}

	value, err := base64.StdEncoding.DecodeString(header[6:])
	if err != nil {
		return nil, errors.New("Failed to decode Authorization Header")
	}

	credentials := string(value)
	i := strings.Index(credentials, ":")
	return &Credentials{Username: credentials[:i], Password: credentials[i+1:]}, nil
}

func Auth(router *gin.Engine, conf *khaos.Config) {
	router.GET("/auth", func(c *gin.Context) {
		h := c.Request.Header.Get("Authorization")

		credentials, err := DecodeAuthorization(h)
		if err != nil {
			c.AbortWithStatusJSON(400, gin.H{"error": err.Error()})
			return
		}

		if success, err := models.Authenticate(credentials.Username, credentials.Password); err != nil || !success {
			c.AbortWithStatusJSON(401, gin.H{"error": "Invalid username or/and password"})
			return
		}

		token, err := util.GenerateToken(credentials.Username, credentials.Password)
		if err != nil {
			c.AbortWithStatusJSON(400, gin.H{"error": "Failed to generate authentication token"})
			return
		}

		c.JSON(http.StatusOK, map[string]string{"token": token})
	})
}
