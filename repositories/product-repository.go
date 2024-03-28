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
	CreateProduct(*models.Product) (*models.Product, error)
	GetProduct(string) (*models.Product, error)
	GetProducts(int, int, string) (*[]models.Product, error)
	GetTotalProduct(string) (int, error)
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

func (r *ProductRepository) GetProduct(uuid string) (*models.Product, error) {

	result := models.Product{}
	err := r.db.Where("uuid = ?", uuid).Take(&result).Error
	if err != nil {
		log.Error().Msg(err.Error())
		return nil, err
	}
	return &result, nil

}

func (r *ProductRepository) GetProducts(offset, limit int, nameFilter string) (*[]models.Product, error) {
	var result []models.Product

	queryStatement := r.db.Offset(offset).Limit(limit)

	if nameFilter != "" {
		queryStatement = queryStatement.Where("name LIKE ?", "%"+nameFilter+"%")
	}

	err := queryStatement.Find(&result).Error
	if err != nil {
		log.Error().Msg(err.Error())
		return nil, err
	}

	return &result, nil
}

func (r *ProductRepository) GetTotalProduct(nameFilter string) (int, error) {

	var totalProduct int64

	if nameFilter != "" {
		err := r.db.Model(&models.Product{}).Where("name LIKE ?", "%"+nameFilter+"%").Count(&totalProduct).Error
		if err != nil {
			log.Error().Msg(err.Error())
			return 0, err
		}
	} else {
		err := r.db.Model(&models.Product{}).Count(&totalProduct).Error
		if err != nil {
			log.Error().Msg(err.Error())
			return 0, err
		}
	}

	return int(totalProduct), nil

}
