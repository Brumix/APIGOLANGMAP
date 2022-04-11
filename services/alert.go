package services

import (
	"APIGOLANGMAP/model"
	"APIGOLANGMAP/repository"
	"fmt"
	"github.com/go-co-op/gocron"
	"log"
	"time"
)

func StartService() {
	cron := gocron.NewScheduler(time.UTC)
	cron.Every(1).Hour().Do(securityConcurrent)
	cron.StartAsync()
}

func securityConcurrent() {
	fmt.Println("LAUNCH!!")
	var results = make(map[string]interface{})
	var positions, errGetAllPositions = repository.NewCrudPositions().GetAllPositions()
	var users, errGetAllUsers = repository.NewCrudPositions().GetAllUsers()
	var auxUsers = make(map[uint]model.User)
	if errGetAllPositions != nil || errGetAllUsers != nil {
		panic("Error service SecurityConcurrent ")
		return
	}
	for _, user := range users {
		auxUsers[user.ID] = user
	}

	defer positions.Close()
	for positions.Next() {
		err := Db.ScanRows(positions, &results)
		if err != nil {
			log.Println("Error Scanning Row")
			continue
		}
		notifyUser := results["user_id"].(int64)
		timeLastUpdate := results["max"].(time.Time)
		currentUser, exist := auxUsers[uint(notifyUser)]

		if !exist {
			log.Println("User reject from Alert ", notifyUser)
			continue
		}

		if !currentUser.SOS && int(time.Now().Sub(timeLastUpdate).Hours()) < currentUser.AlertTime {
			continue
		}
		fmt.Println("Notify user", notifyUser)
		alertUser(uint(notifyUser))
	}
}

func alertUser(user uint) {
	var followers []model.Follower
	Db.Where("user_id = ?", user).Find(&followers)
	msg := fmt.Sprintf("Alert User %d maybe in Danger", user)
	for _, follower := range followers {
		fmt.Printf("MESSAGE SEND TO %d FROM %d \t", follower.FollowerUserID, user)
		sender(follower.FollowerUserID, msg)
	}
}
