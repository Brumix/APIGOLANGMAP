package services

import (
	"APIGOLANGMAP/model"
	"fmt"
	"io/ioutil"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// Create the JWT key used to create the signature
var JwtKey = GetSecretKey()

func GetSecretKey() []byte {
	b, err := ioutil.ReadFile("config/secretKey.key")
	if err != nil {
		fmt.Print(err)
	}
	secretKey := string(b)
	return []byte(secretKey)
}

func GenerateTokenJWT(credentials model.User) string {
	// Set expiration time of the token
	// 43800 is the duration of a month in minutes
	expirationTime := time.Now().Add(43800 * time.Minute)

	// Create the JWT claims, which includes the username and expiry time
	claims := &model.Claims{
		UserID:         credentials.ID,
		Username:       credentials.Username,
		AccessMode:     credentials.AccessMode,
		IsSOSActivated: credentials.IsSOSActivated,
		StandardClaims: jwt.StandardClaims{
			// In JWT, the expiry time is expressed as unix milliseconds
			ExpiresAt: expirationTime.Unix(),
		},
	}
	// Declare the token with the algorithm used for signing, and the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Create the JWT string
	tokenString, err := token.SignedString(JwtKey)

	if err != nil {
		return ""
	}
	return tokenString
}

func InvalidateTokenJWT(c *gin.Context) string {
	token, _, _ := GetAuthorizationToken(c)

	claims := &model.Claims{}
	tkn, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return JwtKey, nil
	})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return ""
		}
	}

	if tkn != nil {
		if !tkn.Valid {
			return ""
		}
	}

	// Create the JWT string
	tokenString, errTkn := tkn.SignedString(JwtKey)
	if errTkn != nil {
		return ""
	}
	return tokenString
}

func ValidateTokenJWT(c *gin.Context) bool {
	token, b, done := GetAuthorizationToken(c)
	if done {
		return b
	}

	claims := &model.Claims{}
	tkn, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return JwtKey, nil
	})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return false
		}
	}

	if tkn != nil {
		if !tkn.Valid {
			return false
		}
	}

	// Create the JWT string
	tokenString, errTkn := tkn.SignedString(JwtKey)
	if errTkn != nil {
		return false
	}

	// Check if token is revoked
	var revokedTkn model.RevokedToken
	if Db.Find(&revokedTkn, "token = ?", tokenString); revokedTkn.Token != "" {
		return false
	}

	return true
}

func GetAuthorizationToken(c *gin.Context) (string, bool, bool) {
	var token string

	reqToken := c.Request.Header.Get("Authorization")
	if reqToken == "" {
		return "", false, true
	}
	if strings.Contains(reqToken, "Bearer") {
		if strings.TrimSpace(reqToken) == "" {
			return "", false, true
		}

		splitToken := strings.Split(reqToken, "Bearer")
		token = strings.TrimSpace(splitToken[1])
	} else {
		token = strings.TrimSpace(reqToken)
	}
	return token, false, false
}
