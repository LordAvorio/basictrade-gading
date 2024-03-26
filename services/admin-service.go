package services

import (
	"basictrade-gading/models"
	"basictrade-gading/repositories"
	"basictrade-gading/utils"
	"basictrade-gading/utils/helpers"
	"errors"
)

type AdminService struct {
	adminRepo repositories.IAdminRepository
}

type IAdminService interface {
	RegisterAdmin(*models.Admin) error
	LoginAdmin(*models.Admin) (string, error)
}

func NewAdminService(adminRepo repositories.IAdminRepository) *AdminService {
	adminService := AdminService{}
	adminService.adminRepo = adminRepo
	return &adminService
}

func (s *AdminService) RegisterAdmin(admin *models.Admin) error {

	generateKsuid, err := utils.GenerateKSUID()
	if err != nil {
		return err
	}

	hashedPassword := utils.HashPass(admin.Password)
	admin.Password = hashedPassword
	admin.UUID = generateKsuid

	err = s.adminRepo.RegisterAdmin(admin)
	if err != nil {
		return err
	}

	return nil
}

func (s *AdminService) LoginAdmin(admin *models.Admin) (string, error) {

	dataAccount, err := s.adminRepo.GetAccount(admin.Email)
	if err != nil {
		return "", err
	}

	passValidation := utils.PassValidation(admin.Password, dataAccount.Password)
	if !passValidation {
		return "", errors.New("login is failed because password is wrong")
	}

	jwtToken, err := helpers.GenerateToken(dataAccount.ID, dataAccount.Email)
	if err != nil {
		return "", err
	}

	return jwtToken, nil

}
