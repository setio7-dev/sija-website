package services

import (
	"sijaku-hebat/dtos"
	"sijaku-hebat/models"
	"sijaku-hebat/repositories"
)

type ItcService interface {
	GetAll() ([]models.Itc, error)
	GetById(id uint) (*models.Itc, error)
	Create(req dtos.CreateItcDTO, imagePath string) (*models.Itc, error)
	Update(req dtos.UpdateItcDTO, imagePath string, id uint) (*models.Itc, error)
	Delete(id uint) error
}

type itcService struct {
	repo repositories.ItcRepository
}

func NewItcService() ItcService {
	repo := repositories.NewItcRepository()
	return &itcService{repo}
}

func (s *itcService) GetAll() ([]models.Itc, error) {
	return s.repo.FindAll()
}

func (s *itcService) GetById(id uint) (*models.Itc, error) {
	return s.repo.FindById(id)
}

func (s *itcService) Create(req dtos.CreateItcDTO, imagePath string) (*models.Itc, error) {
	itc := models.Itc{
		Name:      req.Name,
		Desc:      req.Desc,
		Image:     imagePath,
		ProjectID: req.ProjectID,
	}

	if err := s.repo.Create(&itc); err != nil {
		return nil, err
	}

	return &itc, nil
}

func (s *itcService) Update(req dtos.UpdateItcDTO, imagePath string, id uint) (*models.Itc, error) {
	itc, err := s.repo.FindById(id)
	if err != nil {
		return nil, err
	}
	if req.Name != nil {
		itc.Name = *req.Name
	}
	if req.Desc != nil {
		itc.Desc = *req.Desc
	}
	if imagePath != "" {
		itc.Image = imagePath
	}
	if req.ProjectID != nil {
		itc.ProjectID = *req.ProjectID
	}
	if err := s.repo.Update(itc); err != nil {
		return nil, err
	}

	return itc, nil
}

func (s *itcService) Delete(id uint) error {
	return s.repo.Delete(id)
}
