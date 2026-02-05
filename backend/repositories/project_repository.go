package repositories

import (
	"sijaku-hebat/configs"
	"sijaku-hebat/models"

	"gorm.io/gorm"
)

type ProjectRepository interface {
	FindAll() ([]models.Project, error)
	FindById(id uint) (*models.Project, error)
	Create(project *models.Project) error
	Update(project *models.Project) error
	Delete(id uint) error
}

type projectRepository struct {
	db *gorm.DB
}

func NewProjectRepository() ProjectRepository {
	db := configs.InitDatabaseConfig()
	return &projectRepository{db}
}

func (r *projectRepository) FindAll() ([]models.Project, error) {
	var projects []models.Project
	err := r.db.Preload("Itc").Find(&projects).Error
	return projects, err
}

func (r *projectRepository) FindById(id uint) (*models.Project, error) {
	var project models.Project
	err := r.db.Preload("Itc").First(&project, id).Error

	if err != nil {
		return nil, err
	}

	return &project, nil
}

func (r *projectRepository) Create(project *models.Project) error {
	return r.db.Create(project).Error
}

func (r *projectRepository) Update(project *models.Project) error {
	return r.db.Save(project).Error
}

func (r *projectRepository) Delete(id uint) error {
	return r.db.Delete(&models.Project{}, id).Error
}
