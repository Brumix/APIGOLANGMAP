package controllers

import (
	"APIGOLANGMAP/model"
	"APIGOLANGMAP/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

func FetchAllFollowers(userID uint) []model.Follower {
	var followers []model.Follower

	services.Db.Find(&followers, userID)

	return followers
}

func GetAllFollowers(c *gin.Context) {
	userID, err := c.Get("userid")

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
	userID, errAuth := c.Get("userid")

	if errAuth == false {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "User Auth Token Malformed!"})
		return
	}

	var follower model.Follower

	if err := c.ShouldBindJSON(&follower); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "Check syntax!"})
		return
	}

	follower.UserID = userID.(uint)
	// Falta verificar se o follower user id existe!
	var user model.User
	if err := services.Db.Where("id = ?", follower.FollowerUserID).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "Follower User ID Not Found"})
		return
	}

	services.Db.Save(&follower)

	followers := FetchAllFollowers(userID.(uint))

	c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "message": "Association Successful!", "followers": followers})

}

func DeassociateFollower(c *gin.Context) {
	userID, errAuth := c.Get("userid")

	if errAuth == false {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "User Auth Token Malformed!"})
		return
	}

	var follower model.Follower
	if err := c.ShouldBindJSON(&follower); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "Bad request!"})
		return
	}

	follower.UserID = userID.(uint)
	services.Db.Where(&model.Follower{UserID: follower.UserID, FollowerUserID: follower.FollowerUserID}).First(&follower)

	if follower.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "None found!"})
		return
	}

	services.Db.Delete(&follower)
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Deassociation Successful!"})

}
