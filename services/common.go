package services

import (
	"APIGOLANGMAP/model"
	"io/ioutil"
	"strings"

	//	"time"

	"golang.org/x/crypto/bcrypt"
	postgres "gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var username string
var password string
var dbHost string
var dbPort string
var dbName string

var Db *gorm.DB

func readProperties() {
	content, _ := ioutil.ReadFile("config/db.config")

	lines := strings.Split(string(content), "\n")

	if len(lines) >= 6 {
		username = lines[1]
		password = lines[2]
		dbHost = lines[3]
		dbPort = lines[4]
		dbName = lines[5]
	}

}

func OpenDatabase() {
	//open a db connection
	readProperties()
	var err error
	dsn := "host=ec2-3-223-213-207.compute-1.amazonaws.com" + " user=prakstyrdtlrgx" + " password=4a31b09f5a753a66f0c4408a35edd91742349736bcc1574e909a88cf7f29fbed" + " dbname=ddqar5cpnhds0f" + " port=5432 " + " TimeZone=Europe/Lisbon"

	//dsn := "host=" + dbHost + " user=" + username + " password=" + password + " dbname=" + dbName + " port=" + dbPort + " sslmode=disable TimeZone=Europe/Lisbon"
	Db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	//sqlDB, _ := Db.DB()
	//sqlDB.SetMaxIdleConns(10)
	//sqlDB.SetMaxOpenConns(100)
	//sqlDB.SetConnMaxLifetime(time.Hour)
	if err != nil {
		panic("failed to connect database")
	}
}

func CloseDatabase() {
	psqlDB, err := Db.DB()
	psqlDB.Close()

	if err != nil {
		panic("failed to close database")
	}
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CreateAdmin() {
	var usr model.User
	if Db.Find(&usr, "username = ?", "admin"); usr.Username != "" {
		return
	}

	creds := model.User{
		Username:   "admin",
		Password:   "admin",
		AccessMode: model.AdminAccess,
	}

	hash, _ := HashPassword(creds.Password)

	creds.Password = hash
	result := Db.Save(&creds)
	if result.RowsAffected == 0 {
		panic("Admin could not be created")
	}
}
