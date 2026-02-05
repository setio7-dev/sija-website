package services

import (
	"sijaku-hebat/dtos"
	"sijaku-hebat/models"
	"sijaku-hebat/repositories"
)

type ProjectService interface {
	GetAll() ([]models.Project, error)
	GetById(id uint) (*models.Project, error)
	Create(req dtos.CreateProjectDTO, imagePath string) (*models.Project, error)
	Update(req dtos.UpdateProjectDTO, imagePath string, id uint) (*models.Project, error)
	Delete(id uint) error
}

type projectService struct {
	repo repositories.ProjectRepository
}

func NewProjectService() ProjectService {
	repo := repositories.NewProjectRepository()
	return &projectService{repo}
}

func (s *projectService) GetAll() ([]models.Project, error) {
	return s.repo.FindAll()
}

func (s *projectService) GetById(id uint) (*models.Project, error) {
	return s.repo.FindById(id)
}

func (s *projectService) Create(req dtos.CreateProjectDTO, imagePath string) (*models.Project, error) {
	project := models.Project{
		Name:  req.Name,
		Desc:  req.Desc,
		Image: imagePath,
		Link:  req.Link,
	}

	if err := s.repo.Create(&project); err != nil {
		return nil, err
	}

	return &project, nil
}

func (s *projectService) Update(req dtos.UpdateProjectDTO, imagePath string, id uint) (*models.Project, error) {
	project, err := s.repo.FindById(id)
	if err != nil {
		return nil, err
	}
	if req.Name != nil {
		project.Name = *req.Name
	}
	if req.Desc != nil {
		project.Desc = *req.Desc
	}
	if imagePath != "" {
		project.Image = imagePath
	}
	if req.Link != nil {
		project.Link = *req.Link
	}
	if err := s.repo.Update(project); err != nil {
		return nil, err
	}

	return project, nil
}

func (s *projectService) Delete(id uint) error {
	_, err := s.repo.FindById(id)
	if err != nil {
		return err
	}
	return s.repo.Delete(id)
}
