package setauthid

import (
	"regexp"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"

	"github.com/rene00/khaos/models"
	"github.com/rene00/khaos/pkg/util"
)

type AuthContext struct {
	ID       int
	Username string
}

// SetAuthID will add AuthID to gins Context if the client has submitted
// an Authorization HTTP header which contains a token which khaos can
// parse and decrypt an authentication username from. SetAuthID takes
// this username and queries the database for the user id which it sets
// as the AuthID gin Context.
func SetAuthID() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		if len(authHeader) == 0 {
			log.Debug("No Authorization Header")
			return
		}
		authRegex := regexp.MustCompile("^Bearer\\s([0-9a-zA-Z\\.\\_\\-]+)$")
		match := authRegex.FindStringSubmatch(authHeader)

		if len(match) != 2 {
			log.Debug("Failed to parse Authorization Header")
			return
		}

		authToken := match[1]
		authTokenData, _ := util.ParseToken(authToken)

		username := util.DecryptString(authTokenData.Username)
		authID, err := models.GetID(username)
		if err != nil {
			log.Debug("Failed to get ID")
			return
		}

		c.Set("AuthID", authID)
		c.Next()
	}
}
