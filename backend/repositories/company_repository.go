package repositories

import (
	"sijaku-hebat/configs"
	"sijaku-hebat/models"

	"gorm.io/gorm"
)

type CompanyRepository interface {
	FindAll() ([]models.Company, error)
	FindById(id uint) (*models.Company, error)
	Create(company *models.Company) error
	Update(company *models.Company) error
	Delete(id uint) error
}

type companyRepository struct {
	db *gorm.DB
}

func NewCompanyRepository() CompanyRepository {
	db := configs.InitDatabaseConfig()
	return &companyRepository{db}
}

func (r *companyRepository) FindAll() ([]models.Company, error) {
	var companies []models.Company
	err := r.db.Find(&companies).Error
	return companies, err
}

func (r *companyRepository) FindById(id uint) (*models.Company, error) {
	var company models.Company
	err := r.db.First(&company, id).Error

	if err != nil {
		return nil, err
	}

	return &company, nil
}

func (r *companyRepository) Create(company *models.Company) error {
	return r.db.Create(company).Error
}

func (r *companyRepository) Update(company *models.Company) error {
	return r.db.Save(company).Error
}

func (r *companyRepository) Delete(id uint) error {
	return r.db.Delete(&models.Company{}, id).Error
}
