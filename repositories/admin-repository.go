package repositories

import (
	"basictrade-gading/models"

	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
)

type AdminRepository struct {
	db *gorm.DB
}

type IAdminRepository interface {
	RegisterAdmin(*models.Admin) error
	GetAccount(string) (*models.Admin, error)
}

func NewAdminRepository(db *gorm.DB) *AdminRepository {
	adminRepository := AdminRepository{}
	adminRepository.db = db
	return &adminRepository
}

func (r *AdminRepository) RegisterAdmin(admin *models.Admin) error {

	err := r.db.Create(&admin).Error
	if err != nil {
		log.Error().Msg(err.Error())
		return err
	}

	return nil

}

func (r *AdminRepository) GetAccount(email string) (*models.Admin, error) {

	dataAdmin := models.Admin{}
	err := r.db.Where("email = ?", email).Take(&dataAdmin).Error
	if err != nil {
		log.Error().Msg(err.Error())
		return nil, err
	}

	return &dataAdmin, nil
}
