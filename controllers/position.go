package controllers

import (
	"APIGOLANGMAP/model"
	"APIGOLANGMAP/repository"
	"APIGOLANGMAP/services"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"time"
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
	var data struct {
		usersId []int `gorm:"not null" json:"users_id"`
		dates []string `gorm:"not null" json:"dates"`
	}

	if err := c.BindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "Bad request!"})
		return
	}

	// var position []model.Position
	// var user model.User


	services.OpenDatabase()
	// VERIFICAR SE O UTILIZADOR É ADMIN (SÓ PODE PERMITIR ADMIN)
	// if err := services.Db.Where("id = ?", id).First(&user).Error; err != nil {
	// 	c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
	// 	return
	// }
	// if !user.IsAdmin(){
	//  	c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "Nao é possivel ver localizacoes porque nao e admin"})
	//  	return
	// }
	
	// VERIFICAR SE O PEDIDO É FEITO COM FILTROS ----------------------------------
	fmt.Println(GenerateQuery(data.usersId,data.dates))
	return;
	
	// t, err := time.Parse("yyyy-mm-dd", filter)
	// t, err := time.Parse("2006-01-02", "2020-10-04")
    // if err != nil {
    //     fmt.Println(err)
    //     return
    // }
    // fmt.Println(t)


	// 	// SE VIER COM FILTROS ENTAO FAZEMOS UMA QUERY COM O FILTRO À BD ----------------------------------
	// 	services.Db.Find(&position, "UserId = ?", id) // EXEMPLO DE FILTRO POR ID, DEPOIS CORRIGIR ISTO, VER QUE TIPO DE FILTROS É QUE POSSO FAZER
	// 	if len(position) == 0 {
	// 		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "None found!"})
	// 		return
	// 	}

	// 	// SE NÃO VIER COM FILTRO ENTÃO ENVIAMOS A LOCALIZAÇÃO DE TODOS O UTILIZADORES
	// 	services.Db.Find(&position)
	// 	if len(position) <= 0 {
	// 		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "None found!"})
	// 		return
	// 	}

	// // defer services.Db.Close()
	// c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": position})
}

func GenerateQuery(users_id[] int, date[]string) string {
	fmt.Println(len(users_id))
	where := "where 1 = 1"
	for i := 0; i < len(users_id); i++ {	
		where += " AND UserId ='" + strconv.Itoa(users_id[i]) + "'"
	}
	fmt.Println(len(date))
	if len(date) == 1{
		where += " AND create_date ='" + date[0] + "'"
	}else if len(date) == 2{
		where += " AND create_date >='" + date[0] + "'"
		where += " AND create_date <='" + date[1] + "'"
	}

	return "select * from position " + where
}
func ValidateDate(dateStr string) (time.Time, error) {
	d, err := time.Parse("2006-01-02", dateStr)
	return d, err
}
