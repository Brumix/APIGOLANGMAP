package model

import "gorm.io/gorm"

// swagger:model
type Follower struct {
	gorm.Model     `swaggerignore:"true"`
	UserID         uint `json:"UserID" gorm:"not null"`
	FollowerUserID uint `json:"FollowerUserID" gorm:"not null"`
}
