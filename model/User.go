package model

import (
	"strconv"

	"gorm.io/gorm"
)

const UserAccess  = 1
const AdminAccess = -1

type User struct {
	gorm.Model `swaggerignore:"true"`
	Username   string `json:"username" gorm:"unique"`
	Password   string `json:"password,omitempty"`
	AccessMode    int `json:"access_mode" gorm:"default:1"`
	UserFriends   []Follower `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	UserPositions []Position `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

func (user User) IsAdmin() bool {
	if user.AccessMode == UserAccess { return false
	} else if user.AccessMode == AdminAccess { return true
	} else { panic("User " + user.Username + " has invalid access mode " + strconv.Itoa(user.AccessMode)) }
}
