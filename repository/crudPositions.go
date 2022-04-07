package repository

import (
	"APIGOLANGMAP/model"
	"gorm.io/gorm"
	"math"
	"time"
)

var DB *gorm.DB

type CrudPositions interface {
	StorePosition(position *model.Position) error
	DeletePosition(position *model.Position) error
	GetAllPositions() ([]model.Position, error)
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
	err := DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(position).Error; err != nil {
			panic("ERROR creating the Position")
			return err
		}
		DB.Exec("update positions set geolocation = 'point(? ?)' where user_id=?", int(position.Longitude), int(position.Latitude), position.UserID)

		return nil
	})

	return err
}

func (p *PositionStruck) DeletePosition(position *model.Position) error {

	err := DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Delete(position).Error; err != nil {
			panic("ERROR Deleting the Position")
			return err
		}
		return nil
	})
	return err
}

func (p *PositionStruck) GetAllPositions() ([]model.Position, error) {
	var positions []model.Position
	err := DB.Transaction(func(tx *gorm.DB) error {
		result := tx.Where("updated_at > ?", time.Now().Add(-(1 * time.Minute))).Find(&positions)
		if result.Error != nil {
			panic("ERROR GETTING the Positions")
			return result.Error
		}
		return nil
	})
	if err != nil {
		return []model.Position{}, err
	}
	return positions, nil

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

	positions, _ = p.GetAllPositions()

	for i:=0; i<len(positions);i++{
		la_position2 := positions[i].Latitude
		lo_position2 := positions[i].Longitude

		distance := Distance(la_position1, lo_position1, la_position2, lo_position2)

		if distance <= 5000 //5000, valor assumido hipoteticamente, correspondente a 5 kms, valor a ser substituido pelo numero proveniente da rota.
				//TODO Verificar qual o user e alertar de seguida para o user principal

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
