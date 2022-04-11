package controllers

import (
	"APIGOLANGMAP/model"
	"APIGOLANGMAP/repository"
	"APIGOLANGMAP/services"
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

	services.Db.Unscoped().Delete(&position)
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Delete succeeded!"})
	return
}

func GetUsersLocationWithFilters(c *gin.Context){
	var positions []model.Position
	var user model.User
	var data struct {
		UsersId []int `gorm:"not null" json:"UserId"`
		Dates []string `gorm:"not null" json:"Dates"`
	}

	if err := c.BindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "Bad request!"})
		return
	}

	// Verificar se os dados foram enviados corretamente, na data posso só receber uma data(pesquisar só em um dia) ou um intervalor de datas
	if len(data.Dates) == 1 {
		var startDate, errStart = ValidateDate(data.Dates[0])
		if errStart != nil{
			c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "Invalid date!"})
			return
		}
		data.Dates[0] = startDate.Format("2006-01-02 15:04:05")
	}else if len(data.Dates) == 2{
		var startDate, errStart = ValidateDate(data.Dates[0])
		var endDate, errEnd = ValidateDate(data.Dates[1])
		if errStart != nil || errEnd != nil{
			c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "Invalid date!"})
			return
		}
		data.Dates[0] = startDate.Format("2006-01-02 15:04:05")
		data.Dates[1] = endDate.Format("2006-01-02 15:04:05")
	}

	// Verificar se os users existem na bd
	for i := 0; i < len(data.UsersId); i++ {	
		services.Db.Find(&user, data.UsersId[i])
		if user.ID == 0 {
			c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "Invalid user!"})
			return
		}
	}
	
	// QUERY
	services.Db.Raw(GenerateQuery(data.UsersId,data.Dates)).Scan(&positions)

	if len(positions) == 0{
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "None found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Users Locations", "locations": positions})
}

func GenerateQuery(users_id[] int, date[]string) string {
	where := "where 1 = 1"
	for i := 0; i < len(users_id); i++ {	
		where += " AND user_id = " + strconv.Itoa(users_id[i]) + ""
	}
	if len(date) == 1{
		where += " AND created_at >='" + date[0] + "' AND created_at <'" + date[0] + "'::date + '1 day'::interval"
	}else if len(date) == 2{
		where += " AND created_at >='" + date[0] + "'"
		where += " AND created_at <='" + date[1] + "'::date + '1 day'::interval"
	}

	return "select * from positions " + where
}


func ValidateDate(dateStr string) (time.Time, error) {
	d, err := time.Parse("2006-01-02", dateStr)
	return d, err
}

func GetAllUsersUnderXKms(c *gin.Context) {
	var position model.Position

	if err := c.Bind(&position); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	x := c.Param("x")

	repo.GE

}
