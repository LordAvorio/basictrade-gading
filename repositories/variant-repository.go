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
	GetVariants(int, int, string) (*[]models.Variant, error)
	GetTotalVariant(string) (int, error)
	UpdateVariant(*models.Variant)(*models.Variant, error)
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

func (r *VariantRepository) GetVariants(offset, limit int, nameFilter string) (*[]models.Variant, error) {
	var result []models.Variant

	queryStatement := r.db.Offset(offset).Limit(limit)

	if nameFilter != "" {
		queryStatement = queryStatement.Where("variant_name LIKE ?", "%"+nameFilter+"%")
	}

	err := queryStatement.Find(&result).Error
	if err != nil {
		log.Error().Msg(err.Error())
		return nil, err
	}

	return &result, nil
}

func (r *VariantRepository) GetTotalVariant(nameFilter string) (int, error) {

	var totalVariant int64

	if nameFilter != "" {
		err := r.db.Model(&models.Variant{}).Where("variant_name LIKE ?", "%"+nameFilter+"%").Count(&totalVariant).Error
		if err != nil {
			log.Error().Msg(err.Error())
			return 0, err
		}
	} else {
		err := r.db.Model(&models.Variant{}).Count(&totalVariant).Error
		if err != nil {
			log.Error().Msg(err.Error())
			return 0, err
		}
	}

	return int(totalVariant), nil

}

func (r *VariantRepository) UpdateVariant(variant *models.Variant) (*models.Variant, error) {

	err := r.db.Model(&variant).Where("id = ?", variant.ID).Updates(variant).Error
	if err != nil {
		return nil, err
	}

	return variant, nil

}