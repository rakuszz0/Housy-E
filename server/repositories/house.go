package repositories

import (
	"housy/models"

	"gorm.io/gorm"
)

type HouseRepository interface {
	FindHouses() ([]models.House, error)
	FindHousesFilter(price int) ([]models.House, error)
	GetHouse(ID int) (models.House, error)
	CreateHouse(house models.House) (models.House, error)
	DeleteHouse(house models.House) (models.House, error)
	UpdateHouse(house models.House) (models.House, error)
}

// type houseRepository struct {
// 	db *gorm.DB
// }

func RepositoryHouse(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) GetHouse(ID int) (models.House, error) {
	var house models.House
	err := r.db.First(&house, ID).Error

	return house, err
}

// CreateHouse implements HouseRepository
func (r *repository) CreateHouse(house models.House) (models.House, error) {
	err := r.db.Create(&house).Error // Using Create method

	return house, err
}

// DeleteHouse implements HouseRepository
func (r *repository) DeleteHouse(house models.House) (models.House, error) {
	err := r.db.Delete(&house).Error // Using Delete method

	return house, err
}

// FindHousesFilter implements HouseRepository
func (r *repository) FindHousesFilter(price int) ([]models.House, error) {
	var houses []models.House
	var err error
	if price != 0 {
		err = r.db.Where("price < ?", price).Find(&houses).Error
	} else {
		err = r.db.Find(&houses).Error
	}
	return houses, err
}

// UpdateHouse implements HouseRepository
func (r *repository) UpdateHouse(house models.House) (models.House, error) {
	err := r.db.Save(&house).Error // Using Save method

	return house, err
}

func (r *repository) FindHouses() ([]models.House, error) {
	var houses []models.House
	err := r.db.Where("sold = ?", false).Find(&houses).Error

	return houses, err
}

// func (r *houseRepository) FindHousesFilter(c echo.Context) error {
// 	price := c.QueryParam("price")
// 	var houses []models.House
// 	var err error
// 	if price != "" {
// 		err = r.db.Where("price < ?", price).Find(&houses).Error
// 	} else {
// 		err = r.db.Find(&houses).Error
// 	}

// 	if err != nil {
// 		return c.JSON(http.StatusInternalServerError, err.Error())
// 	}

// 	return c.JSON(http.StatusOK, houses)
// }
