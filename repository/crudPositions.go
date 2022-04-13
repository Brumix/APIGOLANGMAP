package repository

import (
	"APIGOLANGMAP/model"
	"database/sql"
	_ "github.com/lib/pq"
	"gorm.io/gorm"

	"log"
	_ "time"
)

var DB *gorm.DB

type CrudPositions interface {
	StorePosition(position *model.Position) error
	GetAllPositions() (*sql.Rows, error)
	GetAllUsers() ([]model.User, error)
}

type PositionStruck struct{}

func NewCrudPositions() CrudPositions {
	return &PositionStruck{}
}

func GetDataBase(database *gorm.DB) {
	DB = database
}

func (p *PositionStruck) StorePosition(position *model.Position) error {
	if err := DB.Create(position).Error; err != nil {
		log.Println("ERROR creating the Position")
		return err
	}

	if errGeoLocation := DB.Exec("UPDATE positions SET geolocation = ST_SetSRID(ST_Point(longitude,latitude),4326)::geography").Error; errGeoLocation != nil {
		log.Println("ERROR updating the Position")
		return errGeoLocation
	}
	return nil
}

func (p *PositionStruck) GetAllPositions() (*sql.Rows, error) {
	rows, err := DB.Table("positions").Distinct("user_id, MAX(created_at)").Group("user_id").Rows()
	if err != nil {
		return nil, err
	}
	return rows, nil
}

func (p *PositionStruck) GetAllUsers() ([]model.User, error) {
	var users []model.User

	err := DB.Transaction(func(tx *gorm.DB) error {
		result := tx.Find(&users)
		if result.Error != nil {
			panic("ERROR GETTING the Positions")
			return result.Error
		}
		return nil
	})
	if err != nil {
		return []model.User{}, err
	}
	return users, nil

}
