package services

import (
	"APIGOLANGMAP/model"
	"APIGOLANGMAP/repository"
	"fmt"
	"github.com/go-co-op/gocron"
	"time"
)

const MAXCONCURENT = 10

func StartService() {
	cron := gocron.NewScheduler(time.UTC)
	cron.Every(1).Hour().Do(securityConcurrent)
	cron.StartAsync()

}

func securityConcurrent() {
	fmt.Println("LAUNCH!!")
	var semaphoreChan = make(chan struct{}, MAXCONCURENT)
	var results = make(map[string]interface{})
	var positions, errGetAllPositions = repository.NewCrudPositions().GetAllPositions()

	if errGetAllPositions != nil {
		panic("Error service SecurityConcurrent ")
		return
	}
	defer positions.Close()
	for positions.Next() {
		err := Db.ScanRows(positions, &results)
		if err != nil {
			fmt.Println("Error Scanning Row")
			continue
		}
		//TODO ADD CUSTOM TIME VERIFICATION fmt.Println(results["max"].(time.Time))
		semaphoreChan <- struct{}{}
		notifyUser := results["user_id"].(int64)
		go func() {
			defer func() {
				<-semaphoreChan
			}()
			alertUser(uint(notifyUser))

		}()
	}
}

func alertUser(user uint) {
	var followers []model.Follower
	Db.Where("user_id = ?", user).Find(&followers)
	msg := fmt.Sprintf("Alert User %d maybe in Danger", user)
	for _, follower := range followers {
		sender(follower.FollowerUserID, msg)
	}
}
