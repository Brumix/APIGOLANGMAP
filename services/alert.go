package services

import (
	"APIGOLANGMAP/model"
	"APIGOLANGMAP/repository"
	"github.com/go-co-op/gocron"
	"log"
	"time"
)

const MAXCONCURENT = 5

func StartService() {
	cron := gocron.NewScheduler(time.UTC)
	_, err := cron.Every(1).Hour().Do(securityConcurrent)
	if err != nil {
		log.Println("ERROR LAUNCHING THE CRON JOB")
	}
}

func securityConcurrent() {
	var semaphoreChan = make(chan struct{}, MAXCONCURENT)
	var positions, errGetAllPositions = repository.NewCrudPositions().GetAllPositions()
	var users, errUsers = repository.NewCrudPositions().GetAllUsers()

	var setUsers = make(map[uint]struct{})
	if errGetAllPositions != nil || errUsers != nil {
		panic("Error service SecurityConcurrent ")
		return
	}
	for _, position := range positions {
		if _, exists := setUsers[position.UserID]; exists {
			continue
		}
		setUsers[position.UserID] = struct{}{}
	}

	for _, user := range users {

		if _, exists := setUsers[user.ID]; exists {
			continue
		}

		semaphoreChan <- struct{}{}
		notifyUser := user
		go func() {
			defer func() {
				<-semaphoreChan
			}()
			alertUser(notifyUser)

		}()
	}
}

func alertUser(user model.User) {
	//TODO FUNCTION FRIENDS
	sender(user.ID, "ALERT!!")
	log.Println("ALERT!!", user.Username)
}
