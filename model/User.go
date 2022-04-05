package model

import "gorm.io/gorm"

type User struct {
	gorm.Model    `swaggerignore:"true"`
	Username      string     `json:"username" gorm:"unique"`
	Password      string     `json:"password,omitempty"`
	IsAdmin       bool       `json:"IsAdmin" gorm:"default:false"`
	UserFriends   []Follower `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	UserPositions []Position `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
