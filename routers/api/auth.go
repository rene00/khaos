package api

import (
	"encoding/base64"
	"net/http"
	"strings"

	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"

	"fmt"
	"github.com/rene00/khaos/pkg/app"
	"github.com/rene00/khaos/pkg/e"
	"github.com/rene00/khaos/pkg/util"
	"github.com/rene00/khaos/service/auth_service"
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

func GetAuth(c *gin.Context) {
	appG := app.Gin{C: c}
	valid := validation.Validation{}

	h := c.Request.Header.Get("Authorization")

	credentials, err := DecodeAuthorization(h)
	if err != nil {
		appG.Response(http.StatusBadRequest, e.BAD_REQUEST, err)
		return
	}

	ok, err := valid.Valid(credentials)
	if err != nil || !ok {
		err = &AuthError{Message: "Failed Validation"}
		appG.Response(http.StatusBadRequest, e.BAD_REQUEST, err)
		return
	}

	authService := auth_service.Auth{Username: credentials.Username, Password: credentials.Password}
	isExist, err := authService.Check()
	if err != nil || !isExist {
		appG.Response(http.StatusUnauthorized, e.ERROR_AUTH, nil)
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
}
