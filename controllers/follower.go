package controllers

import (
	"APIGOLANGMAP/model"
	"APIGOLANGMAP/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func FetchAllFollowers(userID uint) []model.Follower {
	var followers []model.Follower

	services.Db.Find(&followers, userID)

	return followers
}

func GetAllFollowers(c *gin.Context) {
	userID, err := c.Get("user_id")

	if err == false {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "User Auth Token Malformed!"})
		return
	}
	followers := FetchAllFollowers(userID.(uint))

	if len(followers) <= 0 {
		c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Empty list!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": followers})
}

func AssociateFollower(c *gin.Context) {
	userID, err_auth := c.Get("user_id")

	if err_auth != false {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "User Auth Token Malformed!"})
		return
	}

	var follower model.Follower

	if err := c.ShouldBindJSON(&follower); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "Check syntax!"})
		return
	}

	follower.UserID = userID.(uint)

	services.Db.Save(&follower)

	followers := FetchAllFollowers(userID.(uint))

	c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "message": "Association Successful!", "followers": followers})

}

func DeassociateFollower(c *gin.Context) {
	userID, err_auth := c.Get("user_id")

	if err_auth == false {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "User Auth Token Malformed!"})
		return
	}

	var follower model.Follower

	followerUserID, err_conv := strconv.ParseUint(c.Param("id"), 10, 64)

	if err_conv != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "User Follower ID Must be an Integer."})
		return
	}

	services.Db.Where(&model.Follower{UserID: userID.(uint), FollowerUserID: uint(followerUserID)}).First(&follower)

	if follower.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "None found!"})
		return
	}

	services.Db.Delete(&follower)
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Deassociation Successful!"})

}
