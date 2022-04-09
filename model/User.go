package model

import (
	"strconv"

	"gorm.io/gorm"
)

const UserAccess = 1
const AdminAccess = -1
const minAlertTime = 1
const maxAlertTime = 48

type User struct {
	gorm.Model    `swaggerignore:"true"`
	AlertTime     int        `json:"alertTime,omitempty" gorm:"default:1"`
	SOS           bool       `json:"sos,omitempty" gorm:"default:0"`
	Username      string     `json:"username" gorm:"unique"`
	Password      string     `json:"password,omitempty"`
	AccessMode    int        `json:"access_mode" gorm:"default:1"`
	UserFriends   []Follower `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;foreignKey:user_id;references:id"`
	UserPositions []Position `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

func (user User) IsAdmin() bool {
	if user.AccessMode == UserAccess {
		return false
	} else if user.AccessMode == AdminAccess {
		return true
	} else {
		panic("User " + user.Username + " has invalid access mode " + strconv.Itoa(user.AccessMode))
	}
}

func (user User) InvalidAlertTime() bool { return user.AlertTime < minAlertTime || user.AlertTime > maxAlertTime }
