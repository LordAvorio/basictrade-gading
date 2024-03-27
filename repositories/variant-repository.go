package repositories

import (
	"basictrade-gading/models"

	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
)

type VariantRepository struct {
	db *gorm.DB
}

type IVariantRepository interface {
	CreateVariant(*models.Variant) (*models.Variant, error)
	GetVariant(string) (*models.Variant, error)
}

func NewVariantRepository(db *gorm.DB) *VariantRepository {
	variantRepository := VariantRepository{}
	variantRepository.db = db
	return &variantRepository
}

func (r *VariantRepository) CreateVariant(variant *models.Variant) (*models.Variant, error) {
	err := r.db.Create(&variant).Error
	if err != nil {
		log.Error().Msg(err.Error())
		return nil, err
	}
	return variant, nil
}

func (r *VariantRepository) GetVariant(uuid string) (*models.Variant, error) {
	result := models.Variant{}
	err := r.db.Where("uuid = ?", uuid).Take(&result).Error
	if err != nil {
		log.Error().Msg(err.Error())
		return nil, err
	}
	return &result, nil
}
