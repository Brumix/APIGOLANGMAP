package controllers

import (
	"APIGOLANGMAP/model"
	"APIGOLANGMAP/repository"
	"APIGOLANGMAP/services"
	"fmt"
	"log"
	"strconv"

	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Location struct {
	gorm.Model `swaggerignore:"true" json:"-"`
	Start      string `json:"start" binding:"required"`
	End        string `json:"end" binding:"required"`
}

var repo = repository.NewCrudPositions()

func RegisterLocation(c *gin.Context) {
	var position model.Position
	userID, errAuth := c.Get("userid")
	if errAuth == false {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "User Auth Token Malformed!"})
		return
	}
	err := c.Bind(&position)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return

	}
	position.UserID = userID.(uint)
	if errStore := repo.StorePosition(&position); errStore != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"msg": "Position register with success!!",
		"Position": position})
	return
}

func GetLastLocation(c *gin.Context) {
	var position model.Position
	userID, errAuth := c.Get("userid")
	if errAuth == false {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "User Auth Token Malformed!"})
		return
	}

	if err := services.Db.Where("user_id = ?", userID).Order("created_at DESC").First(&position).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "User ID Not Found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "msg": "Got My Current Location", "location": position})
	return
}

func GetLocationHistory(c *gin.Context) {
	var location Location
	var positions []model.Position
	userID, errAuth := c.Get("userid")
	if errAuth == false {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "User Auth Token Malformed!"})
		return
	}

	if err := c.ShouldBindJSON(&location); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "Check Syntax!"})
		return
	}

	var startDate, errStart = ValidateDate(location.Start)
	var endDate, errEnd = ValidateDate(location.End)

	// Datas invalidas retorna todas as posições do utilizador
	if errStart != nil || errEnd != nil {
		if err := services.Db.Where("user_id = ?", userID).Order("created_at DESC").Find(&positions).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "User ID Not Found"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "extra": "Invalid date, showing all locations", "message": "My Locations History", "locations": positions})
		return
	}

	// Retorna as localizações entre datas caso as datas do body estejam formatadas corretamente
	if startDate.Before(endDate) != true {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "End Date Must Occur After Start Date"})
		return
	}

	if err := services.Db.Where("user_id = ? AND created_at > ? AND created_at < ?", userID, startDate, endDate).Order("created_at DESC").Find(&positions).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "User ID Not Found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "My Locations History Filtered", "locations": positions})
	return

}

func DeleteLocation(c *gin.Context) {
	var position model.Position

	id := c.Param("id")
	services.Db.First(&position, id)

	if position.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "None found!"})
		return
	}

	services.Db.Delete(&position)
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Delete succeeded!"})
	return
}

func GetUsersLocationWithFilters(c *gin.Context){
	var positions []model.Position
	var data struct {
		UsersId []int `gorm:"not null" json:"UserId"`
		Dates []string `gorm:"not null" json:"Dates"`
	}

	if err := c.BindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "Bad request!"})
		return
	}

	fmt.Println(data)

	services.OpenDatabase()

	// VERIFICAR SE FORAM ENVIADOS FILTROS E SE FORAM VERIFICAR SE POR EXEMPLO A STRING DATA CORRESPONDE MESMO A UMA DATA

	//
	// QUERY
	fmt.Println(GenerateQuery(data.UsersId,data.Dates))
	// services.Db.Exec(GenerateQuery(data.UsersId,data.Dates))
	if positions := services.Db.Exec(GenerateQuery(data.UsersId,data.Dates)).Error; positions != nil {
		log.Println("ERROR updating the Position")
		return 
	}
	fmt.Println(positions)
	fmt.Println(services.Db.Exec("select * from positions"))
	return;
	// c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "My Locations History Filtered", "locations": positions})
}

func GenerateQuery(users_id[] int, date[]string) string {
	fmt.Println(len(users_id))
	where := "where 1 = 1"
	for i := 0; i < len(users_id); i++ {	
		where += " AND user_id = " + strconv.Itoa(users_id[i]) + ""
	}
	fmt.Println(len(date))
	if len(date) == 1{
		where += " AND created_at ='" + date[0] + "'"
	}else if len(date) == 2{
		where += " AND created_at >='" + date[0] + "'"
		where += " AND created_at <='" + date[1] + "'"
	}

	return "select * from positions " + where
}


func ValidateDate(dateStr string) (time.Time, error) {
	d, err := time.Parse("2006-01-02", dateStr)
	return d, err
}
