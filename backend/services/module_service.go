package services

import (
	"sijaku-hebat/dtos"
	"sijaku-hebat/models"
	"sijaku-hebat/repositories"
)

type ModuleService interface {
	GetAll() ([]models.Module, error)
	GetById(id uint) (*models.Module, error)
	Create(req dtos.CreateModuleDTO, imagePath string, filePath string) (*models.Module, error)
	Update(req dtos.UpdateModuleDTO, imagePath string, filePath string, id uint) (*models.Module, error)
	Delete(id uint) error
}

type moduleService struct {
	repo repositories.ModuleRepository
}

func NewModuleService() ModuleService {
	repo := repositories.NewModuleRepository()
	return &moduleService{repo}
}

func (s *moduleService) GetAll() ([]models.Module, error) {
	return s.repo.FindAll()
}

func (s *moduleService) GetById(id uint) (*models.Module, error) {
	return s.repo.FindById(id)
}

func (s *moduleService) Create(req dtos.CreateModuleDTO, imagePath string, filePath string) (*models.Module, error) {
	module := models.Module{
		Title: req.Title,
		Desc:  req.Desc,
		ItcID: req.ItcID,
		Image: imagePath,
		File:  filePath,
	}

	if err := s.repo.Create(&module); err != nil {
		return nil, err
	}

	return &module, nil
}

func (s *moduleService) Update(req dtos.UpdateModuleDTO, imagePath string, filePath string, id uint) (*models.Module, error) {
	module, err := s.repo.FindById(id)
	if err != nil {
		return nil, err
	}
	if req.Title != nil {
		module.Title = *req.Title
	}
	if req.Desc != nil {
		module.Desc = *req.Desc
	}
	if req.ItcID != nil {
		module.ItcID = *req.ItcID
	}
	if imagePath != "" {
		module.Image = imagePath
	}
	if filePath != "" {
		module.File = filePath
	}

	if err := s.repo.Update(module); err != nil {
		return nil, err
	}

	return module, nil
}

func (s *moduleService) Delete(id uint) error {
	_, err := s.repo.FindById(id)
	if err != nil {
		return err
	}

	return s.repo.Delete(id)
}
