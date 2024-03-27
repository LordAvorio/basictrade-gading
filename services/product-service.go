package services

import (
	"basictrade-gading/models"
	"basictrade-gading/repositories"
	"basictrade-gading/utils"
	"basictrade-gading/utils/helpers"
)

type ProductService struct {
	productRepo repositories.IProductRepository
}

type IProductService interface {
	CreateProduct(*models.ProductRequest) (*models.Product, error)
	GetProduct(string) (*models.Product, error)
}

func NewProductService(productRepo repositories.IProductRepository) *ProductService {
	productService := ProductService{}
	productService.productRepo = productRepo
	return &productService
}

func (s *ProductService) CreateProduct(dataRequest *models.ProductRequest) (*models.Product, error) {

	fileName := helpers.RemoveExtension(dataRequest.Image.Filename)

	uploadResult, err := helpers.UploadFile(&dataRequest.Image, fileName)
	if err != nil {
		return nil, err
	}

	generateKsuid, err := utils.GenerateKSUID()
	if err != nil {
		return nil, err
	}

	data := models.Product{
		UUID:     generateKsuid,
		Name:     dataRequest.Name,
		AdminID:  dataRequest.AdminId,
		ImageUrl: uploadResult,
	}

	resultProduct, err := s.productRepo.CreateProduct(&data)
	if err != nil {
		return nil, err
	}

	return resultProduct, nil

}

func (s *ProductService) GetProduct(uuid string) (*models.Product, error) {

	resultProduct, err := s.productRepo.GetProduct(uuid)
	if err != nil {
		return nil, err
	}

	return resultProduct, nil
}
