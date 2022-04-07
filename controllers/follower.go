package controllers

import (
	"APIGOLANGMAP/model"
	"APIGOLANGMAP/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

func FetchAllFollowers(userID uint) []model.Follower {
	var followers []model.Follower
	services.Db.Where("user_id = ?", userID).Find(&followers)

	// Get Follower Users Info
	//var user model.User
	//services.Db.First(&user, userID)
	//fmt.Println(">>> ", user.UserFriends)

	//var users []model.User
	//for _, ar := range followers {
	//	services.Db.Where("id = ?", ar.FollowerUserID).Find(&users)
	//}
	//fmt.Println(">>> ", users)
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
	// Verify userID in body exists
	var user model.User
	if err := services.Db.Where("id = ?", follower.FollowerUserID).First(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "Follower User ID Not Found"})
		return
	}
	// Verify if row already exists
	var tmpFollower model.Follower
	services.Db.Where(&model.Follower{UserID: follower.UserID, FollowerUserID: follower.FollowerUserID}).First(&tmpFollower)
	if tmpFollower.ID > 0 {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "Follower User ID Already Associated"})
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
