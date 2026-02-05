package services

import (
	"sijaku-hebat/dtos"
	"sijaku-hebat/models"
	"sijaku-hebat/repositories"
)

type CompanyService interface {
	GetAll() ([]models.Company, error)
	GetById(id uint) (*models.Company, error)
	Create(req dtos.CreateCompanyDTO, imagePath string) (*models.Company, error)
	Update(req dtos.UpdateCompanyDTO, imagePath string, id uint) (*models.Company, error)
	Delete(id uint) error
}

type companyService struct {
	repo repositories.CompanyRepository
}

func NewCompanyService() CompanyService {
	repo := repositories.NewCompanyRepository()
	return &companyService{repo}
}

func (s *companyService) GetAll() ([]models.Company, error) {
	return s.repo.FindAll()
}

func (s *companyService) GetById(id uint) (*models.Company, error) {
	return s.repo.FindById(id)
}

func (s *companyService) Create(req dtos.CreateCompanyDTO, imagePath string) (*models.Company, error) {
	company := models.Company{
		Name:  req.Name,
		Desc:  req.Desc,
		Phone: req.Phone,
		Email: req.Email,
		Image: imagePath,
		Link:  req.Link,
	}

	if err := s.repo.Create(&company); err != nil {
		return nil, err
	}

	return &company, nil
}

func (s *companyService) Update(req dtos.UpdateCompanyDTO, imagePath string, id uint) (*models.Company, error) {
	company, err := s.repo.FindById(id)
	if err != nil {
		return nil, err
	}

	if req.Name != nil {
		company.Name = *req.Name
	}
	if req.Desc != nil {
		company.Desc = *req.Desc
	}
	if req.Phone != nil {
		company.Phone = *req.Phone
	}
	if req.Email != nil {
		company.Email = *req.Email
	}
	if req.Link != nil {
		company.Link = *req.Link
	}
	if imagePath != "" {
		company.Image = imagePath
	}

	if err := s.repo.Update(company); err != nil {
		return nil, err
	}

	return company, nil
}

func (s *companyService) Delete(id uint) error {
	_, err := s.repo.FindById(id)
	if err != nil {
		return err
	}
	return s.repo.Delete(id)
}
