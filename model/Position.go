package model

import (
	"gorm.io/gorm"
)

// swagger:model
type Position struct {
	gorm.Model `swaggerignore:"true"`
	Latitude   float64 `json:"Latitude" binding:"required"`
	Longitude  float64 `json:"Longitude" binding:"required"`
	UserID     uint    `json:"UserId" gorm:"not null"`
}
