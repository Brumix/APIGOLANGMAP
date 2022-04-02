package models

import "gorm.io/gorm"

type User struct {
	gorm.Model `swaggerignore:"true"`
	Username   string `json:"username" gorm:"unique"`
	Password   string `json:"password,omitempty"`
	Admin        bool `json:"admin" gorm:"type:bool;default:false"`
}
