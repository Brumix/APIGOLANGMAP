package repository

import (
	"APIGOLANGMAP/model"
	"gorm.io/gorm"
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
