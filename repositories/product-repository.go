package repositories

import (
	"basictrade-gading/models"

	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
)

type ProductRepository struct {
	db *gorm.DB
}

type IProductRepository interface {
	CreateProduct(*models.Product)(*models.Product, error)
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
	productRepository := ProductRepository{}
	productRepository.db = db
	return &productRepository
}

func (r *ProductRepository) CreateProduct(product *models.Product) (*models.Product, error) {

	err := r.db.Create(&product).Error
	if err != nil {
		log.Error().Msg(err.Error())
		return nil, err
	}

	return product, nil

}
