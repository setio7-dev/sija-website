package repositories

import (
	"sijaku-hebat/configs"
	"sijaku-hebat/models"

	"gorm.io/gorm"
)

type ModuleRepository interface {
	FindAll() ([]models.Module, error)
	FindById(id uint) (*models.Module, error)
	Create(module *models.Module) error
	Update(module *models.Module) error
	Delete(id uint) error
}

type moduleRepository struct {
	db *gorm.DB
}

func NewModuleRepository() ModuleRepository {
	db := configs.InitDatabaseConfig()
	return &moduleRepository{db}
}

func (r *moduleRepository) FindAll() ([]models.Module, error) {
	var modules []models.Module
	err := r.db.Preload("Itc").Find(&modules).Error
	return modules, err
}

func (r *moduleRepository) FindById(id uint) (*models.Module, error) {
	var module models.Module
	err := r.db.Preload("Itc").First(&module, id).Error
	if err != nil {
		return nil, err
	}
	return &module, nil
}

func (r *moduleRepository) Create(module *models.Module) error {
	return r.db.Create(module).Error
}

func (r *moduleRepository) Update(module *models.Module) error {
	return r.db.Save(module).Error
}

func (r *moduleRepository) Delete(id uint) error {
	return r.db.Delete(&models.Module{}, id).Error
}
