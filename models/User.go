package models

import (
	"gorm.io/gorm"
	"strconv"
)

const UserAccess  = 1
const AdminAccess = -1

type User struct {
	gorm.Model `swaggerignore:"true"`
	Username   string `json:"username" gorm:"unique"`
	Password   string `json:"password,omitempty"`
	AccessMode    int `json:"access_mode" gorm:"default:1"`
}

func (user User) IsAdmin() bool {
	if user.AccessMode == UserAccess { return false
	} else if user.AccessMode == AdminAccess { return true
	} else { panic("User " + user.Username + " has invalid access mode " + strconv.Itoa(user.AccessMode)) }
}
