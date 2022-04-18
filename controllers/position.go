package controllers

import (
	"APIGOLANGMAP/model"
	"APIGOLANGMAP/repository"
	"APIGOLANGMAP/services"
	"fmt"
	"math"
	"net/http"
	"strconv"
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
	if err := c.ShouldBindJSON(&position); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return

	}
	position.UserID = userID.(uint)
	if errStore := repo.StorePosition(&position); errStore != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"msg": errStore.Error()})
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

func GetUsersLocationWithFilters(c *gin.Context) {
	var positions []model.Position
	var user model.User
	var data struct {
		UsersId []int    `gorm:"not null" json:"UserId"`
		Dates   []string `gorm:"not null" json:"Dates"`
	}

	if err := c.BindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "Bad request!"})
		return
	}

	// Verificar se os dados foram enviados corretamente, na data posso só receber uma data(pesquisar só em um dia) ou um intervalor de datas
	if len(data.Dates) == 1 {
		var startDate, errStart = ValidateDate(data.Dates[0])
		if errStart != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "Invalid date!"})
			return
		}
		data.Dates[0] = startDate.Format("2006-01-02 15:04:05")
	} else if len(data.Dates) == 2 {
		var startDate, errStart = ValidateDate(data.Dates[0])
		var endDate, errEnd = ValidateDate(data.Dates[1])
		if errStart != nil || errEnd != nil {
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
	services.Db.Raw(GenerateQuery(data.UsersId, data.Dates)).Scan(&positions)

	if len(positions) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "None found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Users Locations", "locations": positions})
}

func GenerateQuery(users_id []int, date []string) string {
	where := "where 1 = 1"
	for i := 0; i < len(users_id); i++ {
		where += " AND user_id = " + strconv.Itoa(users_id[i]) + ""
	}
	if len(date) == 1 {
		where += " AND created_at >='" + date[0] + "' AND created_at <'" + date[0] + "'::date + '1 day'::interval"
	} else if len(date) == 2 {
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

	var users []int

	var data struct {
		Latitude  float64 `gorm:"not null" json:"Latitude"`
		Longitude float64 `gorm:"not null" json:"Longitude"`
		Meters    float64 `gorm:"not null" json:"Meters"`
	}

	if err := c.BindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "Bad request!"})
		return
	}

	la_position := data.Latitude
	lo_position := data.Longitude
	meters := data.Meters

	userid, errAuth := c.Get("userid")
	if errAuth == false {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "User Auth Token Malformed!"})
		return
	}

	//UserId 2
	if userid != 2 {
		var position2 model.Position

		//Vai buscar a posiçao mais recente de um determinado user
		if err := services.Db.Where("user_id = ?", 2).Order("created_at DESC").First(&position2).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "User ID Not Found"})
			return
		}

		la_position2 := position2.Latitude
		lo_position2 := position2.Longitude

		distance2 := Distance(la_position, lo_position, la_position2, lo_position2)
		if distance2 < meters {
			users = append(users, 2)
		}
	}

	//UserId 5

	if userid != 5 {
		var position5 model.Position

		//Vai buscar a posiçao mais recente de um determinado user
		if err := services.Db.Where("user_id = ?", 5).Order("created_at DESC").First(&position5).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "User ID Not Found"})
			return
		}

		la_position5 := position5.Latitude
		lo_position5 := position5.Longitude

		distance5 := Distance(la_position, lo_position, la_position5, lo_position5)
		if distance5 < meters {
			users = append(users, 5)
		}
	}

	//UserId 6
	if userid != 6 {
		var position6 model.Position

		//Vai buscar a posiçao mais recente de um determinado user
		if err := services.Db.Where("user_id = ?", 6).Order("created_at DESC").First(&position6).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "User ID Not Found"})
			return
		}

		la_position6 := position6.Latitude
		lo_position6 := position6.Longitude

		distance6 := Distance(la_position, lo_position, la_position6, lo_position6)
		if distance6 < meters {
			users = append(users, 6)
		}
	}

	for i := range users {
		c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Users closers than 5 kms", "userid": users[i]})
		msg := fmt.Sprintf("Alert user %d is closer to user %d", users[i], userid)
		fmt.Printf("MESSAGE SEND TO %d FROM %d \t", users[i], userid)
		services.Sender(uint(users[i]), msg)

	}

}
func hsin(theta float64) float64 {
	return math.Pow(math.Sin(theta/2), 2)
}

func Distance(lat1, lon1, lat2, lon2 float64) float64 {

	var la1, lo1, la2, lo2, r float64
	la1 = lat1 * math.Pi / 180
	lo1 = lon1 * math.Pi / 180
	la2 = lat2 * math.Pi / 180
	lo2 = lon2 * math.Pi / 180

	r = 6378100 // Earth radius in METERS

	h := hsin(la2-la1) + math.Cos(la1)*math.Cos(la2)*hsin(lo2-lo1)

	return 2 * r * math.Asin(math.Sqrt(h))
}

/*
func GetAllUsersUnderKms(c *gin.Context) {
	var position model.Position
	var results = make(map[string]interface{})
	var users = make(map[string]interface{})
	var locations = make(map[string]interface{})

	type User struct {
		id string
	}

	//var user []User

	userid, errAuth := c.Get("userid")
	if errAuth == false {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "User Auth Token Malformed!"})
		return
	}
	log.Println(userid)

	if err := services.Db.Where("user_id = ?", userid).Order("created_at DESC").First(&position).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "User ID Not Found"})
		return
	}

	usersTable, errorTable := services.Db.Raw("SELECT id from users u where id != ? and id !=1;", userid).Rows()
	if errorTable != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"msg": errorTable.Error()})
		return
	}

	defer usersTable.Close()

	for usersTable.Next() {
		err := services.Db.ScanRows(usersTable, &users)
		if err != nil {
			log.Println("Error Scanning Row")
			continue
		}
		//jsonStr, err := json.Marshal(users)
	}

	locationUsers, errorTable := services.Db.Raw("SELECT geolocation, user_id from positions p where user_id != ? and user_id != 1 order by created_at DESC limit 1;", userid).Rows()
	if errorTable != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"msg": errorTable.Error()})
		return
	}

	defer locationUsers.Close()

	for locationUsers.Next() {
		err := services.Db.ScanRows(locationUsers, &locations)
		if err != nil {
			log.Println("Error Scanning Row")
			continue
		}
		log.Println(locations)
	}

	err := c.Bind(&position)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	positions, errorTable := services.Db.Raw("select user_id as id, ST_Distance(p.geolocation, ST_Point(?,?)) as distance from positions p where  ST_Distance(p.geolocation, ST_Point(?,?)) < 10000000 and user_id != 1 and user_id != ? order by distance;", position.Longitude, position.Latitude, position.Longitude, position.Latitude, userid).Rows()
	if errorTable != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	defer positions.Close()

	for positions.Next() {
		err := services.Db.ScanRows(positions, &results)
		if err != nil {
			log.Println("Error Scanning Row")
			continue
		}
		log.Println(results)

	}

}*/
