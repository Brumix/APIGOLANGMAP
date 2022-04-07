package controllers

import (
	"APIGOLANGMAP/model"
	"APIGOLANGMAP/repository"
	"APIGOLANGMAP/services"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var repo = repository.NewCrudPositions()

func RegisterLocation(c *gin.Context) {
	var position model.Position

	err := c.Bind(&position)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return

	}

	if errStore := repo.StorePosition(&position); errStore != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"msg": "Position register with success!!",
		"Position": position})
	return
}
func DeleteLocation(c *gin.Context) {
	var position model.Position
	err := c.Bind(&position)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return

	}
	if errDelete := repo.DeletePosition(&position); errDelete != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"msg": "Position deleted with success!!",
		"Position": position})
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
