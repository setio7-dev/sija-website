package services

import (
	"errors"
	"sijaku-hebat/dtos"
	"sijaku-hebat/helpers"
	"sijaku-hebat/models"
	"sijaku-hebat/repositories"
)

type UserService interface {
	GetAll() ([]models.User, error)
	GetById(id uint) (*models.User, error)
	Create(req dtos.CreateUserDTO) (*models.User, error)
	Update(req dtos.UpdateUserDTO, id uint) (*models.User, error)
	Delete(id uint) error
	Login(req dtos.LoginUserDTO) (string, error)
}

type userService struct {
	repo repositories.UserRepository
}

func NewUserService() UserService {
	repo := repositories.NewUserRepository()
	return &userService{repo}
}

func (s *userService) GetAll() ([]models.User, error) {
	return s.repo.FindAll()
}

func (s *userService) GetById(id uint) (*models.User, error) {
	return s.repo.FindById(id)
}

func (s *userService) Create(req dtos.CreateUserDTO) (*models.User, error) {
	hashedPassword, err := helpers.HashPassword(req.Password)
	if err != nil {
		return nil, err
	}

	user := models.User{
		Nis:        req.Nis,
		Name:       req.Name,
		Password:   hashedPassword,
		Class:      req.Class,
		Phone:      req.Phone,
		CompanyID:  req.CompanyID,
		CategoryID: req.CategoryID,
		Status:     req.Status,
		IsAdmin:    req.IsAdmin,
	}

	if err := s.repo.Create(&user); err != nil {
		return nil, err
	}

	return &user, nil
}

func (s *userService) Update(req dtos.UpdateUserDTO, id uint) (*models.User, error) {
	user, err := s.repo.FindById(id)
	if err != nil {
		return nil, err
	}

	if req.Name != nil {
		user.Name = *req.Name
	}
	if req.Password != nil {
		hashedPassword, err := helpers.HashPassword(*req.Password)
		if err != nil {
			return nil, err
		}
		user.Password = hashedPassword
	}
	if req.Class != nil {
		user.Class = *req.Class
	}
	if req.Phone != nil {
		user.Phone = *req.Phone
	}
	if req.CompanyID != nil {
		user.CompanyID = req.CompanyID
	}
	if req.CategoryID != nil {
		user.CategoryID = req.CategoryID
	}
	if req.Status != nil {
		user.Status = *req.Status
	}
	if req.IsAdmin != nil {
		user.IsAdmin = *req.IsAdmin
	}
	if err := s.repo.Update(user); err != nil {
		return nil, err
	}

	return user, nil
}

func (s *userService) Delete(id uint) error {
	_, err := s.repo.FindById(id)
	if err != nil {
		return err
	}
	return s.repo.Delete(id)
}

func (s *userService) Login(req dtos.LoginUserDTO) (string, error) {
	user, err := s.repo.FindByNis(req.Nis)
	if err != nil {
		return "", err
	}

	if !helpers.CheckPasswordHash(req.Password, user.Password) {
		return "", errors.New("invalid password")
	}

	token, err := helpers.GenerateToken(user.ID)
	if err != nil {
		return "", err
	}

	return token, nil
}
