package services

import (
	"basictrade-gading/models"
	"basictrade-gading/repositories"
	"basictrade-gading/utils"

	"github.com/google/uuid"
)

type AdminService struct {
	adminRepo repositories.IAdminRepository
}

type IAdminService interface {
	RegisterAdmin(*models.Admin) error
}

func NewAdminService(adminRepo repositories.IAdminRepository) *AdminService {
	adminService := AdminService{}
	adminService.adminRepo = adminRepo
	return &adminService
}

func (s *AdminService) RegisterAdmin(admin *models.Admin) error {

	uuidAdmin := uuid.NewString()
	hashedPassword := utils.HashPass(admin.Password)
	admin.Password = hashedPassword
	admin.UUID = uuidAdmin

	err := s.adminRepo.RegisterAdmin(admin)
	if err != nil {
		return err
	}

	return nil
}
