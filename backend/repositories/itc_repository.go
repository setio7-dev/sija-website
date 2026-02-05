package repositories

import (
	"sijaku-hebat/configs"
	"sijaku-hebat/models"

	"gorm.io/gorm"
)

type ItcRepository interface {
	FindAll() ([]models.Itc, error)
	FindById(id uint) (*models.Itc, error)
	Create(itc *models.Itc) error
	Update(itc *models.Itc) error
	Delete(id uint) error
}

type itcRepository struct {
	db *gorm.DB
}

func NewItcRepository() ItcRepository {
	db := configs.InitDatabaseConfig()
	return &itcRepository{db}
}

func (r *itcRepository) FindAll() ([]models.Itc, error) {
	var itcs []models.Itc
	err := r.db.Preload("Project").Preload("Module").Find(&itcs).Error
	return itcs, err
}

func (r *itcRepository) FindById(id uint) (*models.Itc, error) {
	var itc models.Itc
	err := r.db.Preload("Project").Preload("Module").First(&itc, id).Error
	if err != nil {
		return nil, err
	}

	return &itc, nil
}

func (r *itcRepository) Create(itc *models.Itc) error {
	return r.db.Create(itc).Error
}

func (r *itcRepository) Update(itc *models.Itc) error {
	return r.db.Save(itc).Error
}

func (r *itcRepository) Delete(id uint) error {
	return r.db.Delete(&models.Itc{}, id).Error
}
