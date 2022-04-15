package services

import (
	"APIGOLANGMAP/model"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func GinMiddleware(allowOrigin string) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:8080")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Request.Header.Del("Origin")

		c.Next()
	}
}

func AuthorizationRequired() gin.HandlerFunc {

	return func(c *gin.Context) {
		if !ValidateTokenJWT(c) {
			c.JSON(http.StatusUnauthorized, gin.H{"status": http.StatusUnauthorized, "message": "Not authorized"})
			c.Abort()
		} else {
			var tokenInput, _, _ = GetAuthorizationToken(c)
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
				c.Set("isSOSActivated", claims.IsSOSActivated)
				c.Set("AccessMode", claims.AccessMode)
			}
			//OpenDatabase()

			//defer CloseDatabase()
			// before request
			c.Next()
		}
	}
}

func AdminAuthorizationRequired() gin.HandlerFunc {

	return func(c *gin.Context) {
		if !ValidateTokenJWT(c) {
			c.JSON(http.StatusUnauthorized, gin.H{"status": http.StatusUnauthorized, "message": "Not authorized"})
			c.Abort()
		} else {
			var tokenInput, _, _ = GetAuthorizationToken(c)
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
				if !claims.IsAdmin() {
					c.JSON(http.StatusUnauthorized, gin.H{"status": http.StatusUnauthorized, "message": "Not authorized"})
					c.Abort()
				}

			}
			//OpenDatabase()

			//defer CloseDatabase()
			// before request
			c.Next()
		}
	}
}
