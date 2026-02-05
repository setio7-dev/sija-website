package repositories

import (
	"sijaku-hebat/configs"
	"sijaku-hebat/models"

	"gorm.io/gorm"
)

type UserRepository interface {
	FindAll() ([]models.User, error)
	FindById(id uint) (*models.User, error)
	FindByNis(nis string) (*models.User, error)
	Create(user *models.User) error
	Update(user *models.User) error
	Delete(id uint) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository() UserRepository {
	db := configs.InitDatabaseConfig()
	return &userRepository{db}
}

func (r *userRepository) FindAll() ([]models.User, error) {
	var users []models.User
	err := r.db.Preload("Company").Preload("Category").Find(&users).Error
	return users, err
}

func (r *userRepository) FindById(id uint) (*models.User, error) {
	var user models.User
	err := r.db.Preload("Company").Preload("Category").First(&user, id).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *userRepository) FindByNis(nis string) (*models.User, error) {
	var user models.User
	err := r.db.Preload("Company").Preload("Category").Where("nis = ?", nis).First(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *userRepository) Create(user *models.User) error {
	return r.db.Create(user).Error
}

func (r *userRepository) Update(user *models.User) error {
	return r.db.Save(user).Error
}

func (r *userRepository) Delete(id uint) error {
	return r.db.Delete(&models.User{}, id).Error
}
