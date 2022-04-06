package services

import (
	"APIGOLANGMAP/model"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AuthorizationRequired(adminAccess bool) gin.HandlerFunc {

	return func(c *gin.Context) {
		if !ValidateTokenJWT(c, adminAccess) {
			c.JSON(http.StatusUnauthorized, gin.H{"status": http.StatusUnauthorized, "message": "Not authorized"})
			c.Abort()
		} else {
			var tokenInput, _, _ = getAuthorizationToken(c)
			token, err := jwt.ParseWithClaims(tokenInput, &model.Claims{}, func(token *jwt.Token) (interface{}, error) {
				return JwtKey, nil
			})

			if err != nil {
				c.JSON(http.StatusUnauthorized, gin.H{"status": http.StatusUnauthorized, "message": "Not authorized"})
				c.Abort()
				return
			}

			if claims, ok := token.Claims.(*model.Claims); ok && token.Valid {
				//fmt.Printf("%v %v", claims.Username, claims.StandardClaims.ExpiresAt)
				c.Set("userid", claims.UserID)
				c.Set("username", claims.Username)
			}

			// before request
			c.Next()
		}
	}
}
