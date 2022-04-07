package model

import (
	"strconv"

	"github.com/dgrijalva/jwt-go"
)

// Create a struct that will be encoded to a JWT.
// We add jwt.StandardClaims as an embedded type, to provide fields like expiry time
type Claims struct {
	UserID             uint   `json:"userid"`
	Username           string `json:"username"`
	AccessMode         int    `json:"access_mode"`
	IsSOSActivated     bool   `json:"isSOSActivated"`
	jwt.StandardClaims `swaggerignore:"true"`
}

func (claims Claims) IsAdmin() bool {
	if claims.AccessMode == UserAccess {
		return false
	} else if claims.AccessMode == AdminAccess {
		return true
	} else {
		panic("User " + claims.Username + " has invalid access mode " + strconv.Itoa(claims.AccessMode))
	}
}
