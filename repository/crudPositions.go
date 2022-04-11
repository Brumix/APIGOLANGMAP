package repository

import (
	"APIGOLANGMAP/model"
	"database/sql"
	_ "github.com/lib/pq"
	"gorm.io/gorm"

	"log"
	"math"
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

func (p *PositionStruck) GetAllUsersUnderXKms(position *model.Position) error {
	var positions []model.Position

	la_position1 := position.Latitude
	lo_position1 := position.Longitude

	//positions, _ = p.GetAllPositions()

	for i := 0; i < len(positions); i++ {
		la_position2 := positions[i].Latitude
		lo_position2 := positions[i].Longitude

		distance := Distance(la_position1, lo_position1, la_position2, lo_position2)

		//if distance <= 5000
		//TODO Verificar qual o user e alertar de seguida para o user principal

	}
	return nil

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
