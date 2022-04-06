package models

import "gorm.io/gorm"

// Create a struct that will be encoded to a JWT.
// We add jwt.StandardClaims as an embedded type, to provide fields like expiry time
type RevokedToken struct {
	gorm.Model `swaggerignore:"true"`
	Token string
}
